// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyCopilotEmbedConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentName(v string) *ModifyCopilotEmbedConfigRequest
	GetAgentName() *string
	SetCopilotId(v string) *ModifyCopilotEmbedConfigRequest
	GetCopilotId() *string
	SetDataRange(v string) *ModifyCopilotEmbedConfigRequest
	GetDataRange() *string
	SetModuleName(v string) *ModifyCopilotEmbedConfigRequest
	GetModuleName() *string
}

type ModifyCopilotEmbedConfigRequest struct {
	// Agent nickname.
	//
	// example:
	//
	// 小Q
	AgentName *string `json:"AgentName,omitempty" xml:"AgentName,omitempty"`
	// Embedding ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// ccd3428c-dd2xxxxxxxxxxxxdffee
	CopilotId *string `json:"CopilotId,omitempty" xml:"CopilotId,omitempty"`
	// Data range.
	//
	// 	Notice: The parameter type is jsonString, and only one switch between analysis themes and query resources can be effective. When the all-select switch is true, it takes precedence. It is recommended to pass only one parameter, with other notes
	//
	// example:
	//
	// 如果客户要授权所有分析主题，则 {allTheme: true}
	//
	// 如果客户要授权所有问数资源，则 {allCube: true}
	//
	// 如果客户要授权部分问数资源，则 {llmCubes: [12314,12345]}
	//
	// 如果客户要授权部分分析主题，则 {themes: [12314,12345]}
	DataRange *string `json:"DataRange,omitempty" xml:"DataRange,omitempty"`
	// Module name.
	//
	// example:
	//
	// 小Q
	ModuleName *string `json:"ModuleName,omitempty" xml:"ModuleName,omitempty"`
}

func (s ModifyCopilotEmbedConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyCopilotEmbedConfigRequest) GoString() string {
	return s.String()
}

func (s *ModifyCopilotEmbedConfigRequest) GetAgentName() *string {
	return s.AgentName
}

func (s *ModifyCopilotEmbedConfigRequest) GetCopilotId() *string {
	return s.CopilotId
}

func (s *ModifyCopilotEmbedConfigRequest) GetDataRange() *string {
	return s.DataRange
}

func (s *ModifyCopilotEmbedConfigRequest) GetModuleName() *string {
	return s.ModuleName
}

func (s *ModifyCopilotEmbedConfigRequest) SetAgentName(v string) *ModifyCopilotEmbedConfigRequest {
	s.AgentName = &v
	return s
}

func (s *ModifyCopilotEmbedConfigRequest) SetCopilotId(v string) *ModifyCopilotEmbedConfigRequest {
	s.CopilotId = &v
	return s
}

func (s *ModifyCopilotEmbedConfigRequest) SetDataRange(v string) *ModifyCopilotEmbedConfigRequest {
	s.DataRange = &v
	return s
}

func (s *ModifyCopilotEmbedConfigRequest) SetModuleName(v string) *ModifyCopilotEmbedConfigRequest {
	s.ModuleName = &v
	return s
}

func (s *ModifyCopilotEmbedConfigRequest) Validate() error {
	return dara.Validate(s)
}
