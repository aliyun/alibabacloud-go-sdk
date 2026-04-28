// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteGatewayRequest interface {
	dara.Model
	String() string
	GoString() string
	SetGwClusterId(v string) *DeleteGatewayRequest
	GetGwClusterId() *string
	SetRegionId(v string) *DeleteGatewayRequest
	GetRegionId() *string
}

type DeleteGatewayRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// pg-2ze2079ueg20****
	GwClusterId *string `json:"GwClusterId,omitempty" xml:"GwClusterId,omitempty"`
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DeleteGatewayRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteGatewayRequest) GoString() string {
	return s.String()
}

func (s *DeleteGatewayRequest) GetGwClusterId() *string {
	return s.GwClusterId
}

func (s *DeleteGatewayRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteGatewayRequest) SetGwClusterId(v string) *DeleteGatewayRequest {
	s.GwClusterId = &v
	return s
}

func (s *DeleteGatewayRequest) SetRegionId(v string) *DeleteGatewayRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteGatewayRequest) Validate() error {
	return dara.Validate(s)
}
