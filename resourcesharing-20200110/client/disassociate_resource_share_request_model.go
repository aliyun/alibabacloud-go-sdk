// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisassociateResourceShareRequest interface {
	dara.Model
	String() string
	GoString() string
	SetResourceArns(v []*string) *DisassociateResourceShareRequest
	GetResourceArns() []*string
	SetResourceOwner(v string) *DisassociateResourceShareRequest
	GetResourceOwner() *string
	SetResourceShareId(v string) *DisassociateResourceShareRequest
	GetResourceShareId() *string
	SetResources(v []*DisassociateResourceShareRequestResources) *DisassociateResourceShareRequest
	GetResources() []*DisassociateResourceShareRequestResources
	SetTargets(v []*string) *DisassociateResourceShareRequest
	GetTargets() []*string
}

type DisassociateResourceShareRequest struct {
	ResourceArns []*string `json:"ResourceArns,omitempty" xml:"ResourceArns,omitempty" type:"Repeated"`
	// The owner of the resource share. Valid values:
	//
	// 	- Self: The resource share belongs to the current account. This is the default value. For resource sharing within a resource directory, if you are a resource owner and you want to disassociate resources or principals from a resource share, set this parameter to Self.
	//
	// 	- OtherAccounts: The resource share belongs to another account. For resource sharing outside a resource directory, if you are a principal and you want to exit a resource share, set this parameter to OtherAccounts.
	//
	// example:
	//
	// Self
	ResourceOwner *string `json:"ResourceOwner,omitempty" xml:"ResourceOwner,omitempty"`
	// The ID of the resource share.
	//
	// This parameter is required.
	//
	// example:
	//
	// rs-6GRmdD3X****
	ResourceShareId *string `json:"ResourceShareId,omitempty" xml:"ResourceShareId,omitempty"`
	// The information about the resources.
	Resources []*DisassociateResourceShareRequestResources `json:"Resources,omitempty" xml:"Resources,omitempty" type:"Repeated"`
	// The information about the principals.
	//
	// example:
	//
	// 172050525300****
	Targets []*string `json:"Targets,omitempty" xml:"Targets,omitempty" type:"Repeated"`
}

func (s DisassociateResourceShareRequest) String() string {
	return dara.Prettify(s)
}

func (s DisassociateResourceShareRequest) GoString() string {
	return s.String()
}

func (s *DisassociateResourceShareRequest) GetResourceArns() []*string {
	return s.ResourceArns
}

func (s *DisassociateResourceShareRequest) GetResourceOwner() *string {
	return s.ResourceOwner
}

func (s *DisassociateResourceShareRequest) GetResourceShareId() *string {
	return s.ResourceShareId
}

func (s *DisassociateResourceShareRequest) GetResources() []*DisassociateResourceShareRequestResources {
	return s.Resources
}

func (s *DisassociateResourceShareRequest) GetTargets() []*string {
	return s.Targets
}

func (s *DisassociateResourceShareRequest) SetResourceArns(v []*string) *DisassociateResourceShareRequest {
	s.ResourceArns = v
	return s
}

func (s *DisassociateResourceShareRequest) SetResourceOwner(v string) *DisassociateResourceShareRequest {
	s.ResourceOwner = &v
	return s
}

func (s *DisassociateResourceShareRequest) SetResourceShareId(v string) *DisassociateResourceShareRequest {
	s.ResourceShareId = &v
	return s
}

func (s *DisassociateResourceShareRequest) SetResources(v []*DisassociateResourceShareRequestResources) *DisassociateResourceShareRequest {
	s.Resources = v
	return s
}

func (s *DisassociateResourceShareRequest) SetTargets(v []*string) *DisassociateResourceShareRequest {
	s.Targets = v
	return s
}

func (s *DisassociateResourceShareRequest) Validate() error {
	if s.Resources != nil {
		for _, item := range s.Resources {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DisassociateResourceShareRequestResources struct {
	// The ID of the shared resource.
	//
	// Valid values of N: 1 to 5. This indicates that a maximum of five shared resources can be specified at a time.
	//
	// >  Resources.N.ResourceId and Resources.N.ResourceType must be used in pairs.
	//
	// example:
	//
	// vsw-bp183p93qs667muql****
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// The type of a shared resource.
	//
	// Valid values of N: 1 to 5. This indicates that a maximum of five shared resources can be specified at a time.
	//
	// For more information about the types of resources that can be shared, see [Services that work with Resource Sharing](https://help.aliyun.com/document_detail/450526.html).
	//
	// >  Resources.N.ResourceId and Resources.N.ResourceType must be used in pairs.
	//
	// example:
	//
	// VSwitch
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
}

func (s DisassociateResourceShareRequestResources) String() string {
	return dara.Prettify(s)
}

func (s DisassociateResourceShareRequestResources) GoString() string {
	return s.String()
}

func (s *DisassociateResourceShareRequestResources) GetResourceId() *string {
	return s.ResourceId
}

func (s *DisassociateResourceShareRequestResources) GetResourceType() *string {
	return s.ResourceType
}

func (s *DisassociateResourceShareRequestResources) SetResourceId(v string) *DisassociateResourceShareRequestResources {
	s.ResourceId = &v
	return s
}

func (s *DisassociateResourceShareRequestResources) SetResourceType(v string) *DisassociateResourceShareRequestResources {
	s.ResourceType = &v
	return s
}

func (s *DisassociateResourceShareRequestResources) Validate() error {
	return dara.Validate(s)
}
