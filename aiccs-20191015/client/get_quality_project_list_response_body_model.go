// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetQualityProjectListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetQualityProjectListResponseBody
	GetCode() *string
	SetData(v *GetQualityProjectListResponseBodyData) *GetQualityProjectListResponseBody
	GetData() *GetQualityProjectListResponseBodyData
	SetMessage(v string) *GetQualityProjectListResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetQualityProjectListResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetQualityProjectListResponseBody
	GetSuccess() *bool
}

type GetQualityProjectListResponseBody struct {
	// Status code. A return value of 200 indicates that the request succeeded.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Quality inspection job information.
	Data *GetQualityProjectListResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// Description of the status code.
	//
	// example:
	//
	// xxxx
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Request ID.
	//
	// example:
	//
	// EE338D98-9BD3-4413-B165
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the API call succeeded.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetQualityProjectListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetQualityProjectListResponseBody) GoString() string {
	return s.String()
}

func (s *GetQualityProjectListResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetQualityProjectListResponseBody) GetData() *GetQualityProjectListResponseBodyData {
	return s.Data
}

func (s *GetQualityProjectListResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetQualityProjectListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetQualityProjectListResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetQualityProjectListResponseBody) SetCode(v string) *GetQualityProjectListResponseBody {
	s.Code = &v
	return s
}

func (s *GetQualityProjectListResponseBody) SetData(v *GetQualityProjectListResponseBodyData) *GetQualityProjectListResponseBody {
	s.Data = v
	return s
}

func (s *GetQualityProjectListResponseBody) SetMessage(v string) *GetQualityProjectListResponseBody {
	s.Message = &v
	return s
}

func (s *GetQualityProjectListResponseBody) SetRequestId(v string) *GetQualityProjectListResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetQualityProjectListResponseBody) SetSuccess(v bool) *GetQualityProjectListResponseBody {
	s.Success = &v
	return s
}

func (s *GetQualityProjectListResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetQualityProjectListResponseBodyData struct {
	// Current page number.
	//
	// example:
	//
	// 1
	PageNo *int32 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// Page size.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// List of quality inspection jobs.
	QualityProjectList []*GetQualityProjectListResponseBodyDataQualityProjectList `json:"QualityProjectList,omitempty" xml:"QualityProjectList,omitempty" type:"Repeated"`
	// Total number of records.
	//
	// example:
	//
	// 35
	Total *int64 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s GetQualityProjectListResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetQualityProjectListResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetQualityProjectListResponseBodyData) GetPageNo() *int32 {
	return s.PageNo
}

func (s *GetQualityProjectListResponseBodyData) GetPageSize() *int32 {
	return s.PageSize
}

func (s *GetQualityProjectListResponseBodyData) GetQualityProjectList() []*GetQualityProjectListResponseBodyDataQualityProjectList {
	return s.QualityProjectList
}

func (s *GetQualityProjectListResponseBodyData) GetTotal() *int64 {
	return s.Total
}

func (s *GetQualityProjectListResponseBodyData) SetPageNo(v int32) *GetQualityProjectListResponseBodyData {
	s.PageNo = &v
	return s
}

func (s *GetQualityProjectListResponseBodyData) SetPageSize(v int32) *GetQualityProjectListResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *GetQualityProjectListResponseBodyData) SetQualityProjectList(v []*GetQualityProjectListResponseBodyDataQualityProjectList) *GetQualityProjectListResponseBodyData {
	s.QualityProjectList = v
	return s
}

func (s *GetQualityProjectListResponseBodyData) SetTotal(v int64) *GetQualityProjectListResponseBodyData {
	s.Total = &v
	return s
}

