// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitDeepWriteTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentOrchestration(v *SubmitDeepWriteTaskRequestAgentOrchestration) *SubmitDeepWriteTaskRequest
	GetAgentOrchestration() *SubmitDeepWriteTaskRequestAgentOrchestration
	SetFiles(v []*SubmitDeepWriteTaskRequestFiles) *SubmitDeepWriteTaskRequest
	GetFiles() []*SubmitDeepWriteTaskRequestFiles
	SetInput(v string) *SubmitDeepWriteTaskRequest
	GetInput() *string
	SetInstructions(v string) *SubmitDeepWriteTaskRequest
	GetInstructions() *string
	SetWorkspaceId(v string) *SubmitDeepWriteTaskRequest
	GetWorkspaceId() *string
}

type SubmitDeepWriteTaskRequest struct {
	// The agent orchestration options.
	AgentOrchestration *SubmitDeepWriteTaskRequestAgentOrchestration `json:"AgentOrchestration,omitempty" xml:"AgentOrchestration,omitempty" type:"Struct"`
	// A list of attachments.
	Files []*SubmitDeepWriteTaskRequestFiles `json:"Files,omitempty" xml:"Files,omitempty" type:"Repeated"`
	// The user\\"s question.
	//
	// This parameter is required.
	//
	// if can be null:
	// false
	//
	// example:
	//
	// 北京2025年新能源汽车发展趋势
	Input *string `json:"Input,omitempty" xml:"Input,omitempty"`
	// The instructions.
	//
	// example:
	//
	// 请根据北京新能源汽车在汽车品牌、新车发布、能源等方面进行分析
	Instructions *string `json:"Instructions,omitempty" xml:"Instructions,omitempty"`
	// [The workspace ID.](https://help.aliyun.com/document_detail/2782167.html)
	//
	// example:
	//
	// llm-1setzb9xb8m11vrc
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s SubmitDeepWriteTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s SubmitDeepWriteTaskRequest) GoString() string {
	return s.String()
}

func (s *SubmitDeepWriteTaskRequest) GetAgentOrchestration() *SubmitDeepWriteTaskRequestAgentOrchestration {
	return s.AgentOrchestration
}

func (s *SubmitDeepWriteTaskRequest) GetFiles() []*SubmitDeepWriteTaskRequestFiles {
	return s.Files
}

func (s *SubmitDeepWriteTaskRequest) GetInput() *string {
	return s.Input
}

func (s *SubmitDeepWriteTaskRequest) GetInstructions() *string {
	return s.Instructions
}

func (s *SubmitDeepWriteTaskRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *SubmitDeepWriteTaskRequest) SetAgentOrchestration(v *SubmitDeepWriteTaskRequestAgentOrchestration) *SubmitDeepWriteTaskRequest {
	s.AgentOrchestration = v
	return s
}

func (s *SubmitDeepWriteTaskRequest) SetFiles(v []*SubmitDeepWriteTaskRequestFiles) *SubmitDeepWriteTaskRequest {
	s.Files = v
	return s
}

func (s *SubmitDeepWriteTaskRequest) SetInput(v string) *SubmitDeepWriteTaskRequest {
	s.Input = &v
	return s
}

func (s *SubmitDeepWriteTaskRequest) SetInstructions(v string) *SubmitDeepWriteTaskRequest {
	s.Instructions = &v
	return s
}

func (s *SubmitDeepWriteTaskRequest) SetWorkspaceId(v string) *SubmitDeepWriteTaskRequest {
	s.WorkspaceId = &v
	return s
}

