// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCrawlersResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPagingInfo(v *ListCrawlersResponseBodyPagingInfo) *ListCrawlersResponseBody
	GetPagingInfo() *ListCrawlersResponseBodyPagingInfo
	SetRequestId(v string) *ListCrawlersResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListCrawlersResponseBody
	GetSuccess() *bool
}

type ListCrawlersResponseBody struct {
	// The pagination information.
	PagingInfo *ListCrawlersResponseBodyPagingInfo `json:"PagingInfo,omitempty" xml:"PagingInfo,omitempty" type:"Struct"`
	// The request ID. Used to locate logs and troubleshoot issues.
	//
	// example:
	//
	// 9252F32F-D855-549E-8898-61CF5A733050
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request succeeded.
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListCrawlersResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListCrawlersResponseBody) GoString() string {
	return s.String()
}

func (s *ListCrawlersResponseBody) GetPagingInfo() *ListCrawlersResponseBodyPagingInfo {
	return s.PagingInfo
}

func (s *ListCrawlersResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListCrawlersResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListCrawlersResponseBody) SetPagingInfo(v *ListCrawlersResponseBodyPagingInfo) *ListCrawlersResponseBody {
	s.PagingInfo = v
	return s
}

func (s *ListCrawlersResponseBody) SetRequestId(v string) *ListCrawlersResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListCrawlersResponseBody) SetSuccess(v bool) *ListCrawlersResponseBody {
	s.Success = &v
	return s
}

func (s *ListCrawlersResponseBody) Validate() error {
	if s.PagingInfo != nil {
		if err := s.PagingInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListCrawlersResponseBodyPagingInfo struct {
	// The list of metadata crawlers.
	Crawlers []*ListCrawlersResponseBodyPagingInfoCrawlers `json:"Crawlers,omitempty" xml:"Crawlers,omitempty" type:"Repeated"`
	// The current page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The total number of records that match the query conditions.
	//
	// example:
	//
	// 1
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListCrawlersResponseBodyPagingInfo) String() string {
	return dara.Prettify(s)
}

func (s ListCrawlersResponseBodyPagingInfo) GoString() string {
	return s.String()
}

func (s *ListCrawlersResponseBodyPagingInfo) GetCrawlers() []*ListCrawlersResponseBodyPagingInfoCrawlers {
	return s.Crawlers
}

func (s *ListCrawlersResponseBodyPagingInfo) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListCrawlersResponseBodyPagingInfo) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListCrawlersResponseBodyPagingInfo) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListCrawlersResponseBodyPagingInfo) SetCrawlers(v []*ListCrawlersResponseBodyPagingInfoCrawlers) *ListCrawlersResponseBodyPagingInfo {
	s.Crawlers = v
	return s
}

func (s *ListCrawlersResponseBodyPagingInfo) SetPageNumber(v int32) *ListCrawlersResponseBodyPagingInfo {
	s.PageNumber = &v
	return s
}

func (s *ListCrawlersResponseBodyPagingInfo) SetPageSize(v int32) *ListCrawlersResponseBodyPagingInfo {
	s.PageSize = &v
	return s
}

func (s *ListCrawlersResponseBodyPagingInfo) SetTotalCount(v int64) *ListCrawlersResponseBodyPagingInfo {
	s.TotalCount = &v
	return s
}

