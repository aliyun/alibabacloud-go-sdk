// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateInstanceAclRequest interface {
	dara.Model
	String() string
	GoString() string
	SetActions(v []*string) *CreateInstanceAclRequest
	GetActions() []*string
	SetDecision(v string) *CreateInstanceAclRequest
	GetDecision() *string
	SetIpWhitelists(v []*string) *CreateInstanceAclRequest
	GetIpWhitelists() []*string
	SetResourceName(v string) *CreateInstanceAclRequest
	GetResourceName() *string
	SetResourceType(v string) *CreateInstanceAclRequest
	GetResourceType() *string
}

type CreateInstanceAclRequest struct {
	// The type of operations that can be performed on the resource.
	//
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
	// This parameter is required.
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
	// This parameter is required.
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

func (s CreateInstanceAclRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateInstanceAclRequest) GoString() string {
	return s.String()
}

func (s *CreateInstanceAclRequest) GetActions() []*string {
	return s.Actions
}

func (s *CreateInstanceAclRequest) GetDecision() *string {
	return s.Decision
}

func (s *CreateInstanceAclRequest) GetIpWhitelists() []*string {
	return s.IpWhitelists
}

func (s *CreateInstanceAclRequest) GetResourceName() *string {
	return s.ResourceName
}

func (s *CreateInstanceAclRequest) GetResourceType() *string {
	return s.ResourceType
}

func (s *CreateInstanceAclRequest) SetActions(v []*string) *CreateInstanceAclRequest {
	s.Actions = v
	return s
}

func (s *CreateInstanceAclRequest) SetDecision(v string) *CreateInstanceAclRequest {
	s.Decision = &v
	return s
}

func (s *CreateInstanceAclRequest) SetIpWhitelists(v []*string) *CreateInstanceAclRequest {
	s.IpWhitelists = v
	return s
}

func (s *CreateInstanceAclRequest) SetResourceName(v string) *CreateInstanceAclRequest {
	s.ResourceName = &v
	return s
}

func (s *CreateInstanceAclRequest) SetResourceType(v string) *CreateInstanceAclRequest {
	s.ResourceType = &v
	return s
}

func (s *CreateInstanceAclRequest) Validate() error {
	return dara.Validate(s)
}
