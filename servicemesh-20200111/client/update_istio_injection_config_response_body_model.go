// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateIstioInjectionConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateIstioInjectionConfigResponseBody
	GetRequestId() *string
}

type UpdateIstioInjectionConfigResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 98B268E6-9381-5A98-8012-6E7E82******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateIstioInjectionConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateIstioInjectionConfigResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateIstioInjectionConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateIstioInjectionConfigResponseBody) SetRequestId(v string) *UpdateIstioInjectionConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateIstioInjectionConfigResponseBody) Validate() error {
	return dara.Validate(s)
}
