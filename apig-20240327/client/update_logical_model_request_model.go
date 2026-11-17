// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateLogicalModelRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCallStrategy(v string) *UpdateLogicalModelRequest
	GetCallStrategy() *string
	SetFallbackTargets(v []*UpdateLogicalModelRequestFallbackTargets) *UpdateLogicalModelRequest
	GetFallbackTargets() []*UpdateLogicalModelRequestFallbackTargets
	SetName(v string) *UpdateLogicalModelRequest
	GetName() *string
	SetResponseModel(v string) *UpdateLogicalModelRequest
	GetResponseModel() *string
	SetRoutingPolicy(v string) *UpdateLogicalModelRequest
	GetRoutingPolicy() *string
	SetSlots(v *UpdateLogicalModelRequestSlots) *UpdateLogicalModelRequest
	GetSlots() *UpdateLogicalModelRequestSlots
	SetTargets(v []*UpdateLogicalModelRequestTargets) *UpdateLogicalModelRequest
	GetTargets() []*UpdateLogicalModelRequestTargets
	SetConfirmationToken(v string) *UpdateLogicalModelRequest
	GetConfirmationToken() *string
	SetExpectedRevision(v int64) *UpdateLogicalModelRequest
	GetExpectedRevision() *int64
}

type UpdateLogicalModelRequest struct {
	// The call strategy. This parameter is required for new Auto write requests. Valid values: balanced, cheap, fast, and quality. The new version uses only the basic invocation name. Do not submit this parameter for Mapping and Weighted. The server performs conditional validation.
	//
	// example:
	//
	// balanced
	CallStrategy *string `json:"callStrategy,omitempty" xml:"callStrategy,omitempty"`
	// The ordered fallback targets. A maximum of 5 targets are supported. The targets are deduplicated by the actual card ID. This can be an empty array.
	//
	// This parameter is required.
	FallbackTargets []*UpdateLogicalModelRequestFallbackTargets `json:"fallbackTargets,omitempty" xml:"fallbackTargets,omitempty" type:"Repeated"`
	// The client logical invocation name. It must be encoded in UTF-8 and can be up to 256 bytes in length. It cannot contain spaces or slashes, and cannot start with auto.
	//
	// This parameter is required.
	//
	// example:
	//
	// beta-chat
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The response model. If set to Actual, the public name of the actual model is returned. If set to Logical, the entry name is returned, including Server-Sent Events (SSE).
	//
	// This parameter is required.
	//
	// example:
	//
	// Actual
	ResponseModel *string `json:"responseModel,omitempty" xml:"responseModel,omitempty"`
	// The routing policy. Valid values: Mapping, Weighted, and Auto. If you set this parameter to Auto, targets must be an empty array and slots is required.
	//
	// This parameter is required.
	//
	// example:
	//
	// Weighted
	RoutingPolicy *string `json:"routingPolicy,omitempty" xml:"routingPolicy,omitempty"`
	// The slots. This parameter is submitted only for Auto. Each of the simple, standard, and complex slots requires 1 to 3 source cards. The three slots are globally deduplicated by modelCardId. Omit this parameter for Mapping and Weighted.
	Slots *UpdateLogicalModelRequestSlots `json:"slots,omitempty" xml:"slots,omitempty" type:"Struct"`
	// The primary targets. For Mapping, exactly one item is allowed. For Weighted, the sum of original relative weights must be greater than 0. For Auto, this must be an empty array.
	//
	// This parameter is required.
	Targets []*UpdateLogicalModelRequestTargets `json:"targets,omitempty" xml:"targets,omitempty" type:"Repeated"`
	// The confirmation token returned by the configuration update dry run. Pass this token as is when you submit the update.
	//
	// example:
	//
	// 6b8e7d61b3a94c7285cf0d9ee1a72c836b8e7d61b3a94c7285cf0d9ee1a72c83
	ConfirmationToken *string `json:"confirmationToken,omitempty" xml:"confirmationToken,omitempty"`
	// The expected configuration revision number. Use the revision returned by the most recent query or dry run.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	ExpectedRevision *int64 `json:"expectedRevision,omitempty" xml:"expectedRevision,omitempty"`
}

func (s UpdateLogicalModelRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateLogicalModelRequest) GoString() string {
	return s.String()
}

func (s *UpdateLogicalModelRequest) GetCallStrategy() *string {
	return s.CallStrategy
}

func (s *UpdateLogicalModelRequest) GetFallbackTargets() []*UpdateLogicalModelRequestFallbackTargets {
	return s.FallbackTargets
}

func (s *UpdateLogicalModelRequest) GetName() *string {
	return s.Name
}

func (s *UpdateLogicalModelRequest) GetResponseModel() *string {
	return s.ResponseModel
}

func (s *UpdateLogicalModelRequest) GetRoutingPolicy() *string {
	return s.RoutingPolicy
}

func (s *UpdateLogicalModelRequest) GetSlots() *UpdateLogicalModelRequestSlots {
	return s.Slots
}

