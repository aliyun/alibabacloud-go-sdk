// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPublishKgSchemaRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOpTenantId(v int64) *PublishKgSchemaRequest
	GetOpTenantId() *int64
	SetOpUserId(v string) *PublishKgSchemaRequest
	GetOpUserId() *string
	SetPublishCommand(v *PublishKgSchemaRequestPublishCommand) *PublishKgSchemaRequest
	GetPublishCommand() *PublishKgSchemaRequestPublishCommand
	SetWorkspaceId(v string) *PublishKgSchemaRequest
	GetWorkspaceId() *string
}

type PublishKgSchemaRequest struct {
	// The tenant ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
	// example:
	//
	// 30001011
	OpUserId *string `json:"OpUserId,omitempty" xml:"OpUserId,omitempty"`
	// The publish command and its details.
	//
	// This parameter is required.
	PublishCommand *PublishKgSchemaRequestPublishCommand `json:"PublishCommand,omitempty" xml:"PublishCommand,omitempty" type:"Struct"`
	// The model ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// f1d4559a4db044158305e2d89bccf81f
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s PublishKgSchemaRequest) String() string {
	return dara.Prettify(s)
}

func (s PublishKgSchemaRequest) GoString() string {
	return s.String()
}

func (s *PublishKgSchemaRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *PublishKgSchemaRequest) GetOpUserId() *string {
	return s.OpUserId
}

func (s *PublishKgSchemaRequest) GetPublishCommand() *PublishKgSchemaRequestPublishCommand {
	return s.PublishCommand
}

func (s *PublishKgSchemaRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *PublishKgSchemaRequest) SetOpTenantId(v int64) *PublishKgSchemaRequest {
	s.OpTenantId = &v
	return s
}

func (s *PublishKgSchemaRequest) SetOpUserId(v string) *PublishKgSchemaRequest {
	s.OpUserId = &v
	return s
}

func (s *PublishKgSchemaRequest) SetPublishCommand(v *PublishKgSchemaRequestPublishCommand) *PublishKgSchemaRequest {
	s.PublishCommand = v
	return s
}

func (s *PublishKgSchemaRequest) SetWorkspaceId(v string) *PublishKgSchemaRequest {
	s.WorkspaceId = &v
	return s
}

