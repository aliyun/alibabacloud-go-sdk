// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCampaignResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetCampaignResponseBody
	GetCode() *string
	SetData(v *GetCampaignResponseBodyData) *GetCampaignResponseBody
	GetData() *GetCampaignResponseBodyData
	SetHttpStatusCode(v int32) *GetCampaignResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetCampaignResponseBody
	GetMessage() *string
	SetParams(v []*string) *GetCampaignResponseBody
	GetParams() []*string
	SetRequestId(v string) *GetCampaignResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetCampaignResponseBody
	GetSuccess() *bool
}

type GetCampaignResponseBody struct {
	// The result code.
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The details of the outbound campaign.
	Data *GetCampaignResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
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
	// None
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

func (s GetCampaignResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetCampaignResponseBody) GoString() string {
	return s.String()
}

func (s *GetCampaignResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetCampaignResponseBody) GetData() *GetCampaignResponseBodyData {
	return s.Data
}

func (s *GetCampaignResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetCampaignResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetCampaignResponseBody) GetParams() []*string {
	return s.Params
}

func (s *GetCampaignResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetCampaignResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetCampaignResponseBody) SetCode(v string) *GetCampaignResponseBody {
	s.Code = &v
	return s
}

func (s *GetCampaignResponseBody) SetData(v *GetCampaignResponseBodyData) *GetCampaignResponseBody {
	s.Data = v
	return s
}

func (s *GetCampaignResponseBody) SetHttpStatusCode(v int32) *GetCampaignResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetCampaignResponseBody) SetMessage(v string) *GetCampaignResponseBody {
	s.Message = &v
	return s
}

func (s *GetCampaignResponseBody) SetParams(v []*string) *GetCampaignResponseBody {
	s.Params = v
	return s
}

func (s *GetCampaignResponseBody) SetRequestId(v string) *GetCampaignResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetCampaignResponseBody) SetSuccess(v bool) *GetCampaignResponseBody {
	s.Success = &v
	return s
}

