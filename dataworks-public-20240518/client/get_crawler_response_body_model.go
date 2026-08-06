// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCrawlerResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCrawler(v *GetCrawlerResponseBodyCrawler) *GetCrawlerResponseBody
	GetCrawler() *GetCrawlerResponseBodyCrawler
	SetRequestId(v string) *GetCrawlerResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetCrawlerResponseBody
	GetSuccess() *bool
}

type GetCrawlerResponseBody struct {
	Crawler *GetCrawlerResponseBodyCrawler `json:"Crawler,omitempty" xml:"Crawler,omitempty" type:"Struct"`
	// example:
	//
	// 9252F32F-D855-549E-8898-61CF5A733050
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetCrawlerResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetCrawlerResponseBody) GoString() string {
	return s.String()
}

func (s *GetCrawlerResponseBody) GetCrawler() *GetCrawlerResponseBodyCrawler {
	return s.Crawler
}

func (s *GetCrawlerResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetCrawlerResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetCrawlerResponseBody) SetCrawler(v *GetCrawlerResponseBodyCrawler) *GetCrawlerResponseBody {
	s.Crawler = v
	return s
}

func (s *GetCrawlerResponseBody) SetRequestId(v string) *GetCrawlerResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetCrawlerResponseBody) SetSuccess(v bool) *GetCrawlerResponseBody {
	s.Success = &v
	return s
}

