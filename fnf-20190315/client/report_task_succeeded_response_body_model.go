// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReportTaskSucceededResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetEventId(v int64) *ReportTaskSucceededResponseBody
	GetEventId() *int64
	SetRequestId(v string) *ReportTaskSucceededResponseBody
	GetRequestId() *string
}

type ReportTaskSucceededResponseBody struct {
	// The ID of the event.
	//
	// example:
	//
	// 1
	EventId *int64 `json:"EventId,omitempty" xml:"EventId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// testRequestId
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ReportTaskSucceededResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ReportTaskSucceededResponseBody) GoString() string {
	return s.String()
}

func (s *ReportTaskSucceededResponseBody) GetEventId() *int64 {
	return s.EventId
}

func (s *ReportTaskSucceededResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ReportTaskSucceededResponseBody) SetEventId(v int64) *ReportTaskSucceededResponseBody {
	s.EventId = &v
	return s
}

func (s *ReportTaskSucceededResponseBody) SetRequestId(v string) *ReportTaskSucceededResponseBody {
	s.RequestId = &v
	return s
}

func (s *ReportTaskSucceededResponseBody) Validate() error {
	return dara.Validate(s)
}
