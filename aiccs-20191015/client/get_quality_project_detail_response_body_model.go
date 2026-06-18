// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetQualityProjectDetailResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetQualityProjectDetailResponseBody
	GetCode() *string
	SetData(v *GetQualityProjectDetailResponseBodyData) *GetQualityProjectDetailResponseBody
	GetData() *GetQualityProjectDetailResponseBodyData
	SetMessage(v string) *GetQualityProjectDetailResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetQualityProjectDetailResponseBody
	GetRequestId() *string
	SetSuccess(v string) *GetQualityProjectDetailResponseBody
	GetSuccess() *string
}

type GetQualityProjectDetailResponseBody struct {
	// The status code. A return value of 200 indicates that the request succeeded.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Quality inspection job information.
	Data *GetQualityProjectDetailResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// Status code description.
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
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetQualityProjectDetailResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetQualityProjectDetailResponseBody) GoString() string {
	return s.String()
}

func (s *GetQualityProjectDetailResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetQualityProjectDetailResponseBody) GetData() *GetQualityProjectDetailResponseBodyData {
	return s.Data
}

func (s *GetQualityProjectDetailResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetQualityProjectDetailResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetQualityProjectDetailResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *GetQualityProjectDetailResponseBody) SetCode(v string) *GetQualityProjectDetailResponseBody {
	s.Code = &v
	return s
}

func (s *GetQualityProjectDetailResponseBody) SetData(v *GetQualityProjectDetailResponseBodyData) *GetQualityProjectDetailResponseBody {
	s.Data = v
	return s
}

func (s *GetQualityProjectDetailResponseBody) SetMessage(v string) *GetQualityProjectDetailResponseBody {
	s.Message = &v
	return s
}

func (s *GetQualityProjectDetailResponseBody) SetRequestId(v string) *GetQualityProjectDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetQualityProjectDetailResponseBody) SetSuccess(v string) *GetQualityProjectDetailResponseBody {
	s.Success = &v
	return s
}

func (s *GetQualityProjectDetailResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetQualityProjectDetailResponseBodyData struct {
	// Inspection frequency type. Valid values:
	//
	// - **1**: Periodic quality inspection
	//
	// - **4**: Ad hoc quality inspection
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
	// Quality inspection job name
	//
	// example:
	//
	// 自动化质检任务
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	// Quality inspection rule IDs.
	QualityRuleIds []*int64 `json:"QualityRuleIds,omitempty" xml:"QualityRuleIds,omitempty" type:"Repeated"`
	// Quality inspection type. Fixed value: **1*	- (Consultation).
	//
	// example:
	//
	// 1
	QualityType *int32 `json:"QualityType,omitempty" xml:"QualityType,omitempty"`
	// Quality inspection sampling scope.
	ServicerList []*int64 `json:"ServicerList,omitempty" xml:"ServicerList,omitempty" type:"Repeated"`
	// Quality inspection job status. Valid values:
	//
	// - **0**: Start
	//
	// - **1**: Shutdown
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

func (s GetQualityProjectDetailResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetQualityProjectDetailResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetQualityProjectDetailResponseBodyData) GetCheckFreqType() *int32 {
	return s.CheckFreqType
}

func (s *GetQualityProjectDetailResponseBodyData) GetCreateTime() *string {
	return s.CreateTime
}

func (s *GetQualityProjectDetailResponseBodyData) GetDepList() []*int64 {
	return s.DepList
}

func (s *GetQualityProjectDetailResponseBodyData) GetGroupList() []*int64 {
	return s.GroupList
}

func (s *GetQualityProjectDetailResponseBodyData) GetId() *int64 {
	return s.Id
}

func (s *GetQualityProjectDetailResponseBodyData) GetModifyTime() *string {
	return s.ModifyTime
}

func (s *GetQualityProjectDetailResponseBodyData) GetProjectName() *string {
	return s.ProjectName
}

func (s *GetQualityProjectDetailResponseBodyData) GetQualityRuleIds() []*int64 {
	return s.QualityRuleIds
}

func (s *GetQualityProjectDetailResponseBodyData) GetQualityType() *int32 {
	return s.QualityType
}

func (s *GetQualityProjectDetailResponseBodyData) GetServicerList() []*int64 {
	return s.ServicerList
}

func (s *GetQualityProjectDetailResponseBodyData) GetStatus() *int32 {
	return s.Status
}

func (s *GetQualityProjectDetailResponseBodyData) GetVersion() *int32 {
	return s.Version
}

func (s *GetQualityProjectDetailResponseBodyData) SetCheckFreqType(v int32) *GetQualityProjectDetailResponseBodyData {
	s.CheckFreqType = &v
	return s
}

func (s *GetQualityProjectDetailResponseBodyData) SetCreateTime(v string) *GetQualityProjectDetailResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *GetQualityProjectDetailResponseBodyData) SetDepList(v []*int64) *GetQualityProjectDetailResponseBodyData {
	s.DepList = v
	return s
}

func (s *GetQualityProjectDetailResponseBodyData) SetGroupList(v []*int64) *GetQualityProjectDetailResponseBodyData {
	s.GroupList = v
	return s
}

func (s *GetQualityProjectDetailResponseBodyData) SetId(v int64) *GetQualityProjectDetailResponseBodyData {
	s.Id = &v
	return s
}

func (s *GetQualityProjectDetailResponseBodyData) SetModifyTime(v string) *GetQualityProjectDetailResponseBodyData {
	s.ModifyTime = &v
	return s
}

func (s *GetQualityProjectDetailResponseBodyData) SetProjectName(v string) *GetQualityProjectDetailResponseBodyData {
	s.ProjectName = &v
	return s
}

func (s *GetQualityProjectDetailResponseBodyData) SetQualityRuleIds(v []*int64) *GetQualityProjectDetailResponseBodyData {
	s.QualityRuleIds = v
	return s
}

func (s *GetQualityProjectDetailResponseBodyData) SetQualityType(v int32) *GetQualityProjectDetailResponseBodyData {
	s.QualityType = &v
	return s
}

func (s *GetQualityProjectDetailResponseBodyData) SetServicerList(v []*int64) *GetQualityProjectDetailResponseBodyData {
	s.ServicerList = v
	return s
}

func (s *GetQualityProjectDetailResponseBodyData) SetStatus(v int32) *GetQualityProjectDetailResponseBodyData {
	s.Status = &v
	return s
}

func (s *GetQualityProjectDetailResponseBodyData) SetVersion(v int32) *GetQualityProjectDetailResponseBodyData {
	s.Version = &v
	return s
}

func (s *GetQualityProjectDetailResponseBodyData) Validate() error {
	return dara.Validate(s)
}
