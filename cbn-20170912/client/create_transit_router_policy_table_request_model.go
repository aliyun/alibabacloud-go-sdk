// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateTransitRouterPolicyTableRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *CreateTransitRouterPolicyTableRequest
	GetClientToken() *string
	SetDescription(v string) *CreateTransitRouterPolicyTableRequest
	GetDescription() *string
	SetDryRun(v bool) *CreateTransitRouterPolicyTableRequest
	GetDryRun() *bool
	SetName(v string) *CreateTransitRouterPolicyTableRequest
	GetName() *string
	SetTag(v []*CreateTransitRouterPolicyTableRequestTag) *CreateTransitRouterPolicyTableRequest
	GetTag() []*CreateTransitRouterPolicyTableRequestTag
	SetTransitRouterId(v string) *CreateTransitRouterPolicyTableRequest
	GetTransitRouterId() *string
}

type CreateTransitRouterPolicyTableRequest struct {
	// example:
	//
	// 02fb3da4-130e-11e9-8e44-001****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// example:
	//
	// descriptiontest
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// example:
	//
	// nametest
	Name *string                                     `json:"Name,omitempty" xml:"Name,omitempty"`
	Tag  []*CreateTransitRouterPolicyTableRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// tr-uf654ttymmljlvh2x****
	TransitRouterId *string `json:"TransitRouterId,omitempty" xml:"TransitRouterId,omitempty"`
}

func (s CreateTransitRouterPolicyTableRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateTransitRouterPolicyTableRequest) GoString() string {
	return s.String()
}

func (s *CreateTransitRouterPolicyTableRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateTransitRouterPolicyTableRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateTransitRouterPolicyTableRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *CreateTransitRouterPolicyTableRequest) GetName() *string {
	return s.Name
}

func (s *CreateTransitRouterPolicyTableRequest) GetTag() []*CreateTransitRouterPolicyTableRequestTag {
	return s.Tag
}

func (s *CreateTransitRouterPolicyTableRequest) GetTransitRouterId() *string {
	return s.TransitRouterId
}

func (s *CreateTransitRouterPolicyTableRequest) SetClientToken(v string) *CreateTransitRouterPolicyTableRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateTransitRouterPolicyTableRequest) SetDescription(v string) *CreateTransitRouterPolicyTableRequest {
	s.Description = &v
	return s
}

func (s *CreateTransitRouterPolicyTableRequest) SetDryRun(v bool) *CreateTransitRouterPolicyTableRequest {
	s.DryRun = &v
	return s
}

func (s *CreateTransitRouterPolicyTableRequest) SetName(v string) *CreateTransitRouterPolicyTableRequest {
	s.Name = &v
	return s
}

func (s *CreateTransitRouterPolicyTableRequest) SetTag(v []*CreateTransitRouterPolicyTableRequestTag) *CreateTransitRouterPolicyTableRequest {
	s.Tag = v
	return s
}

func (s *CreateTransitRouterPolicyTableRequest) SetTransitRouterId(v string) *CreateTransitRouterPolicyTableRequest {
	s.TransitRouterId = &v
	return s
}

func (s *CreateTransitRouterPolicyTableRequest) Validate() error {
	if s.Tag != nil {
		for _, item := range s.Tag {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateTransitRouterPolicyTableRequestTag struct {
}

func (s CreateTransitRouterPolicyTableRequestTag) String() string {
	return dara.Prettify(s)
}

func (s CreateTransitRouterPolicyTableRequestTag) GoString() string {
	return s.String()
}

func (s *CreateTransitRouterPolicyTableRequestTag) Validate() error {
	return dara.Validate(s)
}
