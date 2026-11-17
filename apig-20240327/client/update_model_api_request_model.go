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
	// The API access prefix. The default value is / during creation. If omitted during an update, the original value is retained.
	//
	// example:
	//
	// /beta
	BasePath *string `json:"basePath,omitempty" xml:"basePath,omitempty"`
	// The description of the model API. If omitted during an update, the original value is retained.
	//
	// example:
	//
	// Model API for business integration
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// Specifies whether to enable direct supply fallback. If omitted during an update, the original value is retained. Explicitly set to false to disable it.
	//
	// if can be null:
	// false
	DirectSupplyFallbackEnabled *bool `json:"directSupplyFallbackEnabled,omitempty" xml:"directSupplyFallbackEnabled,omitempty"`
	// The IDs of the selected custom or default domain names of the instance. This parameter must be explicitly non-empty during creation. If omitted during an update, the original value is retained. When submitted, the entire list is replaced and cannot be empty.
	//
	// if can be null:
	// false
	DomainIds []*string `json:"domainIds,omitempty" xml:"domainIds,omitempty" type:"Repeated"`
	// Specifies whether to enable consumer API key authentication. If omitted during an update, the original value is retained. The type is fixed to Apikey or Custom.
	EnableAuth *bool `json:"enableAuth,omitempty" xml:"enableAuth,omitempty"`
	// The time to first byte.
	//
	// example:
	//
	// 30000
	FirstByteTimeout *int32 `json:"firstByteTimeout,omitempty" xml:"firstByteTimeout,omitempty"`
	// The model selection. This parameter is required during creation. If omitted during an update, the original value is retained. When submitted, the entire selection is replaced.
	//
	// if can be null:
	// false
	ModelSelection *ModelSelection `json:"modelSelection,omitempty" xml:"modelSelection,omitempty"`
	// Supported only by the new version of ModelAPI in minimalist mode: merges the actual source model IDs added to or removed from the model source in this request into the latest selection within the lock. This parameter is mutually exclusive with modelSelection and cannot be null or an empty delta. The standard update and publish process applies.
	//
	// if can be null:
	// false
	ModelSelectionDelta *ModelSelectionDelta `json:"modelSelectionDelta,omitempty" xml:"modelSelectionDelta,omitempty"`
	// The name of the model API. If omitted or set to null, the original value is retained. When submitted, the value is trimmed of leading and trailing whitespace, must be non-empty, contain a maximum of 128 UTF-8 bytes, and cannot contain @ or /. Renaming the API does not change the API ID or access prefix.
	//
	// example:
	//
	// beta-model-api
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The configurations for AI request logs and AI security protection. If omitted during an update, the original values are retained. When submitted, the inline policies of both types are entirely replaced. An empty array clears both types. Independent token quotas are not affected.
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
