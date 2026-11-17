// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModelSelection interface {
	dara.Model
	String() string
	GoString() string
	SetConvenienceAliasEnabled(v bool) *ModelSelection
	GetConvenienceAliasEnabled() *bool
	SetEndpoints(v []*string) *ModelSelection
	GetEndpoints() []*string
	SetFallbackChains(v []*ModelSelectionFallbackChains) *ModelSelection
	GetFallbackChains() []*ModelSelectionFallbackChains
	SetModelGroupIds(v []*string) *ModelSelection
	GetModelGroupIds() []*string
	SetModels(v []*ModelSelectionModels) *ModelSelection
	GetModels() []*ModelSelectionModels
}

type ModelSelection struct {
	// Specifies whether to enable convenience aliases. When enabled, all selected model names must be unique.
	//
	// example:
	//
	// false
	ConvenienceAliasEnabled *bool `json:"convenienceAliasEnabled,omitempty" xml:"convenienceAliasEnabled,omitempty"`
	// The collection of explicitly selected built-in endpoint (route) names. This parameter is required and cannot be empty.
	Endpoints []*string `json:"endpoints,omitempty" xml:"endpoints,omitempty" type:"Repeated"`
	// The ordered fallback chains configured by endpoint and initial model.
	FallbackChains []*ModelSelectionFallbackChains `json:"fallbackChains,omitempty" xml:"fallbackChains,omitempty" type:"Repeated"`
	// The selected model group IDs within the instance. Group identity is retained, the union of current members and single-selected models is calculated, and empty groups can be saved as drafts.
	ModelGroupIds []*string `json:"modelGroupIds,omitempty" xml:"modelGroupIds,omitempty" type:"Repeated"`
	// The selected model collection.
	Models []*ModelSelectionModels `json:"models,omitempty" xml:"models,omitempty" type:"Repeated"`
}

func (s ModelSelection) String() string {
	return dara.Prettify(s)
}

func (s ModelSelection) GoString() string {
	return s.String()
}

func (s *ModelSelection) GetConvenienceAliasEnabled() *bool {
	return s.ConvenienceAliasEnabled
}

func (s *ModelSelection) GetEndpoints() []*string {
	return s.Endpoints
}

func (s *ModelSelection) GetFallbackChains() []*ModelSelectionFallbackChains {
	return s.FallbackChains
}

func (s *ModelSelection) GetModelGroupIds() []*string {
	return s.ModelGroupIds
}

func (s *ModelSelection) GetModels() []*ModelSelectionModels {
	return s.Models
}

func (s *ModelSelection) SetConvenienceAliasEnabled(v bool) *ModelSelection {
	s.ConvenienceAliasEnabled = &v
	return s
}

func (s *ModelSelection) SetEndpoints(v []*string) *ModelSelection {
	s.Endpoints = v
	return s
}

func (s *ModelSelection) SetFallbackChains(v []*ModelSelectionFallbackChains) *ModelSelection {
	s.FallbackChains = v
	return s
}

func (s *ModelSelection) SetModelGroupIds(v []*string) *ModelSelection {
	s.ModelGroupIds = v
	return s
}

func (s *ModelSelection) SetModels(v []*ModelSelectionModels) *ModelSelection {
	s.Models = v
	return s
}

func (s *ModelSelection) Validate() error {
	if s.FallbackChains != nil {
		for _, item := range s.FallbackChains {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Models != nil {
		for _, item := range s.Models {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ModelSelectionFallbackChains struct {
	// The endpoint name to which the chain belongs. It must be one of the currently selected endpoints.
	//
	// example:
	//
	// create-chat-completion
	Endpoint *string `json:"endpoint,omitempty" xml:"endpoint,omitempty"`
	// The ordered collection of fallback model card IDs.
	FallbackModelCardIds []*string `json:"fallbackModelCardIds,omitempty" xml:"fallbackModelCardIds,omitempty" type:"Repeated"`
	// The source model card ID.
	//
	// example:
	//
	// msc-d00000000000000000000
	SourceModelCardId *string `json:"sourceModelCardId,omitempty" xml:"sourceModelCardId,omitempty"`
}

func (s ModelSelectionFallbackChains) String() string {
	return dara.Prettify(s)
}

func (s ModelSelectionFallbackChains) GoString() string {
	return s.String()
}

func (s *ModelSelectionFallbackChains) GetEndpoint() *string {
	return s.Endpoint
}

func (s *ModelSelectionFallbackChains) GetFallbackModelCardIds() []*string {
	return s.FallbackModelCardIds
}

func (s *ModelSelectionFallbackChains) GetSourceModelCardId() *string {
	return s.SourceModelCardId
}

func (s *ModelSelectionFallbackChains) SetEndpoint(v string) *ModelSelectionFallbackChains {
	s.Endpoint = &v
	return s
}

func (s *ModelSelectionFallbackChains) SetFallbackModelCardIds(v []*string) *ModelSelectionFallbackChains {
	s.FallbackModelCardIds = v
	return s
}

func (s *ModelSelectionFallbackChains) SetSourceModelCardId(v string) *ModelSelectionFallbackChains {
	s.SourceModelCardId = &v
	return s
}

func (s *ModelSelectionFallbackChains) Validate() error {
	return dara.Validate(s)
}

type ModelSelectionModels struct {
	// The model card ID.
	//
	// example:
	//
	// msc-d00000000000000000000
	ModelCardId *string `json:"modelCardId,omitempty" xml:"modelCardId,omitempty"`
	// The explicit execution target AI service ID for the model. This parameter is optional and takes precedence over the provider default.
	//
	// example:
	//
	// svc-d00000000000000000000
	ServiceId *string `json:"serviceId,omitempty" xml:"serviceId,omitempty"`
}

func (s ModelSelectionModels) String() string {
	return dara.Prettify(s)
}

func (s ModelSelectionModels) GoString() string {
	return s.String()
}

func (s *ModelSelectionModels) GetModelCardId() *string {
	return s.ModelCardId
}

func (s *ModelSelectionModels) GetServiceId() *string {
	return s.ServiceId
}

func (s *ModelSelectionModels) SetModelCardId(v string) *ModelSelectionModels {
	s.ModelCardId = &v
	return s
}

func (s *ModelSelectionModels) SetServiceId(v string) *ModelSelectionModels {
	s.ServiceId = &v
	return s
}

func (s *ModelSelectionModels) Validate() error {
	return dara.Validate(s)
}
