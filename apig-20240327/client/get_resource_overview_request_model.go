// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetResourceOverviewRequest interface {
	dara.Model
	String() string
	GoString() string
	SetGatewayType(v string) *GetResourceOverviewRequest
	GetGatewayType() *string
}

type GetResourceOverviewRequest struct {
	// The gateway type.
	//
	// example:
	//
	// AI
	GatewayType *string `json:"gatewayType,omitempty" xml:"gatewayType,omitempty"`
}

func (s GetResourceOverviewRequest) String() string {
	return dara.Prettify(s)
}

func (s GetResourceOverviewRequest) GoString() string {
	return s.String()
}

func (s *GetResourceOverviewRequest) GetGatewayType() *string {
	return s.GatewayType
}

func (s *GetResourceOverviewRequest) SetGatewayType(v string) *GetResourceOverviewRequest {
	s.GatewayType = &v
	return s
}

func (s *GetResourceOverviewRequest) Validate() error {
	return dara.Validate(s)
}