func (s *UpdateLogicalModelRequest) GetTargets() []*UpdateLogicalModelRequestTargets {
	return s.Targets
}

func (s *UpdateLogicalModelRequest) GetConfirmationToken() *string {
	return s.ConfirmationToken
}

func (s *UpdateLogicalModelRequest) GetExpectedRevision() *int64 {
	return s.ExpectedRevision
}

func (s *UpdateLogicalModelRequest) SetCallStrategy(v string) *UpdateLogicalModelRequest {
	s.CallStrategy = &v
	return s
}

func (s *UpdateLogicalModelRequest) SetFallbackTargets(v []*UpdateLogicalModelRequestFallbackTargets) *UpdateLogicalModelRequest {
	s.FallbackTargets = v
	return s
}

func (s *UpdateLogicalModelRequest) SetName(v string) *UpdateLogicalModelRequest {
	s.Name = &v
	return s
}

func (s *UpdateLogicalModelRequest) SetResponseModel(v string) *UpdateLogicalModelRequest {
	s.ResponseModel = &v
	return s
}

func (s *UpdateLogicalModelRequest) SetRoutingPolicy(v string) *UpdateLogicalModelRequest {
	s.RoutingPolicy = &v
	return s
}

func (s *UpdateLogicalModelRequest) SetSlots(v *UpdateLogicalModelRequestSlots) *UpdateLogicalModelRequest {
	s.Slots = v
	return s
}

func (s *UpdateLogicalModelRequest) SetTargets(v []*UpdateLogicalModelRequestTargets) *UpdateLogicalModelRequest {
	s.Targets = v
	return s
}

func (s *UpdateLogicalModelRequest) SetConfirmationToken(v string) *UpdateLogicalModelRequest {
	s.ConfirmationToken = &v
	return s
}

func (s *UpdateLogicalModelRequest) SetExpectedRevision(v int64) *UpdateLogicalModelRequest {
	s.ExpectedRevision = &v
	return s
}

