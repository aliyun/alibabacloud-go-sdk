// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyModelApiRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConfig(v string) *ModifyModelApiRequest
	GetConfig() *string
	SetGwClusterId(v string) *ModifyModelApiRequest
	GetGwClusterId() *string
	SetModelApiId(v string) *ModifyModelApiRequest
	GetModelApiId() *string
	SetModelCategory(v string) *ModifyModelApiRequest
	GetModelCategory() *string
	SetModelMappings(v string) *ModifyModelApiRequest
	GetModelMappings() *string
	SetModelRouting(v string) *ModifyModelApiRequest
	GetModelRouting() *string
	SetName(v string) *ModifyModelApiRequest
	GetName() *string
	SetPathPrefix(v string) *ModifyModelApiRequest
	GetPathPrefix() *string
	SetProtocol(v string) *ModifyModelApiRequest
	GetProtocol() *string
	SetRecordInput(v string) *ModifyModelApiRequest
	GetRecordInput() *string
	SetRecordOutput(v string) *ModifyModelApiRequest
	GetRecordOutput() *string
	SetRegionId(v string) *ModifyModelApiRequest
	GetRegionId() *string
	SetRouteRules(v string) *ModifyModelApiRequest
	GetRouteRules() *string
}

type ModifyModelApiRequest struct {
	// The gateway retry configuration.
	//
	// example:
	//
	// {
	//
	// 	"failover": {
	//
	// 		"enabled": true,
	//
	// 		"max_provider_retries": 2,
	//
	// 		"max_failover_providers": 1,
	//
	// 		"retryable_status_codes": [429, 500, 502, 503, 504],
	//
	// 		"retry_delay": 0.5
	//
	// 	}
	//
	// }
	Config *string `json:"Config,omitempty" xml:"Config,omitempty"`
	// The gateway instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// pg-xxxxxxxxxx
	GwClusterId *string `json:"GwClusterId,omitempty" xml:"GwClusterId,omitempty"`
	// The ID of the model API.
	//
	// This parameter is required.
	//
	// example:
	//
	// mi-xxxxx
	ModelApiId *string `json:"ModelApiId,omitempty" xml:"ModelApiId,omitempty"`
	// The category. Valid values:
	//
	// 	- **text**
	//
	// 	- **embedding**
	//
	// 	- **rerank**
	//
	// example:
	//
	// text
	ModelCategory *string `json:"ModelCategory,omitempty" xml:"ModelCategory,omitempty"`
	// The model name mappings, specified as a JSON array string. Each element is in the format of {"Source":"client model name","Target":"upstream model name"}. If you do not specify this parameter, the current mappings are retained. If you specify "[]", the mappings are cleared. The Source field is matched in a case-insensitive manner and must be unique.
	//
	// example:
	//
	// [{"Source":"client-model-a","Target":"upstream-model-a"},{"Source":"client-model-b","Target":"upstream-model-b"}]
	ModelMappings *string `json:"ModelMappings,omitempty" xml:"ModelMappings,omitempty"`
	// The intelligent routing candidate list, specified as a JSON object string in the format of {"AllowedModels":["candidate model name"]}. If you do not specify this parameter, the current list is retained. If you specify this parameter, the existing list is fully replaced with the new list. You cannot use an empty list to disable the feature. Saving candidates does not enable the feature. You must call EnableModelRouting first.
	//
	// example:
	//
	// {"AllowedModels":["deepseek-v4-pro","glm-5.3","glm-5.3-flash"]}
	ModelRouting *string `json:"ModelRouting,omitempty" xml:"ModelRouting,omitempty"`
	// The name of the model API. This parameter is modifiable. If you do not specify this parameter, the original name is retained.
	//
	// example:
	//
	// jev-chat-api
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The API path prefix.
	//
	// example:
	//
	// /test
	PathPrefix *string `json:"PathPrefix,omitempty" xml:"PathPrefix,omitempty"`
	// The protocol. Valid values:
	//
	// 	- **openai**
	//
	// 	- **anthropic**
	//
	// 	- **bailian**
	//
	// 	- **vllm**
	//
	// example:
	//
	// openai
	Protocol *string `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
	// The number of input tokens.
	//
	// example:
	//
	// 10
	RecordInput *string `json:"RecordInput,omitempty" xml:"RecordInput,omitempty"`
	// The number of output tokens.
	//
	// example:
	//
	// 10
	RecordOutput *string `json:"RecordOutput,omitempty" xml:"RecordOutput,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The list of routing rules, specified as a JSON array string. Usage note: Correct the value of FallbackMode in the RouteRules parameter example by replacing it with a valid value or removing the field, and add the list of valid values and their descriptions for the FallbackMode subfield in the description of RouteRules.
	//
	// example:
	//
	// [
	//
	//   {
	//
	//     "RuleName": "string",
	//
	//     "FallbackMode": "failover",
	//
	//     "MatchModelListJson": "[]",
	//
	//     "providerBalancerAlgorithm": "round-robin",
	//
	//     "Providers": [
	//
	//       {
	//
	//         "ModelServiceName": "string",
	//
	//         "Weight": "0",
	//
	//         "model_protocol": "vllm"
	//
	//         "ModelList": "[]"
	//
	//       }
	//
	//     ],
	//
	//     "FallbackProviders": [
	//
	//       {
	//
	//         "ModelServiceName": "string",
	//
	//         "model_protocol": "anthropic",
	//
	//         "Weight": "10",
	//
	//         "ModelList": "[]"
	//
	//       }
	//
	//     ]
	//
	//   }
	//
	// ]
	RouteRules *string `json:"RouteRules,omitempty" xml:"RouteRules,omitempty"`
}

