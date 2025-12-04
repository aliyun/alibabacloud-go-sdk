// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteVpdRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRegionId(v string) *DeleteVpdRequest
	GetRegionId() *string
	SetVpdId(v string) *DeleteVpdRequest
	GetVpdId() *string
}

type DeleteVpdRequest struct {
	// The ID of the region in which the instance resides.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the Lingjun CIDR block.
	//
	// This parameter is required.
	//
	// example:
	//
	// vpd-zr0farea
	VpdId *string `json:"VpdId,omitempty" xml:"VpdId,omitempty"`
}

func (s DeleteVpdRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteVpdRequest) GoString() string {
	return s.String()
}

func (s *DeleteVpdRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteVpdRequest) GetVpdId() *string {
	return s.VpdId
}

func (s *DeleteVpdRequest) SetRegionId(v string) *DeleteVpdRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteVpdRequest) SetVpdId(v string) *DeleteVpdRequest {
	s.VpdId = &v
	return s
}

func (s *DeleteVpdRequest) Validate() error {
	return dara.Validate(s)
}
