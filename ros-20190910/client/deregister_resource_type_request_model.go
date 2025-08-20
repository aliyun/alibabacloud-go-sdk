// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeregisterResourceTypeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetResourceType(v string) *DeregisterResourceTypeRequest
	GetResourceType() *string
	SetVersionId(v string) *DeregisterResourceTypeRequest
	GetVersionId() *string
}

type DeregisterResourceTypeRequest struct {
	// The resource type.
	//
	// This parameter is required.
	//
	// example:
	//
	// MODULE::MyOrganization::MyService::MyUsecase
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The version ID. If you want to delete a version of the resource type, you must specify this parameter.
	//
	// example:
	//
	// v1
	VersionId *string `json:"VersionId,omitempty" xml:"VersionId,omitempty"`
}

func (s DeregisterResourceTypeRequest) String() string {
	return dara.Prettify(s)
}

func (s DeregisterResourceTypeRequest) GoString() string {
	return s.String()
}

func (s *DeregisterResourceTypeRequest) GetResourceType() *string {
	return s.ResourceType
}

func (s *DeregisterResourceTypeRequest) GetVersionId() *string {
	return s.VersionId
}

func (s *DeregisterResourceTypeRequest) SetResourceType(v string) *DeregisterResourceTypeRequest {
	s.ResourceType = &v
	return s
}

func (s *DeregisterResourceTypeRequest) SetVersionId(v string) *DeregisterResourceTypeRequest {
	s.VersionId = &v
	return s
}

func (s *DeregisterResourceTypeRequest) Validate() error {
	return dara.Validate(s)
}
