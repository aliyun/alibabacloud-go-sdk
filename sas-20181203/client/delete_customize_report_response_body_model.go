// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteCustomizeReportResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteCustomizeReportResponseBody
	GetRequestId() *string
}

type DeleteCustomizeReportResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 90593A3B-85CE-5D87-A430-726D0B87****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteCustomizeReportResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteCustomizeReportResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteCustomizeReportResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteCustomizeReportResponseBody) SetRequestId(v string) *DeleteCustomizeReportResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteCustomizeReportResponseBody) Validate() error {
	return dara.Validate(s)
}
