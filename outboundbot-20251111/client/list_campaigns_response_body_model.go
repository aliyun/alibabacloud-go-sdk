// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCampaignsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListCampaignsResponseBody
	GetCode() *string
	SetData(v *ListCampaignsResponseBodyData) *ListCampaignsResponseBody
	GetData() *ListCampaignsResponseBodyData
	SetHttpStatusCode(v int32) *ListCampaignsResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ListCampaignsResponseBody
	GetMessage() *string
	SetParams(v []*string) *ListCampaignsResponseBody
	GetParams() []*string
	SetRequestId(v string) *ListCampaignsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListCampaignsResponseBody
	GetSuccess() *bool
}

type ListCampaignsResponseBody struct {
	// The result code.
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The paged query result.
	Data *ListCampaignsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The error message.
	//
	// example:
	//
	// Success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The list of error message parameters.
	Params []*string `json:"Params,omitempty" xml:"Params,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 254EB995-DEDF-48A4-9101-9CA5B72FFBCC
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListCampaignsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListCampaignsResponseBody) GoString() string {
	return s.String()
}

func (s *ListCampaignsResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListCampaignsResponseBody) GetData() *ListCampaignsResponseBodyData {
	return s.Data
}

func (s *ListCampaignsResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListCampaignsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListCampaignsResponseBody) GetParams() []*string {
	return s.Params
}

func (s *ListCampaignsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListCampaignsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListCampaignsResponseBody) SetCode(v string) *ListCampaignsResponseBody {
	s.Code = &v
	return s
}

func (s *ListCampaignsResponseBody) SetData(v *ListCampaignsResponseBodyData) *ListCampaignsResponseBody {
	s.Data = v
	return s
}

func (s *ListCampaignsResponseBody) SetHttpStatusCode(v int32) *ListCampaignsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListCampaignsResponseBody) SetMessage(v string) *ListCampaignsResponseBody {
	s.Message = &v
	return s
}

func (s *ListCampaignsResponseBody) SetParams(v []*string) *ListCampaignsResponseBody {
	s.Params = v
	return s
}

func (s *ListCampaignsResponseBody) SetRequestId(v string) *ListCampaignsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListCampaignsResponseBody) SetSuccess(v bool) *ListCampaignsResponseBody {
	s.Success = &v
	return s
}

func (s *ListCampaignsResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListCampaignsResponseBodyData struct {
	// The list of outbound campaigns.
	List []*ListCampaignsResponseBodyDataList `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The total number of records.
	//
	// example:
	//
	// 0
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListCampaignsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListCampaignsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListCampaignsResponseBodyData) GetList() []*ListCampaignsResponseBodyDataList {
	return s.List
}

func (s *ListCampaignsResponseBodyData) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListCampaignsResponseBodyData) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListCampaignsResponseBodyData) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListCampaignsResponseBodyData) SetList(v []*ListCampaignsResponseBodyDataList) *ListCampaignsResponseBodyData {
	s.List = v
	return s
}

