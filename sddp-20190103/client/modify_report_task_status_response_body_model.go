// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyReportTaskStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyReportTaskStatusResponseBody
	GetRequestId() *string
}

type ModifyReportTaskStatusResponseBody struct {
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// 208B016D-4CB9-4A85-96A5-0B8ED1EBF271
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyReportTaskStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyReportTaskStatusResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyReportTaskStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyReportTaskStatusResponseBody) SetRequestId(v string) *ModifyReportTaskStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyReportTaskStatusResponseBody) Validate() error {
	return dara.Validate(s)
}
