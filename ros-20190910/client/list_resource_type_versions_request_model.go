// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListResourceTypeVersionsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetResourceType(v string) *ListResourceTypeVersionsRequest
	GetResourceType() *string
}

type ListResourceTypeVersionsRequest struct {
	// The resource type.
	//
	// This parameter is required.
	//
	// example:
	//
	// MODULE::MyOrganization::MyService::MyUsecase
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
}

func (s ListResourceTypeVersionsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListResourceTypeVersionsRequest) GoString() string {
	return s.String()
}

func (s *ListResourceTypeVersionsRequest) GetResourceType() *string {
	return s.ResourceType
}

func (s *ListResourceTypeVersionsRequest) SetResourceType(v string) *ListResourceTypeVersionsRequest {
	s.ResourceType = &v
	return s
}

func (s *ListResourceTypeVersionsRequest) Validate() error {
	return dara.Validate(s)
}
