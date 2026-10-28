// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPodItem interface {
	dara.Model
	String() string
	GoString() string
	SetGmtCreateTime(v string) *PodItem
	GetGmtCreateTime() *string
	SetGmtFinishTime(v string) *PodItem
	GetGmtFinishTime() *string
	SetGmtStartTime(v string) *PodItem
	GetGmtStartTime() *string
	SetHistoryPods(v []*PodItem) *PodItem
	GetHistoryPods() []*PodItem
	SetIp(v string) *PodItem
	GetIp() *string
	SetNodeName(v string) *PodItem
	GetNodeName() *string
	SetPodId(v string) *PodItem
	GetPodId() *string
	SetPodIp(v string) *PodItem
	GetPodIp() *string
	SetPodIps(v []*PodNetworkInterface) *PodItem
	GetPodIps() []*PodNetworkInterface
	SetPodUid(v string) *PodItem
	GetPodUid() *string
	SetStatus(v string) *PodItem
	GetStatus() *string
	SetSubStatus(v string) *PodItem
	GetSubStatus() *string
	SetType(v string) *PodItem
	GetType() *string
	SetSupportedProfilingTypes(v string) *PodItem
	GetSupportedProfilingTypes() *string
}

type PodItem struct {
	// The time when the pod was created (UTC).
	//
	// example:
	//
	// 2021-01-12T14:36:01Z
	GmtCreateTime *string `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	// The time when the pod was completed (UTC).
	//
	// example:
	//
	// 2021-01-12T15:36:05Z
	GmtFinishTime *string `json:"GmtFinishTime,omitempty" xml:"GmtFinishTime,omitempty"`
	// The time when the pod was started (UTC).
	//
	// example:
	//
	// 2021-01-12T14:36:05Z
	GmtStartTime *string `json:"GmtStartTime,omitempty" xml:"GmtStartTime,omitempty"`
	// The historical pods.
	HistoryPods []*PodItem `json:"HistoryPods,omitempty" xml:"HistoryPods,omitempty" type:"Repeated"`
	// Pod Ip
	//
	// example:
	//
	// 10.0.1.2
	Ip *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	// The name of the compute node where the pod is located.
	//
	// example:
	//
	// i-a
	NodeName *string `json:"NodeName,omitempty" xml:"NodeName,omitempty"`
	// Pod Id
	//
	// example:
	//
	// dlc-20210126170216-mtl37ge7gkvdz-worker-0
	PodId *string `json:"PodId,omitempty" xml:"PodId,omitempty"`
	// The IP address of the pod.
	//
	// example:
	//
	// 10.0.1.2
	PodIp *string `json:"PodIp,omitempty" xml:"PodIp,omitempty"`
	// The IP addresses of the pod. For Lingjun AI Computing Service: user VPC net0, management VPC net, and pod-to-pod communication ethx. For general-purpose computing: user VPC eth1 with dedicated gateway default route, management VPC eth0 with public gateway default route, and pod-to-pod communication net eRDMA.
	PodIps []*PodNetworkInterface `json:"PodIps,omitempty" xml:"PodIps,omitempty" type:"Repeated"`
	// Pod Uid
	//
	// example:
	//
	// fe846462-af2c-4521-bd6f-96787a57591d
	PodUid *string `json:"PodUid,omitempty" xml:"PodUid,omitempty"`
	// The status of the pod.
	//
	// example:
	//
	// Stopped
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The substatus of the pod.
	//
	// example:
	//
	// Restarting
	SubStatus *string `json:"SubStatus,omitempty" xml:"SubStatus,omitempty"`
	// The type of the pod.
	//
	// example:
	//
	// Worker
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The list of supported profiling types, separated by commas, such as sysom. An empty string indicates that profiling is not supported.
	//
	// example:
	//
	// sysom
	SupportedProfilingTypes *string `json:"supportedProfilingTypes,omitempty" xml:"supportedProfilingTypes,omitempty"`
}

func (s PodItem) String() string {
	return dara.Prettify(s)
}

func (s PodItem) GoString() string {
	return s.String()
}

func (s *PodItem) GetGmtCreateTime() *string {
	return s.GmtCreateTime
}

func (s *PodItem) GetGmtFinishTime() *string {
	return s.GmtFinishTime
}

func (s *PodItem) GetGmtStartTime() *string {
	return s.GmtStartTime
}

func (s *PodItem) GetHistoryPods() []*PodItem {
	return s.HistoryPods
}

func (s *PodItem) GetIp() *string {
	return s.Ip
}

func (s *PodItem) GetNodeName() *string {
	return s.NodeName
}

func (s *PodItem) GetPodId() *string {
	return s.PodId
}

func (s *PodItem) GetPodIp() *string {
	return s.PodIp
}

func (s *PodItem) GetPodIps() []*PodNetworkInterface {
	return s.PodIps
}

func (s *PodItem) GetPodUid() *string {
	return s.PodUid
}

func (s *PodItem) GetStatus() *string {
	return s.Status
}

func (s *PodItem) GetSubStatus() *string {
	return s.SubStatus
}

func (s *PodItem) GetType() *string {
	return s.Type
}

func (s *PodItem) GetSupportedProfilingTypes() *string {
	return s.SupportedProfilingTypes
}

func (s *PodItem) SetGmtCreateTime(v string) *PodItem {
	s.GmtCreateTime = &v
	return s
}

func (s *PodItem) SetGmtFinishTime(v string) *PodItem {
	s.GmtFinishTime = &v
	return s
}

func (s *PodItem) SetGmtStartTime(v string) *PodItem {
	s.GmtStartTime = &v
	return s
}

func (s *PodItem) SetHistoryPods(v []*PodItem) *PodItem {
	s.HistoryPods = v
	return s
}

func (s *PodItem) SetIp(v string) *PodItem {
	s.Ip = &v
	return s
}

func (s *PodItem) SetNodeName(v string) *PodItem {
	s.NodeName = &v
	return s
}

func (s *PodItem) SetPodId(v string) *PodItem {
	s.PodId = &v
	return s
}

func (s *PodItem) SetPodIp(v string) *PodItem {
	s.PodIp = &v
	return s
}

func (s *PodItem) SetPodIps(v []*PodNetworkInterface) *PodItem {
	s.PodIps = v
	return s
}

func (s *PodItem) SetPodUid(v string) *PodItem {
	s.PodUid = &v
	return s
}

func (s *PodItem) SetStatus(v string) *PodItem {
	s.Status = &v
	return s
}

func (s *PodItem) SetSubStatus(v string) *PodItem {
	s.SubStatus = &v
	return s
}

func (s *PodItem) SetType(v string) *PodItem {
	s.Type = &v
	return s
}

func (s *PodItem) SetSupportedProfilingTypes(v string) *PodItem {
	s.SupportedProfilingTypes = &v
	return s
}

func (s *PodItem) Validate() error {
	if s.HistoryPods != nil {
		for _, item := range s.HistoryPods {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.PodIps != nil {
		for _, item := range s.PodIps {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
