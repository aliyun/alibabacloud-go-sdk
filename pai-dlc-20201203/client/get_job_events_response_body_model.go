// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetJobEventsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetEvents(v []*string) *GetJobEventsResponseBody
	GetEvents() []*string
	SetJobId(v string) *GetJobEventsResponseBody
	GetJobId() *string
	SetRequestId(v string) *GetJobEventsResponseBody
	GetRequestId() *string
}

type GetJobEventsResponseBody struct {
	// The events.
	Events []*string `json:"Events,omitempty" xml:"Events,omitempty" type:"Repeated"`
	// The job ID.
	//
	// example:
	//
	// dlc-20210126170216-******
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The request ID, which can be used for troubleshooting.
	//
	// example:
	//
	// 78F6FCE2-278F-4C4A-A6B7-DD8ECEA9C456
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetJobEventsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetJobEventsResponseBody) GoString() string {
	return s.String()
}

func (s *GetJobEventsResponseBody) GetEvents() []*string {
	return s.Events
}

func (s *GetJobEventsResponseBody) GetJobId() *string {
	return s.JobId
}

func (s *GetJobEventsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetJobEventsResponseBody) SetEvents(v []*string) *GetJobEventsResponseBody {
	s.Events = v
	return s
}

func (s *GetJobEventsResponseBody) SetJobId(v string) *GetJobEventsResponseBody {
	s.JobId = &v
	return s
}

func (s *GetJobEventsResponseBody) SetRequestId(v string) *GetJobEventsResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetJobEventsResponseBody) Validate() error {
	return dara.Validate(s)
}
