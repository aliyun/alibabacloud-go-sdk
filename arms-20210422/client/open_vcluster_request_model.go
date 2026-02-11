// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOpenVClusterRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterType(v string) *OpenVClusterRequest
	GetClusterType() *string
	SetLength(v int32) *OpenVClusterRequest
	GetLength() *int32
	SetProduct(v string) *OpenVClusterRequest
	GetProduct() *string
	SetRecreateSwitch(v bool) *OpenVClusterRequest
	GetRecreateSwitch() *bool
	SetRegionId(v string) *OpenVClusterRequest
	GetRegionId() *string
}

type OpenVClusterRequest struct {
	// This parameter is required.
	ClusterType    *string `json:"ClusterType,omitempty" xml:"ClusterType,omitempty"`
	Length         *int32  `json:"Length,omitempty" xml:"Length,omitempty"`
	Product        *string `json:"Product,omitempty" xml:"Product,omitempty"`
	RecreateSwitch *bool   `json:"RecreateSwitch,omitempty" xml:"RecreateSwitch,omitempty"`
	// This parameter is required.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s OpenVClusterRequest) String() string {
	return dara.Prettify(s)
}

func (s OpenVClusterRequest) GoString() string {
	return s.String()
}

func (s *OpenVClusterRequest) GetClusterType() *string {
	return s.ClusterType
}

func (s *OpenVClusterRequest) GetLength() *int32 {
	return s.Length
}

func (s *OpenVClusterRequest) GetProduct() *string {
	return s.Product
}

func (s *OpenVClusterRequest) GetRecreateSwitch() *bool {
	return s.RecreateSwitch
}

func (s *OpenVClusterRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *OpenVClusterRequest) SetClusterType(v string) *OpenVClusterRequest {
	s.ClusterType = &v
	return s
}

func (s *OpenVClusterRequest) SetLength(v int32) *OpenVClusterRequest {
	s.Length = &v
	return s
}

func (s *OpenVClusterRequest) SetProduct(v string) *OpenVClusterRequest {
	s.Product = &v
	return s
}

func (s *OpenVClusterRequest) SetRecreateSwitch(v bool) *OpenVClusterRequest {
	s.RecreateSwitch = &v
	return s
}

func (s *OpenVClusterRequest) SetRegionId(v string) *OpenVClusterRequest {
	s.RegionId = &v
	return s
}

func (s *OpenVClusterRequest) Validate() error {
	return dara.Validate(s)
}