func (s *PublishKgSchemaRequest) Validate() error {
	if s.PublishCommand != nil {
		if err := s.PublishCommand.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type PublishKgSchemaRequestPublishCommand struct {
	// The data adjustment policies.
	DataAdjustmentPolicies []*PublishKgSchemaRequestPublishCommandDataAdjustmentPolicies `json:"DataAdjustmentPolicies,omitempty" xml:"DataAdjustmentPolicies,omitempty" type:"Repeated"`
	// The description.
	//
	// This parameter is required.
	//
	// example:
	//
	// test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
}

func (s PublishKgSchemaRequestPublishCommand) String() string {
	return dara.Prettify(s)
}

func (s PublishKgSchemaRequestPublishCommand) GoString() string {
	return s.String()
}

func (s *PublishKgSchemaRequestPublishCommand) GetDataAdjustmentPolicies() []*PublishKgSchemaRequestPublishCommandDataAdjustmentPolicies {
	return s.DataAdjustmentPolicies
}

func (s *PublishKgSchemaRequestPublishCommand) GetDescription() *string {
	return s.Description
}

func (s *PublishKgSchemaRequestPublishCommand) SetDataAdjustmentPolicies(v []*PublishKgSchemaRequestPublishCommandDataAdjustmentPolicies) *PublishKgSchemaRequestPublishCommand {
	s.DataAdjustmentPolicies = v
	return s
}

func (s *PublishKgSchemaRequestPublishCommand) SetDescription(v string) *PublishKgSchemaRequestPublishCommand {
	s.Description = &v
	return s
}

func (s *PublishKgSchemaRequestPublishCommand) Validate() error {
	if s.DataAdjustmentPolicies != nil {
		for _, item := range s.DataAdjustmentPolicies {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type PublishKgSchemaRequestPublishCommandDataAdjustmentPolicies struct {
	// The backfill property default value policy. This parameter takes effect only when PolicyType is set to BackFillDefault.
	//
	// This parameter is required.
	BackFillDefaultValuePolicy *PublishKgSchemaRequestPublishCommandDataAdjustmentPoliciesBackFillDefaultValuePolicy `json:"BackFillDefaultValuePolicy,omitempty" xml:"BackFillDefaultValuePolicy,omitempty" type:"Struct"`
	// The policy type. Valid values:
	//
	// - BackFillDefault: backfills default values when a property changes from optional to required.
	//
	// This parameter is required.
	//
	// example:
	//
	// BackFillDefault
	PolicyType *string `json:"PolicyType,omitempty" xml:"PolicyType,omitempty"`
	// The type to which the policy applies. Valid values:
	//
	// - ENTITY: applies to entity types.
	//
	// - RELATION: applies to relation types.
	//
	// This parameter is required.
	//
	// example:
	//
	// ENTITY
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The code of the entity type or relation type.
	//
	// This parameter is required.
	//
	// example:
	//
	// Product
	TypeCode *string `json:"TypeCode,omitempty" xml:"TypeCode,omitempty"`
}

func (s PublishKgSchemaRequestPublishCommandDataAdjustmentPolicies) String() string {
	return dara.Prettify(s)
}

func (s PublishKgSchemaRequestPublishCommandDataAdjustmentPolicies) GoString() string {
	return s.String()
}

func (s *PublishKgSchemaRequestPublishCommandDataAdjustmentPolicies) GetBackFillDefaultValuePolicy() *PublishKgSchemaRequestPublishCommandDataAdjustmentPoliciesBackFillDefaultValuePolicy {
	return s.BackFillDefaultValuePolicy
}

func (s *PublishKgSchemaRequestPublishCommandDataAdjustmentPolicies) GetPolicyType() *string {
	return s.PolicyType
}

func (s *PublishKgSchemaRequestPublishCommandDataAdjustmentPolicies) GetType() *string {
	return s.Type
}

func (s *PublishKgSchemaRequestPublishCommandDataAdjustmentPolicies) GetTypeCode() *string {
	return s.TypeCode
}

func (s *PublishKgSchemaRequestPublishCommandDataAdjustmentPolicies) SetBackFillDefaultValuePolicy(v *PublishKgSchemaRequestPublishCommandDataAdjustmentPoliciesBackFillDefaultValuePolicy) *PublishKgSchemaRequestPublishCommandDataAdjustmentPolicies {
	s.BackFillDefaultValuePolicy = v
	return s
}

func (s *PublishKgSchemaRequestPublishCommandDataAdjustmentPolicies) SetPolicyType(v string) *PublishKgSchemaRequestPublishCommandDataAdjustmentPolicies {
	s.PolicyType = &v
	return s
}

func (s *PublishKgSchemaRequestPublishCommandDataAdjustmentPolicies) SetType(v string) *PublishKgSchemaRequestPublishCommandDataAdjustmentPolicies {
	s.Type = &v
	return s
}

func (s *PublishKgSchemaRequestPublishCommandDataAdjustmentPolicies) SetTypeCode(v string) *PublishKgSchemaRequestPublishCommandDataAdjustmentPolicies {
	s.TypeCode = &v
	return s
}

func (s *PublishKgSchemaRequestPublishCommandDataAdjustmentPolicies) Validate() error {
	if s.BackFillDefaultValuePolicy != nil {
		if err := s.BackFillDefaultValuePolicy.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type PublishKgSchemaRequestPublishCommandDataAdjustmentPoliciesBackFillDefaultValuePolicy struct {
	// The default value to backfill for the property.
	//
	// This parameter is required.
	//
	// example:
	//
	// 123
	DefaultValue *string `json:"DefaultValue,omitempty" xml:"DefaultValue,omitempty"`
	// The property code.
	//
	// This parameter is required.
	//
	// example:
	//
	// name
	PropertyCode *string `json:"PropertyCode,omitempty" xml:"PropertyCode,omitempty"`
}

func (s PublishKgSchemaRequestPublishCommandDataAdjustmentPoliciesBackFillDefaultValuePolicy) String() string {
	return dara.Prettify(s)
}

func (s PublishKgSchemaRequestPublishCommandDataAdjustmentPoliciesBackFillDefaultValuePolicy) GoString() string {
	return s.String()
}

func (s *PublishKgSchemaRequestPublishCommandDataAdjustmentPoliciesBackFillDefaultValuePolicy) GetDefaultValue() *string {
	return s.DefaultValue
}

func (s *PublishKgSchemaRequestPublishCommandDataAdjustmentPoliciesBackFillDefaultValuePolicy) GetPropertyCode() *string {
	return s.PropertyCode
}

func (s *PublishKgSchemaRequestPublishCommandDataAdjustmentPoliciesBackFillDefaultValuePolicy) SetDefaultValue(v string) *PublishKgSchemaRequestPublishCommandDataAdjustmentPoliciesBackFillDefaultValuePolicy {
	s.DefaultValue = &v
	return s
}

func (s *PublishKgSchemaRequestPublishCommandDataAdjustmentPoliciesBackFillDefaultValuePolicy) SetPropertyCode(v string) *PublishKgSchemaRequestPublishCommandDataAdjustmentPoliciesBackFillDefaultValuePolicy {
	s.PropertyCode = &v
	return s
}

func (s *PublishKgSchemaRequestPublishCommandDataAdjustmentPoliciesBackFillDefaultValuePolicy) Validate() error {
	return dara.Validate(s)
}
