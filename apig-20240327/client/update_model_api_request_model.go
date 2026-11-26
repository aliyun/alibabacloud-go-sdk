// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateModelApiRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBasePath(v string) *UpdateModelApiRequest
	GetBasePath() *string
	SetDescription(v string) *UpdateModelApiRequest
	GetDescription() *string
	SetDirectSupplyFallbackEnabled(v bool) *UpdateModelApiRequest
	GetDirectSupplyFallbackEnabled() *bool
	SetDomainIds(v []*string) *UpdateModelApiRequest
	GetDomainIds() []*string
	SetEnableAuth(v bool) *UpdateModelApiRequest
	GetEnableAuth() *bool
	SetFirstByteTimeout(v int32) *UpdateModelApiRequest
	GetFirstByteTimeout() *int32
	SetModelSelection(v *ModelSelection) *UpdateModelApiRequest
	GetModelSelection() *ModelSelection
	SetModelSelectionDelta(v *ModelSelectionDelta) *UpdateModelApiRequest
	GetModelSelectionDelta() *ModelSelectionDelta
	SetName(v string) *UpdateModelApiRequest
	GetName() *string
	SetPolicyConfigs(v []*ModelApiPolicyConfig) *UpdateModelApiRequest
	GetPolicyConfigs() []*ModelApiPolicyConfig
}

type UpdateModelApiRequest struct {
	// The access prefix of the API. The default value is / when created. If omitted in an update request, the original value is retained.
	//
	// example:
	//
	// /beta
	BasePath *string `json:"basePath,omitempty" xml:"basePath,omitempty"`
	// The description of the model API. If omitted in an update request, the original value is retained.
	//
	// example:
	//
	// Model API for business integration
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// Specifies whether to enable the direct supply fallback. If omitted in an update request, the original value is retained. Set to false to explicitly disable it.
	//
	// if can be null:
	// false
	DirectSupplyFallbackEnabled *bool `json:"directSupplyFallbackEnabled,omitempty" xml:"directSupplyFallbackEnabled,omitempty"`
	// The IDs of the selected custom or instance default domain names. This parameter must be explicitly set to a non-empty value when creating an API. If omitted in an update request, the original value is retained. When submitted, the entire list is replaced and cannot be empty.
	//
	// if can be null:
	// false
	DomainIds []*string `json:"domainIds,omitempty" xml:"domainIds,omitempty" type:"Repeated"`
	// Specifies whether to enable consumer API key authentication. If omitted in an update request, the original value is retained. The type is fixed to Apikey or Custom.
	EnableAuth *bool `json:"enableAuth,omitempty" xml:"enableAuth,omitempty"`
	// The first byte timeout.
	//
	// example:
	//
	// 30000
	FirstByteTimeout *int32 `json:"firstByteTimeout,omitempty" xml:"firstByteTimeout,omitempty"`
	// The model selection. This parameter is required when creating an API. If omitted in an update request, the original value is retained. When submitted, the entire selection is replaced.
	//
	// if can be null:
	// false
	ModelSelection *ModelSelection `json:"modelSelection,omitempty" xml:"modelSelection,omitempty"`
	// Supported only by the new version of Model API in minimalist mode. Merges the actual source model IDs added or removed in this request into the latest locked selection. This parameter is mutually exclusive with modelSelection and cannot be null or an empty delta. The standard update and publish process applies.
	//
	// if can be null:
	// false
	ModelSelectionDelta *ModelSelectionDelta `json:"modelSelectionDelta,omitempty" xml:"modelSelectionDelta,omitempty"`
	// The name of the model API. If omitted or set to null, the original value is retained. When submitted, the name is trimmed of leading and trailing whitespace, must not be empty, can be up to 128 UTF-8 bytes in length, and cannot contain the at sign (@) or forward slash (/). Renaming the API does not change the API ID or the access prefix.
	//
	// example:
	//
	// beta-model-api
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The policy configurations for AI request logging and AI security protection only. If omitted in an update request, the original values are retained. When submitted, the inline policies of both types are entirely replaced. An empty array clears both types of policies. Independent token quotas are not affected.
	//
	// if can be null:
	// false
	PolicyConfigs []*ModelApiPolicyConfig `json:"policyConfigs,omitempty" xml:"policyConfigs,omitempty" type:"Repeated"`
}

func (s UpdateModelApiRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateModelApiRequest) GoString() string {
	return s.String()
}

func (s *UpdateModelApiRequest) GetBasePath() *string {
	return s.BasePath
}

func (s *UpdateModelApiRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateModelApiRequest) GetDirectSupplyFallbackEnabled() *bool {
	return s.DirectSupplyFallbackEnabled
}

func (s *UpdateModelApiRequest) GetDomainIds() []*string {
	return s.DomainIds
}

func (s *UpdateModelApiRequest) GetEnableAuth() *bool {
	return s.EnableAuth
}

func (s *UpdateModelApiRequest) GetFirstByteTimeout() *int32 {
	return s.FirstByteTimeout
}

func (s *UpdateModelApiRequest) GetModelSelection() *ModelSelection {
	return s.ModelSelection
}

func (s *UpdateModelApiRequest) GetModelSelectionDelta() *ModelSelectionDelta {
	return s.ModelSelectionDelta
}

func (s *UpdateModelApiRequest) GetName() *string {
	return s.Name
}

func (s *UpdateModelApiRequest) GetPolicyConfigs() []*ModelApiPolicyConfig {
	return s.PolicyConfigs
}

func (s *UpdateModelApiRequest) SetBasePath(v string) *UpdateModelApiRequest {
	s.BasePath = &v
	return s
}

func (s *UpdateModelApiRequest) SetDescription(v string) *UpdateModelApiRequest {
	s.Description = &v
	return s
}

func (s *UpdateModelApiRequest) SetDirectSupplyFallbackEnabled(v bool) *UpdateModelApiRequest {
	s.DirectSupplyFallbackEnabled = &v
	return s
}

func (s *UpdateModelApiRequest) SetDomainIds(v []*string) *UpdateModelApiRequest {
	s.DomainIds = v
	return s
}

func (s *UpdateModelApiRequest) SetEnableAuth(v bool) *UpdateModelApiRequest {
	s.EnableAuth = &v
	return s
}

func (s *UpdateModelApiRequest) SetFirstByteTimeout(v int32) *UpdateModelApiRequest {
	s.FirstByteTimeout = &v
	return s
}

func (s *UpdateModelApiRequest) SetModelSelection(v *ModelSelection) *UpdateModelApiRequest {
	s.ModelSelection = v
	return s
}

func (s *UpdateModelApiRequest) SetModelSelectionDelta(v *ModelSelectionDelta) *UpdateModelApiRequest {
	s.ModelSelectionDelta = v
	return s
}

func (s *UpdateModelApiRequest) SetName(v string) *UpdateModelApiRequest {
	s.Name = &v
	return s
}

func (s *UpdateModelApiRequest) SetPolicyConfigs(v []*ModelApiPolicyConfig) *UpdateModelApiRequest {
	s.PolicyConfigs = v
	return s
}

func (s *UpdateModelApiRequest) Validate() error {
	if s.ModelSelection != nil {
		if err := s.ModelSelection.Validate(); err != nil {
			return err
		}
	}
	if s.ModelSelectionDelta != nil {
		if err := s.ModelSelectionDelta.Validate(); err != nil {
			return err
		}
	}
	if s.PolicyConfigs != nil {
		for _, item := range s.PolicyConfigs {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
