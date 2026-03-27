// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListServicesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListServicesRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListServicesRequest
	GetNextToken() *string
	SetResourceGroupId(v string) *ListServicesRequest
	GetResourceGroupId() *string
	SetServiceName(v string) *ListServicesRequest
	GetServiceName() *string
	SetServiceType(v string) *ListServicesRequest
	GetServiceType() *string
	SetTags(v []*ListServicesRequestTags) *ListServicesRequest
	GetTags() []*ListServicesRequestTags
}

type ListServicesRequest struct {
	// The maximum number of records to return in this request.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// Token for the next query, an empty value indicates the last page.
	//
	// example:
	//
	// 7-b81a-4bc9-bbfa-a50cc6988667
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// example:
	//
	// rg-aekxxzuad5zzzz
	ResourceGroupId *string `json:"resourceGroupId,omitempty" xml:"resourceGroupId,omitempty"`
	// example:
	//
	// app-demo
	ServiceName *string `json:"serviceName,omitempty" xml:"serviceName,omitempty"`
	// Service type
	//
	// example:
	//
	// apm
	ServiceType *string `json:"serviceType,omitempty" xml:"serviceType,omitempty"`
	// if can be null:
	// true
	Tags []*ListServicesRequestTags `json:"tags,omitempty" xml:"tags,omitempty" type:"Repeated"`
}

func (s ListServicesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListServicesRequest) GoString() string {
	return s.String()
}

func (s *ListServicesRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListServicesRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListServicesRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ListServicesRequest) GetServiceName() *string {
	return s.ServiceName
}

func (s *ListServicesRequest) GetServiceType() *string {
	return s.ServiceType
}

func (s *ListServicesRequest) GetTags() []*ListServicesRequestTags {
	return s.Tags
}

func (s *ListServicesRequest) SetMaxResults(v int32) *ListServicesRequest {
	s.MaxResults = &v
	return s
}

func (s *ListServicesRequest) SetNextToken(v string) *ListServicesRequest {
	s.NextToken = &v
	return s
}

func (s *ListServicesRequest) SetResourceGroupId(v string) *ListServicesRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ListServicesRequest) SetServiceName(v string) *ListServicesRequest {
	s.ServiceName = &v
	return s
}

func (s *ListServicesRequest) SetServiceType(v string) *ListServicesRequest {
	s.ServiceType = &v
	return s
}

func (s *ListServicesRequest) SetTags(v []*ListServicesRequestTags) *ListServicesRequest {
	s.Tags = v
	return s
}

func (s *ListServicesRequest) Validate() error {
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

type ListServicesRequestTags struct {
	// example:
	//
	// evn
	Key *string `json:"key,omitempty" xml:"key,omitempty"`
	// example:
	//
	// prod
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s ListServicesRequestTags) String() string {
	return dara.Prettify(s)
}

func (s ListServicesRequestTags) GoString() string {
	return s.String()
}

func (s *ListServicesRequestTags) GetKey() *string {
	return s.Key
}

func (s *ListServicesRequestTags) GetValue() *string {
	return s.Value
}

func (s *ListServicesRequestTags) SetKey(v string) *ListServicesRequestTags {
	s.Key = &v
	return s
}

func (s *ListServicesRequestTags) SetValue(v string) *ListServicesRequestTags {
	s.Value = &v
	return s
}

func (s *ListServicesRequestTags) Validate() error {
	return dara.Validate(s)
}
