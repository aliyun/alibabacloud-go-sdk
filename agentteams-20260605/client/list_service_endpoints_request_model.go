// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListServiceEndpointsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetComponent(v string) *ListServiceEndpointsRequest
	GetComponent() *string
	SetDomainType(v string) *ListServiceEndpointsRequest
	GetDomainType() *string
	SetInstanceId(v string) *ListServiceEndpointsRequest
	GetInstanceId() *string
	SetMaxResults(v int32) *ListServiceEndpointsRequest
	GetMaxResults() *int32
	SetNetworkType(v string) *ListServiceEndpointsRequest
	GetNetworkType() *string
	SetNextToken(v string) *ListServiceEndpointsRequest
	GetNextToken() *string
	SetResourceName(v string) *ListServiceEndpointsRequest
	GetResourceName() *string
	SetSkip(v string) *ListServiceEndpointsRequest
	GetSkip() *string
}

type ListServiceEndpointsRequest struct {
	// example:
	//
	// MATRIX
	Component *string `json:"Component,omitempty" xml:"Component,omitempty"`
	// example:
	//
	// CUSTOM
	DomainType *string `json:"DomainType,omitempty" xml:"DomainType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// agentteams-cn-xxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	MaxResults *int32  `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// INTERNET
	NetworkType *string `json:"NetworkType,omitempty" xml:"NetworkType,omitempty"`
	NextToken   *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// example:
	//
	// matrix-service
	ResourceName *string `json:"ResourceName,omitempty" xml:"ResourceName,omitempty"`
	Skip         *string `json:"Skip,omitempty" xml:"Skip,omitempty"`
}

func (s ListServiceEndpointsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListServiceEndpointsRequest) GoString() string {
	return s.String()
}

func (s *ListServiceEndpointsRequest) GetComponent() *string {
	return s.Component
}

func (s *ListServiceEndpointsRequest) GetDomainType() *string {
	return s.DomainType
}

func (s *ListServiceEndpointsRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListServiceEndpointsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListServiceEndpointsRequest) GetNetworkType() *string {
	return s.NetworkType
}

func (s *ListServiceEndpointsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListServiceEndpointsRequest) GetResourceName() *string {
	return s.ResourceName
}

func (s *ListServiceEndpointsRequest) GetSkip() *string {
	return s.Skip
}

func (s *ListServiceEndpointsRequest) SetComponent(v string) *ListServiceEndpointsRequest {
	s.Component = &v
	return s
}

func (s *ListServiceEndpointsRequest) SetDomainType(v string) *ListServiceEndpointsRequest {
	s.DomainType = &v
	return s
}

func (s *ListServiceEndpointsRequest) SetInstanceId(v string) *ListServiceEndpointsRequest {
	s.InstanceId = &v
	return s
}

func (s *ListServiceEndpointsRequest) SetMaxResults(v int32) *ListServiceEndpointsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListServiceEndpointsRequest) SetNetworkType(v string) *ListServiceEndpointsRequest {
	s.NetworkType = &v
	return s
}

func (s *ListServiceEndpointsRequest) SetNextToken(v string) *ListServiceEndpointsRequest {
	s.NextToken = &v
	return s
}

func (s *ListServiceEndpointsRequest) SetResourceName(v string) *ListServiceEndpointsRequest {
	s.ResourceName = &v
	return s
}

func (s *ListServiceEndpointsRequest) SetSkip(v string) *ListServiceEndpointsRequest {
	s.Skip = &v
	return s
}

func (s *ListServiceEndpointsRequest) Validate() error {
	return dara.Validate(s)
}
