// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetResourceTypePropertiesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetResourceType(v string) *GetResourceTypePropertiesRequest
	GetResourceType() *string
}

type GetResourceTypePropertiesRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// ACS::ECS::Instance
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
}

func (s GetResourceTypePropertiesRequest) String() string {
	return dara.Prettify(s)
}

func (s GetResourceTypePropertiesRequest) GoString() string {
	return s.String()
}

func (s *GetResourceTypePropertiesRequest) GetResourceType() *string {
	return s.ResourceType
}

func (s *GetResourceTypePropertiesRequest) SetResourceType(v string) *GetResourceTypePropertiesRequest {
	s.ResourceType = &v
	return s
}

func (s *GetResourceTypePropertiesRequest) Validate() error {
	return dara.Validate(s)
}
