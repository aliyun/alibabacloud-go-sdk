// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCrawlerRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDataSourceId(v int64) *CreateCrawlerRequest
	GetDataSourceId() *int64
	SetEnableAiComment(v bool) *CreateCrawlerRequest
	GetEnableAiComment() *bool
	SetName(v string) *CreateCrawlerRequest
	GetName() *string
	SetOptions(v map[string]*string) *CreateCrawlerRequest
	GetOptions() map[string]*string
	SetResourceGroupId(v string) *CreateCrawlerRequest
	GetResourceGroupId() *string
	SetScheduleConfig(v *CreateCrawlerRequestScheduleConfig) *CreateCrawlerRequest
	GetScheduleConfig() *CreateCrawlerRequestScheduleConfig
	SetScope(v *CreateCrawlerRequestScope) *CreateCrawlerRequest
	GetScope() *CreateCrawlerRequestScope
	SetType(v string) *CreateCrawlerRequest
	GetType() *string
}

type CreateCrawlerRequest struct {
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
	Name    *string            `json:"Name,omitempty" xml:"Name,omitempty"`
	Options map[string]*string `json:"Options,omitempty" xml:"Options,omitempty"`
	// example:
	//
	// Serverless_res_group_1234567890123456_1234567890
	ResourceGroupId *string                             `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ScheduleConfig  *CreateCrawlerRequestScheduleConfig `json:"ScheduleConfig,omitempty" xml:"ScheduleConfig,omitempty" type:"Struct"`
	Scope           *CreateCrawlerRequestScope          `json:"Scope,omitempty" xml:"Scope,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// starrocks
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s CreateCrawlerRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateCrawlerRequest) GoString() string {
	return s.String()
}

func (s *CreateCrawlerRequest) GetDataSourceId() *int64 {
	return s.DataSourceId
}

func (s *CreateCrawlerRequest) GetEnableAiComment() *bool {
	return s.EnableAiComment
}

func (s *CreateCrawlerRequest) GetName() *string {
	return s.Name
}

func (s *CreateCrawlerRequest) GetOptions() map[string]*string {
	return s.Options
}

func (s *CreateCrawlerRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *CreateCrawlerRequest) GetScheduleConfig() *CreateCrawlerRequestScheduleConfig {
	return s.ScheduleConfig
}

func (s *CreateCrawlerRequest) GetScope() *CreateCrawlerRequestScope {
	return s.Scope
}

func (s *CreateCrawlerRequest) GetType() *string {
	return s.Type
}

func (s *CreateCrawlerRequest) SetDataSourceId(v int64) *CreateCrawlerRequest {
	s.DataSourceId = &v
	return s
}

func (s *CreateCrawlerRequest) SetEnableAiComment(v bool) *CreateCrawlerRequest {
	s.EnableAiComment = &v
	return s
}

func (s *CreateCrawlerRequest) SetName(v string) *CreateCrawlerRequest {
	s.Name = &v
	return s
}

func (s *CreateCrawlerRequest) SetOptions(v map[string]*string) *CreateCrawlerRequest {
	s.Options = v
	return s
}

func (s *CreateCrawlerRequest) SetResourceGroupId(v string) *CreateCrawlerRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateCrawlerRequest) SetScheduleConfig(v *CreateCrawlerRequestScheduleConfig) *CreateCrawlerRequest {
	s.ScheduleConfig = v
	return s
}

func (s *CreateCrawlerRequest) SetScope(v *CreateCrawlerRequestScope) *CreateCrawlerRequest {
	s.Scope = v
	return s
}

func (s *CreateCrawlerRequest) SetType(v string) *CreateCrawlerRequest {
	s.Type = &v
	return s
}

func (s *CreateCrawlerRequest) Validate() error {
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

type CreateCrawlerRequestScheduleConfig struct {
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

func (s CreateCrawlerRequestScheduleConfig) String() string {
	return dara.Prettify(s)
}

func (s CreateCrawlerRequestScheduleConfig) GoString() string {
	return s.String()
}

func (s *CreateCrawlerRequestScheduleConfig) GetCronExpress() *string {
	return s.CronExpress
}

func (s *CreateCrawlerRequestScheduleConfig) GetType() *string {
	return s.Type
}

func (s *CreateCrawlerRequestScheduleConfig) SetCronExpress(v string) *CreateCrawlerRequestScheduleConfig {
	s.CronExpress = &v
	return s
}

func (s *CreateCrawlerRequestScheduleConfig) SetType(v string) *CreateCrawlerRequestScheduleConfig {
	s.Type = &v
	return s
}

func (s *CreateCrawlerRequestScheduleConfig) Validate() error {
	return dara.Validate(s)
}

type CreateCrawlerRequestScope struct {
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

func (s CreateCrawlerRequestScope) String() string {
	return dara.Prettify(s)
}

func (s CreateCrawlerRequestScope) GoString() string {
	return s.String()
}

func (s *CreateCrawlerRequestScope) GetExcludeRegex() *string {
	return s.ExcludeRegex
}

func (s *CreateCrawlerRequestScope) GetItems() []*string {
	return s.Items
}

func (s *CreateCrawlerRequestScope) GetUnit() *string {
	return s.Unit
}

func (s *CreateCrawlerRequestScope) SetExcludeRegex(v string) *CreateCrawlerRequestScope {
	s.ExcludeRegex = &v
	return s
}

func (s *CreateCrawlerRequestScope) SetItems(v []*string) *CreateCrawlerRequestScope {
	s.Items = v
	return s
}

func (s *CreateCrawlerRequestScope) SetUnit(v string) *CreateCrawlerRequestScope {
	s.Unit = &v
	return s
}

func (s *CreateCrawlerRequestScope) Validate() error {
	return dara.Validate(s)
}
