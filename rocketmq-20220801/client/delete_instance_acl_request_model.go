// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteInstanceAclRequest interface {
	dara.Model
	String() string
	GoString() string
	SetResourceName(v string) *DeleteInstanceAclRequest
	GetResourceName() *string
	SetResourceType(v string) *DeleteInstanceAclRequest
	GetResourceType() *string
}

type DeleteInstanceAclRequest struct {
	// The name of the resource on which the permissions are granted.
	//
	// This parameter is required.
	//
	// example:
	//
	// test
	ResourceName *string `json:"resourceName,omitempty" xml:"resourceName,omitempty"`
	// The type of the resource on which the permissions are granted.
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

func (s DeleteInstanceAclRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteInstanceAclRequest) GoString() string {
	return s.String()
}

func (s *DeleteInstanceAclRequest) GetResourceName() *string {
	return s.ResourceName
}

func (s *DeleteInstanceAclRequest) GetResourceType() *string {
	return s.ResourceType
}

func (s *DeleteInstanceAclRequest) SetResourceName(v string) *DeleteInstanceAclRequest {
	s.ResourceName = &v
	return s
}

func (s *DeleteInstanceAclRequest) SetResourceType(v string) *DeleteInstanceAclRequest {
	s.ResourceType = &v
	return s
}

func (s *DeleteInstanceAclRequest) Validate() error {
	return dara.Validate(s)
}
