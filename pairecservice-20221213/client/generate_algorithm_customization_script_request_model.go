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
	// example:
	//
	// EasyDeploy
	DeployMode *string `json:"DeployMode,omitempty" xml:"DeployMode,omitempty"`
	// example:
	//
	// pairec-test-xxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
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
