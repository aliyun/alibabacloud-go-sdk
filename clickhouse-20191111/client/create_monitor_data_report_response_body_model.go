// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateMonitorDataReportResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateMonitorDataReportResponseBody
	GetRequestId() *string
}

type CreateMonitorDataReportResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 94F92113-FF63-5E57-8401-6FE123AD11DD
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateMonitorDataReportResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateMonitorDataReportResponseBody) GoString() string {
	return s.String()
}

func (s *CreateMonitorDataReportResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateMonitorDataReportResponseBody) SetRequestId(v string) *CreateMonitorDataReportResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateMonitorDataReportResponseBody) Validate() error {
	return dara.Validate(s)
}
