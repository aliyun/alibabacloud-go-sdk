// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPodLogsResponseBody interface {
	dara.Model
	String() string
	GoString() string
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
	// The job ID.
	//
	// example:
	//
	// dlc-20210126170216-******
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The logs.
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
	// The request ID which is used for diagnostics and Q\\&A.
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
	return dara.Validate(s)
}
