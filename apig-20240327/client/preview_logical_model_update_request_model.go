// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPreviewLogicalModelUpdateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCallStrategy(v string) *PreviewLogicalModelUpdateRequest
	GetCallStrategy() *string
	SetFallbackTargets(v []*PreviewLogicalModelUpdateRequestFallbackTargets) *PreviewLogicalModelUpdateRequest
	GetFallbackTargets() []*PreviewLogicalModelUpdateRequestFallbackTargets
	SetName(v string) *PreviewLogicalModelUpdateRequest
	GetName() *string
	SetResponseModel(v string) *PreviewLogicalModelUpdateRequest
	GetResponseModel() *string
	SetRoutingPolicy(v string) *PreviewLogicalModelUpdateRequest
	GetRoutingPolicy() *string
	SetSlots(v *PreviewLogicalModelUpdateRequestSlots) *PreviewLogicalModelUpdateRequest
	GetSlots() *PreviewLogicalModelUpdateRequestSlots
	SetTargets(v []*PreviewLogicalModelUpdateRequestTargets) *PreviewLogicalModelUpdateRequest
	GetTargets() []*PreviewLogicalModelUpdateRequestTargets
}

type PreviewLogicalModelUpdateRequest struct {
	// The call strategy, which is required for new Auto write requests. Valid values: balanced, cheap, fast, and quality. The new version uses only the basic call name. Do not submit this parameter for Mapping or Weighted. The server performs conditional validation.
	//
	// example:
	//
	// balanced
	CallStrategy *string `json:"callStrategy,omitempty" xml:"callStrategy,omitempty"`
	// An ordered list of up to 5 fallback targets. The targets are deduplicated by the actual card ID. This can be an empty array.
	//
	// This parameter is required.
	FallbackTargets []*PreviewLogicalModelUpdateRequestFallbackTargets `json:"fallbackTargets,omitempty" xml:"fallbackTargets,omitempty" type:"Repeated"`
	// The client logical call name. The value can be up to 256 bytes in UTF-8 encoding, must not contain spaces or slashes, and must not start with auto.
	//
	// This parameter is required.
	//
	// example:
	//
	// beta-chat
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The response model type. If set to Actual, the public name of the actual model is returned. If set to Logical, the entry name is returned, including SSE.
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
	// The slots configuration, which is submitted only for Auto. Each of the simple, standard, and complex slots requires 1 to 3 source cards. Cards across the three slots are globally deduplicated by modelCardId. Omit this parameter for Mapping or Weighted.
	Slots *PreviewLogicalModelUpdateRequestSlots `json:"slots,omitempty" xml:"slots,omitempty" type:"Struct"`
	// The primary targets. For Mapping, exactly one item is required. For Weighted, the sum of the original relative weights must be greater than 0. For Auto, this must be an empty array.
	//
	// This parameter is required.
	Targets []*PreviewLogicalModelUpdateRequestTargets `json:"targets,omitempty" xml:"targets,omitempty" type:"Repeated"`
}

func (s PreviewLogicalModelUpdateRequest) String() string {
	return dara.Prettify(s)
}

func (s PreviewLogicalModelUpdateRequest) GoString() string {
	return s.String()
}

func (s *PreviewLogicalModelUpdateRequest) GetCallStrategy() *string {
	return s.CallStrategy
}

func (s *PreviewLogicalModelUpdateRequest) GetFallbackTargets() []*PreviewLogicalModelUpdateRequestFallbackTargets {
	return s.FallbackTargets
}

func (s *PreviewLogicalModelUpdateRequest) GetName() *string {
	return s.Name
}

func (s *PreviewLogicalModelUpdateRequest) GetResponseModel() *string {
	return s.ResponseModel
}

func (s *PreviewLogicalModelUpdateRequest) GetRoutingPolicy() *string {
	return s.RoutingPolicy
}

func (s *PreviewLogicalModelUpdateRequest) GetSlots() *PreviewLogicalModelUpdateRequestSlots {
	return s.Slots
}

func (s *PreviewLogicalModelUpdateRequest) GetTargets() []*PreviewLogicalModelUpdateRequestTargets {
	return s.Targets
}

func (s *PreviewLogicalModelUpdateRequest) SetCallStrategy(v string) *PreviewLogicalModelUpdateRequest {
	s.CallStrategy = &v
	return s
}

func (s *PreviewLogicalModelUpdateRequest) SetFallbackTargets(v []*PreviewLogicalModelUpdateRequestFallbackTargets) *PreviewLogicalModelUpdateRequest {
	s.FallbackTargets = v
	return s
}