func (s *UpdateLogicalModelRequest) Validate() error {
	if s.FallbackTargets != nil {
		for _, item := range s.FallbackTargets {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Slots != nil {
		if err := s.Slots.Validate(); err != nil {
			return err
		}
	}
	if s.Targets != nil {
		for _, item := range s.Targets {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type UpdateLogicalModelRequestFallbackTargets struct {
	// The ID of the actual source card in the same instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// msc-d00000000000000000000
	ModelCardId *string `json:"modelCardId,omitempty" xml:"modelCardId,omitempty"`
}

func (s UpdateLogicalModelRequestFallbackTargets) String() string {
	return dara.Prettify(s)
}

func (s UpdateLogicalModelRequestFallbackTargets) GoString() string {
	return s.String()
}

func (s *UpdateLogicalModelRequestFallbackTargets) GetModelCardId() *string {
	return s.ModelCardId
}

func (s *UpdateLogicalModelRequestFallbackTargets) SetModelCardId(v string) *UpdateLogicalModelRequestFallbackTargets {
	s.ModelCardId = &v
	return s
}

func (s *UpdateLogicalModelRequestFallbackTargets) Validate() error {
	return dara.Validate(s)
}

type UpdateLogicalModelRequestSlots struct {
	// The complex slot. It requires 1 to 3 source cards. The three slots are globally deduplicated by modelCardId.
	//
	// This parameter is required.
	Complex []*UpdateLogicalModelRequestSlotsComplex `json:"complex,omitempty" xml:"complex,omitempty" type:"Repeated"`
	// The simple slot. It requires 1 to 3 source cards. The three slots are globally deduplicated by modelCardId.
	//
	// This parameter is required.
	Simple []*UpdateLogicalModelRequestSlotsSimple `json:"simple,omitempty" xml:"simple,omitempty" type:"Repeated"`
	// The standard slot. It requires 1 to 3 source cards. The three slots are globally deduplicated by modelCardId.
	//
	// This parameter is required.
	Standard []*UpdateLogicalModelRequestSlotsStandard `json:"standard,omitempty" xml:"standard,omitempty" type:"Repeated"`
}

func (s UpdateLogicalModelRequestSlots) String() string {
	return dara.Prettify(s)
}

func (s UpdateLogicalModelRequestSlots) GoString() string {
	return s.String()
}

func (s *UpdateLogicalModelRequestSlots) GetComplex() []*UpdateLogicalModelRequestSlotsComplex {
	return s.Complex
}

func (s *UpdateLogicalModelRequestSlots) GetSimple() []*UpdateLogicalModelRequestSlotsSimple {
	return s.Simple
}

func (s *UpdateLogicalModelRequestSlots) GetStandard() []*UpdateLogicalModelRequestSlotsStandard {
	return s.Standard
}

func (s *UpdateLogicalModelRequestSlots) SetComplex(v []*UpdateLogicalModelRequestSlotsComplex) *UpdateLogicalModelRequestSlots {
	s.Complex = v
	return s
}

func (s *UpdateLogicalModelRequestSlots) SetSimple(v []*UpdateLogicalModelRequestSlotsSimple) *UpdateLogicalModelRequestSlots {
	s.Simple = v
	return s
}

func (s *UpdateLogicalModelRequestSlots) SetStandard(v []*UpdateLogicalModelRequestSlotsStandard) *UpdateLogicalModelRequestSlots {
	s.Standard = v
	return s
}

func (s *UpdateLogicalModelRequestSlots) Validate() error {
	if s.Complex != nil {
		for _, item := range s.Complex {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Simple != nil {
		for _, item := range s.Simple {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Standard != nil {
		for _, item := range s.Standard {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type UpdateLogicalModelRequestSlotsComplex struct {
	// The ID of the actual source card in the same instance. The ID starts with msc-. VMs or service overrides are not supported.
	//
	// This parameter is required.
	//
	// example:
	//
	// msc-d00000000000000000000
	ModelCardId *string `json:"modelCardId,omitempty" xml:"modelCardId,omitempty"`
}

func (s UpdateLogicalModelRequestSlotsComplex) String() string {
	return dara.Prettify(s)
}

func (s UpdateLogicalModelRequestSlotsComplex) GoString() string {
	return s.String()
}

func (s *UpdateLogicalModelRequestSlotsComplex) GetModelCardId() *string {
	return s.ModelCardId
}

func (s *UpdateLogicalModelRequestSlotsComplex) SetModelCardId(v string) *UpdateLogicalModelRequestSlotsComplex {
	s.ModelCardId = &v
	return s
}

func (s *UpdateLogicalModelRequestSlotsComplex) Validate() error {
	return dara.Validate(s)
}

type UpdateLogicalModelRequestSlotsSimple struct {
	// The ID of the actual source card in the same instance. The ID starts with msc-. VMs or service overrides are not supported.
	//
	// This parameter is required.
	//
	// example:
	//
	// msc-d00000000000000000000
	ModelCardId *string `json:"modelCardId,omitempty" xml:"modelCardId,omitempty"`
}

func (s UpdateLogicalModelRequestSlotsSimple) String() string {
	return dara.Prettify(s)
}

func (s UpdateLogicalModelRequestSlotsSimple) GoString() string {
	return s.String()
}

func (s *UpdateLogicalModelRequestSlotsSimple) GetModelCardId() *string {
	return s.ModelCardId
}

func (s *UpdateLogicalModelRequestSlotsSimple) SetModelCardId(v string) *UpdateLogicalModelRequestSlotsSimple {
	s.ModelCardId = &v
	return s
}

func (s *UpdateLogicalModelRequestSlotsSimple) Validate() error {
	return dara.Validate(s)
}

type UpdateLogicalModelRequestSlotsStandard struct {
	// The ID of the actual source card in the same instance. The ID starts with msc-. VMs or service overrides are not supported.
	//
	// This parameter is required.
	//
	// example:
	//
	// msc-d00000000000000000000
	ModelCardId *string `json:"modelCardId,omitempty" xml:"modelCardId,omitempty"`
}

func (s UpdateLogicalModelRequestSlotsStandard) String() string {
	return dara.Prettify(s)
}

func (s UpdateLogicalModelRequestSlotsStandard) GoString() string {
	return s.String()
}

func (s *UpdateLogicalModelRequestSlotsStandard) GetModelCardId() *string {
	return s.ModelCardId
}

func (s *UpdateLogicalModelRequestSlotsStandard) SetModelCardId(v string) *UpdateLogicalModelRequestSlotsStandard {
	s.ModelCardId = &v
	return s
}

func (s *UpdateLogicalModelRequestSlotsStandard) Validate() error {
	return dara.Validate(s)
}

type UpdateLogicalModelRequestTargets struct {
	// The ID of the actual source model card in the same instance. The ID starts with msc-. VMs or service overrides are not supported.
	//
	// This parameter is required.
	//
	// example:
	//
	// msc-d00000000000000000000
	ModelCardId *string `json:"modelCardId,omitempty" xml:"modelCardId,omitempty"`
	// The weight. Valid values: integers from 0 to 10000. The original value is preserved. For Mapping, 0 is allowed.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	Weight *int32 `json:"weight,omitempty" xml:"weight,omitempty"`
}

func (s UpdateLogicalModelRequestTargets) String() string {
	return dara.Prettify(s)
}

func (s UpdateLogicalModelRequestTargets) GoString() string {
	return s.String()
}

func (s *UpdateLogicalModelRequestTargets) GetModelCardId() *string {
	return s.ModelCardId
}

func (s *UpdateLogicalModelRequestTargets) GetWeight() *int32 {
	return s.Weight
}

func (s *UpdateLogicalModelRequestTargets) SetModelCardId(v string) *UpdateLogicalModelRequestTargets {
	s.ModelCardId = &v
	return s
}

func (s *UpdateLogicalModelRequestTargets) SetWeight(v int32) *UpdateLogicalModelRequestTargets {
	s.Weight = &v
	return s
}

func (s *UpdateLogicalModelRequestTargets) Validate() error {
	return dara.Validate(s)
}
