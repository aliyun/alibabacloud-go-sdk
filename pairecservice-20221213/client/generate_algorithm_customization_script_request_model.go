// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGenerateAlgorithmCustomizationScriptRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDeployMode(v string) *GenerateAlgorithmCustomizationScriptRequest
	GetDeployMode() *string
	SetInstanceId(v string) *GenerateAlgorithmCustomizationScriptRequest
	GetInstanceId() *string
	SetModuleFieldTypes(v map[string]interface{}) *GenerateAlgorithmCustomizationScriptRequest
	GetModuleFieldTypes() map[string]interface{}
}

type GenerateAlgorithmCustomizationScriptRequest struct {
	// Specifies whether to generate a deployment script or perform one-click deployment. Valid values:
	//
	// - EasyDeploy: one-click deployment
	//
	// - GenerateScript: script generation
	//
	// example:
	//
	// EasyDeploy
	DeployMode *string `json:"DeployMode,omitempty" xml:"DeployMode,omitempty"`
	// The instance ID. For information about how to obtain the instance ID, see [ListInstances](https://help.aliyun.com/document_detail/2411819.html).
	//
	// example:
	//
	// pairec-test-xxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The types of fields in the JSON configuration.
	//
	// example:
	//
	// {"item_table":"array"}
	ModuleFieldTypes map[string]interface{} `json:"ModuleFieldTypes,omitempty" xml:"ModuleFieldTypes,omitempty"`
}

func (s GenerateAlgorithmCustomizationScriptRequest) String() string {
	return dara.Prettify(s)
}

func (s GenerateAlgorithmCustomizationScriptRequest) GoString() string {
	return s.String()
}

func (s *GenerateAlgorithmCustomizationScriptRequest) GetDeployMode() *string {
	return s.DeployMode
}

func (s *GenerateAlgorithmCustomizationScriptRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GenerateAlgorithmCustomizationScriptRequest) GetModuleFieldTypes() map[string]interface{} {
	return s.ModuleFieldTypes
}

func (s *GenerateAlgorithmCustomizationScriptRequest) SetDeployMode(v string) *GenerateAlgorithmCustomizationScriptRequest {
	s.DeployMode = &v
	return s
}

func (s *GenerateAlgorithmCustomizationScriptRequest) SetInstanceId(v string) *GenerateAlgorithmCustomizationScriptRequest {
	s.InstanceId = &v
	return s
}

func (s *GenerateAlgorithmCustomizationScriptRequest) SetModuleFieldTypes(v map[string]interface{}) *GenerateAlgorithmCustomizationScriptRequest {
	s.ModuleFieldTypes = v
	return s
}

func (s *GenerateAlgorithmCustomizationScriptRequest) Validate() error {
	return dara.Validate(s)
}
