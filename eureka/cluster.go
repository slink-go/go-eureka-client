package eureka

import (
	"go.slink.ws/logging"
	"net/url"
	"strings"
)

type Cluster struct {
	Leader   string   `json:"leader"`
	Machines []string `json:"machines"`
	logger   logging.Logger
}

func NewCluster(machines []string) *Cluster {
	// if an empty slice was sent in then just assume HTTP 4001 on localhost
	if len(machines) == 0 {
		machines = []string{"http://127.0.0.1:4001"}
	}

	// default leader and machines
	return &Cluster{
		Leader:   machines[0],
		Machines: machines,
		logger:   logging.GetLogger("eureka-cluster"),
	}
}

// switchLeader switch the current leader to machines[num]
func (cl *Cluster) switchLeader(num int) {
	cl.logger.Debug("switch.leader[from %v to %v]",
		cl.Leader, cl.Machines[num])

	cl.Leader = cl.Machines[num]
}

func (cl *Cluster) updateFromStr(machines string) {
	cl.Machines = strings.Split(machines, ", ")
}

func (cl *Cluster) updateLeader(leader string) {
	cl.logger.Debug("update.leader[%s,%s]", cl.Leader, leader)
	cl.Leader = leader
}

func (cl *Cluster) updateLeaderFromURL(u *url.URL) {
	var leader string
	if u.Scheme == "" {
		leader = "http://" + u.Host
	} else {
		leader = u.Scheme + "://" + u.Host
	}
	cl.updateLeader(leader)
}
