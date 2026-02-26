// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMultiAccountResourceCountsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFilter(v []*GetMultiAccountResourceCountsRequestFilter) *GetMultiAccountResourceCountsRequest
	GetFilter() []*GetMultiAccountResourceCountsRequestFilter
	SetGroupByKey(v string) *GetMultiAccountResourceCountsRequest
	GetGroupByKey() *string
	SetScope(v string) *GetMultiAccountResourceCountsRequest
	GetScope() *string
}

type GetMultiAccountResourceCountsRequest struct {
	// The filter condition.
	Filter []*GetMultiAccountResourceCountsRequestFilter `json:"Filter,omitempty" xml:"Filter,omitempty" type:"Repeated"`
	// The dimension by which resources are queried. Valid values:
	//
	// - ResourceType: resource type
	//
	// - RegionId: region
	//
	// - ResourceGroupId: resource group
	//
	// > If this parameter is not configured, the total number of resources that meet the conditions is returned.
	//
	// example:
	//
	// ResourceType
	GroupByKey *string `json:"GroupByKey,omitempty" xml:"GroupByKey,omitempty"`
	// The search scope. Valid values:
	//
	// - ID of a resource directory: Resources within the management account and all members of the resource directory are searched.
	//
	// - ID of the Root folder: Resources within all members in the Root folder and the subfolders of the Root folder are searched.
	//
	// - ID of a folder: Resources within all members in the folder are searched.
	//
	// - ID of a member: Resources within the member are searched.
	//
	// For information about how to obtain the ID of a resource directory, the Root folder, a folder, or a member, see [GetResourceDirectory](https://help.aliyun.com/document_detail/159995.html), [ListFoldersForParent](https://help.aliyun.com/document_detail/159997.html), or [ListAccounts](https://help.aliyun.com/document_detail/160016.html).
	//
	// example:
	//
	// rd-r4****
	Scope *string `json:"Scope,omitempty" xml:"Scope,omitempty"`
}

func (s GetMultiAccountResourceCountsRequest) String() string {
	return dara.Prettify(s)
}

func (s GetMultiAccountResourceCountsRequest) GoString() string {
	return s.String()
}

func (s *GetMultiAccountResourceCountsRequest) GetFilter() []*GetMultiAccountResourceCountsRequestFilter {
	return s.Filter
}

func (s *GetMultiAccountResourceCountsRequest) GetGroupByKey() *string {
	return s.GroupByKey
}

func (s *GetMultiAccountResourceCountsRequest) GetScope() *string {
	return s.Scope
}

func (s *GetMultiAccountResourceCountsRequest) SetFilter(v []*GetMultiAccountResourceCountsRequestFilter) *GetMultiAccountResourceCountsRequest {
	s.Filter = v
	return s
}

func (s *GetMultiAccountResourceCountsRequest) SetGroupByKey(v string) *GetMultiAccountResourceCountsRequest {
	s.GroupByKey = &v
	return s
}

func (s *GetMultiAccountResourceCountsRequest) SetScope(v string) *GetMultiAccountResourceCountsRequest {
	s.Scope = &v
	return s
}

func (s *GetMultiAccountResourceCountsRequest) Validate() error {
	if s.Filter != nil {
		for _, item := range s.Filter {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetMultiAccountResourceCountsRequestFilter struct {
	// The key of the filter condition. For more information, see `Supported filter parameters`.
	//
	// example:
	//
	// RegionId
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The matching method.
	//
	// Set the value to Equals, which indicates an exact match.
	//
	// example:
	//
	// Equals
	MatchType *string `json:"MatchType,omitempty" xml:"MatchType,omitempty"`
	// The values of the filter condition.
	Value []*string `json:"Value,omitempty" xml:"Value,omitempty" type:"Repeated"`
}

func (s GetMultiAccountResourceCountsRequestFilter) String() string {
	return dara.Prettify(s)
}

func (s GetMultiAccountResourceCountsRequestFilter) GoString() string {
	return s.String()
}

func (s *GetMultiAccountResourceCountsRequestFilter) GetKey() *string {
	return s.Key
}

func (s *GetMultiAccountResourceCountsRequestFilter) GetMatchType() *string {
	return s.MatchType
}

func (s *GetMultiAccountResourceCountsRequestFilter) GetValue() []*string {
	return s.Value
}

func (s *GetMultiAccountResourceCountsRequestFilter) SetKey(v string) *GetMultiAccountResourceCountsRequestFilter {
	s.Key = &v
	return s
}

func (s *GetMultiAccountResourceCountsRequestFilter) SetMatchType(v string) *GetMultiAccountResourceCountsRequestFilter {
	s.MatchType = &v
	return s
}

func (s *GetMultiAccountResourceCountsRequestFilter) SetValue(v []*string) *GetMultiAccountResourceCountsRequestFilter {
	s.Value = v
	return s
}

func (s *GetMultiAccountResourceCountsRequestFilter) Validate() error {
	return dara.Validate(s)
}
