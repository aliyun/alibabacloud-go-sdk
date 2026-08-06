// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateCrawlerShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEnableAiComment(v bool) *UpdateCrawlerShrinkRequest
	GetEnableAiComment() *bool
	SetId(v int64) *UpdateCrawlerShrinkRequest
	GetId() *int64
	SetOptionsShrink(v string) *UpdateCrawlerShrinkRequest
	GetOptionsShrink() *string
	SetResourceGroupId(v string) *UpdateCrawlerShrinkRequest
	GetResourceGroupId() *string
	SetScheduleConfigShrink(v string) *UpdateCrawlerShrinkRequest
	GetScheduleConfigShrink() *string
	SetScopeShrink(v string) *UpdateCrawlerShrinkRequest
	GetScopeShrink() *string
}

type UpdateCrawlerShrinkRequest struct {
	EnableAiComment *bool `json:"EnableAiComment,omitempty" xml:"EnableAiComment,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1234
	Id            *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	OptionsShrink *string `json:"Options,omitempty" xml:"Options,omitempty"`
	// example:
	//
	// Serverless_res_group_1234567890123456_1234567890
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ScheduleConfigShrink *string `json:"ScheduleConfig,omitempty" xml:"ScheduleConfig,omitempty"`
	ScopeShrink          *string `json:"Scope,omitempty" xml:"Scope,omitempty"`
}

func (s UpdateCrawlerShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateCrawlerShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateCrawlerShrinkRequest) GetEnableAiComment() *bool {
	return s.EnableAiComment
}

func (s *UpdateCrawlerShrinkRequest) GetId() *int64 {
	return s.Id
}

func (s *UpdateCrawlerShrinkRequest) GetOptionsShrink() *string {
	return s.OptionsShrink
}

func (s *UpdateCrawlerShrinkRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *UpdateCrawlerShrinkRequest) GetScheduleConfigShrink() *string {
	return s.ScheduleConfigShrink
}

func (s *UpdateCrawlerShrinkRequest) GetScopeShrink() *string {
	return s.ScopeShrink
}

func (s *UpdateCrawlerShrinkRequest) SetEnableAiComment(v bool) *UpdateCrawlerShrinkRequest {
	s.EnableAiComment = &v
	return s
}

func (s *UpdateCrawlerShrinkRequest) SetId(v int64) *UpdateCrawlerShrinkRequest {
	s.Id = &v
	return s
}

func (s *UpdateCrawlerShrinkRequest) SetOptionsShrink(v string) *UpdateCrawlerShrinkRequest {
	s.OptionsShrink = &v
	return s
}

func (s *UpdateCrawlerShrinkRequest) SetResourceGroupId(v string) *UpdateCrawlerShrinkRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *UpdateCrawlerShrinkRequest) SetScheduleConfigShrink(v string) *UpdateCrawlerShrinkRequest {
	s.ScheduleConfigShrink = &v
	return s
}

func (s *UpdateCrawlerShrinkRequest) SetScopeShrink(v string) *UpdateCrawlerShrinkRequest {
	s.ScopeShrink = &v
	return s
}

func (s *UpdateCrawlerShrinkRequest) Validate() error {
	return dara.Validate(s)
}
