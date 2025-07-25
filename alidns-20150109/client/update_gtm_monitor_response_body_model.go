// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateGtmMonitorResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateGtmMonitorResponseBody
	GetRequestId() *string
}

type UpdateGtmMonitorResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 6856BCF6-11D6-4D7E-AC53-FD579933522B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateGtmMonitorResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateGtmMonitorResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateGtmMonitorResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateGtmMonitorResponseBody) SetRequestId(v string) *UpdateGtmMonitorResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateGtmMonitorResponseBody) Validate() error {
	return dara.Validate(s)
}
