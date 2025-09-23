// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateModelServiceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAiNodes(v []*string) *CreateModelServiceRequest
	GetAiNodes() []*string
	SetClientToken(v string) *CreateModelServiceRequest
	GetClientToken() *string
	SetDBInstanceId(v string) *CreateModelServiceRequest
	GetDBInstanceId() *string
	SetDescription(v string) *CreateModelServiceRequest
	GetDescription() *string
	SetEnablePublicConnection(v bool) *CreateModelServiceRequest
	GetEnablePublicConnection() *bool
	SetInferenceEngine(v string) *CreateModelServiceRequest
	GetInferenceEngine() *string
	SetModelName(v string) *CreateModelServiceRequest
	GetModelName() *string
	SetModelParams(v map[string]interface{}) *CreateModelServiceRequest
	GetModelParams() map[string]interface{}
	SetReplicas(v int32) *CreateModelServiceRequest
	GetReplicas() *int32
	SetResourceGroupId(v string) *CreateModelServiceRequest
	GetResourceGroupId() *string
	SetSecurityIPList(v string) *CreateModelServiceRequest
	GetSecurityIPList() *string
}

type CreateModelServiceRequest struct {
	// This parameter is required.
	AiNodes []*string `json:"AiNodes,omitempty" xml:"AiNodes,omitempty" type:"Repeated"`
	// example:
	//
	// 0c593ea1-3bea-11e9-b96b-88**********
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// gp-xxxxxxxxx
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// example:
	//
	// test
	Description            *string `json:"Description,omitempty" xml:"Description,omitempty"`
	EnablePublicConnection *bool   `json:"EnablePublicConnection,omitempty" xml:"EnablePublicConnection,omitempty"`
	// example:
	//
	// vllm
	InferenceEngine *string `json:"InferenceEngine,omitempty" xml:"InferenceEngine,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// Qwen3-Embedding-8B
	ModelName   *string                `json:"ModelName,omitempty" xml:"ModelName,omitempty"`
	ModelParams map[string]interface{} `json:"ModelParams,omitempty" xml:"ModelParams,omitempty"`
	// example:
	//
	// 1
	Replicas *int32 `json:"Replicas,omitempty" xml:"Replicas,omitempty"`
	// example:
	//
	// rg-bp67acfmxazb4p****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// example:
	//
	// 127.0.0.1
	SecurityIPList *string `json:"SecurityIPList,omitempty" xml:"SecurityIPList,omitempty"`
}

func (s CreateModelServiceRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateModelServiceRequest) GoString() string {
	return s.String()
}

func (s *CreateModelServiceRequest) GetAiNodes() []*string {
	return s.AiNodes
}

func (s *CreateModelServiceRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateModelServiceRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *CreateModelServiceRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateModelServiceRequest) GetEnablePublicConnection() *bool {
	return s.EnablePublicConnection
}

func (s *CreateModelServiceRequest) GetInferenceEngine() *string {
	return s.InferenceEngine
}

func (s *CreateModelServiceRequest) GetModelName() *string {
	return s.ModelName
}

func (s *CreateModelServiceRequest) GetModelParams() map[string]interface{} {
	return s.ModelParams
}

func (s *CreateModelServiceRequest) GetReplicas() *int32 {
	return s.Replicas
}

func (s *CreateModelServiceRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *CreateModelServiceRequest) GetSecurityIPList() *string {
	return s.SecurityIPList
}

func (s *CreateModelServiceRequest) SetAiNodes(v []*string) *CreateModelServiceRequest {
	s.AiNodes = v
	return s
}

func (s *CreateModelServiceRequest) SetClientToken(v string) *CreateModelServiceRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateModelServiceRequest) SetDBInstanceId(v string) *CreateModelServiceRequest {
	s.DBInstanceId = &v
	return s
}

func (s *CreateModelServiceRequest) SetDescription(v string) *CreateModelServiceRequest {
	s.Description = &v
	return s
}

func (s *CreateModelServiceRequest) SetEnablePublicConnection(v bool) *CreateModelServiceRequest {
	s.EnablePublicConnection = &v
	return s
}

func (s *CreateModelServiceRequest) SetInferenceEngine(v string) *CreateModelServiceRequest {
	s.InferenceEngine = &v
	return s
}

func (s *CreateModelServiceRequest) SetModelName(v string) *CreateModelServiceRequest {
	s.ModelName = &v
	return s
}

func (s *CreateModelServiceRequest) SetModelParams(v map[string]interface{}) *CreateModelServiceRequest {
	s.ModelParams = v
	return s
}

func (s *CreateModelServiceRequest) SetReplicas(v int32) *CreateModelServiceRequest {
	s.Replicas = &v
	return s
}

func (s *CreateModelServiceRequest) SetResourceGroupId(v string) *CreateModelServiceRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateModelServiceRequest) SetSecurityIPList(v string) *CreateModelServiceRequest {
	s.SecurityIPList = &v
	return s
}

func (s *CreateModelServiceRequest) Validate() error {
	return dara.Validate(s)
}
