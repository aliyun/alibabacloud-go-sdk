// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeEnterpriseSnapshotPolicyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *DescribeEnterpriseSnapshotPolicyRequest
	GetClientToken() *string
	SetDiskIds(v []*string) *DescribeEnterpriseSnapshotPolicyRequest
	GetDiskIds() []*string
	SetMaxResults(v int32) *DescribeEnterpriseSnapshotPolicyRequest
	GetMaxResults() *int32
	SetNextToken(v string) *DescribeEnterpriseSnapshotPolicyRequest
	GetNextToken() *string
	SetPageNumber(v int32) *DescribeEnterpriseSnapshotPolicyRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeEnterpriseSnapshotPolicyRequest
	GetPageSize() *int32
	SetPolicyIds(v []*string) *DescribeEnterpriseSnapshotPolicyRequest
	GetPolicyIds() []*string
	SetRegionId(v string) *DescribeEnterpriseSnapshotPolicyRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *DescribeEnterpriseSnapshotPolicyRequest
	GetResourceGroupId() *string
	SetTag(v []*DescribeEnterpriseSnapshotPolicyRequestTag) *DescribeEnterpriseSnapshotPolicyRequest
	GetTag() []*DescribeEnterpriseSnapshotPolicyRequestTag
}

type DescribeEnterpriseSnapshotPolicyRequest struct {
	// The client token that is used to ensure the idempotence of the request.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-42665544****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The IDs of disks.
	DiskIds []*string `json:"DiskIds,omitempty" xml:"DiskIds,omitempty" type:"Repeated"`
	// The maximum number of entries per page.
	//
	// example:
	//
	// 100
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results. You do not need to specify this parameter for the first request. You must specify the token that is obtained from the previous query as the value of NextToken. If you specify NextToken, the PageSize and PageNumber request parameters do not take effect, and the TotalCount response parameter is invalid.
	//
	// example:
	//
	// xxx
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The page number.
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
	// The IDs of enterprise-level snapshot policies.
	PolicyIds []*string `json:"PolicyIds,omitempty" xml:"PolicyIds,omitempty" type:"Repeated"`
	// The region ID. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/25609.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// xxx
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The tags of the enterprise-level snapshot policies. Valid values of N: 1 to 20.
	Tag []*DescribeEnterpriseSnapshotPolicyRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s DescribeEnterpriseSnapshotPolicyRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeEnterpriseSnapshotPolicyRequest) GoString() string {
	return s.String()
}

func (s *DescribeEnterpriseSnapshotPolicyRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DescribeEnterpriseSnapshotPolicyRequest) GetDiskIds() []*string {
	return s.DiskIds
}

func (s *DescribeEnterpriseSnapshotPolicyRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *DescribeEnterpriseSnapshotPolicyRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeEnterpriseSnapshotPolicyRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeEnterpriseSnapshotPolicyRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeEnterpriseSnapshotPolicyRequest) GetPolicyIds() []*string {
	return s.PolicyIds
}

func (s *DescribeEnterpriseSnapshotPolicyRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeEnterpriseSnapshotPolicyRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeEnterpriseSnapshotPolicyRequest) GetTag() []*DescribeEnterpriseSnapshotPolicyRequestTag {
	return s.Tag
}

func (s *DescribeEnterpriseSnapshotPolicyRequest) SetClientToken(v string) *DescribeEnterpriseSnapshotPolicyRequest {
	s.ClientToken = &v
	return s
}

func (s *DescribeEnterpriseSnapshotPolicyRequest) SetDiskIds(v []*string) *DescribeEnterpriseSnapshotPolicyRequest {
	s.DiskIds = v
	return s
}

func (s *DescribeEnterpriseSnapshotPolicyRequest) SetMaxResults(v int32) *DescribeEnterpriseSnapshotPolicyRequest {
	s.MaxResults = &v
	return s
}

func (s *DescribeEnterpriseSnapshotPolicyRequest) SetNextToken(v string) *DescribeEnterpriseSnapshotPolicyRequest {
	s.NextToken = &v
	return s
}

func (s *DescribeEnterpriseSnapshotPolicyRequest) SetPageNumber(v int32) *DescribeEnterpriseSnapshotPolicyRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeEnterpriseSnapshotPolicyRequest) SetPageSize(v int32) *DescribeEnterpriseSnapshotPolicyRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeEnterpriseSnapshotPolicyRequest) SetPolicyIds(v []*string) *DescribeEnterpriseSnapshotPolicyRequest {
	s.PolicyIds = v
	return s
}

func (s *DescribeEnterpriseSnapshotPolicyRequest) SetRegionId(v string) *DescribeEnterpriseSnapshotPolicyRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeEnterpriseSnapshotPolicyRequest) SetResourceGroupId(v string) *DescribeEnterpriseSnapshotPolicyRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeEnterpriseSnapshotPolicyRequest) SetTag(v []*DescribeEnterpriseSnapshotPolicyRequestTag) *DescribeEnterpriseSnapshotPolicyRequest {
	s.Tag = v
	return s
}

func (s *DescribeEnterpriseSnapshotPolicyRequest) Validate() error {
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

type DescribeEnterpriseSnapshotPolicyRequestTag struct {
	// The key of tag N of the enterprise-level snapshot policy.
	//
	// This parameter is required.
	//
	// example:
	//
	// tag-key
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of tag N of the enterprise-level snapshot policy.
	//
	// This parameter is required.
	//
	// example:
	//
	// tag-value
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeEnterpriseSnapshotPolicyRequestTag) String() string {
	return dara.Prettify(s)
}

func (s DescribeEnterpriseSnapshotPolicyRequestTag) GoString() string {
	return s.String()
}

func (s *DescribeEnterpriseSnapshotPolicyRequestTag) GetKey() *string {
	return s.Key
}

func (s *DescribeEnterpriseSnapshotPolicyRequestTag) GetValue() *string {
	return s.Value
}

func (s *DescribeEnterpriseSnapshotPolicyRequestTag) SetKey(v string) *DescribeEnterpriseSnapshotPolicyRequestTag {
	s.Key = &v
	return s
}

func (s *DescribeEnterpriseSnapshotPolicyRequestTag) SetValue(v string) *DescribeEnterpriseSnapshotPolicyRequestTag {
	s.Value = &v
	return s
}

func (s *DescribeEnterpriseSnapshotPolicyRequestTag) Validate() error {
	return dara.Validate(s)
}
