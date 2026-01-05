// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListHealthCheckTemplatesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetHealthCheckTemplateIds(v []*string) *ListHealthCheckTemplatesRequest
	GetHealthCheckTemplateIds() []*string
	SetHealthCheckTemplateNames(v []*string) *ListHealthCheckTemplatesRequest
	GetHealthCheckTemplateNames() []*string
	SetMaxResults(v int32) *ListHealthCheckTemplatesRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListHealthCheckTemplatesRequest
	GetNextToken() *string
	SetResourceGroupId(v string) *ListHealthCheckTemplatesRequest
	GetResourceGroupId() *string
	SetTag(v []*ListHealthCheckTemplatesRequestTag) *ListHealthCheckTemplatesRequest
	GetTag() []*ListHealthCheckTemplatesRequestTag
}

type ListHealthCheckTemplatesRequest struct {
	// The IDs of health check templates.
	HealthCheckTemplateIds []*string `json:"HealthCheckTemplateIds,omitempty" xml:"HealthCheckTemplateIds,omitempty" type:"Repeated"`
	// The health check templates.
	HealthCheckTemplateNames []*string `json:"HealthCheckTemplateNames,omitempty" xml:"HealthCheckTemplateNames,omitempty" type:"Repeated"`
	// The number of entries to return in each call. Valid values: **1*	- to **100**. Default value: **20**
	//
	// example:
	//
	// 50
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results. Valid values:
	//
	// 	- You do not need to specify this parameter for the first request.
	//
	// 	- You must specify the token that is obtained from the previous query as the value of **NextToken**.
	//
	// example:
	//
	// FFmyTO70tTpLG6I3FmYAXGKPd****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The resource group ID. You can filter the query results based on the specified ID.
	//
	// example:
	//
	// rg-atstuj3rtop****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The tags.
	Tag []*ListHealthCheckTemplatesRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s ListHealthCheckTemplatesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListHealthCheckTemplatesRequest) GoString() string {
	return s.String()
}

func (s *ListHealthCheckTemplatesRequest) GetHealthCheckTemplateIds() []*string {
	return s.HealthCheckTemplateIds
}

func (s *ListHealthCheckTemplatesRequest) GetHealthCheckTemplateNames() []*string {
	return s.HealthCheckTemplateNames
}

func (s *ListHealthCheckTemplatesRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListHealthCheckTemplatesRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListHealthCheckTemplatesRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ListHealthCheckTemplatesRequest) GetTag() []*ListHealthCheckTemplatesRequestTag {
	return s.Tag
}

func (s *ListHealthCheckTemplatesRequest) SetHealthCheckTemplateIds(v []*string) *ListHealthCheckTemplatesRequest {
	s.HealthCheckTemplateIds = v
	return s
}

func (s *ListHealthCheckTemplatesRequest) SetHealthCheckTemplateNames(v []*string) *ListHealthCheckTemplatesRequest {
	s.HealthCheckTemplateNames = v
	return s
}

func (s *ListHealthCheckTemplatesRequest) SetMaxResults(v int32) *ListHealthCheckTemplatesRequest {
	s.MaxResults = &v
	return s
}

func (s *ListHealthCheckTemplatesRequest) SetNextToken(v string) *ListHealthCheckTemplatesRequest {
	s.NextToken = &v
	return s
}

func (s *ListHealthCheckTemplatesRequest) SetResourceGroupId(v string) *ListHealthCheckTemplatesRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ListHealthCheckTemplatesRequest) SetTag(v []*ListHealthCheckTemplatesRequestTag) *ListHealthCheckTemplatesRequest {
	s.Tag = v
	return s
}

func (s *ListHealthCheckTemplatesRequest) Validate() error {
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

type ListHealthCheckTemplatesRequestTag struct {
	// The tag key. The tag key can be up to 128 characters in length and cannot start with `acs:` or `aliyun`. It cannot contain `http://` or `https://`.
	//
	// example:
	//
	// env
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value. The tag value can be up to 128 characters in length and cannot start with `acs:` or `aliyun`. It cannot contain `http://` or `https://`.
	//
	// example:
	//
	// product
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListHealthCheckTemplatesRequestTag) String() string {
	return dara.Prettify(s)
}

func (s ListHealthCheckTemplatesRequestTag) GoString() string {
	return s.String()
}

func (s *ListHealthCheckTemplatesRequestTag) GetKey() *string {
	return s.Key
}

func (s *ListHealthCheckTemplatesRequestTag) GetValue() *string {
	return s.Value
}

func (s *ListHealthCheckTemplatesRequestTag) SetKey(v string) *ListHealthCheckTemplatesRequestTag {
	s.Key = &v
	return s
}

func (s *ListHealthCheckTemplatesRequestTag) SetValue(v string) *ListHealthCheckTemplatesRequestTag {
	s.Value = &v
	return s
}

func (s *ListHealthCheckTemplatesRequestTag) Validate() error {
	return dara.Validate(s)
}
