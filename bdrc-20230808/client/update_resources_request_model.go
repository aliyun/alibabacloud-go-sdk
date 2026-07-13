// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateResourcesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetResourceType(v string) *UpdateResourcesRequest
	GetResourceType() *string
}

type UpdateResourcesRequest struct {
	// The resource type. If this parameter is not specified, all types of resources are updated.
	//
	// example:
	//
	// ACS::ECS::Instance
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
}

func (s UpdateResourcesRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateResourcesRequest) GoString() string {
	return s.String()
}

func (s *UpdateResourcesRequest) GetResourceType() *string {
	return s.ResourceType
}

func (s *UpdateResourcesRequest) SetResourceType(v string) *UpdateResourcesRequest {
	s.ResourceType = &v
	return s
}

func (s *UpdateResourcesRequest) Validate() error {
	return dara.Validate(s)
}
