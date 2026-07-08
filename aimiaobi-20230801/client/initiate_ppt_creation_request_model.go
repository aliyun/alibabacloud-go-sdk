// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInitiatePptCreationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetExternalUserId(v string) *InitiatePptCreationRequest
	GetExternalUserId() *string
	SetOutline(v string) *InitiatePptCreationRequest
	GetOutline() *string
	SetTaskId(v string) *InitiatePptCreationRequest
	GetTaskId() *string
	SetWorkspaceId(v string) *InitiatePptCreationRequest
	GetWorkspaceId() *string
}

type InitiatePptCreationRequest struct {
	// example:
	//
	// abc
	ExternalUserId *string `json:"ExternalUserId,omitempty" xml:"ExternalUserId,omitempty"`
	// The outline.
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
	// The task ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 95c2fbe6-5a20-4fc2-8a93-376ed05fbe13
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// The unique ID of the Alibaba Cloud Model Studio workspace. For more information, see [Obtain a workspace ID](https://help.aliyun.com/document_detail/2782167.html).
	//
	// example:
	//
	// llm-3fy94b2rtadt01qa
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s InitiatePptCreationRequest) String() string {
	return dara.Prettify(s)
}

func (s InitiatePptCreationRequest) GoString() string {
	return s.String()
}

func (s *InitiatePptCreationRequest) GetExternalUserId() *string {
	return s.ExternalUserId
}

func (s *InitiatePptCreationRequest) GetOutline() *string {
	return s.Outline
}

func (s *InitiatePptCreationRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *InitiatePptCreationRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *InitiatePptCreationRequest) SetExternalUserId(v string) *InitiatePptCreationRequest {
	s.ExternalUserId = &v
	return s
}

func (s *InitiatePptCreationRequest) SetOutline(v string) *InitiatePptCreationRequest {
	s.Outline = &v
	return s
}

func (s *InitiatePptCreationRequest) SetTaskId(v string) *InitiatePptCreationRequest {
	s.TaskId = &v
	return s
}

func (s *InitiatePptCreationRequest) SetWorkspaceId(v string) *InitiatePptCreationRequest {
	s.WorkspaceId = &v
	return s
}

func (s *InitiatePptCreationRequest) Validate() error {
	return dara.Validate(s)
}
