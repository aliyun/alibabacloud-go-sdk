// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifySmartAccessGatewayResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifySmartAccessGatewayResponseBody
	GetRequestId() *string
}

type ModifySmartAccessGatewayResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// CE6642D4-21EB-4168-9BF9-F217953F9892
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifySmartAccessGatewayResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifySmartAccessGatewayResponseBody) GoString() string {
	return s.String()
}

func (s *ModifySmartAccessGatewayResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifySmartAccessGatewayResponseBody) SetRequestId(v string) *ModifySmartAccessGatewayResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifySmartAccessGatewayResponseBody) Validate() error {
	return dara.Validate(s)
}
