// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeStorageVolumeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEnsRegionId(v string) *DescribeStorageVolumeRequest
	GetEnsRegionId() *string
	SetGatewayId(v string) *DescribeStorageVolumeRequest
	GetGatewayId() *string
	SetIsEnable(v int32) *DescribeStorageVolumeRequest
	GetIsEnable() *int32
	SetPageNumber(v int32) *DescribeStorageVolumeRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeStorageVolumeRequest
	GetPageSize() *int32
	SetStorageId(v string) *DescribeStorageVolumeRequest
	GetStorageId() *string
	SetVolumeId(v string) *DescribeStorageVolumeRequest
	GetVolumeId() *string
}

type DescribeStorageVolumeRequest struct {
	// The ID of the node.
	//
	// example:
	//
	// cn-shenzhen-3
	EnsRegionId *string `json:"EnsRegionId,omitempty" xml:"EnsRegionId,omitempty"`
	// The ID of the gateway.
	//
	// example:
	//
	// sgw-****
	GatewayId *string `json:"GatewayId,omitempty" xml:"GatewayId,omitempty"`
	// Specifies whether to enable the volume. Valid values:
	//
	// 	- **1*	- (default): enables the volume.
	//
	// 	- **0**: disables the volume.
	//
	// example:
	//
	// 1
	IsEnable *int32 `json:"IsEnable,omitempty" xml:"IsEnable,omitempty"`
	// The page number. Default value: **1**.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the storage medium.
	//
	// example:
	//
	// d-***
	StorageId *string `json:"StorageId,omitempty" xml:"StorageId,omitempty"`
	// The ID of the volume.
	//
	// example:
	//
	// sv-***
	VolumeId *string `json:"VolumeId,omitempty" xml:"VolumeId,omitempty"`
}

func (s DescribeStorageVolumeRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeStorageVolumeRequest) GoString() string {
	return s.String()
}

func (s *DescribeStorageVolumeRequest) GetEnsRegionId() *string {
	return s.EnsRegionId
}

func (s *DescribeStorageVolumeRequest) GetGatewayId() *string {
	return s.GatewayId
}

func (s *DescribeStorageVolumeRequest) GetIsEnable() *int32 {
	return s.IsEnable
}

func (s *DescribeStorageVolumeRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeStorageVolumeRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeStorageVolumeRequest) GetStorageId() *string {
	return s.StorageId
}

func (s *DescribeStorageVolumeRequest) GetVolumeId() *string {
	return s.VolumeId
}

func (s *DescribeStorageVolumeRequest) SetEnsRegionId(v string) *DescribeStorageVolumeRequest {
	s.EnsRegionId = &v
	return s
}

func (s *DescribeStorageVolumeRequest) SetGatewayId(v string) *DescribeStorageVolumeRequest {
	s.GatewayId = &v
	return s
}

func (s *DescribeStorageVolumeRequest) SetIsEnable(v int32) *DescribeStorageVolumeRequest {
	s.IsEnable = &v
	return s
}

func (s *DescribeStorageVolumeRequest) SetPageNumber(v int32) *DescribeStorageVolumeRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeStorageVolumeRequest) SetPageSize(v int32) *DescribeStorageVolumeRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeStorageVolumeRequest) SetStorageId(v string) *DescribeStorageVolumeRequest {
	s.StorageId = &v
	return s
}

func (s *DescribeStorageVolumeRequest) SetVolumeId(v string) *DescribeStorageVolumeRequest {
	s.VolumeId = &v
	return s
}

func (s *DescribeStorageVolumeRequest) Validate() error {
	return dara.Validate(s)
}
