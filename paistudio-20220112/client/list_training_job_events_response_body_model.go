// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTrainingJobEventsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetEvents(v []*string) *ListTrainingJobEventsResponseBody
	GetEvents() []*string
	SetRequestId(v string) *ListTrainingJobEventsResponseBody
	GetRequestId() *string
	SetTotalCount(v string) *ListTrainingJobEventsResponseBody
	GetTotalCount() *string
}

type ListTrainingJobEventsResponseBody struct {
	Events []*string `json:"Events,omitempty" xml:"Events,omitempty" type:"Repeated"`
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3C83E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 23
	TotalCount *string `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListTrainingJobEventsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListTrainingJobEventsResponseBody) GoString() string {
	return s.String()
}

func (s *ListTrainingJobEventsResponseBody) GetEvents() []*string {
	return s.Events
}

func (s *ListTrainingJobEventsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListTrainingJobEventsResponseBody) GetTotalCount() *string {
	return s.TotalCount
}

func (s *ListTrainingJobEventsResponseBody) SetEvents(v []*string) *ListTrainingJobEventsResponseBody {
	s.Events = v
	return s
}

func (s *ListTrainingJobEventsResponseBody) SetRequestId(v string) *ListTrainingJobEventsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListTrainingJobEventsResponseBody) SetTotalCount(v string) *ListTrainingJobEventsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListTrainingJobEventsResponseBody) Validate() error {
	return dara.Validate(s)
}
