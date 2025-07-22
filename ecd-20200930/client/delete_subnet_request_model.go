// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteSubnetRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRegionId(v string) *DeleteSubnetRequest
	GetRegionId() *string
	SetSubnetId(v string) *DeleteSubnetRequest
	GetSubnetId() *string
}

type DeleteSubnetRequest struct {
	// This parameter is required.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// This parameter is required.
	SubnetId *string `json:"SubnetId,omitempty" xml:"SubnetId,omitempty"`
}

func (s DeleteSubnetRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteSubnetRequest) GoString() string {
	return s.String()
}

func (s *DeleteSubnetRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteSubnetRequest) GetSubnetId() *string {
	return s.SubnetId
}

func (s *DeleteSubnetRequest) SetRegionId(v string) *DeleteSubnetRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteSubnetRequest) SetSubnetId(v string) *DeleteSubnetRequest {
	s.SubnetId = &v
	return s
}

func (s *DeleteSubnetRequest) Validate() error {
	return dara.Validate(s)
}
