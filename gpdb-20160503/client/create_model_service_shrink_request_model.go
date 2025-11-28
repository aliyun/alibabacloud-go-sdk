// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateModelServiceShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAiNodesShrink(v string) *CreateModelServiceShrinkRequest
	GetAiNodesShrink() *string
	SetClientToken(v string) *CreateModelServiceShrinkRequest
	GetClientToken() *string
	SetDBInstanceId(v string) *CreateModelServiceShrinkRequest
	GetDBInstanceId() *string
	SetDescription(v string) *CreateModelServiceShrinkRequest
	GetDescription() *string
	SetEnablePublicConnection(v bool) *CreateModelServiceShrinkRequest
	GetEnablePublicConnection() *bool
	SetInferenceEngine(v string) *CreateModelServiceShrinkRequest
	GetInferenceEngine() *string
	SetModelName(v string) *CreateModelServiceShrinkRequest
	GetModelName() *string
	SetModelParamsShrink(v string) *CreateModelServiceShrinkRequest
	GetModelParamsShrink() *string
	SetReplicas(v int32) *CreateModelServiceShrinkRequest
	GetReplicas() *int32
	SetResourceGroupId(v string) *CreateModelServiceShrinkRequest
	GetResourceGroupId() *string
	SetSecurityIPList(v string) *CreateModelServiceShrinkRequest
	GetSecurityIPList() *string
}

type CreateModelServiceShrinkRequest struct {
	// A list of AI nodes for model deployment.
	//
	// This parameter is required.
	AiNodesShrink *string `json:"AiNodes,omitempty" xml:"AiNodes,omitempty"`
	// The client token that is used to ensure the idempotence of the request. For more information, see [How do I ensure the idempotence?](https://help.aliyun.com/document_detail/327176.html)
	//
	// example:
	//
	// 0c593ea1-3bea-11e9-b96b-88**********
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The cluster ID.
	//
	// > You can call the [DescribeDBInstances](https://help.aliyun.com/document_detail/86911.html) operation to query the IDs of all AnalyticDB for PostgreSQL instances in the specified region.
	//
	// This parameter is required.
	//
	// example:
	//
	// gp-xxxxxxxxx
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The description.
	//
	// example:
	//
	// test
	Description            *string `json:"Description,omitempty" xml:"Description,omitempty"`
	EnablePublicConnection *bool   `json:"EnablePublicConnection,omitempty" xml:"EnablePublicConnection,omitempty"`
	// The inference engine. Only vllm is supported.
	//
	// example:
	//
	// vllm
	InferenceEngine *string `json:"InferenceEngine,omitempty" xml:"InferenceEngine,omitempty"`
	// The name of the model.
	//
	// This parameter is required.
	//
	// example:
	//
	// Qwen3-Embedding-8B
	ModelName *string `json:"ModelName,omitempty" xml:"ModelName,omitempty"`
	// Model parameters (to be supported).
	ModelParamsShrink *string `json:"ModelParams,omitempty" xml:"ModelParams,omitempty"`
	// The number of model service replicas.
	//
	// example:
	//
	// 1
	Replicas *int32 `json:"Replicas,omitempty" xml:"Replicas,omitempty"`
	// The ID of the resource group to which the instance belongs. For more information about how to get the ID of a resource group, see [View the basic information of a resource group](https://help.aliyun.com/document_detail/151181.html).
	//
	// example:
	//
	// rg-bp67acfmxazb4p****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The IP address whitelist.
	//
	// 127.0.0.1 indicates that access from any external IP address is prohibited. You can call the [ModifySecurityIps](https://help.aliyun.com/document_detail/86928.html) operation to modify the IP address whitelist after the instance is created.
	//
	// example:
	//
	// 127.0.0.1
	SecurityIPList *string `json:"SecurityIPList,omitempty" xml:"SecurityIPList,omitempty"`
}

func (s CreateModelServiceShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateModelServiceShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateModelServiceShrinkRequest) GetAiNodesShrink() *string {
	return s.AiNodesShrink
}

func (s *CreateModelServiceShrinkRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateModelServiceShrinkRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *CreateModelServiceShrinkRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateModelServiceShrinkRequest) GetEnablePublicConnection() *bool {
	return s.EnablePublicConnection
}

func (s *CreateModelServiceShrinkRequest) GetInferenceEngine() *string {
	return s.InferenceEngine
}

func (s *CreateModelServiceShrinkRequest) GetModelName() *string {
	return s.ModelName
}

func (s *CreateModelServiceShrinkRequest) GetModelParamsShrink() *string {
	return s.ModelParamsShrink
}

func (s *CreateModelServiceShrinkRequest) GetReplicas() *int32 {
	return s.Replicas
}

func (s *CreateModelServiceShrinkRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *CreateModelServiceShrinkRequest) GetSecurityIPList() *string {
	return s.SecurityIPList
}

func (s *CreateModelServiceShrinkRequest) SetAiNodesShrink(v string) *CreateModelServiceShrinkRequest {
	s.AiNodesShrink = &v
	return s
}

func (s *CreateModelServiceShrinkRequest) SetClientToken(v string) *CreateModelServiceShrinkRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateModelServiceShrinkRequest) SetDBInstanceId(v string) *CreateModelServiceShrinkRequest {
	s.DBInstanceId = &v
	return s
}

func (s *CreateModelServiceShrinkRequest) SetDescription(v string) *CreateModelServiceShrinkRequest {
	s.Description = &v
	return s
}

func (s *CreateModelServiceShrinkRequest) SetEnablePublicConnection(v bool) *CreateModelServiceShrinkRequest {
	s.EnablePublicConnection = &v
	return s
}

func (s *CreateModelServiceShrinkRequest) SetInferenceEngine(v string) *CreateModelServiceShrinkRequest {
	s.InferenceEngine = &v
	return s
}

func (s *CreateModelServiceShrinkRequest) SetModelName(v string) *CreateModelServiceShrinkRequest {
	s.ModelName = &v
	return s
}

func (s *CreateModelServiceShrinkRequest) SetModelParamsShrink(v string) *CreateModelServiceShrinkRequest {
	s.ModelParamsShrink = &v
	return s
}

func (s *CreateModelServiceShrinkRequest) SetReplicas(v int32) *CreateModelServiceShrinkRequest {
	s.Replicas = &v
	return s
}

func (s *CreateModelServiceShrinkRequest) SetResourceGroupId(v string) *CreateModelServiceShrinkRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateModelServiceShrinkRequest) SetSecurityIPList(v string) *CreateModelServiceShrinkRequest {
	s.SecurityIPList = &v
	return s
}

func (s *CreateModelServiceShrinkRequest) Validate() error {
	return dara.Validate(s)
}
