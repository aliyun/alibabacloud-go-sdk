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
	// The type of the cluster. For cloud services, set this parameter to `cloud-product-prometheus`.
	//
	// This parameter is required.
	//
	// example:
	//
	// cloud-product-prometheus
	ClusterType *string `json:"ClusterType,omitempty" xml:"ClusterType,omitempty"`
	// The length of the cluster ID. Default value: 10.
	//
	// example:
	//
	// 10
	Length *int32 `json:"Length,omitempty" xml:"Length,omitempty"`
	// The name of the cloud service. This parameter must be specified when ClusterType is set to `cloud-product-prometheus`. Valid values: influxdb, mongodb, and DLA. You cannot specify multiple service names.
	//
	// example:
	//
	// influxdb
	Product *string `json:"Product,omitempty" xml:"Product,omitempty"`
	// Specifies whether to create or query a virtual cluster. This parameter provides backward compatibility.
	//
	// example:
	//
	// false
	RecreateSwitch *bool `json:"RecreateSwitch,omitempty" xml:"RecreateSwitch,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
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
