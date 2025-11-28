// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListBaaSFabricOrganizationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBaaSFabricChannelId(v string) *ListBaaSFabricOrganizationRequest
	GetBaaSFabricChannelId() *string
	SetRegionId(v string) *ListBaaSFabricOrganizationRequest
	GetRegionId() *string
}

type ListBaaSFabricOrganizationRequest struct {
	// This parameter is required.
	BaaSFabricChannelId *string `json:"BaaSFabricChannelId,omitempty" xml:"BaaSFabricChannelId,omitempty"`
	RegionId            *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ListBaaSFabricOrganizationRequest) String() string {
	return dara.Prettify(s)
}

func (s ListBaaSFabricOrganizationRequest) GoString() string {
	return s.String()
}

func (s *ListBaaSFabricOrganizationRequest) GetBaaSFabricChannelId() *string {
	return s.BaaSFabricChannelId
}

func (s *ListBaaSFabricOrganizationRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListBaaSFabricOrganizationRequest) SetBaaSFabricChannelId(v string) *ListBaaSFabricOrganizationRequest {
	s.BaaSFabricChannelId = &v
	return s
}

func (s *ListBaaSFabricOrganizationRequest) SetRegionId(v string) *ListBaaSFabricOrganizationRequest {
	s.RegionId = &v
	return s
}

func (s *ListBaaSFabricOrganizationRequest) Validate() error {
	return dara.Validate(s)
}
