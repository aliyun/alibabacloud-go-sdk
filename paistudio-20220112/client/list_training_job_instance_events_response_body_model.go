// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTrainingJobInstanceEventsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetEvents(v []*string) *ListTrainingJobInstanceEventsResponseBody
	GetEvents() []*string
	SetRequestId(v string) *ListTrainingJobInstanceEventsResponseBody
	GetRequestId() *string
	SetTotalCount(v string) *ListTrainingJobInstanceEventsResponseBody
	GetTotalCount() *string
}

type ListTrainingJobInstanceEventsResponseBody struct {
	Events []*string `json:"Events,omitempty" xml:"Events,omitempty" type:"Repeated"`
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3C83E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 1
	TotalCount *string `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListTrainingJobInstanceEventsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListTrainingJobInstanceEventsResponseBody) GoString() string {
	return s.String()
}

func (s *ListTrainingJobInstanceEventsResponseBody) GetEvents() []*string {
	return s.Events
}

func (s *ListTrainingJobInstanceEventsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListTrainingJobInstanceEventsResponseBody) GetTotalCount() *string {
	return s.TotalCount
}

func (s *ListTrainingJobInstanceEventsResponseBody) SetEvents(v []*string) *ListTrainingJobInstanceEventsResponseBody {
	s.Events = v
	return s
}

func (s *ListTrainingJobInstanceEventsResponseBody) SetRequestId(v string) *ListTrainingJobInstanceEventsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListTrainingJobInstanceEventsResponseBody) SetTotalCount(v string) *ListTrainingJobInstanceEventsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListTrainingJobInstanceEventsResponseBody) Validate() error {
	return dara.Validate(s)
}
