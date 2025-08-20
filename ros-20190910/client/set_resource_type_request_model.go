// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetResourceTypeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDefaultVersionId(v string) *SetResourceTypeRequest
	GetDefaultVersionId() *string
	SetDescription(v string) *SetResourceTypeRequest
	GetDescription() *string
	SetResourceType(v string) *SetResourceTypeRequest
	GetResourceType() *string
	SetVersionId(v string) *SetResourceTypeRequest
	GetVersionId() *string
}

type SetResourceTypeRequest struct {
	// The ID of the default version. You can use this parameter to specify the default version of the resource type.
	//
	// > You can specify only one of the VersionId and DefaultVersionId parameters.
	//
	// example:
	//
	// v1
	DefaultVersionId *string `json:"DefaultVersionId,omitempty" xml:"DefaultVersionId,omitempty"`
	// The description of the resource type or resource type version. The description can be up to 512 characters in length.
	//
	// example:
	//
	// It is a demo.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The resource type.
	//
	// This parameter is required.
	//
	// example:
	//
	// MODULE::MyOrganization::MyService::MyUsecase
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The version ID. If you want to modify a version of the resource type, you must specify this parameter. If you do not specify this parameter, only the resource type is modified.
	//
	// > You can specify only one of the VersionId and DefaultVersionId parameters.
	//
	// example:
	//
	// v1
	VersionId *string `json:"VersionId,omitempty" xml:"VersionId,omitempty"`
}

func (s SetResourceTypeRequest) String() string {
	return dara.Prettify(s)
}

func (s SetResourceTypeRequest) GoString() string {
	return s.String()
}

func (s *SetResourceTypeRequest) GetDefaultVersionId() *string {
	return s.DefaultVersionId
}

func (s *SetResourceTypeRequest) GetDescription() *string {
	return s.Description
}

func (s *SetResourceTypeRequest) GetResourceType() *string {
	return s.ResourceType
}

func (s *SetResourceTypeRequest) GetVersionId() *string {
	return s.VersionId
}

func (s *SetResourceTypeRequest) SetDefaultVersionId(v string) *SetResourceTypeRequest {
	s.DefaultVersionId = &v
	return s
}

func (s *SetResourceTypeRequest) SetDescription(v string) *SetResourceTypeRequest {
	s.Description = &v
	return s
}

func (s *SetResourceTypeRequest) SetResourceType(v string) *SetResourceTypeRequest {
	s.ResourceType = &v
	return s
}

func (s *SetResourceTypeRequest) SetVersionId(v string) *SetResourceTypeRequest {
	s.VersionId = &v
	return s
}

func (s *SetResourceTypeRequest) Validate() error {
	return dara.Validate(s)
}
