// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateEcdReportTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFilterList(v []*CreateEcdReportTaskRequestFilterList) *CreateEcdReportTaskRequest
	GetFilterList() []*CreateEcdReportTaskRequestFilterList
	SetLangType(v string) *CreateEcdReportTaskRequest
	GetLangType() *string
	SetReportFileName(v string) *CreateEcdReportTaskRequest
	GetReportFileName() *string
	SetSubType(v string) *CreateEcdReportTaskRequest
	GetSubType() *string
	SetTaskType(v string) *CreateEcdReportTaskRequest
	GetTaskType() *string
}

type CreateEcdReportTaskRequest struct {
	// The filter conditions for filtering query results. The logical relationship between each filter condition is "and" (&). Each filter condition contains FilterKey and FilterValues, which indicate the key and value for the filter condition.
	FilterList []*CreateEcdReportTaskRequestFilterList `json:"FilterList,omitempty" xml:"FilterList,omitempty" type:"Repeated"`
	// The language of the report. An enumerated type.
	//
	// Valid values:
	//
	// 	- zh-CN (default): Chinese
	//
	// 	- en-GB: English
	//
	// example:
	//
	// zh-CN
	LangType *string `json:"LangType,omitempty" xml:"LangType,omitempty"`
	// The name of the report file.
	ReportFileName *string `json:"ReportFileName,omitempty" xml:"ReportFileName,omitempty"`
	// The sub-type of the report export task.
	//
	// Valid value:
	//
	// 	- DESKTOP: cloud computer
	//
	// This parameter is required.
	//
	// example:
	//
	// DESKTOP
	SubType *string `json:"SubType,omitempty" xml:"SubType,omitempty"`
	// The type of the report task.
	//
	// Valid value:
	//
	// 	- RESOURCE_REPORT
	//
	// This parameter is required.
	//
	// example:
	//
	// RESOURCE_REPORT
	TaskType *string `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
}

func (s CreateEcdReportTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateEcdReportTaskRequest) GoString() string {
	return s.String()
}

func (s *CreateEcdReportTaskRequest) GetFilterList() []*CreateEcdReportTaskRequestFilterList {
	return s.FilterList
}

func (s *CreateEcdReportTaskRequest) GetLangType() *string {
	return s.LangType
}

func (s *CreateEcdReportTaskRequest) GetReportFileName() *string {
	return s.ReportFileName
}

func (s *CreateEcdReportTaskRequest) GetSubType() *string {
	return s.SubType
}

func (s *CreateEcdReportTaskRequest) GetTaskType() *string {
	return s.TaskType
}

func (s *CreateEcdReportTaskRequest) SetFilterList(v []*CreateEcdReportTaskRequestFilterList) *CreateEcdReportTaskRequest {
	s.FilterList = v
	return s
}

func (s *CreateEcdReportTaskRequest) SetLangType(v string) *CreateEcdReportTaskRequest {
	s.LangType = &v
	return s
}

func (s *CreateEcdReportTaskRequest) SetReportFileName(v string) *CreateEcdReportTaskRequest {
	s.ReportFileName = &v
	return s
}

func (s *CreateEcdReportTaskRequest) SetSubType(v string) *CreateEcdReportTaskRequest {
	s.SubType = &v
	return s
}

func (s *CreateEcdReportTaskRequest) SetTaskType(v string) *CreateEcdReportTaskRequest {
	s.TaskType = &v
	return s
}

func (s *CreateEcdReportTaskRequest) Validate() error {
	if s.FilterList != nil {
		for _, item := range s.FilterList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateEcdReportTaskRequestFilterList struct {
	// The key of the filter condition for filtering query results. When SubType is set to:
	//
	// 1.  DESKTOP (indicating a cloud computer report), the following filter conditions are available:
	//
	// 	- KeyWord: cloud computer keyword (supports automatic recognition)
	//
	// 	- RegionId: region ID
	//
	// 	- DesktopId: cloud computer ID
	//
	// 	- DesktopName: cloud computer name (supports fuzzy matching)
	//
	// 	- OfficeSiteId: office network ID
	//
	// 	- OfficeSiteName: office network name (supports fuzzy matching)
	//
	// 	- Status: cloud computer status
	//
	// 	- DesktopType: desktop type
	//
	// 	- DesktopIP: cloud computer IP address
	//
	// 	- SubPayType: billing method
	//
	// 	- EndUserId: user name (supports fuzzy matching)
	//
	// 	- ExpireTime: expiration date and time, in the yyyy-MM-dd\\"T\\"HH:mm:ss\\"Z\\" format
	//
	// 	- IncludeAssignedUser: indicates whether the cloud computer is assigned to users or not
	//
	// 	- ResourceGroupId: resource group ID
	//
	// 	- PolicyId: policy ID
	//
	// 	- Tag:{Tag Key value}: cloud computer tag (To filter data using multiple tags, specify multiple filter condition objects.)
	//
	// example:
	//
	// Tag:TestKey
	FilterKey *string `json:"FilterKey,omitempty" xml:"FilterKey,omitempty"`
	// The value of the filter condition. Only the first value of the FilterValues parameter is used, if FilterKey is set to one of the following values:
	//
	// 	- KeyWord
	//
	// 	- DesktopName
	//
	// 	- OfficeSiteName
	//
	// 	- DesktopIP
	//
	// 	- EndUserId
	//
	// 	- ExpireTime
	//
	// 	- IncludeAssignedUser
	FilterValues []*string `json:"FilterValues,omitempty" xml:"FilterValues,omitempty" type:"Repeated"`
}

func (s CreateEcdReportTaskRequestFilterList) String() string {
	return dara.Prettify(s)
}

func (s CreateEcdReportTaskRequestFilterList) GoString() string {
	return s.String()
}

func (s *CreateEcdReportTaskRequestFilterList) GetFilterKey() *string {
	return s.FilterKey
}

func (s *CreateEcdReportTaskRequestFilterList) GetFilterValues() []*string {
	return s.FilterValues
}

func (s *CreateEcdReportTaskRequestFilterList) SetFilterKey(v string) *CreateEcdReportTaskRequestFilterList {
	s.FilterKey = &v
	return s
}

func (s *CreateEcdReportTaskRequestFilterList) SetFilterValues(v []*string) *CreateEcdReportTaskRequestFilterList {
	s.FilterValues = v
	return s
}

func (s *CreateEcdReportTaskRequestFilterList) Validate() error {
	return dara.Validate(s)
}
