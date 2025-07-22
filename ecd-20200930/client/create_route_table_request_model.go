// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateRouteTableRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *CreateRouteTableRequest
	GetClientToken() *string
	SetDescription(v string) *CreateRouteTableRequest
	GetDescription() *string
	SetRegionId(v string) *CreateRouteTableRequest
	GetRegionId() *string
	SetRouteTableName(v string) *CreateRouteTableRequest
	GetRouteTableName() *string
	SetVpcId(v string) *CreateRouteTableRequest
	GetVpcId() *string
}

type CreateRouteTableRequest struct {
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// This parameter is required.
	RegionId       *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	RouteTableName *string `json:"RouteTableName,omitempty" xml:"RouteTableName,omitempty"`
	// This parameter is required.
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s CreateRouteTableRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateRouteTableRequest) GoString() string {
	return s.String()
}

func (s *CreateRouteTableRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateRouteTableRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateRouteTableRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateRouteTableRequest) GetRouteTableName() *string {
	return s.RouteTableName
}

func (s *CreateRouteTableRequest) GetVpcId() *string {
	return s.VpcId
}

func (s *CreateRouteTableRequest) SetClientToken(v string) *CreateRouteTableRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateRouteTableRequest) SetDescription(v string) *CreateRouteTableRequest {
	s.Description = &v
	return s
}

func (s *CreateRouteTableRequest) SetRegionId(v string) *CreateRouteTableRequest {
	s.RegionId = &v
	return s
}

func (s *CreateRouteTableRequest) SetRouteTableName(v string) *CreateRouteTableRequest {
	s.RouteTableName = &v
	return s
}

func (s *CreateRouteTableRequest) SetVpcId(v string) *CreateRouteTableRequest {
	s.VpcId = &v
	return s
}

func (s *CreateRouteTableRequest) Validate() error {
	return dara.Validate(s)
}
