// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetConsumerAccessOverviewRequest interface {
	dara.Model
	String() string
	GoString() string
	SetGatewayId(v string) *GetConsumerAccessOverviewRequest
	GetGatewayId() *string
}

type GetConsumerAccessOverviewRequest struct {
	// The target AI gateway instance ID. Quotas and resources are isolated by instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// gw-123456
	GatewayId *string `json:"gatewayId,omitempty" xml:"gatewayId,omitempty"`
}

func (s GetConsumerAccessOverviewRequest) String() string {
	return dara.Prettify(s)
}

func (s GetConsumerAccessOverviewRequest) GoString() string {
	return s.String()
}

func (s *GetConsumerAccessOverviewRequest) GetGatewayId() *string {
	return s.GatewayId
}

func (s *GetConsumerAccessOverviewRequest) SetGatewayId(v string) *GetConsumerAccessOverviewRequest {
	s.GatewayId = &v
	return s
}

func (s *GetConsumerAccessOverviewRequest) Validate() error {
	return dara.Validate(s)
}
