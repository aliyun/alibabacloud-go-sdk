// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetResourceTypeTemplateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetResourceType(v string) *GetResourceTypeTemplateRequest
	GetResourceType() *string
	SetVersionId(v string) *GetResourceTypeTemplateRequest
	GetVersionId() *string
}

type GetResourceTypeTemplateRequest struct {
	// The resource type.
	//
	// This parameter is required.
	//
	// example:
	//
	// ALIYUN::ECS::VPC
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The version ID. If you want to query a specific version of the resource type, you must specify this parameter. If you do not specify this parameter, only the resource type is queried.
	//
	// > This parameter is supported only for modules.
	//
	// example:
	//
	// v1
	VersionId *string `json:"VersionId,omitempty" xml:"VersionId,omitempty"`
}

func (s GetResourceTypeTemplateRequest) String() string {
	return dara.Prettify(s)
}

func (s GetResourceTypeTemplateRequest) GoString() string {
	return s.String()
}

func (s *GetResourceTypeTemplateRequest) GetResourceType() *string {
	return s.ResourceType
}

func (s *GetResourceTypeTemplateRequest) GetVersionId() *string {
	return s.VersionId
}

func (s *GetResourceTypeTemplateRequest) SetResourceType(v string) *GetResourceTypeTemplateRequest {
	s.ResourceType = &v
	return s
}

func (s *GetResourceTypeTemplateRequest) SetVersionId(v string) *GetResourceTypeTemplateRequest {
	s.VersionId = &v
	return s
}

func (s *GetResourceTypeTemplateRequest) Validate() error {
	return dara.Validate(s)
}
