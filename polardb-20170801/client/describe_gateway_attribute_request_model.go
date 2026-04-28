// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeGatewayAttributeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetGwClusterId(v string) *DescribeGatewayAttributeRequest
	GetGwClusterId() *string
	SetRegionId(v string) *DescribeGatewayAttributeRequest
	GetRegionId() *string
}

type DescribeGatewayAttributeRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// pg-xxxxxxx
	GwClusterId *string `json:"GwClusterId,omitempty" xml:"GwClusterId,omitempty"`
	RegionId    *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeGatewayAttributeRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeGatewayAttributeRequest) GoString() string {
	return s.String()
}

func (s *DescribeGatewayAttributeRequest) GetGwClusterId() *string {
	return s.GwClusterId
}

func (s *DescribeGatewayAttributeRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeGatewayAttributeRequest) SetGwClusterId(v string) *DescribeGatewayAttributeRequest {
	s.GwClusterId = &v
	return s
}

func (s *DescribeGatewayAttributeRequest) SetRegionId(v string) *DescribeGatewayAttributeRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeGatewayAttributeRequest) Validate() error {
	return dara.Validate(s)
}
