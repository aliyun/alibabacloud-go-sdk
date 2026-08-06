// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCrawlerShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDataSourceId(v int64) *CreateCrawlerShrinkRequest
	GetDataSourceId() *int64
	SetEnableAiComment(v bool) *CreateCrawlerShrinkRequest
	GetEnableAiComment() *bool
	SetName(v string) *CreateCrawlerShrinkRequest
	GetName() *string
	SetOptionsShrink(v string) *CreateCrawlerShrinkRequest
	GetOptionsShrink() *string
	SetResourceGroupId(v string) *CreateCrawlerShrinkRequest
	GetResourceGroupId() *string
	SetScheduleConfigShrink(v string) *CreateCrawlerShrinkRequest
	GetScheduleConfigShrink() *string
	SetScopeShrink(v string) *CreateCrawlerShrinkRequest
	GetScopeShrink() *string
	SetType(v string) *CreateCrawlerShrinkRequest
	GetType() *string
}

type CreateCrawlerShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 12345
	DataSourceId    *int64 `json:"DataSourceId,omitempty" xml:"DataSourceId,omitempty"`
	EnableAiComment *bool  `json:"EnableAiComment,omitempty" xml:"EnableAiComment,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// example_crawler
	Name          *string `json:"Name,omitempty" xml:"Name,omitempty"`
	OptionsShrink *string `json:"Options,omitempty" xml:"Options,omitempty"`
	// example:
	//
	// Serverless_res_group_1234567890123456_1234567890
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ScheduleConfigShrink *string `json:"ScheduleConfig,omitempty" xml:"ScheduleConfig,omitempty"`
	ScopeShrink          *string `json:"Scope,omitempty" xml:"Scope,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// starrocks
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s CreateCrawlerShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateCrawlerShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateCrawlerShrinkRequest) GetDataSourceId() *int64 {
	return s.DataSourceId
}

func (s *CreateCrawlerShrinkRequest) GetEnableAiComment() *bool {
	return s.EnableAiComment
}

func (s *CreateCrawlerShrinkRequest) GetName() *string {
	return s.Name
}

func (s *CreateCrawlerShrinkRequest) GetOptionsShrink() *string {
	return s.OptionsShrink
}

func (s *CreateCrawlerShrinkRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *CreateCrawlerShrinkRequest) GetScheduleConfigShrink() *string {
	return s.ScheduleConfigShrink
}

func (s *CreateCrawlerShrinkRequest) GetScopeShrink() *string {
	return s.ScopeShrink
}

func (s *CreateCrawlerShrinkRequest) GetType() *string {
	return s.Type
}

func (s *CreateCrawlerShrinkRequest) SetDataSourceId(v int64) *CreateCrawlerShrinkRequest {
	s.DataSourceId = &v
	return s
}

func (s *CreateCrawlerShrinkRequest) SetEnableAiComment(v bool) *CreateCrawlerShrinkRequest {
	s.EnableAiComment = &v
	return s
}

func (s *CreateCrawlerShrinkRequest) SetName(v string) *CreateCrawlerShrinkRequest {
	s.Name = &v
	return s
}

func (s *CreateCrawlerShrinkRequest) SetOptionsShrink(v string) *CreateCrawlerShrinkRequest {
	s.OptionsShrink = &v
	return s
}

func (s *CreateCrawlerShrinkRequest) SetResourceGroupId(v string) *CreateCrawlerShrinkRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateCrawlerShrinkRequest) SetScheduleConfigShrink(v string) *CreateCrawlerShrinkRequest {
	s.ScheduleConfigShrink = &v
	return s
}

func (s *CreateCrawlerShrinkRequest) SetScopeShrink(v string) *CreateCrawlerShrinkRequest {
	s.ScopeShrink = &v
	return s
}

func (s *CreateCrawlerShrinkRequest) SetType(v string) *CreateCrawlerShrinkRequest {
	s.Type = &v
	return s
}

func (s *CreateCrawlerShrinkRequest) Validate() error {
	return dara.Validate(s)
}
