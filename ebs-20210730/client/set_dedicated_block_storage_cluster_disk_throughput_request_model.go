// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetDedicatedBlockStorageClusterDiskThroughputRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBps(v int32) *SetDedicatedBlockStorageClusterDiskThroughputRequest
	GetBps() *int32
	SetClientToken(v string) *SetDedicatedBlockStorageClusterDiskThroughputRequest
	GetClientToken() *string
	SetDiskId(v string) *SetDedicatedBlockStorageClusterDiskThroughputRequest
	GetDiskId() *string
	SetRegionId(v string) *SetDedicatedBlockStorageClusterDiskThroughputRequest
	GetRegionId() *string
}

type SetDedicatedBlockStorageClusterDiskThroughputRequest struct {
	// Target throughput.
	//
	// This parameter is required.
	//
	// example:
	//
	// 100
	Bps *int32 `json:"Bps,omitempty" xml:"Bps,omitempty"`
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the value, but you must make sure that the value is unique among different requests. The ClientToken value can contain only ASCII characters and cannot exceed 64 characters in length. For more information, see [How to ensure idempotence](https://help.aliyun.com/document_detail/25693.html).
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-42665544****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The ID of the disk.
	//
	// This parameter is required.
	//
	// example:
	//
	// d-bp67acfmxazb4p****
	DiskId *string `json:"DiskId,omitempty" xml:"DiskId,omitempty"`
	// The region ID of disk.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s SetDedicatedBlockStorageClusterDiskThroughputRequest) String() string {
	return dara.Prettify(s)
}

func (s SetDedicatedBlockStorageClusterDiskThroughputRequest) GoString() string {
	return s.String()
}

func (s *SetDedicatedBlockStorageClusterDiskThroughputRequest) GetBps() *int32 {
	return s.Bps
}

func (s *SetDedicatedBlockStorageClusterDiskThroughputRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *SetDedicatedBlockStorageClusterDiskThroughputRequest) GetDiskId() *string {
	return s.DiskId
}

func (s *SetDedicatedBlockStorageClusterDiskThroughputRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *SetDedicatedBlockStorageClusterDiskThroughputRequest) SetBps(v int32) *SetDedicatedBlockStorageClusterDiskThroughputRequest {
	s.Bps = &v
	return s
}

func (s *SetDedicatedBlockStorageClusterDiskThroughputRequest) SetClientToken(v string) *SetDedicatedBlockStorageClusterDiskThroughputRequest {
	s.ClientToken = &v
	return s
}

func (s *SetDedicatedBlockStorageClusterDiskThroughputRequest) SetDiskId(v string) *SetDedicatedBlockStorageClusterDiskThroughputRequest {
	s.DiskId = &v
	return s
}

func (s *SetDedicatedBlockStorageClusterDiskThroughputRequest) SetRegionId(v string) *SetDedicatedBlockStorageClusterDiskThroughputRequest {
	s.RegionId = &v
	return s
}

func (s *SetDedicatedBlockStorageClusterDiskThroughputRequest) Validate() error {
	return dara.Validate(s)
}
