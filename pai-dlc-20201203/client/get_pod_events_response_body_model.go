// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPodEventsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetEvents(v []*string) *GetPodEventsResponseBody
	GetEvents() []*string
	SetJobId(v string) *GetPodEventsResponseBody
	GetJobId() *string
	SetPodId(v string) *GetPodEventsResponseBody
	GetPodId() *string
	SetPodUid(v string) *GetPodEventsResponseBody
	GetPodUid() *string
	SetRequestId(v string) *GetPodEventsResponseBody
	GetRequestId() *string
}

type GetPodEventsResponseBody struct {
	// The events returned.
	Events []*string `json:"Events,omitempty" xml:"Events,omitempty" type:"Repeated"`
	// The job ID.
	//
	// example:
	//
	// dlc-20210126170216-*****
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The node ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// dlc-20210126170216-*****-chief-0
	PodId *string `json:"PodId,omitempty" xml:"PodId,omitempty"`
	// The node UID.
	//
	// example:
	//
	// 94a7cc7c-0033-48b5-85bd-71c63592c268
	PodUid *string `json:"PodUid,omitempty" xml:"PodUid,omitempty"`
	// The request ID, which can be used for troubleshooting.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3C83E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetPodEventsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetPodEventsResponseBody) GoString() string {
	return s.String()
}

func (s *GetPodEventsResponseBody) GetEvents() []*string {
	return s.Events
}

func (s *GetPodEventsResponseBody) GetJobId() *string {
	return s.JobId
}

func (s *GetPodEventsResponseBody) GetPodId() *string {
	return s.PodId
}

func (s *GetPodEventsResponseBody) GetPodUid() *string {
	return s.PodUid
}

func (s *GetPodEventsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetPodEventsResponseBody) SetEvents(v []*string) *GetPodEventsResponseBody {
	s.Events = v
	return s
}

func (s *GetPodEventsResponseBody) SetJobId(v string) *GetPodEventsResponseBody {
	s.JobId = &v
	return s
}

func (s *GetPodEventsResponseBody) SetPodId(v string) *GetPodEventsResponseBody {
	s.PodId = &v
	return s
}

func (s *GetPodEventsResponseBody) SetPodUid(v string) *GetPodEventsResponseBody {
	s.PodUid = &v
	return s
}

func (s *GetPodEventsResponseBody) SetRequestId(v string) *GetPodEventsResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetPodEventsResponseBody) Validate() error {
	return dara.Validate(s)
}
