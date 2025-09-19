// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateCustomizeReportStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateCustomizeReportStatusResponseBody
	GetRequestId() *string
}

type UpdateCustomizeReportStatusResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 64C76BEE-6A47-54D9-BD91-BD3E8A1B****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateCustomizeReportStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateCustomizeReportStatusResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateCustomizeReportStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateCustomizeReportStatusResponseBody) SetRequestId(v string) *UpdateCustomizeReportStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateCustomizeReportStatusResponseBody) Validate() error {
	return dara.Validate(s)
}
