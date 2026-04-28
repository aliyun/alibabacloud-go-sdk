// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyModelApiRequest interface {
	dara.Model
	String() string
	GoString() string
	SetGwClusterId(v string) *ModifyModelApiRequest
	GetGwClusterId() *string
	SetModelApiId(v string) *ModifyModelApiRequest
	GetModelApiId() *string
	SetModelCategory(v string) *ModifyModelApiRequest
	GetModelCategory() *string
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
	// mi-xxxxx
	ModelApiId *string `json:"ModelApiId,omitempty" xml:"ModelApiId,omitempty"`
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

func (s ModifyModelApiRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyModelApiRequest) GoString() string {
	return s.String()
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
