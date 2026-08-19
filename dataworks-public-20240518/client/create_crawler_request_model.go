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
	// The ID of the data source associated with the crawler. The data source must be bound to a DataWorks workspace, and the data source type must match the Type value.
	//
	// This parameter is required.
	//
	// example:
	//
	// 12345
	DataSourceId *int64 `json:"DataSourceId,omitempty" xml:"DataSourceId,omitempty"`
	// Specifies whether to enable AI metadata descriptions. This parameter is supported only when the SupportAiComment value returned by GetCrawlerTypeCapabilities is true.
	EnableAiComment *bool `json:"EnableAiComment,omitempty" xml:"EnableAiComment,omitempty"`
	// The name of the metadata crawler. The name can be up to 128 characters in length.
	//
	// This parameter is required.
	//
	// example:
	//
	// example_crawler
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The extended configuration for the crawler type. The key names, value types, required fields, default values, and valid values are determined by the SupportedOptionKeys value returned by GetCrawlerTypeCapabilities.
	Options map[string]*string `json:"Options,omitempty" xml:"Options,omitempty"`
	// The ID of the Serverless 2.0 resource group used to run the collection task. Whether this parameter is required depends on the RequireResourceGroup value returned by GetCrawlerTypeCapabilities.
	//
	// example:
	//
	// Serverless_res_group_1234567890123456_1234567890
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The scheduling configuration. If this parameter is not specified, manual scheduling is used.
	ScheduleConfig *CreateCrawlerRequestScheduleConfig `json:"ScheduleConfig,omitempty" xml:"ScheduleConfig,omitempty" type:"Struct"`
	// The collection scope configuration. If this parameter is not specified, the DefaultScopeUnit value returned by GetCrawlerTypeCapabilities is used.
	Scope *CreateCrawlerRequestScope `json:"Scope,omitempty" xml:"Scope,omitempty" type:"Struct"`
	// The crawler type. Call GetCrawlerTypeCapabilities to query the valid values supported in the current region.
	//
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
	// The six-field cron expression for periodic scheduling. This parameter is required when Type is set to NORMAL. The seconds field must be 0, and the scheduling frequency cannot exceed once per hour.
	//
	// example:
	//
	// 0 0 2 ? 	- *
	CronExpress *string `json:"CronExpress,omitempty" xml:"CronExpress,omitempty"`
	// The scheduling type. MANUAL indicates manual execution, and NORMAL indicates periodic scheduling. Data sources in the development environment support only MANUAL. Whether NORMAL is available depends on the SupportSchedule value returned by GetCrawlerTypeCapabilities.
	//
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
	// The regular expression used to exclude objects from the collection scope. This parameter is supported only when the SupportExcludeRegex value returned by GetCrawlerTypeCapabilities is true.
	//
	// example:
	//
	// ^tmp_.*
	ExcludeRegex *string `json:"ExcludeRegex,omitempty" xml:"ExcludeRegex,omitempty"`
	// The list of database names. This parameter is supported only when Unit is set to DATABASE. A maximum of 1000 entries are allowed. Names cannot be empty or duplicated.
	Items []*string `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	// The collection scope granularity. Valid values are determined by the SupportedScopeUnits value returned by GetCrawlerTypeCapabilities.
	//
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
