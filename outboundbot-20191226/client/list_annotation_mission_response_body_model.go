// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAnnotationMissionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListAnnotationMissionResponseBody
	GetCode() *string
	SetData(v *ListAnnotationMissionResponseBodyData) *ListAnnotationMissionResponseBody
	GetData() *ListAnnotationMissionResponseBodyData
	SetHttpStatusCode(v int32) *ListAnnotationMissionResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ListAnnotationMissionResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListAnnotationMissionResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListAnnotationMissionResponseBody
	GetSuccess() *bool
}

type ListAnnotationMissionResponseBody struct {
	// example:
	//
	// OK
	Code *string                                `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *ListAnnotationMissionResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// CDR \\"job-c7b8a817-b8e8-40f3-b7ad-f28dcea218ff\\" doesn\\"t exists.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 254EB995-DEDF-48A4-9101-9CA5B72FFBCC
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListAnnotationMissionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListAnnotationMissionResponseBody) GoString() string {
	return s.String()
}

func (s *ListAnnotationMissionResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListAnnotationMissionResponseBody) GetData() *ListAnnotationMissionResponseBodyData {
	return s.Data
}

func (s *ListAnnotationMissionResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListAnnotationMissionResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListAnnotationMissionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListAnnotationMissionResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListAnnotationMissionResponseBody) SetCode(v string) *ListAnnotationMissionResponseBody {
	s.Code = &v
	return s
}

func (s *ListAnnotationMissionResponseBody) SetData(v *ListAnnotationMissionResponseBodyData) *ListAnnotationMissionResponseBody {
	s.Data = v
	return s
}

func (s *ListAnnotationMissionResponseBody) SetHttpStatusCode(v int32) *ListAnnotationMissionResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListAnnotationMissionResponseBody) SetMessage(v string) *ListAnnotationMissionResponseBody {
	s.Message = &v
	return s
}

func (s *ListAnnotationMissionResponseBody) SetRequestId(v string) *ListAnnotationMissionResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAnnotationMissionResponseBody) SetSuccess(v bool) *ListAnnotationMissionResponseBody {
	s.Success = &v
	return s
}

func (s *ListAnnotationMissionResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListAnnotationMissionResponseBodyData struct {
	AnnotationMissionList []*ListAnnotationMissionResponseBodyDataAnnotationMissionList `json:"AnnotationMissionList,omitempty" xml:"AnnotationMissionList,omitempty" type:"Repeated"`
	// example:
	//
	// CDR \\"job-c7b8a817-b8e8-40f3-b7ad-f28dcea218ff\\" doesn\\"t exists.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 1
	PageIndex *int64 `json:"PageIndex,omitempty" xml:"PageIndex,omitempty"`
	// example:
	//
	// 10
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// example:
	//
	// 30
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// example:
	//
	// 4
	TotalPageCount *int64 `json:"TotalPageCount,omitempty" xml:"TotalPageCount,omitempty"`
}

func (s ListAnnotationMissionResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListAnnotationMissionResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListAnnotationMissionResponseBodyData) GetAnnotationMissionList() []*ListAnnotationMissionResponseBodyDataAnnotationMissionList {
	return s.AnnotationMissionList
}

func (s *ListAnnotationMissionResponseBodyData) GetMessage() *string {
	return s.Message
}

func (s *ListAnnotationMissionResponseBodyData) GetPageIndex() *int64 {
	return s.PageIndex
}

func (s *ListAnnotationMissionResponseBodyData) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListAnnotationMissionResponseBodyData) GetSuccess() *bool {
	return s.Success
}

func (s *ListAnnotationMissionResponseBodyData) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListAnnotationMissionResponseBodyData) GetTotalPageCount() *int64 {
	return s.TotalPageCount
}

func (s *ListAnnotationMissionResponseBodyData) SetAnnotationMissionList(v []*ListAnnotationMissionResponseBodyDataAnnotationMissionList) *ListAnnotationMissionResponseBodyData {
	s.AnnotationMissionList = v
	return s
}

func (s *ListAnnotationMissionResponseBodyData) SetMessage(v string) *ListAnnotationMissionResponseBodyData {
	s.Message = &v
	return s
}

func (s *ListAnnotationMissionResponseBodyData) SetPageIndex(v int64) *ListAnnotationMissionResponseBodyData {
	s.PageIndex = &v
	return s
}

func (s *ListAnnotationMissionResponseBodyData) SetPageSize(v int64) *ListAnnotationMissionResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListAnnotationMissionResponseBodyData) SetSuccess(v bool) *ListAnnotationMissionResponseBodyData {
	s.Success = &v
	return s
}

func (s *ListAnnotationMissionResponseBodyData) SetTotalCount(v int64) *ListAnnotationMissionResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *ListAnnotationMissionResponseBodyData) SetTotalPageCount(v int64) *ListAnnotationMissionResponseBodyData {
	s.TotalPageCount = &v
	return s
}

func (s *ListAnnotationMissionResponseBodyData) Validate() error {
	if s.AnnotationMissionList != nil {
		for _, item := range s.AnnotationMissionList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListAnnotationMissionResponseBodyDataAnnotationMissionList struct {
	// example:
	//
	// 1
	AnnotationMissionDataSourceType      *int32   `json:"AnnotationMissionDataSourceType,omitempty" xml:"AnnotationMissionDataSourceType,omitempty"`
	AnnotationMissionDebugDataSourceList []*int32 `json:"AnnotationMissionDebugDataSourceList,omitempty" xml:"AnnotationMissionDebugDataSourceList,omitempty" type:"Repeated"`
	// example:
	//
	// ddce607f-f537-4ebd-9914-cf45671defb9
	AnnotationMissionId   *string `json:"AnnotationMissionId,omitempty" xml:"AnnotationMissionId,omitempty"`
	AnnotationMissionName *string `json:"AnnotationMissionName,omitempty" xml:"AnnotationMissionName,omitempty"`
	// example:
	//
	// 1
	AnnotationStatus *int32 `json:"AnnotationStatus,omitempty" xml:"AnnotationStatus,omitempty"`
	// example:
	//
	// 1684511999000
	ConversationTimeEndFilter *int64 `json:"ConversationTimeEndFilter,omitempty" xml:"ConversationTimeEndFilter,omitempty"`
	// example:
	//
	// 1683216000000
	ConversationTimeStartFilter *int64 `json:"ConversationTimeStartFilter,omitempty" xml:"ConversationTimeStartFilter,omitempty"`
	// example:
	//
	// 1676170339515
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// false
	ExcludeOtherMissionSession *bool `json:"ExcludeOtherMissionSession,omitempty" xml:"ExcludeOtherMissionSession,omitempty"`
	// example:
	//
	// 1683443903785
	FinishTime *int64 `json:"FinishTime,omitempty" xml:"FinishTime,omitempty"`
	// example:
	//
	// 32be9d94-1346-4c4a-a4d0-ccd379f87013
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// 1
	SamplingCount       *int32  `json:"SamplingCount,omitempty" xml:"SamplingCount,omitempty"`
	SamplingDescription *string `json:"SamplingDescription,omitempty" xml:"SamplingDescription,omitempty"`
	// example:
	//
	// 1
	SamplingRate *int32 `json:"SamplingRate,omitempty" xml:"SamplingRate,omitempty"`
	// example:
	//
	// 1
	SamplingType               *int32   `json:"SamplingType,omitempty" xml:"SamplingType,omitempty"`
	SessionEndReasonFilterList []*int32 `json:"SessionEndReasonFilterList,omitempty" xml:"SessionEndReasonFilterList,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	SessionFinishCount *int32 `json:"SessionFinishCount,omitempty" xml:"SessionFinishCount,omitempty"`
	// example:
	//
	// 1
	SessionTotalCount *int32 `json:"SessionTotalCount,omitempty" xml:"SessionTotalCount,omitempty"`
}