func (s *SubmitDeepWriteTaskRequest) Validate() error {
	if s.AgentOrchestration != nil {
		if err := s.AgentOrchestration.Validate(); err != nil {
			return err
		}
	}
	if s.Files != nil {
		for _, item := range s.Files {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type SubmitDeepWriteTaskRequestAgentOrchestration struct {
	// The data analysis agent.
	DataAnalystAgent *SubmitDeepWriteTaskRequestAgentOrchestrationDataAnalystAgent `json:"DataAnalystAgent,omitempty" xml:"DataAnalystAgent,omitempty" type:"Struct"`
	// The data collection agent.
	DataCollectorAgent *SubmitDeepWriteTaskRequestAgentOrchestrationDataCollectorAgent `json:"DataCollectorAgent,omitempty" xml:"DataCollectorAgent,omitempty" type:"Struct"`
	// The reporter agent.
	ReporterAgent *SubmitDeepWriteTaskRequestAgentOrchestrationReporterAgent `json:"ReporterAgent,omitempty" xml:"ReporterAgent,omitempty" type:"Struct"`
}

func (s SubmitDeepWriteTaskRequestAgentOrchestration) String() string {
	return dara.Prettify(s)
}

func (s SubmitDeepWriteTaskRequestAgentOrchestration) GoString() string {
	return s.String()
}

func (s *SubmitDeepWriteTaskRequestAgentOrchestration) GetDataAnalystAgent() *SubmitDeepWriteTaskRequestAgentOrchestrationDataAnalystAgent {
	return s.DataAnalystAgent
}

func (s *SubmitDeepWriteTaskRequestAgentOrchestration) GetDataCollectorAgent() *SubmitDeepWriteTaskRequestAgentOrchestrationDataCollectorAgent {
	return s.DataCollectorAgent
}

func (s *SubmitDeepWriteTaskRequestAgentOrchestration) GetReporterAgent() *SubmitDeepWriteTaskRequestAgentOrchestrationReporterAgent {
	return s.ReporterAgent
}

func (s *SubmitDeepWriteTaskRequestAgentOrchestration) SetDataAnalystAgent(v *SubmitDeepWriteTaskRequestAgentOrchestrationDataAnalystAgent) *SubmitDeepWriteTaskRequestAgentOrchestration {
	s.DataAnalystAgent = v
	return s
}

func (s *SubmitDeepWriteTaskRequestAgentOrchestration) SetDataCollectorAgent(v *SubmitDeepWriteTaskRequestAgentOrchestrationDataCollectorAgent) *SubmitDeepWriteTaskRequestAgentOrchestration {
	s.DataCollectorAgent = v
	return s
}

func (s *SubmitDeepWriteTaskRequestAgentOrchestration) SetReporterAgent(v *SubmitDeepWriteTaskRequestAgentOrchestrationReporterAgent) *SubmitDeepWriteTaskRequestAgentOrchestration {
	s.ReporterAgent = v
	return s
}

func (s *SubmitDeepWriteTaskRequestAgentOrchestration) Validate() error {
	if s.DataAnalystAgent != nil {
		if err := s.DataAnalystAgent.Validate(); err != nil {
			return err
		}
	}
	if s.DataCollectorAgent != nil {
		if err := s.DataCollectorAgent.Validate(); err != nil {
			return err
		}
	}
	if s.ReporterAgent != nil {
		if err := s.ReporterAgent.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type SubmitDeepWriteTaskRequestAgentOrchestrationDataAnalystAgent struct {
	// Specifies whether to enable retrieval.
	//
	// example:
	//
	// true
	EnableSearch *bool `json:"EnableSearch,omitempty" xml:"EnableSearch,omitempty"`
	// The name.
	//
	// example:
	//
	// DataAnalystAgent
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s SubmitDeepWriteTaskRequestAgentOrchestrationDataAnalystAgent) String() string {
	return dara.Prettify(s)
}

func (s SubmitDeepWriteTaskRequestAgentOrchestrationDataAnalystAgent) GoString() string {
	return s.String()
}

func (s *SubmitDeepWriteTaskRequestAgentOrchestrationDataAnalystAgent) GetEnableSearch() *bool {
	return s.EnableSearch
}

func (s *SubmitDeepWriteTaskRequestAgentOrchestrationDataAnalystAgent) GetName() *string {
	return s.Name
}

func (s *SubmitDeepWriteTaskRequestAgentOrchestrationDataAnalystAgent) SetEnableSearch(v bool) *SubmitDeepWriteTaskRequestAgentOrchestrationDataAnalystAgent {
	s.EnableSearch = &v
	return s
}

func (s *SubmitDeepWriteTaskRequestAgentOrchestrationDataAnalystAgent) SetName(v string) *SubmitDeepWriteTaskRequestAgentOrchestrationDataAnalystAgent {
	s.Name = &v
	return s
}

func (s *SubmitDeepWriteTaskRequestAgentOrchestrationDataAnalystAgent) Validate() error {
	return dara.Validate(s)
}

type SubmitDeepWriteTaskRequestAgentOrchestrationDataCollectorAgent struct {
	// The name.
	//
	// example:
	//
	// DataCollectorAgent
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s SubmitDeepWriteTaskRequestAgentOrchestrationDataCollectorAgent) String() string {
	return dara.Prettify(s)
}

func (s SubmitDeepWriteTaskRequestAgentOrchestrationDataCollectorAgent) GoString() string {
	return s.String()
}

func (s *SubmitDeepWriteTaskRequestAgentOrchestrationDataCollectorAgent) GetName() *string {
	return s.Name
}

func (s *SubmitDeepWriteTaskRequestAgentOrchestrationDataCollectorAgent) SetName(v string) *SubmitDeepWriteTaskRequestAgentOrchestrationDataCollectorAgent {
	s.Name = &v
	return s
}

func (s *SubmitDeepWriteTaskRequestAgentOrchestrationDataCollectorAgent) Validate() error {
	return dara.Validate(s)
}

type SubmitDeepWriteTaskRequestAgentOrchestrationReporterAgent struct {
	// Specifies whether to enable citations.
	EnableCitation *bool `json:"EnableCitation,omitempty" xml:"EnableCitation,omitempty"`
	// The name.
	//
	// example:
	//
	// ReporterAgent
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s SubmitDeepWriteTaskRequestAgentOrchestrationReporterAgent) String() string {
	return dara.Prettify(s)
}

func (s SubmitDeepWriteTaskRequestAgentOrchestrationReporterAgent) GoString() string {
	return s.String()
}

func (s *SubmitDeepWriteTaskRequestAgentOrchestrationReporterAgent) GetEnableCitation() *bool {
	return s.EnableCitation
}

func (s *SubmitDeepWriteTaskRequestAgentOrchestrationReporterAgent) GetName() *string {
	return s.Name
}

func (s *SubmitDeepWriteTaskRequestAgentOrchestrationReporterAgent) SetEnableCitation(v bool) *SubmitDeepWriteTaskRequestAgentOrchestrationReporterAgent {
	s.EnableCitation = &v
	return s
}

func (s *SubmitDeepWriteTaskRequestAgentOrchestrationReporterAgent) SetName(v string) *SubmitDeepWriteTaskRequestAgentOrchestrationReporterAgent {
	s.Name = &v
	return s
}

func (s *SubmitDeepWriteTaskRequestAgentOrchestrationReporterAgent) Validate() error {
	return dara.Validate(s)
}

type SubmitDeepWriteTaskRequestFiles struct {
	// A description of the attachment.
	//
	// example:
	//
	// 附件的说明
	FileDescription *string `json:"FileDescription,omitempty" xml:"FileDescription,omitempty"`
	// The Object Storage Service (OSS) address of the attachment.
	//
	// example:
	//
	// oss://default/aimiaobi-poc/aimiaobi/deep-write-upload/1_1/xxx.txt
	FileKey *string `json:"FileKey,omitempty" xml:"FileKey,omitempty"`
	// The name of the attachment.
	//
	// example:
	//
	// 附件的名称
	FileName *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
}

func (s SubmitDeepWriteTaskRequestFiles) String() string {
	return dara.Prettify(s)
}

func (s SubmitDeepWriteTaskRequestFiles) GoString() string {
	return s.String()
}

func (s *SubmitDeepWriteTaskRequestFiles) GetFileDescription() *string {
	return s.FileDescription
}

func (s *SubmitDeepWriteTaskRequestFiles) GetFileKey() *string {
	return s.FileKey
}

func (s *SubmitDeepWriteTaskRequestFiles) GetFileName() *string {
	return s.FileName
}

func (s *SubmitDeepWriteTaskRequestFiles) SetFileDescription(v string) *SubmitDeepWriteTaskRequestFiles {
	s.FileDescription = &v
	return s
}

func (s *SubmitDeepWriteTaskRequestFiles) SetFileKey(v string) *SubmitDeepWriteTaskRequestFiles {
	s.FileKey = &v
	return s
}

func (s *SubmitDeepWriteTaskRequestFiles) SetFileName(v string) *SubmitDeepWriteTaskRequestFiles {
	s.FileName = &v
	return s
}

func (s *SubmitDeepWriteTaskRequestFiles) Validate() error {
	return dara.Validate(s)
}
