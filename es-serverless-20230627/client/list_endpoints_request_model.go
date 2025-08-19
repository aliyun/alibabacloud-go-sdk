// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListEndpointsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int32) *ListEndpointsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListEndpointsRequest
	GetPageSize() *int32
	SetResourceId(v string) *ListEndpointsRequest
	GetResourceId() *string
	SetType(v string) *ListEndpointsRequest
	GetType() *string
	SetVpcId(v string) *ListEndpointsRequest
	GetVpcId() *string
}

type ListEndpointsRequest struct {
	// example:
	//
	// 1
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// example:
	//
	// ep-bp1id41dd116e52e****
	ResourceId *string `json:"resourceId,omitempty" xml:"resourceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// VPC
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// example:
	//
	// vpc-bp1212sb7cj2j4e6x****
	VpcId *string `json:"vpcId,omitempty" xml:"vpcId,omitempty"`
}

func (s ListEndpointsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListEndpointsRequest) GoString() string {
	return s.String()
}

func (s *ListEndpointsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListEndpointsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListEndpointsRequest) GetResourceId() *string {
	return s.ResourceId
}

func (s *ListEndpointsRequest) GetType() *string {
	return s.Type
}

func (s *ListEndpointsRequest) GetVpcId() *string {
	return s.VpcId
}

func (s *ListEndpointsRequest) SetPageNumber(v int32) *ListEndpointsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListEndpointsRequest) SetPageSize(v int32) *ListEndpointsRequest {
	s.PageSize = &v
	return s
}

func (s *ListEndpointsRequest) SetResourceId(v string) *ListEndpointsRequest {
	s.ResourceId = &v
	return s
}

func (s *ListEndpointsRequest) SetType(v string) *ListEndpointsRequest {
	s.Type = &v
	return s
}

func (s *ListEndpointsRequest) SetVpcId(v string) *ListEndpointsRequest {
	s.VpcId = &v
	return s
}

func (s *ListEndpointsRequest) Validate() error {
	return dara.Validate(s)
}
