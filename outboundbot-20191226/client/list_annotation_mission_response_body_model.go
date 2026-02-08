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
	// Response code
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Data
	Data *ListAnnotationMissionResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// HTTP status code
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// Additional information. The value is as follows: If the request is normal, returns "success". If the request is abnormal, returns the specific error code.
	//
	// example:
	//
	// CDR \\"job-c7b8a817-b8e8-40f3-b7ad-f28dcea218ff\\" doesn\\"t exists.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Request ID
	//
	// example:
	//
	// 254EB995-DEDF-48A4-9101-9CA5B72FFBCC
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the invocation succeeded. true: The invocation succeeded. false: Failed to Invocate.
	//
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
	// Annotation job list
	AnnotationMissionList []*ListAnnotationMissionResponseBodyDataAnnotationMissionList `json:"AnnotationMissionList,omitempty" xml:"AnnotationMissionList,omitempty" type:"Repeated"`
	// Additional information. The value is as follows: If the request is normal, returns "success". If the request is abnormal, returns the specific error code.
	//
	// example:
	//
	// CDR \\"job-c7b8a817-b8e8-40f3-b7ad-f28dcea218ff\\" doesn\\"t exists.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Page index
	//
	// example:
	//
	// 1
	PageIndex *int64 `json:"PageIndex,omitempty" xml:"PageIndex,omitempty"`
	// Number of entries displayed per page
	//
	// example:
	//
	// 10
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Indicates whether the invocation succeeded. true: The invocation succeeded. false: Failed to Invocate.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// Total count
	//
	// example:
	//
	// 30
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// Total number of pages
	//
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
	// Annotation data source. 1: Outbound call, 2: Navigation
	//
	// example:
	//
	// 1
	AnnotationMissionDataSourceType *int32 `json:"AnnotationMissionDataSourceType,omitempty" xml:"AnnotationMissionDataSourceType,omitempty"`
	// List of annotation debug data sources
	//
	// > Note: An additional category 0 will be included in the return value, for example: [0,1], [0,2], or [0,1,2].
	AnnotationMissionDebugDataSourceList []*int32 `json:"AnnotationMissionDebugDataSourceList,omitempty" xml:"AnnotationMissionDebugDataSourceList,omitempty" type:"Repeated"`
	// Job ID
	//
	// example:
	//
	// ddce607f-f537-4ebd-9914-cf45671defb9
	AnnotationMissionId *string `json:"AnnotationMissionId,omitempty" xml:"AnnotationMissionId,omitempty"`
	// Task Name
	//
	// example:
	//
	// 体验场景--移车-标注任务-20230220-141347
	AnnotationMissionName *string `json:"AnnotationMissionName,omitempty" xml:"AnnotationMissionName,omitempty"`
	// Annotation status
	//
	// - 1: In progress
	//
	// - 2: Completed
	//
	// - 3: Shutdown
	//
	// example:
	//
	// 1
	AnnotationStatus *int32 `json:"AnnotationStatus,omitempty" xml:"AnnotationStatus,omitempty"`
	// End time of the conversation time filter condition for the annotation job
	//
	// example:
	//
	// 1684511999000
	ConversationTimeEndFilter *int64 `json:"ConversationTimeEndFilter,omitempty" xml:"ConversationTimeEndFilter,omitempty"`
	// Start time of the conversation time filter condition for the annotation job
	//
	// example:
	//
	// 1683216000000
	ConversationTimeStartFilter *int64 `json:"ConversationTimeStartFilter,omitempty" xml:"ConversationTimeStartFilter,omitempty"`
	// Creation Time
	//
	// example:
	//
	// 1676170339515
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// Whether to exclude call records already annotated by other jobs
	//
	// example:
	//
	// false
	ExcludeOtherMissionSession *bool `json:"ExcludeOtherMissionSession,omitempty" xml:"ExcludeOtherMissionSession,omitempty"`
	// Completion time of the annotation job
	//
	// example:
	//
	// 1683443903785
	FinishTime *int64 `json:"FinishTime,omitempty" xml:"FinishTime,omitempty"`
	// instance ID
	//
	// example:
	//
	// 32be9d94-1346-4c4a-a4d0-ccd379f87013
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// Maximum limit for sampling quantity
	//
	// example:
	//
	// 1
	SamplingCount *int32 `json:"SamplingCount,omitempty" xml:"SamplingCount,omitempty"`
	// Sampling description
	//
	// example:
	//
	// 标注
	SamplingDescription *string `json:"SamplingDescription,omitempty" xml:"SamplingDescription,omitempty"`
	// Sampling ratio
	//
	// example:
	//
	// 1
	SamplingRate *int32 `json:"SamplingRate,omitempty" xml:"SamplingRate,omitempty"`
	// Sampling type
	//
	// example:
	//
	// 1
	SamplingType *int32 `json:"SamplingType,omitempty" xml:"SamplingType,omitempty"`
	// List of call termination types used when creating the job:
	//
	// - 1: Normal completion
	//
	// - 2: Bot hang-up after failed detection
	//
	// - 3: Hang-up due to silence timeout
	//
	// - 4: User hang-up after failed detection
	//
	// - 5: User hang-up without reason
	//
	// - 6: Hit intent and transferred to human agent
	//
	// - 7: Failed detection and transferred to human agent
	//
	// - 8: No interaction from user side
	//
	// - 9: System Exception break
	//
	// - 10: Hit intent and transferred to IVR
	//
	// - 11: Failed detection and transferred to IVR
	SessionEndReasonFilterList []*int32 `json:"SessionEndReasonFilterList,omitempty" xml:"SessionEndReasonFilterList,omitempty" type:"Repeated"`
	// Number of completed annotation sessions
	//
	// example:
	//
	// 1
	SessionFinishCount *int32 `json:"SessionFinishCount,omitempty" xml:"SessionFinishCount,omitempty"`
	// Total number of annotation jobs
	//
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
