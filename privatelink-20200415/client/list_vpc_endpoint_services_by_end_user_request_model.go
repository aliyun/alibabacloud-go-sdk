// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListVpcEndpointServicesByEndUserRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListVpcEndpointServicesByEndUserRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListVpcEndpointServicesByEndUserRequest
	GetNextToken() *string
	SetRegionId(v string) *ListVpcEndpointServicesByEndUserRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *ListVpcEndpointServicesByEndUserRequest
	GetResourceGroupId() *string
	SetServiceId(v string) *ListVpcEndpointServicesByEndUserRequest
	GetServiceId() *string
	SetServiceName(v string) *ListVpcEndpointServicesByEndUserRequest
	GetServiceName() *string
	SetServiceRegionId(v string) *ListVpcEndpointServicesByEndUserRequest
	GetServiceRegionId() *string
	SetServiceType(v string) *ListVpcEndpointServicesByEndUserRequest
	GetServiceType() *string
	SetTag(v []*ListVpcEndpointServicesByEndUserRequestTag) *ListVpcEndpointServicesByEndUserRequest
	GetTag() []*ListVpcEndpointServicesByEndUserRequestTag
}

type ListVpcEndpointServicesByEndUserRequest struct {
	// The number of entries to return on each page. Valid values: **1*	- to **1000**. Default value: **50**.
	//
	// example:
	//
	// 50
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token.
	//
	// - If this is your first request, do not specify this parameter.
	//
	// - If more results are available, set this parameter to the **NextToken*	- value from the previous response to retrieve the next page.
	//
	// example:
	//
	// FFmyTO70tTpLG6I3FmYAXGKPd****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The region ID.
	//
	// You can call the [DescribeRegions](https://help.aliyun.com/document_detail/120468.html) operation to obtain the region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-huhehaote
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The resource group ID.
	//
	// example:
	//
	// rg-acfmy*****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The ID of the endpoint service.
	//
	// example:
	//
	// epsrv-hp3vpx8yqxblby3i****
	ServiceId *string `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
	// The name of the endpoint service.
	//
	// example:
	//
	// com.aliyuncs.privatelink.cn-huhehaote.epsrv-hp3xdsq46ael67lo****
	ServiceName *string `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
	// The region ID of the endpoint service.
	//
	// example:
	//
	// cn-beijing
	ServiceRegionId *string `json:"ServiceRegionId,omitempty" xml:"ServiceRegionId,omitempty"`
	// The type of the endpoint service.
	//
	// Only **Interface*	- is supported. An interface endpoint allows you to use Application Load Balancers (ALBs), Classic Load Balancers (CLBs), and Network Load Balancers (NLBs) as service resources.
	//
	// example:
	//
	// Interface
	ServiceType *string `json:"ServiceType,omitempty" xml:"ServiceType,omitempty"`
	// A list of tags to filter resources. You can specify up to 20 tags.
	Tag []*ListVpcEndpointServicesByEndUserRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s ListVpcEndpointServicesByEndUserRequest) String() string {
	return dara.Prettify(s)
}

func (s ListVpcEndpointServicesByEndUserRequest) GoString() string {
	return s.String()
}

func (s *ListVpcEndpointServicesByEndUserRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListVpcEndpointServicesByEndUserRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListVpcEndpointServicesByEndUserRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListVpcEndpointServicesByEndUserRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ListVpcEndpointServicesByEndUserRequest) GetServiceId() *string {
	return s.ServiceId
}

func (s *ListVpcEndpointServicesByEndUserRequest) GetServiceName() *string {
	return s.ServiceName
}

func (s *ListVpcEndpointServicesByEndUserRequest) GetServiceRegionId() *string {
	return s.ServiceRegionId
}

func (s *ListVpcEndpointServicesByEndUserRequest) GetServiceType() *string {
	return s.ServiceType
}

func (s *ListVpcEndpointServicesByEndUserRequest) GetTag() []*ListVpcEndpointServicesByEndUserRequestTag {
	return s.Tag
}

func (s *ListVpcEndpointServicesByEndUserRequest) SetMaxResults(v int32) *ListVpcEndpointServicesByEndUserRequest {
	s.MaxResults = &v
	return s
}

func (s *ListVpcEndpointServicesByEndUserRequest) SetNextToken(v string) *ListVpcEndpointServicesByEndUserRequest {
	s.NextToken = &v
	return s
}

func (s *ListVpcEndpointServicesByEndUserRequest) SetRegionId(v string) *ListVpcEndpointServicesByEndUserRequest {
	s.RegionId = &v
	return s
}

func (s *ListVpcEndpointServicesByEndUserRequest) SetResourceGroupId(v string) *ListVpcEndpointServicesByEndUserRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ListVpcEndpointServicesByEndUserRequest) SetServiceId(v string) *ListVpcEndpointServicesByEndUserRequest {
	s.ServiceId = &v
	return s
}

func (s *ListVpcEndpointServicesByEndUserRequest) SetServiceName(v string) *ListVpcEndpointServicesByEndUserRequest {
	s.ServiceName = &v
	return s
}

func (s *ListVpcEndpointServicesByEndUserRequest) SetServiceRegionId(v string) *ListVpcEndpointServicesByEndUserRequest {
	s.ServiceRegionId = &v
	return s
}

func (s *ListVpcEndpointServicesByEndUserRequest) SetServiceType(v string) *ListVpcEndpointServicesByEndUserRequest {
	s.ServiceType = &v
	return s
}

func (s *ListVpcEndpointServicesByEndUserRequest) SetTag(v []*ListVpcEndpointServicesByEndUserRequestTag) *ListVpcEndpointServicesByEndUserRequest {
	s.Tag = v
	return s
}

func (s *ListVpcEndpointServicesByEndUserRequest) Validate() error {
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

type ListVpcEndpointServicesByEndUserRequestTag struct {
	// The key of the tag. The tag key cannot be an empty string.
	//
	// The tag key can be up to 64 characters in length. It cannot start with `aliyun` or `acs:` and cannot contain `http://` or `https://`.
	//
	// example:
	//
	// FinanceDept
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of the tag. The tag value can be an empty string.
	//
	// The tag value can be up to 128 characters in length. It cannot start with `aliyun` or `acs:` and cannot contain `http://` or `https://`.
	//
	// example:
	//
	// FinanceJoshua
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListVpcEndpointServicesByEndUserRequestTag) String() string {
	return dara.Prettify(s)
}

func (s ListVpcEndpointServicesByEndUserRequestTag) GoString() string {
	return s.String()
}

func (s *ListVpcEndpointServicesByEndUserRequestTag) GetKey() *string {
	return s.Key
}

func (s *ListVpcEndpointServicesByEndUserRequestTag) GetValue() *string {
	return s.Value
}

func (s *ListVpcEndpointServicesByEndUserRequestTag) SetKey(v string) *ListVpcEndpointServicesByEndUserRequestTag {
	s.Key = &v
	return s
}

func (s *ListVpcEndpointServicesByEndUserRequestTag) SetValue(v string) *ListVpcEndpointServicesByEndUserRequestTag {
	s.Value = &v
	return s
}

func (s *ListVpcEndpointServicesByEndUserRequestTag) Validate() error {
	return dara.Validate(s)
}
