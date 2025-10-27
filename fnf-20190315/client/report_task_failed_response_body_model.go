// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReportTaskFailedResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetEventId(v int64) *ReportTaskFailedResponseBody
	GetEventId() *int64
	SetRequestId(v string) *ReportTaskFailedResponseBody
	GetRequestId() *string
}

type ReportTaskFailedResponseBody struct {
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

func (s ReportTaskFailedResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ReportTaskFailedResponseBody) GoString() string {
	return s.String()
}

func (s *ReportTaskFailedResponseBody) GetEventId() *int64 {
	return s.EventId
}

func (s *ReportTaskFailedResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ReportTaskFailedResponseBody) SetEventId(v int64) *ReportTaskFailedResponseBody {
	s.EventId = &v
	return s
}

func (s *ReportTaskFailedResponseBody) SetRequestId(v string) *ReportTaskFailedResponseBody {
	s.RequestId = &v
	return s
}

func (s *ReportTaskFailedResponseBody) Validate() error {
	return dara.Validate(s)
}
