// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDnsGtmMonitorResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateDnsGtmMonitorResponseBody
	GetRequestId() *string
}

type UpdateDnsGtmMonitorResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 6856BCF6-11D6-4D7E-AC53-FD579933522B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateDnsGtmMonitorResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateDnsGtmMonitorResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateDnsGtmMonitorResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateDnsGtmMonitorResponseBody) SetRequestId(v string) *UpdateDnsGtmMonitorResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateDnsGtmMonitorResponseBody) Validate() error {
	return dara.Validate(s)
}