func (s *PreviewLogicalModelUpdateRequest) SetName(v string) *PreviewLogicalModelUpdateRequest {
	s.Name = &v
	return s
}

func (s *PreviewLogicalModelUpdateRequest) SetResponseModel(v string) *PreviewLogicalModelUpdateRequest {
	s.ResponseModel = &v
	return s
}

func (s *PreviewLogicalModelUpdateRequest) SetRoutingPolicy(v string) *PreviewLogicalModelUpdateRequest {
	s.RoutingPolicy = &v
	return s
}

func (s *PreviewLogicalModelUpdateRequest) SetSlots(v *PreviewLogicalModelUpdateRequestSlots) *PreviewLogicalModelUpdateRequest {
	s.Slots = v
	return s
}

func (s *PreviewLogicalModelUpdateRequest) SetTargets(v []*PreviewLogicalModelUpdateRequestTargets) *PreviewLogicalModelUpdateRequest {
	s.Targets = v
	return s
}

func (s *PreviewLogicalModelUpdateRequest) Validate() error {
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

type PreviewLogicalModelUpdateRequestFallbackTargets struct {
	// The actual source card ID within the same instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// msc-d00000000000000000000
	ModelCardId *string `json:"modelCardId,omitempty" xml:"modelCardId,omitempty"`
}

func (s PreviewLogicalModelUpdateRequestFallbackTargets) String() string {
	return dara.Prettify(s)
}

func (s PreviewLogicalModelUpdateRequestFallbackTargets) GoString() string {
	return s.String()
}

func (s *PreviewLogicalModelUpdateRequestFallbackTargets) GetModelCardId() *string {
	return s.ModelCardId
}

func (s *PreviewLogicalModelUpdateRequestFallbackTargets) SetModelCardId(v string) *PreviewLogicalModelUpdateRequestFallbackTargets {
	s.ModelCardId = &v
	return s
}

func (s *PreviewLogicalModelUpdateRequestFallbackTargets) Validate() error {
	return dara.Validate(s)
}

type PreviewLogicalModelUpdateRequestSlots struct {
	// The complex slot, which requires 1 to 3 source cards. Cards across the three slots are globally deduplicated by modelCardId.
	//
	// This parameter is required.
	Complex []*PreviewLogicalModelUpdateRequestSlotsComplex `json:"complex,omitempty" xml:"complex,omitempty" type:"Repeated"`
	// The simple slot, which requires 1 to 3 source cards. Cards across the three slots are globally deduplicated by modelCardId.
	//
	// This parameter is required.
	Simple []*PreviewLogicalModelUpdateRequestSlotsSimple `json:"simple,omitempty" xml:"simple,omitempty" type:"Repeated"`
	// The standard slot, which requires 1 to 3 source cards. Cards across the three slots are globally deduplicated by modelCardId.
	//
	// This parameter is required.
	Standard []*PreviewLogicalModelUpdateRequestSlotsStandard `json:"standard,omitempty" xml:"standard,omitempty" type:"Repeated"`
}

func (s PreviewLogicalModelUpdateRequestSlots) String() string {
	return dara.Prettify(s)
}

func (s PreviewLogicalModelUpdateRequestSlots) GoString() string {
	return s.String()
}

func (s *PreviewLogicalModelUpdateRequestSlots) GetComplex() []*PreviewLogicalModelUpdateRequestSlotsComplex {
	return s.Complex
}

func (s *PreviewLogicalModelUpdateRequestSlots) GetSimple() []*PreviewLogicalModelUpdateRequestSlotsSimple {
	return s.Simple
}

func (s *PreviewLogicalModelUpdateRequestSlots) GetStandard() []*PreviewLogicalModelUpdateRequestSlotsStandard {
	return s.Standard
}

func (s *PreviewLogicalModelUpdateRequestSlots) SetComplex(v []*PreviewLogicalModelUpdateRequestSlotsComplex) *PreviewLogicalModelUpdateRequestSlots {
	s.Complex = v
	return s
}

func (s *PreviewLogicalModelUpdateRequestSlots) SetSimple(v []*PreviewLogicalModelUpdateRequestSlotsSimple) *PreviewLogicalModelUpdateRequestSlots {
	s.Simple = v
	return s
}

func (s *PreviewLogicalModelUpdateRequestSlots) SetStandard(v []*PreviewLogicalModelUpdateRequestSlotsStandard) *PreviewLogicalModelUpdateRequestSlots {
	s.Standard = v
	return s
}

func (s *PreviewLogicalModelUpdateRequestSlots) Validate() error {
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

type PreviewLogicalModelUpdateRequestSlotsComplex struct {
	// The actual source card ID within the same instance, which starts with msc-. VM or service overrides are not supported.
	//
	// This parameter is required.
	//
	// example:
	//
	// msc-d00000000000000000000
	ModelCardId *string `json:"modelCardId,omitempty" xml:"modelCardId,omitempty"`
}

func (s PreviewLogicalModelUpdateRequestSlotsComplex) String() string {
	return dara.Prettify(s)
}

func (s PreviewLogicalModelUpdateRequestSlotsComplex) GoString() string {
	return s.String()
}

func (s *PreviewLogicalModelUpdateRequestSlotsComplex) GetModelCardId() *string {
	return s.ModelCardId
}

func (s *PreviewLogicalModelUpdateRequestSlotsComplex) SetModelCardId(v string) *PreviewLogicalModelUpdateRequestSlotsComplex {
	s.ModelCardId = &v
	return s
}

func (s *PreviewLogicalModelUpdateRequestSlotsComplex) Validate() error {
	return dara.Validate(s)
}

type PreviewLogicalModelUpdateRequestSlotsSimple struct {
	// The actual source card ID within the same instance, which starts with msc-. VM or service overrides are not supported.
	//
	// This parameter is required.
	//
	// example:
	//
	// msc-d00000000000000000000
	ModelCardId *string `json:"modelCardId,omitempty" xml:"modelCardId,omitempty"`
}

func (s PreviewLogicalModelUpdateRequestSlotsSimple) String() string {
	return dara.Prettify(s)
}

func (s PreviewLogicalModelUpdateRequestSlotsSimple) GoString() string {
	return s.String()
}

func (s *PreviewLogicalModelUpdateRequestSlotsSimple) GetModelCardId() *string {
	return s.ModelCardId
}

func (s *PreviewLogicalModelUpdateRequestSlotsSimple) SetModelCardId(v string) *PreviewLogicalModelUpdateRequestSlotsSimple {
	s.ModelCardId = &v
	return s
}

func (s *PreviewLogicalModelUpdateRequestSlotsSimple) Validate() error {
	return dara.Validate(s)
}

type PreviewLogicalModelUpdateRequestSlotsStandard struct {
	// The actual source card ID within the same instance, which starts with msc-. VM or service overrides are not supported.
	//
	// This parameter is required.
	//
	// example:
	//
	// msc-d00000000000000000000
	ModelCardId *string `json:"modelCardId,omitempty" xml:"modelCardId,omitempty"`
}

func (s PreviewLogicalModelUpdateRequestSlotsStandard) String() string {
	return dara.Prettify(s)
}

func (s PreviewLogicalModelUpdateRequestSlotsStandard) GoString() string {
	return s.String()
}

func (s *PreviewLogicalModelUpdateRequestSlotsStandard) GetModelCardId() *string {
	return s.ModelCardId
}

func (s *PreviewLogicalModelUpdateRequestSlotsStandard) SetModelCardId(v string) *PreviewLogicalModelUpdateRequestSlotsStandard {
	s.ModelCardId = &v
	return s
}

func (s *PreviewLogicalModelUpdateRequestSlotsStandard) Validate() error {
	return dara.Validate(s)
}

type PreviewLogicalModelUpdateRequestTargets struct {
	// The actual source model card ID within the same instance, which starts with msc-. VM or service overrides are not supported.
	//
	// This parameter is required.
	//
	// example:
	//
	// msc-d00000000000000000000
	ModelCardId *string `json:"modelCardId,omitempty" xml:"modelCardId,omitempty"`
	// The weight value. Valid values: integers from 0 to 10000. The original value is preserved. For Mapping, 0 is allowed.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	Weight *int32 `json:"weight,omitempty" xml:"weight,omitempty"`
}

func (s PreviewLogicalModelUpdateRequestTargets) String() string {
	return dara.Prettify(s)
}

func (s PreviewLogicalModelUpdateRequestTargets) GoString() string {
	return s.String()
}

func (s *PreviewLogicalModelUpdateRequestTargets) GetModelCardId() *string {
	return s.ModelCardId
}

func (s *PreviewLogicalModelUpdateRequestTargets) GetWeight() *int32 {
	return s.Weight
}

func (s *PreviewLogicalModelUpdateRequestTargets) SetModelCardId(v string) *PreviewLogicalModelUpdateRequestTargets {
	s.ModelCardId = &v
	return s
}

func (s *PreviewLogicalModelUpdateRequestTargets) SetWeight(v int32) *PreviewLogicalModelUpdateRequestTargets {
	s.Weight = &v
	return s
}

func (s *PreviewLogicalModelUpdateRequestTargets) Validate() error {
	return dara.Validate(s)
}
