// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateModelApiRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBasePath(v string) *CreateModelApiRequest
	GetBasePath() *string
	SetDescription(v string) *CreateModelApiRequest
	GetDescription() *string
	SetDirectSupplyFallbackEnabled(v bool) *CreateModelApiRequest
	GetDirectSupplyFallbackEnabled() *bool
	SetDomainIds(v []*string) *CreateModelApiRequest
	GetDomainIds() []*string
	SetEnableAuth(v bool) *CreateModelApiRequest
	GetEnableAuth() *bool
	SetFirstByteTimeout(v int32) *CreateModelApiRequest
	GetFirstByteTimeout() *int32
	SetGatewayId(v string) *CreateModelApiRequest
	GetGatewayId() *string
	SetInitialConfiguration(v bool) *CreateModelApiRequest
	GetInitialConfiguration() *bool
	SetModelSelection(v *ModelSelection) *CreateModelApiRequest
	GetModelSelection() *ModelSelection
	SetName(v string) *CreateModelApiRequest
	GetName() *string
	SetPolicyConfigs(v []*ModelApiPolicyConfig) *CreateModelApiRequest
	GetPolicyConfigs() []*ModelApiPolicyConfig
	SetResourceGroupId(v string) *CreateModelApiRequest
	GetResourceGroupId() *string
	SetClientToken(v string) *CreateModelApiRequest
	GetClientToken() *string
}

