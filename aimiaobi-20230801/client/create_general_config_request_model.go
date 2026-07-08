// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateGeneralConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConfigKey(v string) *CreateGeneralConfigRequest
	GetConfigKey() *string
	SetConfigValue(v string) *CreateGeneralConfigRequest
	GetConfigValue() *string
	SetWorkspaceId(v string) *CreateGeneralConfigRequest
	GetWorkspaceId() *string
}

type CreateGeneralConfigRequest struct {
	// Unique identifier of the configuration item. Supported keys include the following:
	//
	// - MiaoSou text search threshold (double): searchGenerate.searchTextMinScore
	//
	// - MiaoSou image search threshold (double): searchGenerate.searchImageMinScore
	//
	// - MiaoSou video search threshold (double): searchGenerate.searchVideoMinScore
	//
	// - MiaoSou audio search threshold (double): searchGenerate.searchAudioMinScore
	//
	// - MiaoSou Q\\&A search general answer summary prompt template (string): searchGenerate.sumQaAgentPrompt
	//
	// - MiaoSou Q\\&A search general answer summary prompt template with text and images (string): searchGenerate.sumQaAgentVlPrompt
	//
	// - MiaoSou Q\\&A search deep answer summary prompt template (string): searchGenerate.sumQaEnhanceAgentPrompt
	//
	// - MiaoSou Q\\&A search deep answer summary prompt template with text and images (string): searchGenerate.sumQaEnhanceAgentVlPrompt
	//
	// This parameter is required.
	//
	// example:
	//
	// xxxx
	ConfigKey *string `json:"ConfigKey,omitempty" xml:"ConfigKey,omitempty"`
	// Value of the configuration item
	//
	// This parameter is required.
	//
	// example:
	//
	// xxx
	ConfigValue *string `json:"ConfigValue,omitempty" xml:"ConfigValue,omitempty"`
	// Unique identifier of the Model Studio workspace. [Get the workspace ID](https://help.aliyun.com/document_detail/2782167.html)
	//
	// This parameter is required.
	//
	// example:
	//
	// llm-
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s CreateGeneralConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateGeneralConfigRequest) GoString() string {
	return s.String()
}

func (s *CreateGeneralConfigRequest) GetConfigKey() *string {
	return s.ConfigKey
}

func (s *CreateGeneralConfigRequest) GetConfigValue() *string {
	return s.ConfigValue
}

func (s *CreateGeneralConfigRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *CreateGeneralConfigRequest) SetConfigKey(v string) *CreateGeneralConfigRequest {
	s.ConfigKey = &v
	return s
}

func (s *CreateGeneralConfigRequest) SetConfigValue(v string) *CreateGeneralConfigRequest {
	s.ConfigValue = &v
	return s
}

func (s *CreateGeneralConfigRequest) SetWorkspaceId(v string) *CreateGeneralConfigRequest {
	s.WorkspaceId = &v
	return s
}

func (s *CreateGeneralConfigRequest) Validate() error {
	return dara.Validate(s)
}
