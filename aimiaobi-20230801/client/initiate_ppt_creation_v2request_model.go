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
	// example:
	//
	// abc
	ExternalUserId *string `json:"ExternalUserId,omitempty" xml:"ExternalUserId,omitempty"`
	// example:
	//
	// true
	IsMobile *bool   `json:"IsMobile,omitempty" xml:"IsMobile,omitempty"`
	Outline  *string `json:"Outline,omitempty" xml:"Outline,omitempty"`
	// example:
	//
	// 500
	PptTemplateId *int32 `json:"PptTemplateId,omitempty" xml:"PptTemplateId,omitempty"`
	// example:
	//
	// 1
	PptTemplateType *int32 `json:"PptTemplateType,omitempty" xml:"PptTemplateType,omitempty"`
	// example:
	//
	// 中国传统文化艺术的魅力
	PptTitle *string `json:"PptTitle,omitempty" xml:"PptTitle,omitempty"`
	// example:
	//
	// 1
	ProcessType *int32 `json:"ProcessType,omitempty" xml:"ProcessType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 8a7dfece-f204-4380-a7d0-a13d37de3924
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
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
