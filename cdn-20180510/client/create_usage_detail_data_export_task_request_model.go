// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateUsageDetailDataExportTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainNames(v string) *CreateUsageDetailDataExportTaskRequest
	GetDomainNames() *string
	SetEndTime(v string) *CreateUsageDetailDataExportTaskRequest
	GetEndTime() *string
	SetGroup(v string) *CreateUsageDetailDataExportTaskRequest
	GetGroup() *string
	SetLanguage(v string) *CreateUsageDetailDataExportTaskRequest
	GetLanguage() *string
	SetStartTime(v string) *CreateUsageDetailDataExportTaskRequest
	GetStartTime() *string
	SetTaskName(v string) *CreateUsageDetailDataExportTaskRequest
	GetTaskName() *string
	SetType(v string) *CreateUsageDetailDataExportTaskRequest
	GetType() *string
}

type CreateUsageDetailDataExportTaskRequest struct {
	// The domain names. If you do not specify the Group parameter, resource usage details of these domain names are exported.
	//
	// If you do not specify this parameter, resource usage details are exported based on accounts.
	//
	// example:
	//
	// example.com
	DomainNames *string `json:"DomainNames,omitempty" xml:"DomainNames,omitempty"`
	// The end of the time range to query.
	//
	// Specify the time in the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC.
	//
	// > The end time must be later than the start time.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2019-12-10T21:00:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The domain name group. If you specify this parameter, the **DomainNames*	- parameter is ignored.
	//
	// example:
	//
	// xxx
	Group *string `json:"Group,omitempty" xml:"Group,omitempty"`
	// The language in which you want to export the file. Valid values:
	//
	// 	- **zh-cn**: Chinese. This is the default value.
	//
	// 	- **en-us**: English
	//
	// example:
	//
	// en-us
	Language *string `json:"Language,omitempty" xml:"Language,omitempty"`
	// The beginning of the time range to query.
	//
	// Specify the time in the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2019-12-10T20:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The name of the task.
	//
	// example:
	//
	// Refresh
	TaskName *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
	// The type of resource usage data to query. Valid values:
	//
	// 	- **flow**: traffic and bandwidth
	//
	// 	- **vas**: requests
	//
	// This parameter is required.
	//
	// example:
	//
	// flow
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s CreateUsageDetailDataExportTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateUsageDetailDataExportTaskRequest) GoString() string {
	return s.String()
}

func (s *CreateUsageDetailDataExportTaskRequest) GetDomainNames() *string {
	return s.DomainNames
}

func (s *CreateUsageDetailDataExportTaskRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *CreateUsageDetailDataExportTaskRequest) GetGroup() *string {
	return s.Group
}

func (s *CreateUsageDetailDataExportTaskRequest) GetLanguage() *string {
	return s.Language
}

func (s *CreateUsageDetailDataExportTaskRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *CreateUsageDetailDataExportTaskRequest) GetTaskName() *string {
	return s.TaskName
}

func (s *CreateUsageDetailDataExportTaskRequest) GetType() *string {
	return s.Type
}

func (s *CreateUsageDetailDataExportTaskRequest) SetDomainNames(v string) *CreateUsageDetailDataExportTaskRequest {
	s.DomainNames = &v
	return s
}

func (s *CreateUsageDetailDataExportTaskRequest) SetEndTime(v string) *CreateUsageDetailDataExportTaskRequest {
	s.EndTime = &v
	return s
}

func (s *CreateUsageDetailDataExportTaskRequest) SetGroup(v string) *CreateUsageDetailDataExportTaskRequest {
	s.Group = &v
	return s
}

func (s *CreateUsageDetailDataExportTaskRequest) SetLanguage(v string) *CreateUsageDetailDataExportTaskRequest {
	s.Language = &v
	return s
}

func (s *CreateUsageDetailDataExportTaskRequest) SetStartTime(v string) *CreateUsageDetailDataExportTaskRequest {
	s.StartTime = &v
	return s
}

func (s *CreateUsageDetailDataExportTaskRequest) SetTaskName(v string) *CreateUsageDetailDataExportTaskRequest {
	s.TaskName = &v
	return s
}

func (s *CreateUsageDetailDataExportTaskRequest) SetType(v string) *CreateUsageDetailDataExportTaskRequest {
	s.Type = &v
	return s
}

func (s *CreateUsageDetailDataExportTaskRequest) Validate() error {
	return dara.Validate(s)
}
