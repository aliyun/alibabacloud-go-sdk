// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeComponentPropetiesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCommodityCode(v string) *DescribeComponentPropetiesRequest
	GetCommodityCode() *string
	SetComponentName(v string) *DescribeComponentPropetiesRequest
	GetComponentName() *string
	SetRegionId(v string) *DescribeComponentPropetiesRequest
	GetRegionId() *string
	SetStorageType(v string) *DescribeComponentPropetiesRequest
	GetStorageType() *string
}

type DescribeComponentPropetiesRequest struct {
	// The commodity code.
	//
	// This parameter is required.
	//
	// example:
	//
	// learn_EasDedicatedPrepay_public_cn
	CommodityCode *string `json:"CommodityCode,omitempty" xml:"CommodityCode,omitempty"`
	// The resource group ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// rg-xxxxx
	ComponentName *string `json:"ComponentName,omitempty" xml:"ComponentName,omitempty"`
	// The region in which the instance resides. > You can call the [describeregions](https://help.aliyun.com/document_detail/196841.html) operation to query the regions supported by PolarDB-X, including region IDs.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The storage type.
	//
	// example:
	//
	// cloud_auto
	StorageType *string `json:"StorageType,omitempty" xml:"StorageType,omitempty"`
}

func (s DescribeComponentPropetiesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeComponentPropetiesRequest) GoString() string {
	return s.String()
}

func (s *DescribeComponentPropetiesRequest) GetCommodityCode() *string {
	return s.CommodityCode
}

func (s *DescribeComponentPropetiesRequest) GetComponentName() *string {
	return s.ComponentName
}

func (s *DescribeComponentPropetiesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeComponentPropetiesRequest) GetStorageType() *string {
	return s.StorageType
}

func (s *DescribeComponentPropetiesRequest) SetCommodityCode(v string) *DescribeComponentPropetiesRequest {
	s.CommodityCode = &v
	return s
}

func (s *DescribeComponentPropetiesRequest) SetComponentName(v string) *DescribeComponentPropetiesRequest {
	s.ComponentName = &v
	return s
}

func (s *DescribeComponentPropetiesRequest) SetRegionId(v string) *DescribeComponentPropetiesRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeComponentPropetiesRequest) SetStorageType(v string) *DescribeComponentPropetiesRequest {
	s.StorageType = &v
	return s
}

func (s *DescribeComponentPropetiesRequest) Validate() error {
	return dara.Validate(s)
}
