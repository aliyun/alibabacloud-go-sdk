// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUnbindSmartAccessGatewayResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UnbindSmartAccessGatewayResponseBody
	GetRequestId() *string
}

type UnbindSmartAccessGatewayResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// CE6642D4-21EB-4168-9BF9-F217953F9892
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UnbindSmartAccessGatewayResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UnbindSmartAccessGatewayResponseBody) GoString() string {
	return s.String()
}

func (s *UnbindSmartAccessGatewayResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UnbindSmartAccessGatewayResponseBody) SetRequestId(v string) *UnbindSmartAccessGatewayResponseBody {
	s.RequestId = &v
	return s
}

func (s *UnbindSmartAccessGatewayResponseBody) Validate() error {
	return dara.Validate(s)
}
