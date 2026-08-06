// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateCrawlerRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEnableAiComment(v bool) *UpdateCrawlerRequest
	GetEnableAiComment() *bool
	SetId(v int64) *UpdateCrawlerRequest
	GetId() *int64
	SetOptions(v map[string]*string) *UpdateCrawlerRequest
	GetOptions() map[string]*string
	SetResourceGroupId(v string) *UpdateCrawlerRequest
	GetResourceGroupId() *string
	SetScheduleConfig(v *UpdateCrawlerRequestScheduleConfig) *UpdateCrawlerRequest
	GetScheduleConfig() *UpdateCrawlerRequestScheduleConfig
	SetScope(v *UpdateCrawlerRequestScope) *UpdateCrawlerRequest
	GetScope() *UpdateCrawlerRequestScope
}

type UpdateCrawlerRequest struct {
	EnableAiComment *bool `json:"EnableAiComment,omitempty" xml:"EnableAiComment,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1234
	Id      *int64             `json:"Id,omitempty" xml:"Id,omitempty"`
	Options map[string]*string `json:"Options,omitempty" xml:"Options,omitempty"`
	// example:
	//
	// Serverless_res_group_1234567890123456_1234567890
	ResourceGroupId *string                             `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ScheduleConfig  *UpdateCrawlerRequestScheduleConfig `json:"ScheduleConfig,omitempty" xml:"ScheduleConfig,omitempty" type:"Struct"`
	Scope           *UpdateCrawlerRequestScope          `json:"Scope,omitempty" xml:"Scope,omitempty" type:"Struct"`
}

func (s UpdateCrawlerRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateCrawlerRequest) GoString() string {
	return s.String()
}

func (s *UpdateCrawlerRequest) GetEnableAiComment() *bool {
	return s.EnableAiComment
}

func (s *UpdateCrawlerRequest) GetId() *int64 {
	return s.Id
}

func (s *UpdateCrawlerRequest) GetOptions() map[string]*string {
	return s.Options
}

func (s *UpdateCrawlerRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *UpdateCrawlerRequest) GetScheduleConfig() *UpdateCrawlerRequestScheduleConfig {
	return s.ScheduleConfig
}

func (s *UpdateCrawlerRequest) GetScope() *UpdateCrawlerRequestScope {
	return s.Scope
}

func (s *UpdateCrawlerRequest) SetEnableAiComment(v bool) *UpdateCrawlerRequest {
	s.EnableAiComment = &v
	return s
}

func (s *UpdateCrawlerRequest) SetId(v int64) *UpdateCrawlerRequest {
	s.Id = &v
	return s
}

func (s *UpdateCrawlerRequest) SetOptions(v map[string]*string) *UpdateCrawlerRequest {
	s.Options = v
	return s
}

func (s *UpdateCrawlerRequest) SetResourceGroupId(v string) *UpdateCrawlerRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *UpdateCrawlerRequest) SetScheduleConfig(v *UpdateCrawlerRequestScheduleConfig) *UpdateCrawlerRequest {
	s.ScheduleConfig = v
	return s
}

func (s *UpdateCrawlerRequest) SetScope(v *UpdateCrawlerRequestScope) *UpdateCrawlerRequest {
	s.Scope = v
	return s
}

func (s *UpdateCrawlerRequest) Validate() error {
	if s.ScheduleConfig != nil {
		if err := s.ScheduleConfig.Validate(); err != nil {
			return err
		}
	}
	if s.Scope != nil {
		if err := s.Scope.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateCrawlerRequestScheduleConfig struct {
	// example:
	//
	// 0 0 2 ? 	- *
	CronExpress *string `json:"CronExpress,omitempty" xml:"CronExpress,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// NORMAL
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s UpdateCrawlerRequestScheduleConfig) String() string {
	return dara.Prettify(s)
}

func (s UpdateCrawlerRequestScheduleConfig) GoString() string {
	return s.String()
}

func (s *UpdateCrawlerRequestScheduleConfig) GetCronExpress() *string {
	return s.CronExpress
}

func (s *UpdateCrawlerRequestScheduleConfig) GetType() *string {
	return s.Type
}

func (s *UpdateCrawlerRequestScheduleConfig) SetCronExpress(v string) *UpdateCrawlerRequestScheduleConfig {
	s.CronExpress = &v
	return s
}

func (s *UpdateCrawlerRequestScheduleConfig) SetType(v string) *UpdateCrawlerRequestScheduleConfig {
	s.Type = &v
	return s
}

func (s *UpdateCrawlerRequestScheduleConfig) Validate() error {
	return dara.Validate(s)
}

type UpdateCrawlerRequestScope struct {
	// example:
	//
	// ^tmp_.*
	ExcludeRegex *string   `json:"ExcludeRegex,omitempty" xml:"ExcludeRegex,omitempty"`
	Items        []*string `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// DATABASE
	Unit *string `json:"Unit,omitempty" xml:"Unit,omitempty"`
}

func (s UpdateCrawlerRequestScope) String() string {
	return dara.Prettify(s)
}

func (s UpdateCrawlerRequestScope) GoString() string {
	return s.String()
}

func (s *UpdateCrawlerRequestScope) GetExcludeRegex() *string {
	return s.ExcludeRegex
}

func (s *UpdateCrawlerRequestScope) GetItems() []*string {
	return s.Items
}

func (s *UpdateCrawlerRequestScope) GetUnit() *string {
	return s.Unit
}

func (s *UpdateCrawlerRequestScope) SetExcludeRegex(v string) *UpdateCrawlerRequestScope {
	s.ExcludeRegex = &v
	return s
}

func (s *UpdateCrawlerRequestScope) SetItems(v []*string) *UpdateCrawlerRequestScope {
	s.Items = v
	return s
}

func (s *UpdateCrawlerRequestScope) SetUnit(v string) *UpdateCrawlerRequestScope {
	s.Unit = &v
	return s
}

func (s *UpdateCrawlerRequestScope) Validate() error {
	return dara.Validate(s)
}
