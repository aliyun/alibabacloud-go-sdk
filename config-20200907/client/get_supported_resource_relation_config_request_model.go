// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSupportedResourceRelationConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetResourceType(v string) *GetSupportedResourceRelationConfigRequest
	GetResourceType() *string
}

type GetSupportedResourceRelationConfigRequest struct {
	// The resource type.
	//
	// This parameter is required.
	//
	// example:
	//
	// ACS::ECS::Instance
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
}

func (s GetSupportedResourceRelationConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s GetSupportedResourceRelationConfigRequest) GoString() string {
	return s.String()
}

func (s *GetSupportedResourceRelationConfigRequest) GetResourceType() *string {
	return s.ResourceType
}

func (s *GetSupportedResourceRelationConfigRequest) SetResourceType(v string) *GetSupportedResourceRelationConfigRequest {
	s.ResourceType = &v
	return s
}

func (s *GetSupportedResourceRelationConfigRequest) Validate() error {
	return dara.Validate(s)
}
