// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteGadInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetGadInstanceName(v string) *DeleteGadInstanceRequest
	GetGadInstanceName() *string
	SetRegionId(v string) *DeleteGadInstanceRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *DeleteGadInstanceRequest
	GetResourceGroupId() *string
}

type DeleteGadInstanceRequest struct {
	// This parameter is required.
	GadInstanceName *string `json:"GadInstanceName,omitempty" xml:"GadInstanceName,omitempty"`
	RegionId        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s DeleteGadInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteGadInstanceRequest) GoString() string {
	return s.String()
}

func (s *DeleteGadInstanceRequest) GetGadInstanceName() *string {
	return s.GadInstanceName
}

func (s *DeleteGadInstanceRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteGadInstanceRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DeleteGadInstanceRequest) SetGadInstanceName(v string) *DeleteGadInstanceRequest {
	s.GadInstanceName = &v
	return s
}

func (s *DeleteGadInstanceRequest) SetRegionId(v string) *DeleteGadInstanceRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteGadInstanceRequest) SetResourceGroupId(v string) *DeleteGadInstanceRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DeleteGadInstanceRequest) Validate() error {
	return dara.Validate(s)
}
