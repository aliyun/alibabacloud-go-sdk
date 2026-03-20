// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListServerGroupsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListServerGroupsRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListServerGroupsRequest
	GetNextToken() *string
	SetResourceGroupId(v string) *ListServerGroupsRequest
	GetResourceGroupId() *string
	SetServerGroupIds(v []*string) *ListServerGroupsRequest
	GetServerGroupIds() []*string
	SetServerGroupNames(v []*string) *ListServerGroupsRequest
	GetServerGroupNames() []*string
	SetServerGroupType(v string) *ListServerGroupsRequest
	GetServerGroupType() *string
	SetSkip(v int32) *ListServerGroupsRequest
	GetSkip() *int32
	SetTag(v []*ListServerGroupsRequestTag) *ListServerGroupsRequest
	GetTag() []*ListServerGroupsRequestTag
	SetVpcId(v string) *ListServerGroupsRequest
	GetVpcId() *string
}

type ListServerGroupsRequest struct {
	// The number of entries per page.
	//
	// Valid values: 1 to 1000.
	//
	// Default value: 20.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results. Valid values:
	//
	// 	- If **NextToken*	- is empty, no next page exists.
	//
	// 	- If a value of **NextToken*	- is returned, the value indicates the token that is used for the next query.
	//
	// example:
	//
	// FFmyTO70tTpLG6I3FmYAXGKPd****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The ID of the resource group to which the instance belongs.
	//
	// example:
	//
	// rg-atstuj3rtop****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The server group IDs.
	//
	// You can specify at most 20 server group IDs in each call.
	ServerGroupIds []*string `json:"ServerGroupIds,omitempty" xml:"ServerGroupIds,omitempty" type:"Repeated"`
	// The server group names.
	//
	// You can specify at most 20 server group names in each call.
	ServerGroupNames []*string `json:"ServerGroupNames,omitempty" xml:"ServerGroupNames,omitempty" type:"Repeated"`
	// The server group type. Valid values:
	//
	// 	- **Instance**: allows you to specify resources of the **Ecs**, **Eni**, or **Eci*	- type.
	//
	// 	- **Ip**: allows you to add servers by specifying IP addresses.
	//
	// example:
	//
	// Instance
	ServerGroupType *string `json:"ServerGroupType,omitempty" xml:"ServerGroupType,omitempty"`
	// The number of entries to be skipped in the call.
	//
	// example:
	//
	// 1
	Skip *int32 `json:"Skip,omitempty" xml:"Skip,omitempty"`
	// The tag keys.
	//
	// You can specify at most 20 tags in each call.
	Tag []*ListServerGroupsRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// The VPC ID.
	//
	// example:
	//
	// vpc-bp15zckdt37pq72zv****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s ListServerGroupsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListServerGroupsRequest) GoString() string {
	return s.String()
}

func (s *ListServerGroupsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListServerGroupsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListServerGroupsRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ListServerGroupsRequest) GetServerGroupIds() []*string {
	return s.ServerGroupIds
}

func (s *ListServerGroupsRequest) GetServerGroupNames() []*string {
	return s.ServerGroupNames
}

func (s *ListServerGroupsRequest) GetServerGroupType() *string {
	return s.ServerGroupType
}

func (s *ListServerGroupsRequest) GetSkip() *int32 {
	return s.Skip
}

func (s *ListServerGroupsRequest) GetTag() []*ListServerGroupsRequestTag {
	return s.Tag
}

func (s *ListServerGroupsRequest) GetVpcId() *string {
	return s.VpcId
}

func (s *ListServerGroupsRequest) SetMaxResults(v int32) *ListServerGroupsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListServerGroupsRequest) SetNextToken(v string) *ListServerGroupsRequest {
	s.NextToken = &v
	return s
}

func (s *ListServerGroupsRequest) SetResourceGroupId(v string) *ListServerGroupsRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ListServerGroupsRequest) SetServerGroupIds(v []*string) *ListServerGroupsRequest {
	s.ServerGroupIds = v
	return s
}

func (s *ListServerGroupsRequest) SetServerGroupNames(v []*string) *ListServerGroupsRequest {
	s.ServerGroupNames = v
	return s
}

func (s *ListServerGroupsRequest) SetServerGroupType(v string) *ListServerGroupsRequest {
	s.ServerGroupType = &v
	return s
}

func (s *ListServerGroupsRequest) SetSkip(v int32) *ListServerGroupsRequest {
	s.Skip = &v
	return s
}

func (s *ListServerGroupsRequest) SetTag(v []*ListServerGroupsRequestTag) *ListServerGroupsRequest {
	s.Tag = v
	return s
}

func (s *ListServerGroupsRequest) SetVpcId(v string) *ListServerGroupsRequest {
	s.VpcId = &v
	return s
}

func (s *ListServerGroupsRequest) Validate() error {
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

type ListServerGroupsRequestTag struct {
	// The tag key. The tag key cannot be an empty string. The tag key can be up to 128 characters in length, and cannot start with `acs:` or `aliyun`. The tag key cannot contain `http://` or `https://`.
	//
	// example:
	//
	// testTagKey
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value. The tag value can be up to 256 characters in length and cannot contain `http://` or `https://`.
	//
	// example:
	//
	// testTagValue
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListServerGroupsRequestTag) String() string {
	return dara.Prettify(s)
}

func (s ListServerGroupsRequestTag) GoString() string {
	return s.String()
}

func (s *ListServerGroupsRequestTag) GetKey() *string {
	return s.Key
}

func (s *ListServerGroupsRequestTag) GetValue() *string {
	return s.Value
}

func (s *ListServerGroupsRequestTag) SetKey(v string) *ListServerGroupsRequestTag {
	s.Key = &v
	return s
}

func (s *ListServerGroupsRequestTag) SetValue(v string) *ListServerGroupsRequestTag {
	s.Value = &v
	return s
}

func (s *ListServerGroupsRequestTag) Validate() error {
	return dara.Validate(s)
}
