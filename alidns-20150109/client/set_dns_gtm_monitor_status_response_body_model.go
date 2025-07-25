// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetDnsGtmMonitorStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *SetDnsGtmMonitorStatusResponseBody
	GetRequestId() *string
}

type SetDnsGtmMonitorStatusResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 6856BCF6-11D6-4D7E-AC53-FD579933522B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SetDnsGtmMonitorStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SetDnsGtmMonitorStatusResponseBody) GoString() string {
	return s.String()
}

func (s *SetDnsGtmMonitorStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SetDnsGtmMonitorStatusResponseBody) SetRequestId(v string) *SetDnsGtmMonitorStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *SetDnsGtmMonitorStatusResponseBody) Validate() error {
	return dara.Validate(s)
}
