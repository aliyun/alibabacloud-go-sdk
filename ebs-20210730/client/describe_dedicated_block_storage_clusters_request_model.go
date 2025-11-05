// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDedicatedBlockStorageClustersRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAzoneId(v string) *DescribeDedicatedBlockStorageClustersRequest
	GetAzoneId() *string
	SetCategory(v string) *DescribeDedicatedBlockStorageClustersRequest
	GetCategory() *string
	SetClientToken(v string) *DescribeDedicatedBlockStorageClustersRequest
	GetClientToken() *string
	SetDedicatedBlockStorageClusterId(v []*string) *DescribeDedicatedBlockStorageClustersRequest
	GetDedicatedBlockStorageClusterId() []*string
	SetMaxResults(v int32) *DescribeDedicatedBlockStorageClustersRequest
	GetMaxResults() *int32
	SetNextToken(v string) *DescribeDedicatedBlockStorageClustersRequest
	GetNextToken() *string
	SetPageNumber(v int32) *DescribeDedicatedBlockStorageClustersRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeDedicatedBlockStorageClustersRequest
	GetPageSize() *int32
	SetRegionId(v string) *DescribeDedicatedBlockStorageClustersRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *DescribeDedicatedBlockStorageClustersRequest
	GetResourceGroupId() *string
	SetStatus(v []*string) *DescribeDedicatedBlockStorageClustersRequest
	GetStatus() []*string
	SetTag(v []*DescribeDedicatedBlockStorageClustersRequestTag) *DescribeDedicatedBlockStorageClustersRequest
	GetTag() []*DescribeDedicatedBlockStorageClustersRequestTag
}

