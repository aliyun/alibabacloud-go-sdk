// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetGtmMonitorStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *SetGtmMonitorStatusResponseBody
	GetRequestId() *string
}

type SetGtmMonitorStatusResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 6856BCF6-11D6-4D7E-AC53-FD579933522B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SetGtmMonitorStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SetGtmMonitorStatusResponseBody) GoString() string {
	return s.String()
}

func (s *SetGtmMonitorStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SetGtmMonitorStatusResponseBody) SetRequestId(v string) *SetGtmMonitorStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *SetGtmMonitorStatusResponseBody) Validate() error {
	return dara.Validate(s)
}
