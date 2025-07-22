// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSubnetRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCidrBlock(v string) *CreateSubnetRequest
	GetCidrBlock() *string
	SetName(v string) *CreateSubnetRequest
	GetName() *string
	SetOfficeSiteId(v string) *CreateSubnetRequest
	GetOfficeSiteId() *string
	SetRegionId(v string) *CreateSubnetRequest
	GetRegionId() *string
	SetZoneId(v string) *CreateSubnetRequest
	GetZoneId() *string
}

type CreateSubnetRequest struct {
	// This parameter is required.
	CidrBlock *string `json:"CidrBlock,omitempty" xml:"CidrBlock,omitempty"`
	Name      *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// This parameter is required.
	OfficeSiteId *string `json:"OfficeSiteId,omitempty" xml:"OfficeSiteId,omitempty"`
	// This parameter is required.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ZoneId   *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s CreateSubnetRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateSubnetRequest) GoString() string {
	return s.String()
}

func (s *CreateSubnetRequest) GetCidrBlock() *string {
	return s.CidrBlock
}

func (s *CreateSubnetRequest) GetName() *string {
	return s.Name
}

func (s *CreateSubnetRequest) GetOfficeSiteId() *string {
	return s.OfficeSiteId
}

func (s *CreateSubnetRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateSubnetRequest) GetZoneId() *string {
	return s.ZoneId
}

func (s *CreateSubnetRequest) SetCidrBlock(v string) *CreateSubnetRequest {
	s.CidrBlock = &v
	return s
}

func (s *CreateSubnetRequest) SetName(v string) *CreateSubnetRequest {
	s.Name = &v
	return s
}

func (s *CreateSubnetRequest) SetOfficeSiteId(v string) *CreateSubnetRequest {
	s.OfficeSiteId = &v
	return s
}

func (s *CreateSubnetRequest) SetRegionId(v string) *CreateSubnetRequest {
	s.RegionId = &v
	return s
}

func (s *CreateSubnetRequest) SetZoneId(v string) *CreateSubnetRequest {
	s.ZoneId = &v
	return s
}

func (s *CreateSubnetRequest) Validate() error {
	return dara.Validate(s)
}
