// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListServiceInstancesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFilter(v []*ListServiceInstancesRequestFilter) *ListServiceInstancesRequest
	GetFilter() []*ListServiceInstancesRequestFilter
	SetMaxResults(v int32) *ListServiceInstancesRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListServiceInstancesRequest
	GetNextToken() *string
	SetRegionId(v string) *ListServiceInstancesRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *ListServiceInstancesRequest
	GetResourceGroupId() *string
	SetTag(v []*ListServiceInstancesRequestTag) *ListServiceInstancesRequest
	GetTag() []*ListServiceInstancesRequestTag
}

type ListServiceInstancesRequest struct {
	// The filter.
	Filter []*ListServiceInstancesRequestFilter `json:"Filter,omitempty" xml:"Filter,omitempty" type:"Repeated"`
	// The number of entries to return on each page. Maximum value: 100. Default value: 20.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The query token. Set it to the **NextToken*	- value returned from the previous API call.
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
	// The resource group ID.
	//
	// example:
	//
	// rg-bp67acfmxazb4p****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The custom tags.
	Tag []*ListServiceInstancesRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s ListServiceInstancesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListServiceInstancesRequest) GoString() string {
	return s.String()
}

func (s *ListServiceInstancesRequest) GetFilter() []*ListServiceInstancesRequestFilter {
	return s.Filter
}

func (s *ListServiceInstancesRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListServiceInstancesRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListServiceInstancesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListServiceInstancesRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ListServiceInstancesRequest) GetTag() []*ListServiceInstancesRequestTag {
	return s.Tag
}

func (s *ListServiceInstancesRequest) SetFilter(v []*ListServiceInstancesRequestFilter) *ListServiceInstancesRequest {
	s.Filter = v
	return s
}

func (s *ListServiceInstancesRequest) SetMaxResults(v int32) *ListServiceInstancesRequest {
	s.MaxResults = &v
	return s
}

func (s *ListServiceInstancesRequest) SetNextToken(v string) *ListServiceInstancesRequest {
	s.NextToken = &v
	return s
}

func (s *ListServiceInstancesRequest) SetRegionId(v string) *ListServiceInstancesRequest {
	s.RegionId = &v
	return s
}

func (s *ListServiceInstancesRequest) SetResourceGroupId(v string) *ListServiceInstancesRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ListServiceInstancesRequest) SetTag(v []*ListServiceInstancesRequestTag) *ListServiceInstancesRequest {
	s.Tag = v
	return s
}

func (s *ListServiceInstancesRequest) Validate() error {
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

type ListServiceInstancesRequestFilter struct {
	// The name of the filter. You can specify one or more filter names to query resources. Valid values:
	//
	// - Name: The name of the service. To perform a fuzzy search, enter the value in the \\*xxx\\	- format. For example, if the service name is My Service, you can enter \\*My\\	- or \\*Service\\	- for a fuzzy search.
	//
	// - ServiceInstanceName: The name of the service instance. A fuzzy query is performed if you enter one service instance name. A term query is performed if you enter multiple service instance names.
	//
	// - ServiceInstanceId: The service instance ID.
	//
	// - ServiceId: The service ID.
	//
	// - Version: The service version.
	//
	// - Status: The instance status.
	//
	// - DeployType: The deployment type.
	//
	// - ServiceType: The service type.
	//
	// - OperationStartTimeBefore: The time before the start of the Alibaba Cloud Managed Services.
	//
	// - OperationStartTimeAfter: The time after the start of the Alibaba Cloud Managed Services.
	//
	// - OperationEndTimeBefore: The time before the end of the Alibaba Cloud Managed Services.
	//
	// - OperationEndTimeAfter: The time after the end of the Alibaba Cloud Managed Services.
	//
	// - OperatedServiceInstanceId: The ID of the managed service instance under a private service.
	//
	// - OperationServiceInstanceId: The ID of the service instance under a pure managed service.
	//
	// - EnableInstanceOps: Indicates whether the Alibaba Cloud Managed Services feature is enabled for the service instance.
	//
	// example:
	//
	// ServiceInstanceId
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The list of filter values.
	Value []*string `json:"Value,omitempty" xml:"Value,omitempty" type:"Repeated"`
}

func (s ListServiceInstancesRequestFilter) String() string {
	return dara.Prettify(s)
}

func (s ListServiceInstancesRequestFilter) GoString() string {
	return s.String()
}

func (s *ListServiceInstancesRequestFilter) GetName() *string {
	return s.Name
}

func (s *ListServiceInstancesRequestFilter) GetValue() []*string {
	return s.Value
}

func (s *ListServiceInstancesRequestFilter) SetName(v string) *ListServiceInstancesRequestFilter {
	s.Name = &v
	return s
}

func (s *ListServiceInstancesRequestFilter) SetValue(v []*string) *ListServiceInstancesRequestFilter {
	s.Value = v
	return s
}

func (s *ListServiceInstancesRequestFilter) Validate() error {
	return dara.Validate(s)
}

type ListServiceInstancesRequestTag struct {
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

func (s ListServiceInstancesRequestTag) String() string {
	return dara.Prettify(s)
}

func (s ListServiceInstancesRequestTag) GoString() string {
	return s.String()
}

func (s *ListServiceInstancesRequestTag) GetKey() *string {
	return s.Key
}

func (s *ListServiceInstancesRequestTag) GetValue() *string {
	return s.Value
}

func (s *ListServiceInstancesRequestTag) SetKey(v string) *ListServiceInstancesRequestTag {
	s.Key = &v
	return s
}

func (s *ListServiceInstancesRequestTag) SetValue(v string) *ListServiceInstancesRequestTag {
	s.Value = &v
	return s
}

func (s *ListServiceInstancesRequestTag) Validate() error {
	return dara.Validate(s)
}
