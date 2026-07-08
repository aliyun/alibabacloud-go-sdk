// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateGeneralConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConfigKey(v string) *UpdateGeneralConfigRequest
	GetConfigKey() *string
	SetConfigValue(v string) *UpdateGeneralConfigRequest
	GetConfigValue() *string
	SetWorkspaceId(v string) *UpdateGeneralConfigRequest
	GetWorkspaceId() *string
}

type UpdateGeneralConfigRequest struct {
	// The unique identifier of the configuration item. The following configurations are supported:
	//
	// - Text search threshold for data sources (double): \\`searchGenerate.searchTextMinScore\\`
	//
	// - Image search threshold for data sources (double): \\`searchGenerate.searchImageMinScore\\`
	//
	// - Video search threshold for data sources (double): \\`searchGenerate.searchVideoMinScore\\`
	//
	// - Audio search threshold for data sources (double): \\`searchGenerate.searchAudioMinScore\\`
	//
	// - Plain text prompt template for answer summarization in general Q\\&A search (string): \\`searchGenerate.sumQaAgentPrompt\\`
	//
	// - Text and image prompt template for answer summarization in general Q\\&A search (string): \\`searchGenerate.sumQaAgentVlPrompt\\`
	//
	// - Plain text prompt template for answer summarization in enhanced Q\\&A search (string): \\`searchGenerate.sumQaEnhanceAgentPrompt\\`
	//
	// - Text and image prompt template for answer summarization in enhanced Q\\&A search (string): \\`searchGenerate.sumQaEnhanceAgentVlPrompt\\`
	//
	// This parameter is required.
	//
	// example:
	//
	// xx
	ConfigKey *string `json:"ConfigKey,omitempty" xml:"ConfigKey,omitempty"`
	// The value of the configuration item.
	//
	// This parameter is required.
	//
	// example:
	//
	// xxx
	ConfigValue *string `json:"ConfigValue,omitempty" xml:"ConfigValue,omitempty"`
	// The unique identifier of the Model Studio workspace. For more information, see [Get a workspaceId](https://help.aliyun.com/document_detail/2782167.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// xxxx
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s UpdateGeneralConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateGeneralConfigRequest) GoString() string {
	return s.String()
}

func (s *UpdateGeneralConfigRequest) GetConfigKey() *string {
	return s.ConfigKey
}

func (s *UpdateGeneralConfigRequest) GetConfigValue() *string {
	return s.ConfigValue
}

func (s *UpdateGeneralConfigRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *UpdateGeneralConfigRequest) SetConfigKey(v string) *UpdateGeneralConfigRequest {
	s.ConfigKey = &v
	return s
}

func (s *UpdateGeneralConfigRequest) SetConfigValue(v string) *UpdateGeneralConfigRequest {
	s.ConfigValue = &v
	return s
}

func (s *UpdateGeneralConfigRequest) SetWorkspaceId(v string) *UpdateGeneralConfigRequest {
	s.WorkspaceId = &v
	return s
}

func (s *UpdateGeneralConfigRequest) Validate() error {
	return dara.Validate(s)
}