func (s ListAnnotationMissionResponseBodyDataAnnotationMissionList) String() string {
	return dara.Prettify(s)
}

func (s ListAnnotationMissionResponseBodyDataAnnotationMissionList) GoString() string {
	return s.String()
}

func (s *ListAnnotationMissionResponseBodyDataAnnotationMissionList) GetAnnotationMissionDataSourceType() *int32 {
	return s.AnnotationMissionDataSourceType
}

func (s *ListAnnotationMissionResponseBodyDataAnnotationMissionList) GetAnnotationMissionDebugDataSourceList() []*int32 {
	return s.AnnotationMissionDebugDataSourceList
}

func (s *ListAnnotationMissionResponseBodyDataAnnotationMissionList) GetAnnotationMissionId() *string {
	return s.AnnotationMissionId
}

func (s *ListAnnotationMissionResponseBodyDataAnnotationMissionList) GetAnnotationMissionName() *string {
	return s.AnnotationMissionName
}

func (s *ListAnnotationMissionResponseBodyDataAnnotationMissionList) GetAnnotationStatus() *int32 {
	return s.AnnotationStatus
}

func (s *ListAnnotationMissionResponseBodyDataAnnotationMissionList) GetConversationTimeEndFilter() *int64 {
	return s.ConversationTimeEndFilter
}

