// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListServicesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAllVersions(v bool) *ListServicesRequest
	GetAllVersions() *bool
	SetFilter(v []*ListServicesRequestFilter) *ListServicesRequest
	GetFilter() []*ListServicesRequestFilter
	SetMaxResults(v int32) *ListServicesRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListServicesRequest
	GetNextToken() *string
	SetRegionId(v string) *ListServicesRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *ListServicesRequest
	GetResourceGroupId() *string
	SetTag(v []*ListServicesRequestTag) *ListServicesRequest
	GetTag() []*ListServicesRequestTag
}

type ListServicesRequest struct {
	// Specifies whether to return all versions of the service. The default value is false. If this parameter is set to false, only the default version of each service is returned.
	//
	// example:
	//
	// false
	AllVersions *bool `json:"AllVersions,omitempty" xml:"AllVersions,omitempty"`
	// The filter.
	Filter []*ListServicesRequestFilter `json:"Filter,omitempty" xml:"Filter,omitempty" type:"Repeated"`
	// The number of entries to return on each page. Maximum value: 100. Default value: 20.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token that is used to retrieve the next page of results. Set this parameter to the NextToken value returned by a previous call.
	//
	// example:
	//
	// BBBAAfu+XtuBE55iRLHEYYuojI4=
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The region ID.
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
	// rg-aekzkt5bu****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The custom tags.
	Tag []*ListServicesRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s ListServicesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListServicesRequest) GoString() string {
	return s.String()
}

func (s *ListServicesRequest) GetAllVersions() *bool {
	return s.AllVersions
}

func (s *ListServicesRequest) GetFilter() []*ListServicesRequestFilter {
	return s.Filter
}

func (s *ListServicesRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListServicesRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListServicesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListServicesRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ListServicesRequest) GetTag() []*ListServicesRequestTag {
	return s.Tag
}

func (s *ListServicesRequest) SetAllVersions(v bool) *ListServicesRequest {
	s.AllVersions = &v
	return s
}

func (s *ListServicesRequest) SetFilter(v []*ListServicesRequestFilter) *ListServicesRequest {
	s.Filter = v
	return s
}

func (s *ListServicesRequest) SetMaxResults(v int32) *ListServicesRequest {
	s.MaxResults = &v
	return s
}

func (s *ListServicesRequest) SetNextToken(v string) *ListServicesRequest {
	s.NextToken = &v
	return s
}

func (s *ListServicesRequest) SetRegionId(v string) *ListServicesRequest {
	s.RegionId = &v
	return s
}

func (s *ListServicesRequest) SetResourceGroupId(v string) *ListServicesRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ListServicesRequest) SetTag(v []*ListServicesRequestTag) *ListServicesRequest {
	s.Tag = v
	return s
}

func (s *ListServicesRequest) Validate() error {
	if s.Filter != nil {
		for _, item := range s.Filter {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
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

type ListServicesRequestFilter struct {
	// The name of the filter. You can query by one or more filter names. Valid values:
	//
	// - ServiceId: The service ID.
	//
	// - Name: The service name.
	//
	// - Status: The service status.
	//
	// - SupplierName: The name of the service provider.
	//
	// example:
	//
	// Status
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The filter values.
	Value []*string `json:"Value,omitempty" xml:"Value,omitempty" type:"Repeated"`
}

func (s ListServicesRequestFilter) String() string {
	return dara.Prettify(s)
}

func (s ListServicesRequestFilter) GoString() string {
	return s.String()
}

func (s *ListServicesRequestFilter) GetName() *string {
	return s.Name
}

func (s *ListServicesRequestFilter) GetValue() []*string {
	return s.Value
}

func (s *ListServicesRequestFilter) SetName(v string) *ListServicesRequestFilter {
	s.Name = &v
	return s
}

func (s *ListServicesRequestFilter) SetValue(v []*string) *ListServicesRequestFilter {
	s.Value = v
	return s
}

func (s *ListServicesRequestFilter) Validate() error {
	return dara.Validate(s)
}

type ListServicesRequestTag struct {
	// The tag key.
	//
	// example:
	//
	// key1
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value.
	//
	// example:
	//
	// value1
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListServicesRequestTag) String() string {
	return dara.Prettify(s)
}

func (s ListServicesRequestTag) GoString() string {
	return s.String()
}

func (s *ListServicesRequestTag) GetKey() *string {
	return s.Key
}

func (s *ListServicesRequestTag) GetValue() *string {
	return s.Value
}

func (s *ListServicesRequestTag) SetKey(v string) *ListServicesRequestTag {
	s.Key = &v
	return s
}

func (s *ListServicesRequestTag) SetValue(v string) *ListServicesRequestTag {
	s.Value = &v
	return s
}

func (s *ListServicesRequestTag) Validate() error {
	return dara.Validate(s)
}
