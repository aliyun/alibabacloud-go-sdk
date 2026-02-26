// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMultiAccountResourceCountsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetFilters(v []*GetMultiAccountResourceCountsResponseBodyFilters) *GetMultiAccountResourceCountsResponseBody
	GetFilters() []*GetMultiAccountResourceCountsResponseBodyFilters
	SetGroupByKey(v string) *GetMultiAccountResourceCountsResponseBody
	GetGroupByKey() *string
	SetRequestId(v string) *GetMultiAccountResourceCountsResponseBody
	GetRequestId() *string
	SetResourceCounts(v []*GetMultiAccountResourceCountsResponseBodyResourceCounts) *GetMultiAccountResourceCountsResponseBody
	GetResourceCounts() []*GetMultiAccountResourceCountsResponseBodyResourceCounts
	SetScope(v string) *GetMultiAccountResourceCountsResponseBody
	GetScope() *string
}

type GetMultiAccountResourceCountsResponseBody struct {
	// The filter condition.
	Filters []*GetMultiAccountResourceCountsResponseBodyFilters `json:"Filters,omitempty" xml:"Filters,omitempty" type:"Repeated"`
	// The dimension by which resources are queried.
	//
	// example:
	//
	// ResourceType
	GroupByKey *string `json:"GroupByKey,omitempty" xml:"GroupByKey,omitempty"`
	// The request ID.
	//
	// example:
	//
	// EFA806B9-7F36-55AB-8B7A-D680C2C5EE57
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The numbers of resources.
	ResourceCounts []*GetMultiAccountResourceCountsResponseBodyResourceCounts `json:"ResourceCounts,omitempty" xml:"ResourceCounts,omitempty" type:"Repeated"`
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

func (s GetMultiAccountResourceCountsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetMultiAccountResourceCountsResponseBody) GoString() string {
	return s.String()
}

func (s *GetMultiAccountResourceCountsResponseBody) GetFilters() []*GetMultiAccountResourceCountsResponseBodyFilters {
	return s.Filters
}

func (s *GetMultiAccountResourceCountsResponseBody) GetGroupByKey() *string {
	return s.GroupByKey
}

func (s *GetMultiAccountResourceCountsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetMultiAccountResourceCountsResponseBody) GetResourceCounts() []*GetMultiAccountResourceCountsResponseBodyResourceCounts {
	return s.ResourceCounts
}

func (s *GetMultiAccountResourceCountsResponseBody) GetScope() *string {
	return s.Scope
}

func (s *GetMultiAccountResourceCountsResponseBody) SetFilters(v []*GetMultiAccountResourceCountsResponseBodyFilters) *GetMultiAccountResourceCountsResponseBody {
	s.Filters = v
	return s
}

func (s *GetMultiAccountResourceCountsResponseBody) SetGroupByKey(v string) *GetMultiAccountResourceCountsResponseBody {
	s.GroupByKey = &v
	return s
}

func (s *GetMultiAccountResourceCountsResponseBody) SetRequestId(v string) *GetMultiAccountResourceCountsResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetMultiAccountResourceCountsResponseBody) SetResourceCounts(v []*GetMultiAccountResourceCountsResponseBodyResourceCounts) *GetMultiAccountResourceCountsResponseBody {
	s.ResourceCounts = v
	return s
}

func (s *GetMultiAccountResourceCountsResponseBody) SetScope(v string) *GetMultiAccountResourceCountsResponseBody {
	s.Scope = &v
	return s
}

func (s *GetMultiAccountResourceCountsResponseBody) Validate() error {
	if s.Filters != nil {
		for _, item := range s.Filters {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.ResourceCounts != nil {
		for _, item := range s.ResourceCounts {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetMultiAccountResourceCountsResponseBodyFilters struct {
	// The key of the filter condition.
	//
	// example:
	//
	// RegionId
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The values of the filter condition.
	Values []*string `json:"Values,omitempty" xml:"Values,omitempty" type:"Repeated"`
}

func (s GetMultiAccountResourceCountsResponseBodyFilters) String() string {
	return dara.Prettify(s)
}

func (s GetMultiAccountResourceCountsResponseBodyFilters) GoString() string {
	return s.String()
}

func (s *GetMultiAccountResourceCountsResponseBodyFilters) GetKey() *string {
	return s.Key
}

func (s *GetMultiAccountResourceCountsResponseBodyFilters) GetValues() []*string {
	return s.Values
}

func (s *GetMultiAccountResourceCountsResponseBodyFilters) SetKey(v string) *GetMultiAccountResourceCountsResponseBodyFilters {
	s.Key = &v
	return s
}

func (s *GetMultiAccountResourceCountsResponseBodyFilters) SetValues(v []*string) *GetMultiAccountResourceCountsResponseBodyFilters {
	s.Values = v
	return s
}

func (s *GetMultiAccountResourceCountsResponseBodyFilters) Validate() error {
	return dara.Validate(s)
}

type GetMultiAccountResourceCountsResponseBodyResourceCounts struct {
	// The number of resources.
	//
	// example:
	//
	// 2
	Count *int64 `json:"Count,omitempty" xml:"Count,omitempty"`
	// The group name.
	//
	// example:
	//
	// ACS::ECS::NetworkInterface
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
}

func (s GetMultiAccountResourceCountsResponseBodyResourceCounts) String() string {
	return dara.Prettify(s)
}

func (s GetMultiAccountResourceCountsResponseBodyResourceCounts) GoString() string {
	return s.String()
}

func (s *GetMultiAccountResourceCountsResponseBodyResourceCounts) GetCount() *int64 {
	return s.Count
}

func (s *GetMultiAccountResourceCountsResponseBodyResourceCounts) GetGroupName() *string {
	return s.GroupName
}

func (s *GetMultiAccountResourceCountsResponseBodyResourceCounts) SetCount(v int64) *GetMultiAccountResourceCountsResponseBodyResourceCounts {
	s.Count = &v
	return s
}

func (s *GetMultiAccountResourceCountsResponseBodyResourceCounts) SetGroupName(v string) *GetMultiAccountResourceCountsResponseBodyResourceCounts {
	s.GroupName = &v
	return s
}

func (s *GetMultiAccountResourceCountsResponseBodyResourceCounts) Validate() error {
	return dara.Validate(s)
}