func (s *ListCrawlersResponseBodyPagingInfo) Validate() error {
	if s.Crawlers != nil {
		for _, item := range s.Crawlers {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListCrawlersResponseBodyPagingInfoCrawlers struct {
	// The creation time, expressed as a millisecond-precision UNIX timestamp.
	//
	// example:
	//
	// 1710239005403
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The data source ID.
	//
	// example:
	//
	// 12345
	DataSourceId *int64 `json:"DataSourceId,omitempty" xml:"DataSourceId,omitempty"`
	// The DataWorks environment type. Valid values: Dev, Prod.
	//
	// example:
	//
	// Prod
	EnvType *string `json:"EnvType,omitempty" xml:"EnvType,omitempty"`
	// The crawler ID.
	//
	// example:
	//
	// 1234
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The most recent run status. Valid values: WAITING, RUNNING, SUCCESS, ERROR, SHUTDOWN. This field may be empty if the crawler has not run yet.
	//
	// example:
	//
	// SUCCESS
	LastRunStatus *string `json:"LastRunStatus,omitempty" xml:"LastRunStatus,omitempty"`
	// The meta entity ID associated with the crawler. You can use this ID to connect to metadata query APIs.
	//
	// example:
	//
	// starrocks:example-instance
	MetaEntityId *string `json:"MetaEntityId,omitempty" xml:"MetaEntityId,omitempty"`
	// The modification time, expressed as a millisecond-precision UNIX timestamp.
	//
	// example:
	//
	// 1710239005403
	ModifyTime *int64 `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	// The crawler name.
	//
	// example:
	//
	// example_crawler
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The DataWorks user ID of the crawler owner.
	//
	// example:
	//
	// 1000
	Owner *string `json:"Owner,omitempty" xml:"Owner,omitempty"`
	// The DataWorks workspace ID.
	//
	// example:
	//
	// 100
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The ID of the Serverless 2.0 resource group used to run the crawl task.
	//
	// example:
	//
	// Serverless_res_group_1234567890123456_1234567890
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The schedule configuration.
	ScheduleConfig *ListCrawlersResponseBodyPagingInfoCrawlersScheduleConfig `json:"ScheduleConfig,omitempty" xml:"ScheduleConfig,omitempty" type:"Struct"`
	// The crawler status. The value is VALID when the crawler configuration is valid and the associated data source exists. Otherwise, the value is INVALID.
	//
	// example:
	//
	// VALID
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The DataWorks scheduling node ID associated with the crawler. You can use this ID to call GetTask to query the node definition.
	//
	// example:
	//
	// 1234
	TaskId *int64 `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// The crawler type.
	//
	// example:
	//
	// starrocks
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListCrawlersResponseBodyPagingInfoCrawlers) String() string {
	return dara.Prettify(s)
}

func (s ListCrawlersResponseBodyPagingInfoCrawlers) GoString() string {
	return s.String()
}

func (s *ListCrawlersResponseBodyPagingInfoCrawlers) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *ListCrawlersResponseBodyPagingInfoCrawlers) GetDataSourceId() *int64 {
	return s.DataSourceId
}

func (s *ListCrawlersResponseBodyPagingInfoCrawlers) GetEnvType() *string {
	return s.EnvType
}

func (s *ListCrawlersResponseBodyPagingInfoCrawlers) GetId() *int64 {
	return s.Id
}

func (s *ListCrawlersResponseBodyPagingInfoCrawlers) GetLastRunStatus() *string {
	return s.LastRunStatus
}

func (s *ListCrawlersResponseBodyPagingInfoCrawlers) GetMetaEntityId() *string {
	return s.MetaEntityId
}

func (s *ListCrawlersResponseBodyPagingInfoCrawlers) GetModifyTime() *int64 {
	return s.ModifyTime
}

func (s *ListCrawlersResponseBodyPagingInfoCrawlers) GetName() *string {
	return s.Name
}

func (s *ListCrawlersResponseBodyPagingInfoCrawlers) GetOwner() *string {
	return s.Owner
}

func (s *ListCrawlersResponseBodyPagingInfoCrawlers) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *ListCrawlersResponseBodyPagingInfoCrawlers) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ListCrawlersResponseBodyPagingInfoCrawlers) GetScheduleConfig() *ListCrawlersResponseBodyPagingInfoCrawlersScheduleConfig {
	return s.ScheduleConfig
}

func (s *ListCrawlersResponseBodyPagingInfoCrawlers) GetStatus() *string {
	return s.Status
}

func (s *ListCrawlersResponseBodyPagingInfoCrawlers) GetTaskId() *int64 {
	return s.TaskId
}

func (s *ListCrawlersResponseBodyPagingInfoCrawlers) GetType() *string {
	return s.Type
}

func (s *ListCrawlersResponseBodyPagingInfoCrawlers) SetCreateTime(v int64) *ListCrawlersResponseBodyPagingInfoCrawlers {
	s.CreateTime = &v
	return s
}

func (s *ListCrawlersResponseBodyPagingInfoCrawlers) SetDataSourceId(v int64) *ListCrawlersResponseBodyPagingInfoCrawlers {
	s.DataSourceId = &v
	return s
}

func (s *ListCrawlersResponseBodyPagingInfoCrawlers) SetEnvType(v string) *ListCrawlersResponseBodyPagingInfoCrawlers {
	s.EnvType = &v
	return s
}

