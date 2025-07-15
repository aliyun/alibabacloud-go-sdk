// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyNatGatewaySpecResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyNatGatewaySpecResponseBody
	GetRequestId() *string
}

type ModifyNatGatewaySpecResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// DBD4E4A2-786E-4BD2-8EB6-107FFC2B5B7D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyNatGatewaySpecResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyNatGatewaySpecResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyNatGatewaySpecResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyNatGatewaySpecResponseBody) SetRequestId(v string) *ModifyNatGatewaySpecResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyNatGatewaySpecResponseBody) Validate() error {
	return dara.Validate(s)
}
