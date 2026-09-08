// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPodLogsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetContainerInfo(v *ContainerInfo) *GetPodLogsResponseBody
	GetContainerInfo() *ContainerInfo
	SetContainers(v string) *GetPodLogsResponseBody
	GetContainers() *string
	SetJobId(v string) *GetPodLogsResponseBody
	GetJobId() *string
	SetLogs(v []*string) *GetPodLogsResponseBody
	GetLogs() []*string
	SetPodId(v string) *GetPodLogsResponseBody
	GetPodId() *string
	SetPodUid(v string) *GetPodLogsResponseBody
	GetPodUid() *string
	SetRequestId(v string) *GetPodLogsResponseBody
	GetRequestId() *string
}

type GetPodLogsResponseBody struct {
	// The container information that may be associated with the node.
	ContainerInfo *ContainerInfo `json:"ContainerInfo,omitempty" xml:"ContainerInfo,omitempty"`
	// The containers used to filter logs. Separate multiple container names with commas (,).
	//
	// example:
	//
	// pytorch,aimaster-worker
	Containers *string `json:"Containers,omitempty" xml:"Containers,omitempty"`
	// The job ID.
	//
	// example:
	//
	// dlc-20210126170216-******
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The log list.
	Logs []*string `json:"Logs,omitempty" xml:"Logs,omitempty" type:"Repeated"`
	// The node ID.
	//
	// example:
	//
	// dlc-20210126170216-****-chief-0
	PodId *string `json:"PodId,omitempty" xml:"PodId,omitempty"`
	// The instance UID.
	//
	// example:
	//
	// 94a7cc7c-0033-48b5-85bd-71c63592c268
	PodUid *string `json:"PodUid,omitempty" xml:"PodUid,omitempty"`
	// The request ID for this call, used for diagnostics and troubleshooting.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3C83E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetPodLogsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetPodLogsResponseBody) GoString() string {
	return s.String()
}

func (s *GetPodLogsResponseBody) GetContainerInfo() *ContainerInfo {
	return s.ContainerInfo
}

func (s *GetPodLogsResponseBody) GetContainers() *string {
	return s.Containers
}

func (s *GetPodLogsResponseBody) GetJobId() *string {
	return s.JobId
}

func (s *GetPodLogsResponseBody) GetLogs() []*string {
	return s.Logs
}

func (s *GetPodLogsResponseBody) GetPodId() *string {
	return s.PodId
}

func (s *GetPodLogsResponseBody) GetPodUid() *string {
	return s.PodUid
}

func (s *GetPodLogsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetPodLogsResponseBody) SetContainerInfo(v *ContainerInfo) *GetPodLogsResponseBody {
	s.ContainerInfo = v
	return s
}

func (s *GetPodLogsResponseBody) SetContainers(v string) *GetPodLogsResponseBody {
	s.Containers = &v
	return s
}

func (s *GetPodLogsResponseBody) SetJobId(v string) *GetPodLogsResponseBody {
	s.JobId = &v
	return s
}

func (s *GetPodLogsResponseBody) SetLogs(v []*string) *GetPodLogsResponseBody {
	s.Logs = v
	return s
}

func (s *GetPodLogsResponseBody) SetPodId(v string) *GetPodLogsResponseBody {
	s.PodId = &v
	return s
}

func (s *GetPodLogsResponseBody) SetPodUid(v string) *GetPodLogsResponseBody {
	s.PodUid = &v
	return s
}

func (s *GetPodLogsResponseBody) SetRequestId(v string) *GetPodLogsResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetPodLogsResponseBody) Validate() error {
	if s.ContainerInfo != nil {
		if err := s.ContainerInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}