func (s *GetQualityProjectListResponseBodyData) Validate() error {
	if s.QualityProjectList != nil {
		for _, item := range s.QualityProjectList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetQualityProjectListResponseBodyDataQualityProjectList struct {
	// Quality inspection job frequency.
	//
	// example:
	//
	// 1
	CheckFreqType *int32 `json:"CheckFreqType,omitempty" xml:"CheckFreqType,omitempty"`
	// Creation Time.
	//
	// example:
	//
	// 2021-04-07 18:07:18
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// Quality inspection sampling scope.
	DepList []*int64 `json:"DepList,omitempty" xml:"DepList,omitempty" type:"Repeated"`
	// Quality inspection sampling scope.
	GroupList []*int64 `json:"GroupList,omitempty" xml:"GroupList,omitempty" type:"Repeated"`
	// Quality inspection job ID.
	//
	// example:
	//
	// 15977801
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// Updated At.
	//
	// example:
	//
	// 2021-04-07 18:07:19
	ModifyTime *string `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	// Quality inspection job name.
	//
	// example:
	//
	// 自动化质检任务
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	// Quality inspection rule IDs.
	QualityRuleIds []*int64 `json:"QualityRuleIds,omitempty" xml:"QualityRuleIds,omitempty" type:"Repeated"`
	// Quality inspection type. Fixed value is **1*	- (Consultation).
	//
	// example:
	//
	// 1
	QualityType *int32 `json:"QualityType,omitempty" xml:"QualityType,omitempty"`
	// Scope of quality inspection sampling.
	ServicerList []*int64 `json:"ServicerList,omitempty" xml:"ServicerList,omitempty" type:"Repeated"`
	// Status of the quality inspection job.
	//
	// example:
	//
	// 0
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// Quality inspection job version number.
	//
	// example:
	//
	// 1
	Version *int32 `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s GetQualityProjectListResponseBodyDataQualityProjectList) String() string {
	return dara.Prettify(s)
}

func (s GetQualityProjectListResponseBodyDataQualityProjectList) GoString() string {
	return s.String()
}

func (s *GetQualityProjectListResponseBodyDataQualityProjectList) GetCheckFreqType() *int32 {
	return s.CheckFreqType
}

func (s *GetQualityProjectListResponseBodyDataQualityProjectList) GetCreateTime() *string {
	return s.CreateTime
}

func (s *GetQualityProjectListResponseBodyDataQualityProjectList) GetDepList() []*int64 {
	return s.DepList
}

func (s *GetQualityProjectListResponseBodyDataQualityProjectList) GetGroupList() []*int64 {
	return s.GroupList
}

func (s *GetQualityProjectListResponseBodyDataQualityProjectList) GetId() *int64 {
	return s.Id
}

func (s *GetQualityProjectListResponseBodyDataQualityProjectList) GetModifyTime() *string {
	return s.ModifyTime
}

func (s *GetQualityProjectListResponseBodyDataQualityProjectList) GetProjectName() *string {
	return s.ProjectName
}

func (s *GetQualityProjectListResponseBodyDataQualityProjectList) GetQualityRuleIds() []*int64 {
	return s.QualityRuleIds
}

func (s *GetQualityProjectListResponseBodyDataQualityProjectList) GetQualityType() *int32 {
	return s.QualityType
}

func (s *GetQualityProjectListResponseBodyDataQualityProjectList) GetServicerList() []*int64 {
	return s.ServicerList
}

func (s *GetQualityProjectListResponseBodyDataQualityProjectList) GetStatus() *int32 {
	return s.Status
}

func (s *GetQualityProjectListResponseBodyDataQualityProjectList) GetVersion() *int32 {
	return s.Version
}

func (s *GetQualityProjectListResponseBodyDataQualityProjectList) SetCheckFreqType(v int32) *GetQualityProjectListResponseBodyDataQualityProjectList {
	s.CheckFreqType = &v
	return s
}

func (s *GetQualityProjectListResponseBodyDataQualityProjectList) SetCreateTime(v string) *GetQualityProjectListResponseBodyDataQualityProjectList {
	s.CreateTime = &v
	return s
}

func (s *GetQualityProjectListResponseBodyDataQualityProjectList) SetDepList(v []*int64) *GetQualityProjectListResponseBodyDataQualityProjectList {
	s.DepList = v
	return s
}

func (s *GetQualityProjectListResponseBodyDataQualityProjectList) SetGroupList(v []*int64) *GetQualityProjectListResponseBodyDataQualityProjectList {
	s.GroupList = v
	return s
}

func (s *GetQualityProjectListResponseBodyDataQualityProjectList) SetId(v int64) *GetQualityProjectListResponseBodyDataQualityProjectList {
	s.Id = &v
	return s
}

func (s *GetQualityProjectListResponseBodyDataQualityProjectList) SetModifyTime(v string) *GetQualityProjectListResponseBodyDataQualityProjectList {
	s.ModifyTime = &v
	return s
}

func (s *GetQualityProjectListResponseBodyDataQualityProjectList) SetProjectName(v string) *GetQualityProjectListResponseBodyDataQualityProjectList {
	s.ProjectName = &v
	return s
}

func (s *GetQualityProjectListResponseBodyDataQualityProjectList) SetQualityRuleIds(v []*int64) *GetQualityProjectListResponseBodyDataQualityProjectList {
	s.QualityRuleIds = v
	return s
}

func (s *GetQualityProjectListResponseBodyDataQualityProjectList) SetQualityType(v int32) *GetQualityProjectListResponseBodyDataQualityProjectList {
	s.QualityType = &v
	return s
}

func (s *GetQualityProjectListResponseBodyDataQualityProjectList) SetServicerList(v []*int64) *GetQualityProjectListResponseBodyDataQualityProjectList {
	s.ServicerList = v
	return s
}

func (s *GetQualityProjectListResponseBodyDataQualityProjectList) SetStatus(v int32) *GetQualityProjectListResponseBodyDataQualityProjectList {
	s.Status = &v
	return s
}

func (s *GetQualityProjectListResponseBodyDataQualityProjectList) SetVersion(v int32) *GetQualityProjectListResponseBodyDataQualityProjectList {
	s.Version = &v
	return s
}

func (s *GetQualityProjectListResponseBodyDataQualityProjectList) Validate() error {
	return dara.Validate(s)
}