func (s *ListCrawlersResponseBodyPagingInfoCrawlers) SetId(v int64) *ListCrawlersResponseBodyPagingInfoCrawlers {
	s.Id = &v
	return s
}

func (s *ListCrawlersResponseBodyPagingInfoCrawlers) SetLastRunStatus(v string) *ListCrawlersResponseBodyPagingInfoCrawlers {
	s.LastRunStatus = &v
	return s
}

func (s *ListCrawlersResponseBodyPagingInfoCrawlers) SetMetaEntityId(v string) *ListCrawlersResponseBodyPagingInfoCrawlers {
	s.MetaEntityId = &v
	return s
}

func (s *ListCrawlersResponseBodyPagingInfoCrawlers) SetModifyTime(v int64) *ListCrawlersResponseBodyPagingInfoCrawlers {
	s.ModifyTime = &v
	return s
}

func (s *ListCrawlersResponseBodyPagingInfoCrawlers) SetName(v string) *ListCrawlersResponseBodyPagingInfoCrawlers {
	s.Name = &v
	return s
}

func (s *ListCrawlersResponseBodyPagingInfoCrawlers) SetOwner(v string) *ListCrawlersResponseBodyPagingInfoCrawlers {
	s.Owner = &v
	return s
}

func (s *ListCrawlersResponseBodyPagingInfoCrawlers) SetProjectId(v int64) *ListCrawlersResponseBodyPagingInfoCrawlers {
	s.ProjectId = &v
	return s
}

func (s *ListCrawlersResponseBodyPagingInfoCrawlers) SetResourceGroupId(v string) *ListCrawlersResponseBodyPagingInfoCrawlers {
	s.ResourceGroupId = &v
	return s
}

func (s *ListCrawlersResponseBodyPagingInfoCrawlers) SetScheduleConfig(v *ListCrawlersResponseBodyPagingInfoCrawlersScheduleConfig) *ListCrawlersResponseBodyPagingInfoCrawlers {
	s.ScheduleConfig = v
	return s
}

func (s *ListCrawlersResponseBodyPagingInfoCrawlers) SetStatus(v string) *ListCrawlersResponseBodyPagingInfoCrawlers {
	s.Status = &v
	return s
}

func (s *ListCrawlersResponseBodyPagingInfoCrawlers) SetTaskId(v int64) *ListCrawlersResponseBodyPagingInfoCrawlers {
	s.TaskId = &v
	return s
}

func (s *ListCrawlersResponseBodyPagingInfoCrawlers) SetType(v string) *ListCrawlersResponseBodyPagingInfoCrawlers {
	s.Type = &v
	return s
}

func (s *ListCrawlersResponseBodyPagingInfoCrawlers) Validate() error {
	if s.ScheduleConfig != nil {
		if err := s.ScheduleConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListCrawlersResponseBodyPagingInfoCrawlersScheduleConfig struct {
	// The cron expression.
	//
	// example:
	//
	// 0 0 2 ? 	- *
	CronExpress *string `json:"CronExpress,omitempty" xml:"CronExpress,omitempty"`
	// The schedule type. Valid values: MANUAL, NORMAL.
	//
	// example:
	//
	// NORMAL
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListCrawlersResponseBodyPagingInfoCrawlersScheduleConfig) String() string {
	return dara.Prettify(s)
}

func (s ListCrawlersResponseBodyPagingInfoCrawlersScheduleConfig) GoString() string {
	return s.String()
}

func (s *ListCrawlersResponseBodyPagingInfoCrawlersScheduleConfig) GetCronExpress() *string {
	return s.CronExpress
}

func (s *ListCrawlersResponseBodyPagingInfoCrawlersScheduleConfig) GetType() *string {
	return s.Type
}

func (s *ListCrawlersResponseBodyPagingInfoCrawlersScheduleConfig) SetCronExpress(v string) *ListCrawlersResponseBodyPagingInfoCrawlersScheduleConfig {
	s.CronExpress = &v
	return s
}

func (s *ListCrawlersResponseBodyPagingInfoCrawlersScheduleConfig) SetType(v string) *ListCrawlersResponseBodyPagingInfoCrawlersScheduleConfig {
	s.Type = &v
	return s
}

func (s *ListCrawlersResponseBodyPagingInfoCrawlersScheduleConfig) Validate() error {
	return dara.Validate(s)
}
