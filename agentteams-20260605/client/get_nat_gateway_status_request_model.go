// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetNatGatewayStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *GetNatGatewayStatusRequest
	GetInstanceId() *string
}

type GetNatGatewayStatusRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// inst-1
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s GetNatGatewayStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s GetNatGatewayStatusRequest) GoString() string {
	return s.String()
}

func (s *GetNatGatewayStatusRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetNatGatewayStatusRequest) SetInstanceId(v string) *GetNatGatewayStatusRequest {
	s.InstanceId = &v
	return s
}

func (s *GetNatGatewayStatusRequest) Validate() error {
	return dara.Validate(s)
}