func (s *ListCampaignsResponseBodyData) SetPageNumber(v int32) *ListCampaignsResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *ListCampaignsResponseBodyData) SetPageSize(v int32) *ListCampaignsResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListCampaignsResponseBodyData) SetTotalCount(v int32) *ListCampaignsResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *ListCampaignsResponseBodyData) Validate() error {
	if s.List != nil {
		for _, item := range s.List {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListCampaignsResponseBodyDataList struct {
	// The actual end time.
	//
	// example:
	//
	// 1634008800000
	ActualEndTime *int64 `json:"ActualEndTime,omitempty" xml:"ActualEndTime,omitempty"`
	// The actual start time.
	//
	// example:
	//
	// 1634008800000
	ActualStartTime *int64 `json:"ActualStartTime,omitempty" xml:"ActualStartTime,omitempty"`
	// The campaign ID.
	//
	// example:
	//
	// 7607dae1-91ad-47ea-ad76-3d81ac34f729
	CampaignId *string `json:"CampaignId,omitempty" xml:"CampaignId,omitempty"`
	// The number of aborted cases.
	//
	// example:
	//
	// 0
	CasesAborted *int32 `json:"CasesAborted,omitempty" xml:"CasesAborted,omitempty"`
	// The number of connected cases.
	//
	// example:
	//
	// 50
	CasesConnected *int32 `json:"CasesConnected,omitempty" xml:"CasesConnected,omitempty"`
	// The number of uncompleted cases.
	//
	// example:
	//
	// 0
	CasesUncompleted *int32 `json:"CasesUncompleted,omitempty" xml:"CasesUncompleted,omitempty"`
	// The number of cases that were attempted but not completed.
	//
	// example:
	//
	// 0
	CasesUncompletedAfterAttempted *int32 `json:"CasesUncompletedAfterAttempted,omitempty" xml:"CasesUncompletedAfterAttempted,omitempty"`
	// The completion rate.
	//
	// example:
	//
	// 100
	CompletedRate *float64 `json:"CompletedRate,omitempty" xml:"CompletedRate,omitempty"`
	// The time when the campaign was created.
	//
	// example:
	//
	// 2025-07-27T11:25:15+08:00
	CreatedTime *int64 `json:"CreatedTime,omitempty" xml:"CreatedTime,omitempty"`
	// The fixed number of concurrent calls.
	//
	// example:
	//
	// 0
	FixedQuota *int32 `json:"FixedQuota,omitempty" xml:"FixedQuota,omitempty"`
	// The maximum number of retries.
	//
	// example:
	//
	// 2
	MaxAttemptCount *int32 `json:"MaxAttemptCount,omitempty" xml:"MaxAttemptCount,omitempty"`
	// The minimum retry interval.
	//
	// example:
	//
	// 5
	MinAttemptInterval *int32 `json:"MinAttemptInterval,omitempty" xml:"MinAttemptInterval,omitempty"`
	// The campaign name.
	//
	// example:
	//
	// Kiaconnect本月到期续费外呼话术-OPIO_20260727_102718
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The planned end time.
	//
	// example:
	//
	// 1634008800000
	PlannedEndTime *int64 `json:"PlannedEndTime,omitempty" xml:"PlannedEndTime,omitempty"`
	// The planned start time.
	//
	// example:
	//
	// 1634008800000
	PlannedStartTime *int64 `json:"PlannedStartTime,omitempty" xml:"PlannedStartTime,omitempty"`
	// Indicates whether the campaign continues to run until the planned end time after all contacts have been called.
	//
	// example:
	//
	// false
	RunUntilEndTime *bool `json:"RunUntilEndTime,omitempty" xml:"RunUntilEndTime,omitempty"`
	// The IVR flow ID.
	//
	// example:
	//
	// 8a988bd4-6c6e-45c6-b3a5-3def5ca3bc6f
	ScriptId *string `json:"ScriptId,omitempty" xml:"ScriptId,omitempty"`
	// The campaign status.
	//
	// example:
	//
	// Executing
	State *string `json:"State,omitempty" xml:"State,omitempty"`
	// The total number of cases.
	//
	// example:
	//
	// 100
	TotalCases *int32 `json:"TotalCases,omitempty" xml:"TotalCases,omitempty"`
	// The time when the campaign was last updated.
	//
	// example:
	//
	// 1760272478
	UpdatedTime *int64 `json:"UpdatedTime,omitempty" xml:"UpdatedTime,omitempty"`
	// The weight of the campaign.
	//
	// example:
	//
	// 1
	Weight *int32 `json:"Weight,omitempty" xml:"Weight,omitempty"`
}

func (s ListCampaignsResponseBodyDataList) String() string {
	return dara.Prettify(s)
}

func (s ListCampaignsResponseBodyDataList) GoString() string {
	return s.String()
}

func (s *ListCampaignsResponseBodyDataList) GetActualEndTime() *int64 {
	return s.ActualEndTime
}

func (s *ListCampaignsResponseBodyDataList) GetActualStartTime() *int64 {
	return s.ActualStartTime
}

func (s *ListCampaignsResponseBodyDataList) GetCampaignId() *string {
	return s.CampaignId
}

func (s *ListCampaignsResponseBodyDataList) GetCasesAborted() *int32 {
	return s.CasesAborted
}

func (s *ListCampaignsResponseBodyDataList) GetCasesConnected() *int32 {
	return s.CasesConnected
}

func (s *ListCampaignsResponseBodyDataList) GetCasesUncompleted() *int32 {
	return s.CasesUncompleted
}

func (s *ListCampaignsResponseBodyDataList) GetCasesUncompletedAfterAttempted() *int32 {
	return s.CasesUncompletedAfterAttempted
}

func (s *ListCampaignsResponseBodyDataList) GetCompletedRate() *float64 {
	return s.CompletedRate
}

func (s *ListCampaignsResponseBodyDataList) GetCreatedTime() *int64 {
	return s.CreatedTime
}

func (s *ListCampaignsResponseBodyDataList) GetFixedQuota() *int32 {
	return s.FixedQuota
}

func (s *ListCampaignsResponseBodyDataList) GetMaxAttemptCount() *int32 {
	return s.MaxAttemptCount
}

func (s *ListCampaignsResponseBodyDataList) GetMinAttemptInterval() *int32 {
	return s.MinAttemptInterval
}

func (s *ListCampaignsResponseBodyDataList) GetName() *string {
	return s.Name
}

func (s *ListCampaignsResponseBodyDataList) GetPlannedEndTime() *int64 {
	return s.PlannedEndTime
}

func (s *ListCampaignsResponseBodyDataList) GetPlannedStartTime() *int64 {
	return s.PlannedStartTime
}

func (s *ListCampaignsResponseBodyDataList) GetRunUntilEndTime() *bool {
	return s.RunUntilEndTime
}

func (s *ListCampaignsResponseBodyDataList) GetScriptId() *string {
	return s.ScriptId
}

func (s *ListCampaignsResponseBodyDataList) GetState() *string {
	return s.State
}

func (s *ListCampaignsResponseBodyDataList) GetTotalCases() *int32 {
	return s.TotalCases
}

func (s *ListCampaignsResponseBodyDataList) GetUpdatedTime() *int64 {
	return s.UpdatedTime
}

func (s *ListCampaignsResponseBodyDataList) GetWeight() *int32 {
	return s.Weight
}

func (s *ListCampaignsResponseBodyDataList) SetActualEndTime(v int64) *ListCampaignsResponseBodyDataList {
	s.ActualEndTime = &v
	return s
}

func (s *ListCampaignsResponseBodyDataList) SetActualStartTime(v int64) *ListCampaignsResponseBodyDataList {
	s.ActualStartTime = &v
	return s
}

func (s *ListCampaignsResponseBodyDataList) SetCampaignId(v string) *ListCampaignsResponseBodyDataList {
	s.CampaignId = &v
	return s
}

func (s *ListCampaignsResponseBodyDataList) SetCasesAborted(v int32) *ListCampaignsResponseBodyDataList {
	s.CasesAborted = &v
	return s
}

func (s *ListCampaignsResponseBodyDataList) SetCasesConnected(v int32) *ListCampaignsResponseBodyDataList {
	s.CasesConnected = &v
	return s
}

func (s *ListCampaignsResponseBodyDataList) SetCasesUncompleted(v int32) *ListCampaignsResponseBodyDataList {
	s.CasesUncompleted = &v
	return s
}

func (s *ListCampaignsResponseBodyDataList) SetCasesUncompletedAfterAttempted(v int32) *ListCampaignsResponseBodyDataList {
	s.CasesUncompletedAfterAttempted = &v
	return s
}

func (s *ListCampaignsResponseBodyDataList) SetCompletedRate(v float64) *ListCampaignsResponseBodyDataList {
	s.CompletedRate = &v
	return s
}

func (s *ListCampaignsResponseBodyDataList) SetCreatedTime(v int64) *ListCampaignsResponseBodyDataList {
	s.CreatedTime = &v
	return s
}

func (s *ListCampaignsResponseBodyDataList) SetFixedQuota(v int32) *ListCampaignsResponseBodyDataList {
	s.FixedQuota = &v
	return s
}

func (s *ListCampaignsResponseBodyDataList) SetMaxAttemptCount(v int32) *ListCampaignsResponseBodyDataList {
	s.MaxAttemptCount = &v
	return s
}

func (s *ListCampaignsResponseBodyDataList) SetMinAttemptInterval(v int32) *ListCampaignsResponseBodyDataList {
	s.MinAttemptInterval = &v
	return s
}

func (s *ListCampaignsResponseBodyDataList) SetName(v string) *ListCampaignsResponseBodyDataList {
	s.Name = &v
	return s
}

func (s *ListCampaignsResponseBodyDataList) SetPlannedEndTime(v int64) *ListCampaignsResponseBodyDataList {
	s.PlannedEndTime = &v
	return s
}

func (s *ListCampaignsResponseBodyDataList) SetPlannedStartTime(v int64) *ListCampaignsResponseBodyDataList {
	s.PlannedStartTime = &v
	return s
}

func (s *ListCampaignsResponseBodyDataList) SetRunUntilEndTime(v bool) *ListCampaignsResponseBodyDataList {
	s.RunUntilEndTime = &v
	return s
}

func (s *ListCampaignsResponseBodyDataList) SetScriptId(v string) *ListCampaignsResponseBodyDataList {
	s.ScriptId = &v
	return s
}

func (s *ListCampaignsResponseBodyDataList) SetState(v string) *ListCampaignsResponseBodyDataList {
	s.State = &v
	return s
}

func (s *ListCampaignsResponseBodyDataList) SetTotalCases(v int32) *ListCampaignsResponseBodyDataList {
	s.TotalCases = &v
	return s
}

func (s *ListCampaignsResponseBodyDataList) SetUpdatedTime(v int64) *ListCampaignsResponseBodyDataList {
	s.UpdatedTime = &v
	return s
}

func (s *ListCampaignsResponseBodyDataList) SetWeight(v int32) *ListCampaignsResponseBodyDataList {
	s.Weight = &v
	return s
}

func (s *ListCampaignsResponseBodyDataList) Validate() error {
	return dara.Validate(s)
}
