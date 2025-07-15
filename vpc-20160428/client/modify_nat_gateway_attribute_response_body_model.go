// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyNatGatewayAttributeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyNatGatewayAttributeResponseBody
	GetRequestId() *string
}

type ModifyNatGatewayAttributeResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// AB5F62CF-2B60-4458-A756-42C9DFE108D1
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyNatGatewayAttributeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyNatGatewayAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyNatGatewayAttributeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyNatGatewayAttributeResponseBody) SetRequestId(v string) *ModifyNatGatewayAttributeResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyNatGatewayAttributeResponseBody) Validate() error {
	return dara.Validate(s)
}