func (s ModifyModelApiRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyModelApiRequest) GoString() string {
	return s.String()
}

func (s *ModifyModelApiRequest) GetConfig() *string {
	return s.Config
}

func (s *ModifyModelApiRequest) GetGwClusterId() *string {
	return s.GwClusterId
}

func (s *ModifyModelApiRequest) GetModelApiId() *string {
	return s.ModelApiId
}

func (s *ModifyModelApiRequest) GetModelCategory() *string {
	return s.ModelCategory
}

func (s *ModifyModelApiRequest) GetModelMappings() *string {
	return s.ModelMappings
}

func (s *ModifyModelApiRequest) GetModelRouting() *string {
	return s.ModelRouting
}

func (s *ModifyModelApiRequest) GetName() *string {
	return s.Name
}

func (s *ModifyModelApiRequest) GetPathPrefix() *string {
	return s.PathPrefix
}

func (s *ModifyModelApiRequest) GetProtocol() *string {
	return s.Protocol
}

func (s *ModifyModelApiRequest) GetRecordInput() *string {
	return s.RecordInput
}

func (s *ModifyModelApiRequest) GetRecordOutput() *string {
	return s.RecordOutput
}

func (s *ModifyModelApiRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyModelApiRequest) GetRouteRules() *string {
	return s.RouteRules
}

func (s *ModifyModelApiRequest) SetConfig(v string) *ModifyModelApiRequest {
	s.Config = &v
	return s
}

func (s *ModifyModelApiRequest) SetGwClusterId(v string) *ModifyModelApiRequest {
	s.GwClusterId = &v
	return s
}

func (s *ModifyModelApiRequest) SetModelApiId(v string) *ModifyModelApiRequest {
	s.ModelApiId = &v
	return s
}

func (s *ModifyModelApiRequest) SetModelCategory(v string) *ModifyModelApiRequest {
	s.ModelCategory = &v
	return s
}

func (s *ModifyModelApiRequest) SetModelMappings(v string) *ModifyModelApiRequest {
	s.ModelMappings = &v
	return s
}

func (s *ModifyModelApiRequest) SetModelRouting(v string) *ModifyModelApiRequest {
	s.ModelRouting = &v
	return s
}

func (s *ModifyModelApiRequest) SetName(v string) *ModifyModelApiRequest {
	s.Name = &v
	return s
}

func (s *ModifyModelApiRequest) SetPathPrefix(v string) *ModifyModelApiRequest {
	s.PathPrefix = &v
	return s
}

func (s *ModifyModelApiRequest) SetProtocol(v string) *ModifyModelApiRequest {
	s.Protocol = &v
	return s
}

func (s *ModifyModelApiRequest) SetRecordInput(v string) *ModifyModelApiRequest {
	s.RecordInput = &v
	return s
}

func (s *ModifyModelApiRequest) SetRecordOutput(v string) *ModifyModelApiRequest {
	s.RecordOutput = &v
	return s
}

func (s *ModifyModelApiRequest) SetRegionId(v string) *ModifyModelApiRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyModelApiRequest) SetRouteRules(v string) *ModifyModelApiRequest {
	s.RouteRules = &v
	return s
}

func (s *ModifyModelApiRequest) Validate() error {
	return dara.Validate(s)
}