func (s *ListAnnotationMissionResponseBodyDataAnnotationMissionList) GetConversationTimeStartFilter() *int64 {
	return s.ConversationTimeStartFilter
}

func (s *ListAnnotationMissionResponseBodyDataAnnotationMissionList) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *ListAnnotationMissionResponseBodyDataAnnotationMissionList) GetExcludeOtherMissionSession() *bool {
	return s.ExcludeOtherMissionSession
}

func (s *ListAnnotationMissionResponseBodyDataAnnotationMissionList) GetFinishTime() *int64 {
	return s.FinishTime
}

func (s *ListAnnotationMissionResponseBodyDataAnnotationMissionList) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListAnnotationMissionResponseBodyDataAnnotationMissionList) GetSamplingCount() *int32 {
	return s.SamplingCount
}

func (s *ListAnnotationMissionResponseBodyDataAnnotationMissionList) GetSamplingDescription() *string {
	return s.SamplingDescription
}

func (s *ListAnnotationMissionResponseBodyDataAnnotationMissionList) GetSamplingRate() *int32 {
	return s.SamplingRate
}

func (s *ListAnnotationMissionResponseBodyDataAnnotationMissionList) GetSamplingType() *int32 {
	return s.SamplingType
}

func (s *ListAnnotationMissionResponseBodyDataAnnotationMissionList) GetSessionEndReasonFilterList() []*int32 {
	return s.SessionEndReasonFilterList
}

func (s *ListAnnotationMissionResponseBodyDataAnnotationMissionList) GetSessionFinishCount() *int32 {
	return s.SessionFinishCount
}

func (s *ListAnnotationMissionResponseBodyDataAnnotationMissionList) GetSessionTotalCount() *int32 {
	return s.SessionTotalCount
}

func (s *ListAnnotationMissionResponseBodyDataAnnotationMissionList) SetAnnotationMissionDataSourceType(v int32) *ListAnnotationMissionResponseBodyDataAnnotationMissionList {
	s.AnnotationMissionDataSourceType = &v
	return s
}

func (s *ListAnnotationMissionResponseBodyDataAnnotationMissionList) SetAnnotationMissionDebugDataSourceList(v []*int32) *ListAnnotationMissionResponseBodyDataAnnotationMissionList {
	s.AnnotationMissionDebugDataSourceList = v
	return s
}

func (s *ListAnnotationMissionResponseBodyDataAnnotationMissionList) SetAnnotationMissionId(v string) *ListAnnotationMissionResponseBodyDataAnnotationMissionList {
	s.AnnotationMissionId = &v
	return s
}