func (s *GetCrawlerResponseBody) Validate() error {
	if s.Crawler != nil {
		if err := s.Crawler.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetCrawlerResponseBodyCrawler struct {
	// example:
	//
	// 1710239005403
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// 12345
	DataSourceId    *int64 `json:"DataSourceId,omitempty" xml:"DataSourceId,omitempty"`
	EnableAiComment *bool  `json:"EnableAiComment,omitempty" xml:"EnableAiComment,omitempty"`
	// example:
	//
	// Prod
	EnvType *string `json:"EnvType,omitempty" xml:"EnvType,omitempty"`
	// example:
	//
	// 1234
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// SUCCESS
	LastRunStatus *string `json:"LastRunStatus,omitempty" xml:"LastRunStatus,omitempty"`
	// example:
	//
	// 1234
	LastRunTaskInstanceId *int64 `json:"LastRunTaskInstanceId,omitempty" xml:"LastRunTaskInstanceId,omitempty"`
	// example:
	//
	// starrocks:example-instance
	MetaEntityId *string `json:"MetaEntityId,omitempty" xml:"MetaEntityId,omitempty"`
	// example:
	//
	// 1710239005403
	ModifyTime *int64 `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	// example:
	//
	// example_crawler
	Name    *string            `json:"Name,omitempty" xml:"Name,omitempty"`
	Options map[string]*string `json:"Options,omitempty" xml:"Options,omitempty"`
	// example:
	//
	// 1000
	Owner *string `json:"Owner,omitempty" xml:"Owner,omitempty"`
	// example:
	//
	// 100
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// example:
	//
	// Serverless_res_group_1234567890123456_1234567890
	ResourceGroupId *string                                      `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ScheduleConfig  *GetCrawlerResponseBodyCrawlerScheduleConfig `json:"ScheduleConfig,omitempty" xml:"ScheduleConfig,omitempty" type:"Struct"`
	Scope           *GetCrawlerResponseBodyCrawlerScope          `json:"Scope,omitempty" xml:"Scope,omitempty" type:"Struct"`
	// example:
	//
	// VALID
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// 1234
	TaskId *int64 `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// example:
	//
	// starrocks
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetCrawlerResponseBodyCrawler) String() string {
	return dara.Prettify(s)
}

func (s GetCrawlerResponseBodyCrawler) GoString() string {
	return s.String()
}

func (s *GetCrawlerResponseBodyCrawler) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *GetCrawlerResponseBodyCrawler) GetDataSourceId() *int64 {
	return s.DataSourceId
}

func (s *GetCrawlerResponseBodyCrawler) GetEnableAiComment() *bool {
	return s.EnableAiComment
}

func (s *GetCrawlerResponseBodyCrawler) GetEnvType() *string {
	return s.EnvType
}

func (s *GetCrawlerResponseBodyCrawler) GetId() *int64 {
	return s.Id
}

func (s *GetCrawlerResponseBodyCrawler) GetLastRunStatus() *string {
	return s.LastRunStatus
}

func (s *GetCrawlerResponseBodyCrawler) GetLastRunTaskInstanceId() *int64 {
	return s.LastRunTaskInstanceId
}

func (s *GetCrawlerResponseBodyCrawler) GetMetaEntityId() *string {
	return s.MetaEntityId
}

func (s *GetCrawlerResponseBodyCrawler) GetModifyTime() *int64 {
	return s.ModifyTime
}

func (s *GetCrawlerResponseBodyCrawler) GetName() *string {
	return s.Name
}

func (s *GetCrawlerResponseBodyCrawler) GetOptions() map[string]*string {
	return s.Options
}

func (s *GetCrawlerResponseBodyCrawler) GetOwner() *string {
	return s.Owner
}

func (s *GetCrawlerResponseBodyCrawler) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *GetCrawlerResponseBodyCrawler) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *GetCrawlerResponseBodyCrawler) GetScheduleConfig() *GetCrawlerResponseBodyCrawlerScheduleConfig {
	return s.ScheduleConfig
}

func (s *GetCrawlerResponseBodyCrawler) GetScope() *GetCrawlerResponseBodyCrawlerScope {
	return s.Scope
}

func (s *GetCrawlerResponseBodyCrawler) GetStatus() *string {
	return s.Status
}

func (s *GetCrawlerResponseBodyCrawler) GetTaskId() *int64 {
	return s.TaskId
}

func (s *GetCrawlerResponseBodyCrawler) GetType() *string {
	return s.Type
}

func (s *GetCrawlerResponseBodyCrawler) SetCreateTime(v int64) *GetCrawlerResponseBodyCrawler {
	s.CreateTime = &v
	return s
}

func (s *GetCrawlerResponseBodyCrawler) SetDataSourceId(v int64) *GetCrawlerResponseBodyCrawler {
	s.DataSourceId = &v
	return s
}

func (s *GetCrawlerResponseBodyCrawler) SetEnableAiComment(v bool) *GetCrawlerResponseBodyCrawler {
	s.EnableAiComment = &v
	return s
}

func (s *GetCrawlerResponseBodyCrawler) SetEnvType(v string) *GetCrawlerResponseBodyCrawler {
	s.EnvType = &v
	return s
}

func (s *GetCrawlerResponseBodyCrawler) SetId(v int64) *GetCrawlerResponseBodyCrawler {
	s.Id = &v
	return s
}

func (s *GetCrawlerResponseBodyCrawler) SetLastRunStatus(v string) *GetCrawlerResponseBodyCrawler {
	s.LastRunStatus = &v
	return s
}

func (s *GetCrawlerResponseBodyCrawler) SetLastRunTaskInstanceId(v int64) *GetCrawlerResponseBodyCrawler {
	s.LastRunTaskInstanceId = &v
	return s
}

func (s *GetCrawlerResponseBodyCrawler) SetMetaEntityId(v string) *GetCrawlerResponseBodyCrawler {
	s.MetaEntityId = &v
	return s
}

func (s *GetCrawlerResponseBodyCrawler) SetModifyTime(v int64) *GetCrawlerResponseBodyCrawler {
	s.ModifyTime = &v
	return s
}

func (s *GetCrawlerResponseBodyCrawler) SetName(v string) *GetCrawlerResponseBodyCrawler {
	s.Name = &v
	return s
}

func (s *GetCrawlerResponseBodyCrawler) SetOptions(v map[string]*string) *GetCrawlerResponseBodyCrawler {
	s.Options = v
	return s
}

func (s *GetCrawlerResponseBodyCrawler) SetOwner(v string) *GetCrawlerResponseBodyCrawler {
	s.Owner = &v
	return s
}

func (s *GetCrawlerResponseBodyCrawler) SetProjectId(v int64) *GetCrawlerResponseBodyCrawler {
	s.ProjectId = &v
	return s
}

func (s *GetCrawlerResponseBodyCrawler) SetResourceGroupId(v string) *GetCrawlerResponseBodyCrawler {
	s.ResourceGroupId = &v
	return s
}

func (s *GetCrawlerResponseBodyCrawler) SetScheduleConfig(v *GetCrawlerResponseBodyCrawlerScheduleConfig) *GetCrawlerResponseBodyCrawler {
	s.ScheduleConfig = v
	return s
}

func (s *GetCrawlerResponseBodyCrawler) SetScope(v *GetCrawlerResponseBodyCrawlerScope) *GetCrawlerResponseBodyCrawler {
	s.Scope = v
	return s
}

func (s *GetCrawlerResponseBodyCrawler) SetStatus(v string) *GetCrawlerResponseBodyCrawler {
	s.Status = &v
	return s
}

func (s *GetCrawlerResponseBodyCrawler) SetTaskId(v int64) *GetCrawlerResponseBodyCrawler {
	s.TaskId = &v
	return s
}

func (s *GetCrawlerResponseBodyCrawler) SetType(v string) *GetCrawlerResponseBodyCrawler {
	s.Type = &v
	return s
}

func (s *GetCrawlerResponseBodyCrawler) Validate() error {
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

type GetCrawlerResponseBodyCrawlerScheduleConfig struct {
	// example:
	//
	// 0 0 2 ? 	- *
	CronExpress *string `json:"CronExpress,omitempty" xml:"CronExpress,omitempty"`
	// example:
	//
	// NORMAL
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetCrawlerResponseBodyCrawlerScheduleConfig) String() string {
	return dara.Prettify(s)
}

func (s GetCrawlerResponseBodyCrawlerScheduleConfig) GoString() string {
	return s.String()
}

func (s *GetCrawlerResponseBodyCrawlerScheduleConfig) GetCronExpress() *string {
	return s.CronExpress
}

func (s *GetCrawlerResponseBodyCrawlerScheduleConfig) GetType() *string {
	return s.Type
}

func (s *GetCrawlerResponseBodyCrawlerScheduleConfig) SetCronExpress(v string) *GetCrawlerResponseBodyCrawlerScheduleConfig {
	s.CronExpress = &v
	return s
}

func (s *GetCrawlerResponseBodyCrawlerScheduleConfig) SetType(v string) *GetCrawlerResponseBodyCrawlerScheduleConfig {
	s.Type = &v
	return s
}

func (s *GetCrawlerResponseBodyCrawlerScheduleConfig) Validate() error {
	return dara.Validate(s)
}

type GetCrawlerResponseBodyCrawlerScope struct {
	// example:
	//
	// ^tmp_.*
	ExcludeRegex *string   `json:"ExcludeRegex,omitempty" xml:"ExcludeRegex,omitempty"`
	Items        []*string `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	// example:
	//
	// DATABASE
	Unit *string `json:"Unit,omitempty" xml:"Unit,omitempty"`
}

func (s GetCrawlerResponseBodyCrawlerScope) String() string {
	return dara.Prettify(s)
}

func (s GetCrawlerResponseBodyCrawlerScope) GoString() string {
	return s.String()
}

func (s *GetCrawlerResponseBodyCrawlerScope) GetExcludeRegex() *string {
	return s.ExcludeRegex
}

func (s *GetCrawlerResponseBodyCrawlerScope) GetItems() []*string {
	return s.Items
}

func (s *GetCrawlerResponseBodyCrawlerScope) GetUnit() *string {
	return s.Unit
}

func (s *GetCrawlerResponseBodyCrawlerScope) SetExcludeRegex(v string) *GetCrawlerResponseBodyCrawlerScope {
	s.ExcludeRegex = &v
	return s
}

func (s *GetCrawlerResponseBodyCrawlerScope) SetItems(v []*string) *GetCrawlerResponseBodyCrawlerScope {
	s.Items = v
	return s
}

func (s *GetCrawlerResponseBodyCrawlerScope) SetUnit(v string) *GetCrawlerResponseBodyCrawlerScope {
	s.Unit = &v
	return s
}

func (s *GetCrawlerResponseBodyCrawlerScope) Validate() error {
	return dara.Validate(s)
}