type DescribeDedicatedBlockStorageClustersRequest struct {
	// The zone ID of the dedicated block storage cluster. You can call the [DescribeZones](https://help.aliyun.com/document_detail/25610.html) operation to query the most recent zone list.
	//
	// example:
	//
	// cn-heyuan-b
	AzoneId *string `json:"AzoneId,omitempty" xml:"AzoneId,omitempty"`
	// The category of disks that can be created in the dedicated block storage cluster.
	//
	// Set the value to cloud_essd. Only enhanced SSDs (ESSDs) can be created in dedicated block storage clusters.
	//
	// example:
	//
	// cloud_essd
	Category                       *string   `json:"Category,omitempty" xml:"Category,omitempty"`
	ClientToken                    *string   `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	DedicatedBlockStorageClusterId []*string `json:"DedicatedBlockStorageClusterId,omitempty" xml:"DedicatedBlockStorageClusterId,omitempty" type:"Repeated"`
	MaxResults                     *int32    `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken                      *string   `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Valid values: 1 to 100.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region ID of the dedicated block storage cluster. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/25609.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-heyuan
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group to which the dedicated block storage cluster belongs.
	//
	// example:
	//
	// rg-acfmvs4****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The states of dedicated block storage clusters. Valid values:
	//
	// 	- Preparing
	//
	// 	- Running
	//
	// 	- Expired
	//
	// 	- Offline
	//
	// Multiple states can be specified. Valid values of N: 1, 2, 3, and 4.
	Status []*string `json:"Status,omitempty" xml:"Status,omitempty" type:"Repeated"`
	// The tags. Up to 20 tags are supported.
	Tag []*DescribeDedicatedBlockStorageClustersRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s DescribeDedicatedBlockStorageClustersRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDedicatedBlockStorageClustersRequest) GoString() string {
	return s.String()
}

func (s *DescribeDedicatedBlockStorageClustersRequest) GetAzoneId() *string {
	return s.AzoneId
}

func (s *DescribeDedicatedBlockStorageClustersRequest) GetCategory() *string {
	return s.Category
}

func (s *DescribeDedicatedBlockStorageClustersRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DescribeDedicatedBlockStorageClustersRequest) GetDedicatedBlockStorageClusterId() []*string {
	return s.DedicatedBlockStorageClusterId
}

func (s *DescribeDedicatedBlockStorageClustersRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *DescribeDedicatedBlockStorageClustersRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeDedicatedBlockStorageClustersRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeDedicatedBlockStorageClustersRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeDedicatedBlockStorageClustersRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeDedicatedBlockStorageClustersRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeDedicatedBlockStorageClustersRequest) GetStatus() []*string {
	return s.Status
}

func (s *DescribeDedicatedBlockStorageClustersRequest) GetTag() []*DescribeDedicatedBlockStorageClustersRequestTag {
	return s.Tag
}

func (s *DescribeDedicatedBlockStorageClustersRequest) SetAzoneId(v string) *DescribeDedicatedBlockStorageClustersRequest {
	s.AzoneId = &v
	return s
}

func (s *DescribeDedicatedBlockStorageClustersRequest) SetCategory(v string) *DescribeDedicatedBlockStorageClustersRequest {
	s.Category = &v
	return s
}

func (s *DescribeDedicatedBlockStorageClustersRequest) SetClientToken(v string) *DescribeDedicatedBlockStorageClustersRequest {
	s.ClientToken = &v
	return s
}

func (s *DescribeDedicatedBlockStorageClustersRequest) SetDedicatedBlockStorageClusterId(v []*string) *DescribeDedicatedBlockStorageClustersRequest {
	s.DedicatedBlockStorageClusterId = v
	return s
}

func (s *DescribeDedicatedBlockStorageClustersRequest) SetMaxResults(v int32) *DescribeDedicatedBlockStorageClustersRequest {
	s.MaxResults = &v
	return s
}

func (s *DescribeDedicatedBlockStorageClustersRequest) SetNextToken(v string) *DescribeDedicatedBlockStorageClustersRequest {
	s.NextToken = &v
	return s
}

func (s *DescribeDedicatedBlockStorageClustersRequest) SetPageNumber(v int32) *DescribeDedicatedBlockStorageClustersRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeDedicatedBlockStorageClustersRequest) SetPageSize(v int32) *DescribeDedicatedBlockStorageClustersRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeDedicatedBlockStorageClustersRequest) SetRegionId(v string) *DescribeDedicatedBlockStorageClustersRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeDedicatedBlockStorageClustersRequest) SetResourceGroupId(v string) *DescribeDedicatedBlockStorageClustersRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeDedicatedBlockStorageClustersRequest) SetStatus(v []*string) *DescribeDedicatedBlockStorageClustersRequest {
	s.Status = v
	return s
}

func (s *DescribeDedicatedBlockStorageClustersRequest) SetTag(v []*DescribeDedicatedBlockStorageClustersRequestTag) *DescribeDedicatedBlockStorageClustersRequest {
	s.Tag = v
	return s
}

func (s *DescribeDedicatedBlockStorageClustersRequest) Validate() error {
	if s.Tag != nil {
		for _, item := range s.Tag {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeDedicatedBlockStorageClustersRequestTag struct {
	// The tag key of the dedicated block storage cluster.
	//
	// example:
	//
	// TestKey
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value of the dedicated block storage cluster.
	//
	// example:
	//
	// TestValue
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeDedicatedBlockStorageClustersRequestTag) String() string {
	return dara.Prettify(s)
}

func (s DescribeDedicatedBlockStorageClustersRequestTag) GoString() string {
	return s.String()
}

func (s *DescribeDedicatedBlockStorageClustersRequestTag) GetKey() *string {
	return s.Key
}

func (s *DescribeDedicatedBlockStorageClustersRequestTag) GetValue() *string {
	return s.Value
}

func (s *DescribeDedicatedBlockStorageClustersRequestTag) SetKey(v string) *DescribeDedicatedBlockStorageClustersRequestTag {
	s.Key = &v
	return s
}

func (s *DescribeDedicatedBlockStorageClustersRequestTag) SetValue(v string) *DescribeDedicatedBlockStorageClustersRequestTag {
	s.Value = &v
	return s
}

func (s *DescribeDedicatedBlockStorageClustersRequestTag) Validate() error {
	return dara.Validate(s)
}
