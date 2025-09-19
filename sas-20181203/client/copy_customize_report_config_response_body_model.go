// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCopyCustomizeReportConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CopyCustomizeReportConfigResponseBody
	GetRequestId() *string
}

type CopyCustomizeReportConfigResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 1EE7B150-D67E-53FD-A52D-3E8E669A****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CopyCustomizeReportConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CopyCustomizeReportConfigResponseBody) GoString() string {
	return s.String()
}

func (s *CopyCustomizeReportConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CopyCustomizeReportConfigResponseBody) SetRequestId(v string) *CopyCustomizeReportConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *CopyCustomizeReportConfigResponseBody) Validate() error {
	return dara.Validate(s)
}
