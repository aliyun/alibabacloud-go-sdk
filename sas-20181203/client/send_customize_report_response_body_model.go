// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSendCustomizeReportResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *SendCustomizeReportResponseBody
	GetRequestId() *string
}

type SendCustomizeReportResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 2BEA397D-1FD0-5C79-AB24-EC051158****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SendCustomizeReportResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SendCustomizeReportResponseBody) GoString() string {
	return s.String()
}

func (s *SendCustomizeReportResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SendCustomizeReportResponseBody) SetRequestId(v string) *SendCustomizeReportResponseBody {
	s.RequestId = &v
	return s
}

func (s *SendCustomizeReportResponseBody) Validate() error {
	return dara.Validate(s)
}
