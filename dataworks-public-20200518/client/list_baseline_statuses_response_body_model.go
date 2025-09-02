// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListBaselineStatusesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *ListBaselineStatusesResponseBodyData) *ListBaselineStatusesResponseBody
	GetData() *ListBaselineStatusesResponseBodyData
	SetErrorCode(v string) *ListBaselineStatusesResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *ListBaselineStatusesResponseBody
	GetErrorMessage() *string
	SetHttpStatusCode(v int32) *ListBaselineStatusesResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *ListBaselineStatusesResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListBaselineStatusesResponseBody
	GetSuccess() *bool
}

type ListBaselineStatusesResponseBody struct {
	// The data returned.
	Data *ListBaselineStatusesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error code returned.
	//
	// example:
	//
	// Invalid.Tenant.ConnectionNotExists
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message returned.
	//
	// example:
	//
	// The specified parameters are invalid.
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The HTTP status code returned.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The ID of the request. You can use the ID to troubleshoot issues.
	//
	// example:
	//
	// 0000-ABCD-EFG****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListBaselineStatusesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListBaselineStatusesResponseBody) GoString() string {
	return s.String()
}

func (s *ListBaselineStatusesResponseBody) GetData() *ListBaselineStatusesResponseBodyData {
	return s.Data
}

func (s *ListBaselineStatusesResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *ListBaselineStatusesResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *ListBaselineStatusesResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListBaselineStatusesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListBaselineStatusesResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListBaselineStatusesResponseBody) SetData(v *ListBaselineStatusesResponseBodyData) *ListBaselineStatusesResponseBody {
	s.Data = v
	return s
}

func (s *ListBaselineStatusesResponseBody) SetErrorCode(v string) *ListBaselineStatusesResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListBaselineStatusesResponseBody) SetErrorMessage(v string) *ListBaselineStatusesResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *ListBaselineStatusesResponseBody) SetHttpStatusCode(v int32) *ListBaselineStatusesResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListBaselineStatusesResponseBody) SetRequestId(v string) *ListBaselineStatusesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListBaselineStatusesResponseBody) SetSuccess(v bool) *ListBaselineStatusesResponseBody {
	s.Success = &v
	return s
}

func (s *ListBaselineStatusesResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListBaselineStatusesResponseBodyData struct {
	// The list of baseline instances.
	BaselineStatuses []*ListBaselineStatusesResponseBodyDataBaselineStatuses `json:"BaselineStatuses,omitempty" xml:"BaselineStatuses,omitempty" type:"Repeated"`
	// The page number of the returned page.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries returned per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The total number of baseline instances.
	//
	// example:
	//
	// 100
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListBaselineStatusesResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListBaselineStatusesResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListBaselineStatusesResponseBodyData) GetBaselineStatuses() []*ListBaselineStatusesResponseBodyDataBaselineStatuses {
	return s.BaselineStatuses
}

func (s *ListBaselineStatusesResponseBodyData) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListBaselineStatusesResponseBodyData) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListBaselineStatusesResponseBodyData) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListBaselineStatusesResponseBodyData) SetBaselineStatuses(v []*ListBaselineStatusesResponseBodyDataBaselineStatuses) *ListBaselineStatusesResponseBodyData {
	s.BaselineStatuses = v
	return s
}

