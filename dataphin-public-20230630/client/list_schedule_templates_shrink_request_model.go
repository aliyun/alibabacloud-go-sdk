// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListScheduleTemplatesShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetListScheduleTemplatesCommandShrink(v string) *ListScheduleTemplatesShrinkRequest
	GetListScheduleTemplatesCommandShrink() *string
	SetOpTenantId(v int64) *ListScheduleTemplatesShrinkRequest
	GetOpTenantId() *int64
	SetOpUserId(v string) *ListScheduleTemplatesShrinkRequest
	GetOpUserId() *string
}

type ListScheduleTemplatesShrinkRequest struct {
	// This parameter is required.
	ListScheduleTemplatesCommandShrink *string `json:"ListScheduleTemplatesCommand,omitempty" xml:"ListScheduleTemplatesCommand,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
	// example:
	//
	// 30001011
	OpUserId *string `json:"OpUserId,omitempty" xml:"OpUserId,omitempty"`
}

func (s ListScheduleTemplatesShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ListScheduleTemplatesShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListScheduleTemplatesShrinkRequest) GetListScheduleTemplatesCommandShrink() *string {
	return s.ListScheduleTemplatesCommandShrink
}

func (s *ListScheduleTemplatesShrinkRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *ListScheduleTemplatesShrinkRequest) GetOpUserId() *string {
	return s.OpUserId
}

func (s *ListScheduleTemplatesShrinkRequest) SetListScheduleTemplatesCommandShrink(v string) *ListScheduleTemplatesShrinkRequest {
	s.ListScheduleTemplatesCommandShrink = &v
	return s
}

func (s *ListScheduleTemplatesShrinkRequest) SetOpTenantId(v int64) *ListScheduleTemplatesShrinkRequest {
	s.OpTenantId = &v
	return s
}

func (s *ListScheduleTemplatesShrinkRequest) SetOpUserId(v string) *ListScheduleTemplatesShrinkRequest {
	s.OpUserId = &v
	return s
}

func (s *ListScheduleTemplatesShrinkRequest) Validate() error {
	return dara.Validate(s)
}
