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
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can generate the parameter value on your client to ensure that the value is unique among different requests. The client token can contain only ASCII characters.
	//
	// > If you do not specify this parameter, the system automatically uses the **RequestId*	- of the API request as the **ClientToken**. The **RequestId*	- of each API request is unique.
	//
	// example:
	//
	// 02fb3da4-130e-11e9-8e44-001****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The description of the policy route table.
	//
	// example:
	//
	// descriptiontest
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Specifies whether to perform a dry run on the creation request, including permission and instance status checks. Valid values:
	//
	// - **false*	- (default): sends a normal request. The system performs a check and creates the policy route table if the check passes.
	//
	// - **true**: sends a dry run request. The system only performs a check and does not create the policy route table. Check items include whether required parameters are specified and whether the request format is valid. If the check fails, the corresponding error is returned. If the check passes, the corresponding request ID is returned.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The name of the policy route table.
	//
	// example:
	//
	// nametest
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The list of tags bound to the policy route table.
	Tag []*CreateTransitRouterPolicyTableRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// The ID of the Enterprise Edition transit router.
	//
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
	// example:
	//
	// tagtest
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// example:
	//
	// tagtest
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateTransitRouterPolicyTableRequestTag) String() string {
	return dara.Prettify(s)
}

func (s CreateTransitRouterPolicyTableRequestTag) GoString() string {
	return s.String()
}

func (s *CreateTransitRouterPolicyTableRequestTag) GetKey() *string {
	return s.Key
}

func (s *CreateTransitRouterPolicyTableRequestTag) GetValue() *string {
	return s.Value
}

func (s *CreateTransitRouterPolicyTableRequestTag) SetKey(v string) *CreateTransitRouterPolicyTableRequestTag {
	s.Key = &v
	return s
}

func (s *CreateTransitRouterPolicyTableRequestTag) SetValue(v string) *CreateTransitRouterPolicyTableRequestTag {
	s.Value = &v
	return s
}

func (s *CreateTransitRouterPolicyTableRequestTag) Validate() error {
	return dara.Validate(s)
}
