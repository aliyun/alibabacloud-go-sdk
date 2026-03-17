// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateSmartAccessGatewayVersionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateSmartAccessGatewayVersionResponseBody
	GetRequestId() *string
}

type UpdateSmartAccessGatewayVersionResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// CE6642D4-21EB-4168-9BF9-F217953F9892
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateSmartAccessGatewayVersionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateSmartAccessGatewayVersionResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateSmartAccessGatewayVersionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateSmartAccessGatewayVersionResponseBody) SetRequestId(v string) *UpdateSmartAccessGatewayVersionResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateSmartAccessGatewayVersionResponseBody) Validate() error {
	return dara.Validate(s)
}
