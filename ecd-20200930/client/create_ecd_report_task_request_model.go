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
	FilterList []*CreateEcdReportTaskRequestFilterList `json:"FilterList,omitempty" xml:"FilterList,omitempty" type:"Repeated"`
	// example:
	//
	// zh-CN
	LangType       *string `json:"LangType,omitempty" xml:"LangType,omitempty"`
	ReportFileName *string `json:"ReportFileName,omitempty" xml:"ReportFileName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// DESKTOP
	SubType *string `json:"SubType,omitempty" xml:"SubType,omitempty"`
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
	return dara.Validate(s)
}

type CreateEcdReportTaskRequestFilterList struct {
	// example:
	//
	// Tag:TestKey
	FilterKey    *string   `json:"FilterKey,omitempty" xml:"FilterKey,omitempty"`
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
