// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateEcdReportTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBusinessChannel(v string) *CreateEcdReportTaskRequest
	GetBusinessChannel() *string
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
	// The business channel. Valid values:
	//
	// Enterprise: Enterprise Edition.
	//
	// Business: Business Edition.
	//
	// example:
	//
	// Enterprise
	BusinessChannel *string `json:"BusinessChannel,omitempty" xml:"BusinessChannel,omitempty"`
	// The list of filter conditions for query results. The filter condition objects have a logical AND (&) relationship.
	//
	// Each filter condition contains FilterKey and FilterValues, which specify the filter condition name and its values.
	FilterList []*CreateEcdReportTaskRequestFilterList `json:"FilterList,omitempty" xml:"FilterList,omitempty" type:"Repeated"`
	// The report language type. This is an enumerated value.
	//
	// example:
	//
	// zh-CN
	LangType *string `json:"LangType,omitempty" xml:"LangType,omitempty"`
	// The report file name.
	//
	// example:
	//
	// TestReportName
	ReportFileName *string `json:"ReportFileName,omitempty" xml:"ReportFileName,omitempty"`
	// The report task subtype.
	//
	// [_single.params.SubType.enum.  DESKTOP]Cloud computer
	//
	// This parameter is required.
	//
	// example:
	//
	// DESKTOP
	SubType *string `json:"SubType,omitempty" xml:"SubType,omitempty"`
	// The report task type.
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

func (s *CreateEcdReportTaskRequest) GetBusinessChannel() *string {
	return s.BusinessChannel
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

func (s *CreateEcdReportTaskRequest) SetBusinessChannel(v string) *CreateEcdReportTaskRequest {
	s.BusinessChannel = &v
	return s
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
	// The report query filter condition. The valid values vary based on the SubType value:
	//
	// 1. Cloud computer report
	//
	// - KeyWord: automatically identifies the cloud computer keyword.
	//
	// - RegionId: the region ID.
	//
	// - DesktopId: the cloud computer ID.
	//
	// - DesktopName: the cloud computer name (fuzzy match).
	//
	// - OfficeSiteId: the office network ID.
	//
	// - OfficeSiteName: the office network name (fuzzy match).
	//
	// - Status: the cloud computer status.
	//
	// - DesktopType: the desktop specifications.
	//
	// - DesktopIP: the cloud computer IP address.
	//
	// - SubPayType: the billing method.
	//
	// - EndUserId: the username (fuzzy match).
	//
	// - ExpireTime: the expiration time, in the yyyy-MM-dd\\"T\\"HH:mm:ss\\"Z\\" format.
	//
	// - IncludeAssignedUser: specifies whether users are assigned.
	//
	// - ResourceGroupId: the resource group ID.
	//
	// - PolicyId: the policy ID.
	//
	// - Tag:{Tag key}: the cloud computer tag. To query by multiple tags, pass in multiple Filter objects.
	//
	// example:
	//
	// Tag:TestKey
	FilterKey *string `json:"FilterKey,omitempty" xml:"FilterKey,omitempty"`
	// The values of the filter condition.
	//
	// When FilterKey is set to one of the following values, only the first value in FilterValues is used:
	//
	// - KeyWord
	//
	// - DesktopName
	//
	// - OfficeSiteName
	//
	// - DesktopIP
	//
	// - EndUserId
	//
	// - ExpireTime
	//
	// - IncludeAssignedUser
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
