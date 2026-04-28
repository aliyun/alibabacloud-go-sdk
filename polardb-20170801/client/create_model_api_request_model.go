// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateModelApiRequest interface {
	dara.Model
	String() string
	GoString() string
	SetForceModel(v string) *CreateModelApiRequest
	GetForceModel() *string
	SetGwClusterId(v string) *CreateModelApiRequest
	GetGwClusterId() *string
	SetModelCategory(v string) *CreateModelApiRequest
	GetModelCategory() *string
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
	ForceModel *string `json:"ForceModel,omitempty" xml:"ForceModel,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// pg-xxxxxxxxxx
	GwClusterId *string `json:"GwClusterId,omitempty" xml:"GwClusterId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// text
	ModelCategory *string `json:"ModelCategory,omitempty" xml:"ModelCategory,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// /test
	PathPrefix *string `json:"PathPrefix,omitempty" xml:"PathPrefix,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// openai
	Protocol *string `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
	// example:
	//
	// 10
	RecordInput *string `json:"RecordInput,omitempty" xml:"RecordInput,omitempty"`
	// example:
	//
	// 10
	RecordOutput *string `json:"RecordOutput,omitempty" xml:"RecordOutput,omitempty"`
	RegionId     *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// This parameter is required.
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

func (s *CreateModelApiRequest) GetForceModel() *string {
	return s.ForceModel
}

func (s *CreateModelApiRequest) GetGwClusterId() *string {
	return s.GwClusterId
}

func (s *CreateModelApiRequest) GetModelCategory() *string {
	return s.ModelCategory
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
