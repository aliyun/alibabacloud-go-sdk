// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateInstanceAclRequest interface {
	dara.Model
	String() string
	GoString() string
	SetActions(v []*string) *UpdateInstanceAclRequest
	GetActions() []*string
	SetDecision(v string) *UpdateInstanceAclRequest
	GetDecision() *string
	SetIpWhitelists(v []*string) *UpdateInstanceAclRequest
	GetIpWhitelists() []*string
	SetResourceName(v string) *UpdateInstanceAclRequest
	GetResourceName() *string
	SetResourceType(v string) *UpdateInstanceAclRequest
	GetResourceType() *string
}

type UpdateInstanceAclRequest struct {
	// The following types of operations are supported based on the resource type:
	//
	// 	- Topic: Pub, Sub, and Pub|Sub
	//
	// 	- Group: Sub
	//
	// Valid values:
	//
	// 	- Sub: subscribe
	//
	// 	- Pub|Sub: publish and subscribe
	//
	// 	- Pub: publish
	//
	// example:
	//
	// Pub
	Actions []*string `json:"actions,omitempty" xml:"actions,omitempty" type:"Repeated"`
	// The decision result of the authorization.
	//
	// Valid values:
	//
	// 	- Deny
	//
	// 	- Allow
	//
	// example:
	//
	// Allow
	Decision *string `json:"decision,omitempty" xml:"decision,omitempty"`
	// The IP addresses in the whitelist.
	IpWhitelists []*string `json:"ipWhitelists,omitempty" xml:"ipWhitelists,omitempty" type:"Repeated"`
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

func (s UpdateInstanceAclRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateInstanceAclRequest) GoString() string {
	return s.String()
}

func (s *UpdateInstanceAclRequest) GetActions() []*string {
	return s.Actions
}

func (s *UpdateInstanceAclRequest) GetDecision() *string {
	return s.Decision
}

func (s *UpdateInstanceAclRequest) GetIpWhitelists() []*string {
	return s.IpWhitelists
}

func (s *UpdateInstanceAclRequest) GetResourceName() *string {
	return s.ResourceName
}

func (s *UpdateInstanceAclRequest) GetResourceType() *string {
	return s.ResourceType
}

func (s *UpdateInstanceAclRequest) SetActions(v []*string) *UpdateInstanceAclRequest {
	s.Actions = v
	return s
}

func (s *UpdateInstanceAclRequest) SetDecision(v string) *UpdateInstanceAclRequest {
	s.Decision = &v
	return s
}

func (s *UpdateInstanceAclRequest) SetIpWhitelists(v []*string) *UpdateInstanceAclRequest {
	s.IpWhitelists = v
	return s
}

func (s *UpdateInstanceAclRequest) SetResourceName(v string) *UpdateInstanceAclRequest {
	s.ResourceName = &v
	return s
}

func (s *UpdateInstanceAclRequest) SetResourceType(v string) *UpdateInstanceAclRequest {
	s.ResourceType = &v
	return s
}

func (s *UpdateInstanceAclRequest) Validate() error {
	return dara.Validate(s)
}
