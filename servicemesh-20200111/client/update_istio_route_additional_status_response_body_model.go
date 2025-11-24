// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateIstioRouteAdditionalStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateIstioRouteAdditionalStatusResponseBody
	GetRequestId() *string
}

type UpdateIstioRouteAdditionalStatusResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// BD65C0AD-D3C6-48D3-8D93-38D2015C****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateIstioRouteAdditionalStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateIstioRouteAdditionalStatusResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateIstioRouteAdditionalStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateIstioRouteAdditionalStatusResponseBody) SetRequestId(v string) *UpdateIstioRouteAdditionalStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateIstioRouteAdditionalStatusResponseBody) Validate() error {
	return dara.Validate(s)
}
