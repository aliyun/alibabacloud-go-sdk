// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetInstanceAclRequest interface {
	dara.Model
	String() string
	GoString() string
	SetResourceName(v string) *GetInstanceAclRequest
	GetResourceName() *string
	SetResourceType(v string) *GetInstanceAclRequest
	GetResourceType() *string
}

type GetInstanceAclRequest struct {
	// The name of the resource on which you want to grant permissions.
	//
	// This parameter is required.
	//
	// example:
	//
	// test
	ResourceName *string `json:"resourceName,omitempty" xml:"resourceName,omitempty"`
	// The type of the resource on which you want to grant permissions.
	//
	// Valid values:
	//
	// 	- Group
	//
	// 	- Topic
	//
	// This parameter is required.
	//
	// example:
	//
	// Topic
	ResourceType *string `json:"resourceType,omitempty" xml:"resourceType,omitempty"`
}

func (s GetInstanceAclRequest) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceAclRequest) GoString() string {
	return s.String()
}

func (s *GetInstanceAclRequest) GetResourceName() *string {
	return s.ResourceName
}

func (s *GetInstanceAclRequest) GetResourceType() *string {
	return s.ResourceType
}

func (s *GetInstanceAclRequest) SetResourceName(v string) *GetInstanceAclRequest {
	s.ResourceName = &v
	return s
}

func (s *GetInstanceAclRequest) SetResourceType(v string) *GetInstanceAclRequest {
	s.ResourceType = &v
	return s
}

func (s *GetInstanceAclRequest) Validate() error {
	return dara.Validate(s)
}