func (s *ListBaselineStatusesResponseBodyData) SetPageNumber(v int32) *ListBaselineStatusesResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *ListBaselineStatusesResponseBodyData) SetPageSize(v int32) *ListBaselineStatusesResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListBaselineStatusesResponseBodyData) SetTotalCount(v int32) *ListBaselineStatusesResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *ListBaselineStatusesResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type ListBaselineStatusesResponseBodyDataBaselineStatuses struct {
	// The baseline ID.
	//
	// example:
	//
	// 1234
	BaselineId *int64 `json:"BaselineId,omitempty" xml:"BaselineId,omitempty"`
	// The name of the baseline.
	//
	// example:
	//
	// Baseline name
	BaselineName *string `json:"BaselineName,omitempty" xml:"BaselineName,omitempty"`
	// The type of the baseline, including DAILY and HOURLY. Separate multiple types with commas (,).
	//
	// example:
	//
	// DAILY,HOURLY
	BaselineType *string `json:"BaselineType,omitempty" xml:"BaselineType,omitempty"`
	// The data timestamp.
	//
	// example:
	//
	// 1553443200000
	Bizdate *int64 `json:"Bizdate,omitempty" xml:"Bizdate,omitempty"`
	// The margin of the baseline instance. Unit: seconds.
	//
	// example:
	//
	// 1800
	Buffer *int64 `json:"Buffer,omitempty" xml:"Buffer,omitempty"`
	// The timestamp of the predicted time when the baseline instance finished running.
	//
	// example:
	//
	// 1553531400000
	EndCast *int64 `json:"EndCast,omitempty" xml:"EndCast,omitempty"`
	// The timestamp of the alerting time of the baseline instance.
	//
	// example:
	//
	// 1553531400000
	ExpTime *int64 `json:"ExpTime,omitempty" xml:"ExpTime,omitempty"`
	// The status of the baseline instance. Valid values: UNFINISH and FINISH.
	//
	// example:
	//
	// UNFINISH
	FinishStatus *string `json:"FinishStatus,omitempty" xml:"FinishStatus,omitempty"`
	// The timestamp of the actual time when the baseline instance finished running. This parameter is returned if the value of the FinishStatus parameter is FINISH.
	//
	// example:
	//
	// 1553531400000
	FinishTime *int64 `json:"FinishTime,omitempty" xml:"FinishTime,omitempty"`
	// The ID of the cycle of the baseline instance. Valid values of the ID of an hour-level cycle: [1,24]. The ID of a day-level cycle is 1.
	//
	// example:
	//
	// 1
	InGroupId *int32 `json:"InGroupId,omitempty" xml:"InGroupId,omitempty"`
	// The ID of the Alibaba Cloud account used by the baseline owner. Multiple IDs are separated by commas (,).
	//
	// example:
	//
	// 9527952795****
	Owner *string `json:"Owner,omitempty" xml:"Owner,omitempty"`
	// The priority of the baseline. Valid values: {1,3,5,7,8}.
	//
	// example:
	//
	// 1
	Priority *int32 `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// The ID of the workspace to which the baseline belongs.
	//
	// example:
	//
	// 1234
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The timestamp of the actual time when the baseline instance finished running.
	//
	// example:
	//
	// 1553531400000
	SlaTime *int64 `json:"SlaTime,omitempty" xml:"SlaTime,omitempty"`
	// The status of the baseline. Valid values: ERROR, SAFE, DANGEROUS, and OVER. The value ERROR indicates that no nodes are associated with the baseline, or all nodes associated with the baseline are suspended. The value SAFE indicates that nodes are run before the alert duration begins. The value DANGEROUS indicates that nodes are still running after the alert duration ends but the committed completion time does not arrive. The value OVER indicates that nodes are still running after the committed completion time.
	//
	// example:
	//
	// SAFE
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListBaselineStatusesResponseBodyDataBaselineStatuses) String() string {
	return dara.Prettify(s)
}

func (s ListBaselineStatusesResponseBodyDataBaselineStatuses) GoString() string {
	return s.String()
}

func (s *ListBaselineStatusesResponseBodyDataBaselineStatuses) GetBaselineId() *int64 {
	return s.BaselineId
}

func (s *ListBaselineStatusesResponseBodyDataBaselineStatuses) GetBaselineName() *string {
	return s.BaselineName
}

func (s *ListBaselineStatusesResponseBodyDataBaselineStatuses) GetBaselineType() *string {
	return s.BaselineType
}

func (s *ListBaselineStatusesResponseBodyDataBaselineStatuses) GetBizdate() *int64 {
	return s.Bizdate
}

func (s *ListBaselineStatusesResponseBodyDataBaselineStatuses) GetBuffer() *int64 {
	return s.Buffer
}

func (s *ListBaselineStatusesResponseBodyDataBaselineStatuses) GetEndCast() *int64 {
	return s.EndCast
}

func (s *ListBaselineStatusesResponseBodyDataBaselineStatuses) GetExpTime() *int64 {
	return s.ExpTime
}

func (s *ListBaselineStatusesResponseBodyDataBaselineStatuses) GetFinishStatus() *string {
	return s.FinishStatus
}

func (s *ListBaselineStatusesResponseBodyDataBaselineStatuses) GetFinishTime() *int64 {
	return s.FinishTime
}

func (s *ListBaselineStatusesResponseBodyDataBaselineStatuses) GetInGroupId() *int32 {
	return s.InGroupId
}

func (s *ListBaselineStatusesResponseBodyDataBaselineStatuses) GetOwner() *string {
	return s.Owner
}

func (s *ListBaselineStatusesResponseBodyDataBaselineStatuses) GetPriority() *int32 {
	return s.Priority
}

func (s *ListBaselineStatusesResponseBodyDataBaselineStatuses) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *ListBaselineStatusesResponseBodyDataBaselineStatuses) GetSlaTime() *int64 {
	return s.SlaTime
}

func (s *ListBaselineStatusesResponseBodyDataBaselineStatuses) GetStatus() *string {
	return s.Status
}

func (s *ListBaselineStatusesResponseBodyDataBaselineStatuses) SetBaselineId(v int64) *ListBaselineStatusesResponseBodyDataBaselineStatuses {
	s.BaselineId = &v
	return s
}

func (s *ListBaselineStatusesResponseBodyDataBaselineStatuses) SetBaselineName(v string) *ListBaselineStatusesResponseBodyDataBaselineStatuses {
	s.BaselineName = &v
	return s
}

func (s *ListBaselineStatusesResponseBodyDataBaselineStatuses) SetBaselineType(v string) *ListBaselineStatusesResponseBodyDataBaselineStatuses {
	s.BaselineType = &v
	return s
}

func (s *ListBaselineStatusesResponseBodyDataBaselineStatuses) SetBizdate(v int64) *ListBaselineStatusesResponseBodyDataBaselineStatuses {
	s.Bizdate = &v
	return s
}

func (s *ListBaselineStatusesResponseBodyDataBaselineStatuses) SetBuffer(v int64) *ListBaselineStatusesResponseBodyDataBaselineStatuses {
	s.Buffer = &v
	return s
}

func (s *ListBaselineStatusesResponseBodyDataBaselineStatuses) SetEndCast(v int64) *ListBaselineStatusesResponseBodyDataBaselineStatuses {
	s.EndCast = &v
	return s
}

func (s *ListBaselineStatusesResponseBodyDataBaselineStatuses) SetExpTime(v int64) *ListBaselineStatusesResponseBodyDataBaselineStatuses {
	s.ExpTime = &v
	return s
}

func (s *ListBaselineStatusesResponseBodyDataBaselineStatuses) SetFinishStatus(v string) *ListBaselineStatusesResponseBodyDataBaselineStatuses {
	s.FinishStatus = &v
	return s
}

func (s *ListBaselineStatusesResponseBodyDataBaselineStatuses) SetFinishTime(v int64) *ListBaselineStatusesResponseBodyDataBaselineStatuses {
	s.FinishTime = &v
	return s
}

func (s *ListBaselineStatusesResponseBodyDataBaselineStatuses) SetInGroupId(v int32) *ListBaselineStatusesResponseBodyDataBaselineStatuses {
	s.InGroupId = &v
	return s
}

func (s *ListBaselineStatusesResponseBodyDataBaselineStatuses) SetOwner(v string) *ListBaselineStatusesResponseBodyDataBaselineStatuses {
	s.Owner = &v
	return s
}

func (s *ListBaselineStatusesResponseBodyDataBaselineStatuses) SetPriority(v int32) *ListBaselineStatusesResponseBodyDataBaselineStatuses {
	s.Priority = &v
	return s
}

func (s *ListBaselineStatusesResponseBodyDataBaselineStatuses) SetProjectId(v int64) *ListBaselineStatusesResponseBodyDataBaselineStatuses {
	s.ProjectId = &v
	return s
}

func (s *ListBaselineStatusesResponseBodyDataBaselineStatuses) SetSlaTime(v int64) *ListBaselineStatusesResponseBodyDataBaselineStatuses {
	s.SlaTime = &v
	return s
}

func (s *ListBaselineStatusesResponseBodyDataBaselineStatuses) SetStatus(v string) *ListBaselineStatusesResponseBodyDataBaselineStatuses {
	s.Status = &v
	return s
}

func (s *ListBaselineStatusesResponseBodyDataBaselineStatuses) Validate() error {
	return dara.Validate(s)
}