func (s *ListAnnotationMissionResponseBodyDataAnnotationMissionList) SetAnnotationMissionName(v string) *ListAnnotationMissionResponseBodyDataAnnotationMissionList {
	s.AnnotationMissionName = &v
	return s
}

func (s *ListAnnotationMissionResponseBodyDataAnnotationMissionList) SetAnnotationStatus(v int32) *ListAnnotationMissionResponseBodyDataAnnotationMissionList {
	s.AnnotationStatus = &v
	return s
}

func (s *ListAnnotationMissionResponseBodyDataAnnotationMissionList) SetConversationTimeEndFilter(v int64) *ListAnnotationMissionResponseBodyDataAnnotationMissionList {
	s.ConversationTimeEndFilter = &v
	return s
}

func (s *ListAnnotationMissionResponseBodyDataAnnotationMissionList) SetConversationTimeStartFilter(v int64) *ListAnnotationMissionResponseBodyDataAnnotationMissionList {
	s.ConversationTimeStartFilter = &v
	return s
}

func (s *ListAnnotationMissionResponseBodyDataAnnotationMissionList) SetCreateTime(v int64) *ListAnnotationMissionResponseBodyDataAnnotationMissionList {
	s.CreateTime = &v
	return s
}

func (s *ListAnnotationMissionResponseBodyDataAnnotationMissionList) SetExcludeOtherMissionSession(v bool) *ListAnnotationMissionResponseBodyDataAnnotationMissionList {
	s.ExcludeOtherMissionSession = &v
	return s
}

func (s *ListAnnotationMissionResponseBodyDataAnnotationMissionList) SetFinishTime(v int64) *ListAnnotationMissionResponseBodyDataAnnotationMissionList {
	s.FinishTime = &v
	return s
}

func (s *ListAnnotationMissionResponseBodyDataAnnotationMissionList) SetInstanceId(v string) *ListAnnotationMissionResponseBodyDataAnnotationMissionList {
	s.InstanceId = &v
	return s
}

func (s *ListAnnotationMissionResponseBodyDataAnnotationMissionList) SetSamplingCount(v int32) *ListAnnotationMissionResponseBodyDataAnnotationMissionList {
	s.SamplingCount = &v
	return s
}

func (s *ListAnnotationMissionResponseBodyDataAnnotationMissionList) SetSamplingDescription(v string) *ListAnnotationMissionResponseBodyDataAnnotationMissionList {
	s.SamplingDescription = &v
	return s
}

func (s *ListAnnotationMissionResponseBodyDataAnnotationMissionList) SetSamplingRate(v int32) *ListAnnotationMissionResponseBodyDataAnnotationMissionList {
	s.SamplingRate = &v
	return s
}

func (s *ListAnnotationMissionResponseBodyDataAnnotationMissionList) SetSamplingType(v int32) *ListAnnotationMissionResponseBodyDataAnnotationMissionList {
	s.SamplingType = &v
	return s
}

func (s *ListAnnotationMissionResponseBodyDataAnnotationMissionList) SetSessionEndReasonFilterList(v []*int32) *ListAnnotationMissionResponseBodyDataAnnotationMissionList {
	s.SessionEndReasonFilterList = v
	return s
}

func (s *ListAnnotationMissionResponseBodyDataAnnotationMissionList) SetSessionFinishCount(v int32) *ListAnnotationMissionResponseBodyDataAnnotationMissionList {
	s.SessionFinishCount = &v
	return s
}

func (s *ListAnnotationMissionResponseBodyDataAnnotationMissionList) SetSessionTotalCount(v int32) *ListAnnotationMissionResponseBodyDataAnnotationMissionList {
	s.SessionTotalCount = &v
	return s
}

func (s *ListAnnotationMissionResponseBodyDataAnnotationMissionList) Validate() error {
	return dara.Validate(s)
}
