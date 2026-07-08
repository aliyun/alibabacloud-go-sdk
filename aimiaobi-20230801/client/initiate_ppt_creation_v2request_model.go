// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInitiatePptCreationV2Request interface {
	dara.Model
	String() string
	GoString() string
	SetExternalUserId(v string) *InitiatePptCreationV2Request
	GetExternalUserId() *string
	SetIsMobile(v bool) *InitiatePptCreationV2Request
	GetIsMobile() *bool
	SetOutline(v string) *InitiatePptCreationV2Request
	GetOutline() *string
	SetPptTemplateId(v int32) *InitiatePptCreationV2Request
	GetPptTemplateId() *int32
	SetPptTemplateType(v int32) *InitiatePptCreationV2Request
	GetPptTemplateType() *int32
	SetPptTitle(v string) *InitiatePptCreationV2Request
	GetPptTitle() *string
	SetProcessType(v int32) *InitiatePptCreationV2Request
	GetProcessType() *int32
	SetTaskId(v string) *InitiatePptCreationV2Request
	GetTaskId() *string
	SetWorkspaceId(v string) *InitiatePptCreationV2Request
	GetWorkspaceId() *string
}

type InitiatePptCreationV2Request struct {
	// The unique ID of the external user.
	//
	// example:
	//
	// abc
	ExternalUserId *string `json:"ExternalUserId,omitempty" xml:"ExternalUserId,omitempty"`
	// Specifies whether the request originates from a mobile client.
	//
	// example:
	//
	// true
	IsMobile *bool `json:"IsMobile,omitempty" xml:"IsMobile,omitempty"`
	// The presentation outline, formatted in Markdown.
	//
	// example:
	//
	// # 中国传统文化艺术的魅力
	//
	// ## 1. 传统文化艺术的源远流长
	//
	// ### 1.1 中国古代艺术发展历程
	//
	// #### 1.1.1 古代绘画艺术的演变
	//
	// - 从新石器时代的彩陶绘画到东汉时期帛画的出现，绘画形式不断丰富，展现了古人对美的独特追求。唐代绘画风格多样，吴道子的《送子天王图》线条流畅，色彩绚丽，体现了唐代绘画的高超技艺。
	//
	// #### 1.1.2 书法艺术的传承与创新
	//
	// - 书法从甲骨文到楷书、行书、草书，历经数千年演变，承载着中华文化的深厚内涵。王羲之的《兰亭序》被誉为“天下第一行书”，其笔法精妙，结构严谨，展现了书法艺术的巅峰。
	Outline *string `json:"Outline,omitempty" xml:"Outline,omitempty"`
	// The ID of the PPT template.
	//
	// example:
	//
	// 500
	PptTemplateId *int32 `json:"PptTemplateId,omitempty" xml:"PptTemplateId,omitempty"`
	// The template type. The default value is `1`. Valid values: `1` (system template) and `2` (enterprise template).
	//
	// example:
	//
	// 1
	PptTemplateType *int32 `json:"PptTemplateType,omitempty" xml:"PptTemplateType,omitempty"`
	// example:
	//
	// 中国传统文化艺术的魅力
	PptTitle *string `json:"PptTitle,omitempty" xml:"PptTitle,omitempty"`
	// The type of process to initiate. Valid values:<br>
	//
	// `0`: Generates only a signature to initialize the front-end SDK for the full creation process.<br>
	//
	// `1`: Generates a signature and a process ID. Use this option if you have a custom front-end page for templates before you initialize the SDK.<br>
	//
	// `2`: Generates an artifact ID, which allows for direct editing of the artifact.<br>
	//
	// `3`: Generates an export task ID. You can poll this ID to retrieve the export result.<br><br><br><br>
	//
	// example:
	//
	// 1
	ProcessType *int32 `json:"ProcessType,omitempty" xml:"ProcessType,omitempty"`
	// The ID of the task.
	//
	// This parameter is required.
	//
	// example:
	//
	// 8a7dfece-f204-4380-a7d0-a13d37de3924
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// The ID of the workspace.
	//
	// example:
	//
	// llm-2setzb9x4ewsd
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s InitiatePptCreationV2Request) String() string {
	return dara.Prettify(s)
}

func (s InitiatePptCreationV2Request) GoString() string {
	return s.String()
}

func (s *InitiatePptCreationV2Request) GetExternalUserId() *string {
	return s.ExternalUserId
}

func (s *InitiatePptCreationV2Request) GetIsMobile() *bool {
	return s.IsMobile
}

func (s *InitiatePptCreationV2Request) GetOutline() *string {
	return s.Outline
}

func (s *InitiatePptCreationV2Request) GetPptTemplateId() *int32 {
	return s.PptTemplateId
}

func (s *InitiatePptCreationV2Request) GetPptTemplateType() *int32 {
	return s.PptTemplateType
}

func (s *InitiatePptCreationV2Request) GetPptTitle() *string {
	return s.PptTitle
}

func (s *InitiatePptCreationV2Request) GetProcessType() *int32 {
	return s.ProcessType
}

func (s *InitiatePptCreationV2Request) GetTaskId() *string {
	return s.TaskId
}

func (s *InitiatePptCreationV2Request) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *InitiatePptCreationV2Request) SetExternalUserId(v string) *InitiatePptCreationV2Request {
	s.ExternalUserId = &v
	return s
}

func (s *InitiatePptCreationV2Request) SetIsMobile(v bool) *InitiatePptCreationV2Request {
	s.IsMobile = &v
	return s
}

func (s *InitiatePptCreationV2Request) SetOutline(v string) *InitiatePptCreationV2Request {
	s.Outline = &v
	return s
}

func (s *InitiatePptCreationV2Request) SetPptTemplateId(v int32) *InitiatePptCreationV2Request {
	s.PptTemplateId = &v
	return s
}

func (s *InitiatePptCreationV2Request) SetPptTemplateType(v int32) *InitiatePptCreationV2Request {
	s.PptTemplateType = &v
	return s
}

func (s *InitiatePptCreationV2Request) SetPptTitle(v string) *InitiatePptCreationV2Request {
	s.PptTitle = &v
	return s
}

func (s *InitiatePptCreationV2Request) SetProcessType(v int32) *InitiatePptCreationV2Request {
	s.ProcessType = &v
	return s
}

func (s *InitiatePptCreationV2Request) SetTaskId(v string) *InitiatePptCreationV2Request {
	s.TaskId = &v
	return s
}

func (s *InitiatePptCreationV2Request) SetWorkspaceId(v string) *InitiatePptCreationV2Request {
	s.WorkspaceId = &v
	return s
}

func (s *InitiatePptCreationV2Request) Validate() error {
	return dara.Validate(s)
}