func (s *GetCampaignResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetCampaignResponseBodyData struct {
	// The actual end time.
	//
	// example:
	//
	// 1634054500000
	ActualEndTime *int64 `json:"ActualEndTime,omitempty" xml:"ActualEndTime,omitempty"`
	// The actual start time.
	//
	// example:
	//
	// 1634054400000
	ActualStartTime *int64 `json:"ActualStartTime,omitempty" xml:"ActualStartTime,omitempty"`
	// The campaign ID.
	//
	// example:
	//
	// 6ac878ab-115b-4170-a5d8-547481273364
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
	// 0.5
	CompletedRate *float64 `json:"CompletedRate,omitempty" xml:"CompletedRate,omitempty"`
	// The time when the campaign was created.
	//
	// example:
	//
	// 1735660800000
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
	// Satisfaction Survey
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The list of caller numbers.
	Numbers []*string `json:"Numbers,omitempty" xml:"Numbers,omitempty" type:"Repeated"`
	// The planned end time.
	//
	// example:
	//
	// 1634054500000
	PlannedEndTime *int64 `json:"PlannedEndTime,omitempty" xml:"PlannedEndTime,omitempty"`
	// The planned start time.
	//
	// example:
	//
	// 1634054400000
	PlannedStartTime *int64 `json:"PlannedStartTime,omitempty" xml:"PlannedStartTime,omitempty"`
	// Indicates whether the campaign continues to run until the planned end time after all contacts have been called.
	//
	// example:
	//
	// false
	RunUntilEndTime *bool `json:"RunUntilEndTime,omitempty" xml:"RunUntilEndTime,omitempty"`
	// The scenario ID.
	//
	// example:
	//
	// d13ad2d3-3fe6-4352-b38b-bd6559047de8
	ScriptId *string `json:"ScriptId,omitempty" xml:"ScriptId,omitempty"`
	// The campaign state.
	//
	// example:
	//
	// Completed
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
	// 1735660800000
	UpdatedTime *int64 `json:"UpdatedTime,omitempty" xml:"UpdatedTime,omitempty"`
	// The campaign weight.
	//
	// example:
	//
	// 1
	Weight *int32 `json:"Weight,omitempty" xml:"Weight,omitempty"`
}

func (s GetCampaignResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetCampaignResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetCampaignResponseBodyData) GetActualEndTime() *int64 {
	return s.ActualEndTime
}

func (s *GetCampaignResponseBodyData) GetActualStartTime() *int64 {
	return s.ActualStartTime
}

func (s *GetCampaignResponseBodyData) GetCampaignId() *string {
	return s.CampaignId
}

func (s *GetCampaignResponseBodyData) GetCasesAborted() *int32 {
	return s.CasesAborted
}

func (s *GetCampaignResponseBodyData) GetCasesConnected() *int32 {
	return s.CasesConnected
}

func (s *GetCampaignResponseBodyData) GetCasesUncompleted() *int32 {
	return s.CasesUncompleted
}

func (s *GetCampaignResponseBodyData) GetCasesUncompletedAfterAttempted() *int32 {
	return s.CasesUncompletedAfterAttempted
}

func (s *GetCampaignResponseBodyData) GetCompletedRate() *float64 {
	return s.CompletedRate
}

func (s *GetCampaignResponseBodyData) GetCreatedTime() *int64 {
	return s.CreatedTime
}

func (s *GetCampaignResponseBodyData) GetFixedQuota() *int32 {
	return s.FixedQuota
}

func (s *GetCampaignResponseBodyData) GetMaxAttemptCount() *int32 {
	return s.MaxAttemptCount
}

func (s *GetCampaignResponseBodyData) GetMinAttemptInterval() *int32 {
	return s.MinAttemptInterval
}

func (s *GetCampaignResponseBodyData) GetName() *string {
	return s.Name
}

func (s *GetCampaignResponseBodyData) GetNumbers() []*string {
	return s.Numbers
}

func (s *GetCampaignResponseBodyData) GetPlannedEndTime() *int64 {
	return s.PlannedEndTime
}

func (s *GetCampaignResponseBodyData) GetPlannedStartTime() *int64 {
	return s.PlannedStartTime
}

func (s *GetCampaignResponseBodyData) GetRunUntilEndTime() *bool {
	return s.RunUntilEndTime
}

func (s *GetCampaignResponseBodyData) GetScriptId() *string {
	return s.ScriptId
}

func (s *GetCampaignResponseBodyData) GetState() *string {
	return s.State
}

func (s *GetCampaignResponseBodyData) GetTotalCases() *int32 {
	return s.TotalCases
}

func (s *GetCampaignResponseBodyData) GetUpdatedTime() *int64 {
	return s.UpdatedTime
}

func (s *GetCampaignResponseBodyData) GetWeight() *int32 {
	return s.Weight
}

func (s *GetCampaignResponseBodyData) SetActualEndTime(v int64) *GetCampaignResponseBodyData {
	s.ActualEndTime = &v
	return s
}

func (s *GetCampaignResponseBodyData) SetActualStartTime(v int64) *GetCampaignResponseBodyData {
	s.ActualStartTime = &v
	return s
}

func (s *GetCampaignResponseBodyData) SetCampaignId(v string) *GetCampaignResponseBodyData {
	s.CampaignId = &v
	return s
}

func (s *GetCampaignResponseBodyData) SetCasesAborted(v int32) *GetCampaignResponseBodyData {
	s.CasesAborted = &v
	return s
}

func (s *GetCampaignResponseBodyData) SetCasesConnected(v int32) *GetCampaignResponseBodyData {
	s.CasesConnected = &v
	return s
}

func (s *GetCampaignResponseBodyData) SetCasesUncompleted(v int32) *GetCampaignResponseBodyData {
	s.CasesUncompleted = &v
	return s
}

func (s *GetCampaignResponseBodyData) SetCasesUncompletedAfterAttempted(v int32) *GetCampaignResponseBodyData {
	s.CasesUncompletedAfterAttempted = &v
	return s
}

func (s *GetCampaignResponseBodyData) SetCompletedRate(v float64) *GetCampaignResponseBodyData {
	s.CompletedRate = &v
	return s
}

func (s *GetCampaignResponseBodyData) SetCreatedTime(v int64) *GetCampaignResponseBodyData {
	s.CreatedTime = &v
	return s
}

func (s *GetCampaignResponseBodyData) SetFixedQuota(v int32) *GetCampaignResponseBodyData {
	s.FixedQuota = &v
	return s
}

func (s *GetCampaignResponseBodyData) SetMaxAttemptCount(v int32) *GetCampaignResponseBodyData {
	s.MaxAttemptCount = &v
	return s
}

func (s *GetCampaignResponseBodyData) SetMinAttemptInterval(v int32) *GetCampaignResponseBodyData {
	s.MinAttemptInterval = &v
	return s
}

func (s *GetCampaignResponseBodyData) SetName(v string) *GetCampaignResponseBodyData {
	s.Name = &v
	return s
}

func (s *GetCampaignResponseBodyData) SetNumbers(v []*string) *GetCampaignResponseBodyData {
	s.Numbers = v
	return s
}

func (s *GetCampaignResponseBodyData) SetPlannedEndTime(v int64) *GetCampaignResponseBodyData {
	s.PlannedEndTime = &v
	return s
}

func (s *GetCampaignResponseBodyData) SetPlannedStartTime(v int64) *GetCampaignResponseBodyData {
	s.PlannedStartTime = &v
	return s
}

func (s *GetCampaignResponseBodyData) SetRunUntilEndTime(v bool) *GetCampaignResponseBodyData {
	s.RunUntilEndTime = &v
	return s
}

func (s *GetCampaignResponseBodyData) SetScriptId(v string) *GetCampaignResponseBodyData {
	s.ScriptId = &v
	return s
}

func (s *GetCampaignResponseBodyData) SetState(v string) *GetCampaignResponseBodyData {
	s.State = &v
	return s
}

func (s *GetCampaignResponseBodyData) SetTotalCases(v int32) *GetCampaignResponseBodyData {
	s.TotalCases = &v
	return s
}

func (s *GetCampaignResponseBodyData) SetUpdatedTime(v int64) *GetCampaignResponseBodyData {
	s.UpdatedTime = &v
	return s
}

func (s *GetCampaignResponseBodyData) SetWeight(v int32) *GetCampaignResponseBodyData {
	s.Weight = &v
	return s
}

func (s *GetCampaignResponseBodyData) Validate() error {
	return dara.Validate(s)
}
