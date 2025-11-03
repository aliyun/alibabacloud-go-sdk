// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTrafficMirrorFiltersRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListTrafficMirrorFiltersRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListTrafficMirrorFiltersRequest
	GetNextToken() *string
	SetOwnerAccount(v string) *ListTrafficMirrorFiltersRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ListTrafficMirrorFiltersRequest
	GetOwnerId() *int64
	SetRegionId(v string) *ListTrafficMirrorFiltersRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *ListTrafficMirrorFiltersRequest
	GetResourceGroupId() *string
	SetResourceOwnerAccount(v string) *ListTrafficMirrorFiltersRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ListTrafficMirrorFiltersRequest
	GetResourceOwnerId() *int64
	SetTags(v []*ListTrafficMirrorFiltersRequestTags) *ListTrafficMirrorFiltersRequest
	GetTags() []*ListTrafficMirrorFiltersRequestTags
	SetTrafficMirrorFilterIds(v []*string) *ListTrafficMirrorFiltersRequest
	GetTrafficMirrorFilterIds() []*string
	SetTrafficMirrorFilterName(v string) *ListTrafficMirrorFiltersRequest
	GetTrafficMirrorFilterName() *string
}

type ListTrafficMirrorFiltersRequest struct {
	// The maximum number of entries to return.
	//
	// Valid values: **1*	- to **100**. Default value: **10**.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results. Valid values:
	//
	// 	- You do not need to specify this parameter for the first request.
	//
	// 	- You must specify the token that is obtained from the previous query as the value of NextToken.
	//
	// example:
	//
	// FFmyTO70tTpLG6I3FmYAXGKPd****
	NextToken    *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the region to which the mirrored traffic belongs.
	//
	// You can call the [DescribeRegions](https://help.aliyun.com/document_detail/36063.html) operation to query the most recent region list. For more information about regions that support traffic mirror, see [Overview of traffic mirror](https://help.aliyun.com/document_detail/207513.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hongkong
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group to which the mirrored traffic belongs.
	//
	// example:
	//
	// rg-bp67acfmxazb4ph****
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The tag list.
	Tags []*ListTrafficMirrorFiltersRequestTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// The ID of the traffic mirror filter. The maximum value of **N*	- is **100**, which specifies that you can query up to 100 filters at a time.
	//
	// example:
	//
	// tmf-j6cmls82xnc86vtpe****
	TrafficMirrorFilterIds []*string `json:"TrafficMirrorFilterIds,omitempty" xml:"TrafficMirrorFilterIds,omitempty" type:"Repeated"`
	// The name of the filter.
	//
	// example:
	//
	// abc
	TrafficMirrorFilterName *string `json:"TrafficMirrorFilterName,omitempty" xml:"TrafficMirrorFilterName,omitempty"`
}

func (s ListTrafficMirrorFiltersRequest) String() string {
	return dara.Prettify(s)
}

func (s ListTrafficMirrorFiltersRequest) GoString() string {
	return s.String()
}

func (s *ListTrafficMirrorFiltersRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListTrafficMirrorFiltersRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListTrafficMirrorFiltersRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ListTrafficMirrorFiltersRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ListTrafficMirrorFiltersRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListTrafficMirrorFiltersRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ListTrafficMirrorFiltersRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ListTrafficMirrorFiltersRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ListTrafficMirrorFiltersRequest) GetTags() []*ListTrafficMirrorFiltersRequestTags {
	return s.Tags
}

func (s *ListTrafficMirrorFiltersRequest) GetTrafficMirrorFilterIds() []*string {
	return s.TrafficMirrorFilterIds
}

func (s *ListTrafficMirrorFiltersRequest) GetTrafficMirrorFilterName() *string {
	return s.TrafficMirrorFilterName
}

func (s *ListTrafficMirrorFiltersRequest) SetMaxResults(v int32) *ListTrafficMirrorFiltersRequest {
	s.MaxResults = &v
	return s
}

func (s *ListTrafficMirrorFiltersRequest) SetNextToken(v string) *ListTrafficMirrorFiltersRequest {
	s.NextToken = &v
	return s
}

func (s *ListTrafficMirrorFiltersRequest) SetOwnerAccount(v string) *ListTrafficMirrorFiltersRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ListTrafficMirrorFiltersRequest) SetOwnerId(v int64) *ListTrafficMirrorFiltersRequest {
	s.OwnerId = &v
	return s
}

func (s *ListTrafficMirrorFiltersRequest) SetRegionId(v string) *ListTrafficMirrorFiltersRequest {
	s.RegionId = &v
	return s
}

func (s *ListTrafficMirrorFiltersRequest) SetResourceGroupId(v string) *ListTrafficMirrorFiltersRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ListTrafficMirrorFiltersRequest) SetResourceOwnerAccount(v string) *ListTrafficMirrorFiltersRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ListTrafficMirrorFiltersRequest) SetResourceOwnerId(v int64) *ListTrafficMirrorFiltersRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ListTrafficMirrorFiltersRequest) SetTags(v []*ListTrafficMirrorFiltersRequestTags) *ListTrafficMirrorFiltersRequest {
	s.Tags = v
	return s
}

func (s *ListTrafficMirrorFiltersRequest) SetTrafficMirrorFilterIds(v []*string) *ListTrafficMirrorFiltersRequest {
	s.TrafficMirrorFilterIds = v
	return s
}

func (s *ListTrafficMirrorFiltersRequest) SetTrafficMirrorFilterName(v string) *ListTrafficMirrorFiltersRequest {
	s.TrafficMirrorFilterName = &v
	return s
}

func (s *ListTrafficMirrorFiltersRequest) Validate() error {
	if s.Tags != nil {
		for _, item := range s.Tags {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListTrafficMirrorFiltersRequestTags struct {
	// The tag key. You can specify at most 20 tag keys. The tag key cannot be an empty string.
	//
	// The tag key can be up to 128 characters in length. It cannot start with `aliyun` or `acs:`, and cannot contain `http://` or `https://`.
	//
	// example:
	//
	// FinanceDept
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value. You can specify at most 20 tag values. The tag value can be an empty string.
	//
	// The tag value can be up to 128 characters in length. It cannot start with `aliyun` or `acs:`, and cannot contain `http://` or `https://`.
	//
	// example:
	//
	// FinanceJoshua
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListTrafficMirrorFiltersRequestTags) String() string {
	return dara.Prettify(s)
}

func (s ListTrafficMirrorFiltersRequestTags) GoString() string {
	return s.String()
}

func (s *ListTrafficMirrorFiltersRequestTags) GetKey() *string {
	return s.Key
}

func (s *ListTrafficMirrorFiltersRequestTags) GetValue() *string {
	return s.Value
}

func (s *ListTrafficMirrorFiltersRequestTags) SetKey(v string) *ListTrafficMirrorFiltersRequestTags {
	s.Key = &v
	return s
}

func (s *ListTrafficMirrorFiltersRequestTags) SetValue(v string) *ListTrafficMirrorFiltersRequestTags {
	s.Value = &v
	return s
}

func (s *ListTrafficMirrorFiltersRequestTags) Validate() error {
	return dara.Validate(s)
}