type CreateModelApiRequest struct {
	// The API access prefix. The default value for creation is /. If omitted during an update, the original value is retained.
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
	// Specifies whether to enable direct supply fallback. If omitted in a create request, the default value is false, which does not implicitly incur pay-as-you-go tail charges. You can only enable this feature by explicitly setting this parameter to true.
	//
	// if can be null:
	// false
	DirectSupplyFallbackEnabled *bool `json:"directSupplyFallbackEnabled,omitempty" xml:"directSupplyFallbackEnabled,omitempty"`
	// The IDs of the selected custom or default instance domain names. This parameter must be explicitly specified and cannot be empty during creation. If omitted during an update, the original value is retained. If submitted, the entire list is replaced and cannot be empty.
	//
	// This parameter is required.
	//
	// if can be null:
	// false
	DomainIds []*string `json:"domainIds,omitempty" xml:"domainIds,omitempty" type:"Repeated"`
	// Specifies whether to enable consumer API key authentication. If omitted during an update, the original value is retained. The type is fixed to Apikey or Custom.
	EnableAuth *bool `json:"enableAuth,omitempty" xml:"enableAuth,omitempty"`
	// The timeout period for the first byte.
	//
	// example:
	//
	// 30000
	FirstByteTimeout *int32 `json:"firstByteTimeout,omitempty" xml:"firstByteTimeout,omitempty"`
	// The target dedicated gateway instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// gw-d00000000000000000000
	GatewayId *string `json:"gatewayId,omitempty" xml:"gatewayId,omitempty"`
	// Specifies whether to explicitly perform the initial configuration. If omitted or set to false, standard multi-API creation is retained. If set to true, you must explicitly set name to default, basePath to /, and provide a non-empty clientToken. Additionally, when saving, multi-API management must be disabled and the number of complete new-version APIs must be zero.
	//
	// if can be null:
	// false
	InitialConfiguration *bool `json:"initialConfiguration,omitempty" xml:"initialConfiguration,omitempty"`
	// The model selection. This parameter is required for creation. If omitted during an update, the original value is retained. If submitted, the entire value is replaced.
	//
	// This parameter is required.
	//
	// if can be null:
	// false
	ModelSelection *ModelSelection `json:"modelSelection,omitempty" xml:"modelSelection,omitempty"`
	// The name of the model API.
	//
	// This parameter is required.
	//
	// example:
	//
	// beta-model-api
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The configurations for AI request logging and AI security protection only. If omitted during an update, the original values are retained. If submitted, the inline policies of both types are entirely replaced. An empty array clears both types. Independent token quotas are not affected.
	//
	// if can be null:
	// false
	PolicyConfigs []*ModelApiPolicyConfig `json:"policyConfigs,omitempty" xml:"policyConfigs,omitempty" type:"Repeated"`
	// The ID of the resource group.
	//
	// example:
	//
	// rg-acf0000000000000000
	ResourceGroupId *string `json:"resourceGroupId,omitempty" xml:"resourceGroupId,omitempty"`
	// The idempotency identity. The caller generates a globally unique value (we recommend a UUID) to invoke the API. The maximum length is 64 characters. Within approximately 24 hours after the initial request succeeds, if you invoke a duplicate request with the same clientToken and identical request parameter values, the system directly returns the initial creation result. If a request parameter is inconsistent or the initial request is still being processed, the system returns the corresponding idempotency fault. If omitted, the original creation behavior is retained.
	//
	// example:
	//
	// create-model-api-client-token-001
	ClientToken *string `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
}

func (s CreateModelApiRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateModelApiRequest) GoString() string {
	return s.String()
}

func (s *CreateModelApiRequest) GetBasePath() *string {
	return s.BasePath
}

func (s *CreateModelApiRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateModelApiRequest) GetDirectSupplyFallbackEnabled() *bool {
	return s.DirectSupplyFallbackEnabled
}

func (s *CreateModelApiRequest) GetDomainIds() []*string {
	return s.DomainIds
}

func (s *CreateModelApiRequest) GetEnableAuth() *bool {
	return s.EnableAuth
}

func (s *CreateModelApiRequest) GetFirstByteTimeout() *int32 {
	return s.FirstByteTimeout
}

func (s *CreateModelApiRequest) GetGatewayId() *string {
	return s.GatewayId
}

func (s *CreateModelApiRequest) GetInitialConfiguration() *bool {
	return s.InitialConfiguration
}

func (s *CreateModelApiRequest) GetModelSelection() *ModelSelection {
	return s.ModelSelection
}

func (s *CreateModelApiRequest) GetName() *string {
	return s.Name
}

func (s *CreateModelApiRequest) GetPolicyConfigs() []*ModelApiPolicyConfig {
	return s.PolicyConfigs
}

func (s *CreateModelApiRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *CreateModelApiRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateModelApiRequest) SetBasePath(v string) *CreateModelApiRequest {
	s.BasePath = &v
	return s
}

func (s *CreateModelApiRequest) SetDescription(v string) *CreateModelApiRequest {
	s.Description = &v
	return s
}

func (s *CreateModelApiRequest) SetDirectSupplyFallbackEnabled(v bool) *CreateModelApiRequest {
	s.DirectSupplyFallbackEnabled = &v
	return s
}

func (s *CreateModelApiRequest) SetDomainIds(v []*string) *CreateModelApiRequest {
	s.DomainIds = v
	return s
}

func (s *CreateModelApiRequest) SetEnableAuth(v bool) *CreateModelApiRequest {
	s.EnableAuth = &v
	return s
}

func (s *CreateModelApiRequest) SetFirstByteTimeout(v int32) *CreateModelApiRequest {
	s.FirstByteTimeout = &v
	return s
}

func (s *CreateModelApiRequest) SetGatewayId(v string) *CreateModelApiRequest {
	s.GatewayId = &v
	return s
}

func (s *CreateModelApiRequest) SetInitialConfiguration(v bool) *CreateModelApiRequest {
	s.InitialConfiguration = &v
	return s
}

func (s *CreateModelApiRequest) SetModelSelection(v *ModelSelection) *CreateModelApiRequest {
	s.ModelSelection = v
	return s
}

func (s *CreateModelApiRequest) SetName(v string) *CreateModelApiRequest {
	s.Name = &v
	return s
}

func (s *CreateModelApiRequest) SetPolicyConfigs(v []*ModelApiPolicyConfig) *CreateModelApiRequest {
	s.PolicyConfigs = v
	return s
}

func (s *CreateModelApiRequest) SetResourceGroupId(v string) *CreateModelApiRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateModelApiRequest) SetClientToken(v string) *CreateModelApiRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateModelApiRequest) Validate() error {
	if s.ModelSelection != nil {
		if err := s.ModelSelection.Validate(); err != nil {
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
