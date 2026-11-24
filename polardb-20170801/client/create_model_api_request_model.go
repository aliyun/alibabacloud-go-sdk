// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateModelApiRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConfig(v string) *CreateModelApiRequest
	GetConfig() *string
	SetForceModel(v string) *CreateModelApiRequest
	GetForceModel() *string
	SetGwClusterId(v string) *CreateModelApiRequest
	GetGwClusterId() *string
	SetModelCategory(v string) *CreateModelApiRequest
	GetModelCategory() *string
	SetModelMappings(v string) *CreateModelApiRequest
	GetModelMappings() *string
	SetModelRouting(v string) *CreateModelApiRequest
	GetModelRouting() *string
	SetName(v string) *CreateModelApiRequest
	GetName() *string
	SetPathPrefix(v string) *CreateModelApiRequest
	GetPathPrefix() *string
	SetProtocol(v string) *CreateModelApiRequest
	GetProtocol() *string
	SetRecordInput(v string) *CreateModelApiRequest
	GetRecordInput() *string
	SetRecordOutput(v string) *CreateModelApiRequest
	GetRecordOutput() *string
	SetRegionId(v string) *CreateModelApiRequest
	GetRegionId() *string
	SetRouteRules(v string) *CreateModelApiRequest
	GetRouteRules() *string
}

type CreateModelApiRequest struct {
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
	// The forced model.
	//
	// example:
	//
	// Qwen3-32B
	ForceModel *string `json:"ForceModel,omitempty" xml:"ForceModel,omitempty"`
	// The ID of the gateway instance. You can obtain this ID from the GwClusterId value returned by the CreateGateway operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// pg-xxxxxxxxxx
	GwClusterId *string `json:"GwClusterId,omitempty" xml:"GwClusterId,omitempty"`
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
	// The model name mappings, specified as a JSON array string. Each element is in the format of {"Source":"client model name","Target":"upstream model name"}. The Source parameter is matched in a case-insensitive manner and must be unique. The Target parameter retains its original case when forwarded. Both parameters cannot be empty and cannot contain leading or trailing spaces.
	//
	// example:
	//
	// [{"Source":"client-model-a","Target":"upstream-model-a"},{"Source":"client-model-b","Target":"upstream-model-b"}]
	ModelMappings *string `json:"ModelMappings,omitempty" xml:"ModelMappings,omitempty"`
	// The intelligent routing candidate list, specified as a JSON object string. Currently, it only contains AllowedModels, which is a non-empty array of unique model names. The order of the array determines the candidate order. Do not pass the Enabled, cost, or quality tier parameters. Saving candidates does not enable the feature. You must first call the EnableModelRouting operation to obtain the gateway entitlements.
	//
	// example:
	//
	// {"AllowedModels":["deepseek-v4-pro","glm-5.3","glm-5.3-flash"]}
	ModelRouting *string `json:"ModelRouting,omitempty" xml:"ModelRouting,omitempty"`
	// The name of the model API.
	//
	// example:
	//
	// test
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
	// The number of input points.
	//
	// example:
	//
	// 10
	RecordInput *string `json:"RecordInput,omitempty" xml:"RecordInput,omitempty"`
	// The number of output points.
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
	// The list of routing rules, specified as a JSON array string.
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

func (s CreateModelApiRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateModelApiRequest) GoString() string {
	return s.String()
}

func (s *CreateModelApiRequest) GetConfig() *string {
	return s.Config
}

func (s *CreateModelApiRequest) GetForceModel() *string {
	return s.ForceModel
}

func (s *CreateModelApiRequest) GetGwClusterId() *string {
	return s.GwClusterId
}

func (s *CreateModelApiRequest) GetModelCategory() *string {
	return s.ModelCategory
}

func (s *CreateModelApiRequest) GetModelMappings() *string {
	return s.ModelMappings
}

func (s *CreateModelApiRequest) GetModelRouting() *string {
	return s.ModelRouting
}

func (s *CreateModelApiRequest) GetName() *string {
	return s.Name
}

func (s *CreateModelApiRequest) GetPathPrefix() *string {
	return s.PathPrefix
}

func (s *CreateModelApiRequest) GetProtocol() *string {
	return s.Protocol
}

func (s *CreateModelApiRequest) GetRecordInput() *string {
	return s.RecordInput
}

func (s *CreateModelApiRequest) GetRecordOutput() *string {
	return s.RecordOutput
}

func (s *CreateModelApiRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateModelApiRequest) GetRouteRules() *string {
	return s.RouteRules
}

func (s *CreateModelApiRequest) SetConfig(v string) *CreateModelApiRequest {
	s.Config = &v
	return s
}

func (s *CreateModelApiRequest) SetForceModel(v string) *CreateModelApiRequest {
	s.ForceModel = &v
	return s
}

func (s *CreateModelApiRequest) SetGwClusterId(v string) *CreateModelApiRequest {
	s.GwClusterId = &v
	return s
}

func (s *CreateModelApiRequest) SetModelCategory(v string) *CreateModelApiRequest {
	s.ModelCategory = &v
	return s
}

func (s *CreateModelApiRequest) SetModelMappings(v string) *CreateModelApiRequest {
	s.ModelMappings = &v
	return s
}

func (s *CreateModelApiRequest) SetModelRouting(v string) *CreateModelApiRequest {
	s.ModelRouting = &v
	return s
}

func (s *CreateModelApiRequest) SetName(v string) *CreateModelApiRequest {
	s.Name = &v
	return s
}

func (s *CreateModelApiRequest) SetPathPrefix(v string) *CreateModelApiRequest {
	s.PathPrefix = &v
	return s
}

func (s *CreateModelApiRequest) SetProtocol(v string) *CreateModelApiRequest {
	s.Protocol = &v
	return s
}

func (s *CreateModelApiRequest) SetRecordInput(v string) *CreateModelApiRequest {
	s.RecordInput = &v
	return s
}

func (s *CreateModelApiRequest) SetRecordOutput(v string) *CreateModelApiRequest {
	s.RecordOutput = &v
	return s
}

func (s *CreateModelApiRequest) SetRegionId(v string) *CreateModelApiRequest {
	s.RegionId = &v
	return s
}

func (s *CreateModelApiRequest) SetRouteRules(v string) *CreateModelApiRequest {
	s.RouteRules = &v
	return s
}

func (s *CreateModelApiRequest) Validate() error {
	return dara.Validate(s)
}
