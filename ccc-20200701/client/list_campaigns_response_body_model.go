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
	SetHttpStatusCode(v int64) *ListCampaignsResponseBody
	GetHttpStatusCode() *int64
	SetMessage(v string) *ListCampaignsResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListCampaignsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListCampaignsResponseBody
	GetSuccess() *bool
}

type ListCampaignsResponseBody struct {
	// The response code.
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The data.
	Data *ListCampaignsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int64 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The response message.
	//
	// example:
	//
	// 无
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 6CCEF32F-8614-535F-A1D9-D85B8C0DC4F0
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

func (s *ListCampaignsResponseBody) GetHttpStatusCode() *int64 {
	return s.HttpStatusCode
}

func (s *ListCampaignsResponseBody) GetMessage() *string {
	return s.Message
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

func (s *ListCampaignsResponseBody) SetHttpStatusCode(v int64) *ListCampaignsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListCampaignsResponseBody) SetMessage(v string) *ListCampaignsResponseBody {
	s.Message = &v
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
	// The list of predictive dialing campaigns.
	List []*ListCampaignsResponseBodyDataList `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
	// The page number. Valid values: 1 to 100.
	//
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The page size. Valid values: 1 to 100.
	//
	// example:
	//
	// 10
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The total count.
	//
	// example:
	//
	// 1
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
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

func (s *ListCampaignsResponseBodyData) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *ListCampaignsResponseBodyData) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListCampaignsResponseBodyData) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListCampaignsResponseBodyData) SetList(v []*ListCampaignsResponseBodyDataList) *ListCampaignsResponseBodyData {
	s.List = v
	return s
}

func (s *ListCampaignsResponseBodyData) SetPageNumber(v int64) *ListCampaignsResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *ListCampaignsResponseBodyData) SetPageSize(v int64) *ListCampaignsResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListCampaignsResponseBodyData) SetTotalCount(v int64) *ListCampaignsResponseBodyData {
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
	// The actual end time of the predictive dialing campaign. The value is a UNIX timestamp in milliseconds.
	//
	// example:
	//
	// 1634008800000
	ActualEndTime *int64 `json:"ActualEndTime,omitempty" xml:"ActualEndTime,omitempty"`
	// The actual start time of the predictive dialing campaign. The value is a UNIX timestamp in milliseconds.
	//
	// example:
	//
	// 1634000460000
	ActualStartTime *int64 `json:"ActualStartTime,omitempty" xml:"ActualStartTime,omitempty"`
	// The ID of the predictive dialing campaign.
	//
	// example:
	//
	// 6badb397-a8b5-40b6-21019d382a09
	CampaignId *string `json:"CampaignId,omitempty" xml:"CampaignId,omitempty"`
	// The number of aborted cases in the predictive dialing campaign. An aborted case indicates that the call to the contact was canceled.
	//
	// example:
	//
	// 0
	CasesAborted *int64 `json:"CasesAborted,omitempty" xml:"CasesAborted,omitempty"`
	// The number of connected cases in the predictive dialing campaign.
	//
	// example:
	//
	// 40
	CasesConnected *int64 `json:"CasesConnected,omitempty" xml:"CasesConnected,omitempty"`
	// The number of uncompleted cases in the predictive dialing campaign. An uncompleted case indicates that the call was not connected and the maximum number of retry attempts was not reached.
	//
	// example:
	//
	// 0
	CasesUncompleted *int64 `json:"CasesUncompleted,omitempty" xml:"CasesUncompleted,omitempty"`
	// The completion rate. This parameter is deprecated. You can calculate the completion rate by using the formula (TotalCases - CasesUnCompleted) / TotalCases.
	//
	// example:
	//
	// 无
	CompletionRate *float32 `json:"CompletionRate,omitempty" xml:"CompletionRate,omitempty"`
	// The ID of the IVR contact flow associated with the phone number.
	//
	// example:
	//
	// a3fb6c62-9b49-4942-ae5b-cf2abd4123ek
	ContactFlowId *string `json:"ContactFlowId,omitempty" xml:"ContactFlowId,omitempty"`
	// The maximum number of attempts for the predictive dialing campaign. This value specifies the maximum number of redial attempts when a call to a number fails.
	//
	// example:
	//
	// 1
	MaxAttemptCount *int64 `json:"MaxAttemptCount,omitempty" xml:"MaxAttemptCount,omitempty"`
	// The minimum redial interval for the predictive dialing campaign. This value specifies the minimum interval between redial attempts after a failure. Unit: seconds.
	//
	// example:
	//
	// 1
	MinAttemptInterval *int64 `json:"MinAttemptInterval,omitempty" xml:"MinAttemptInterval,omitempty"`
	// The name of the predictive dialing campaign.
	//
	// example:
	//
	// test-campaign
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The planned end time of the predictive dialing campaign. The value is a UNIX timestamp in milliseconds.
	//
	// example:
	//
	// 1634054400000
	PlanedEndTime *int64 `json:"PlanedEndTime,omitempty" xml:"PlanedEndTime,omitempty"`
	// The planned start time of the predictive dialing campaign. The value is a UNIX timestamp in milliseconds.
	//
	// example:
	//
	// 1633968000000
	PlanedStartTime *int64 `json:"PlanedStartTime,omitempty" xml:"PlanedStartTime,omitempty"`
	// The ID of the associated skill group.
	//
	// example:
	//
	// skillgroup@ccc-test
	QueueId *string `json:"QueueId,omitempty" xml:"QueueId,omitempty"`
	// The name of the skill group.
	//
	// example:
	//
	// 测试技能组
	QueueName *string `json:"QueueName,omitempty" xml:"QueueName,omitempty"`
	// Indicates whether the campaign is a simulated campaign.
	//
	// example:
	//
	// false
	Simulation *bool `json:"Simulation,omitempty" xml:"Simulation,omitempty"`
	// The state of the predictive dialing campaign.
	//
	// example:
	//
	// Completed
	State *string `json:"State,omitempty" xml:"State,omitempty"`
	// The strategy parameters of the predictive dialing campaign. Example for the PID strategy: {"abandonRate":"5","historicalConnectedRate":"35"}. Example for the PACING strategy: {"ratio":1}. abandonRate specifies the expected call abandon rate. historicalConnectedRate specifies the historical reference connection rate. ratio specifies the fixed dialing ratio.
	//
	// example:
	//
	// {"ratio":1}
	StrategyParameters *string `json:"StrategyParameters,omitempty" xml:"StrategyParameters,omitempty"`
	// The strategy mode of the predictive dialing campaign.
	//
	// example:
	//
	// PACING
	StrategyType *string `json:"StrategyType,omitempty" xml:"StrategyType,omitempty"`
	// The total number of phone numbers.
	//
	// example:
	//
	// 100
	TotalCases *int64 `json:"TotalCases,omitempty" xml:"TotalCases,omitempty"`
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

func (s *ListCampaignsResponseBodyDataList) GetCasesAborted() *int64 {
	return s.CasesAborted
}

func (s *ListCampaignsResponseBodyDataList) GetCasesConnected() *int64 {
	return s.CasesConnected
}

func (s *ListCampaignsResponseBodyDataList) GetCasesUncompleted() *int64 {
	return s.CasesUncompleted
}

func (s *ListCampaignsResponseBodyDataList) GetCompletionRate() *float32 {
	return s.CompletionRate
}

func (s *ListCampaignsResponseBodyDataList) GetContactFlowId() *string {
	return s.ContactFlowId
}

func (s *ListCampaignsResponseBodyDataList) GetMaxAttemptCount() *int64 {
	return s.MaxAttemptCount
}

func (s *ListCampaignsResponseBodyDataList) GetMinAttemptInterval() *int64 {
	return s.MinAttemptInterval
}

func (s *ListCampaignsResponseBodyDataList) GetName() *string {
	return s.Name
}

func (s *ListCampaignsResponseBodyDataList) GetPlanedEndTime() *int64 {
	return s.PlanedEndTime
}

func (s *ListCampaignsResponseBodyDataList) GetPlanedStartTime() *int64 {
	return s.PlanedStartTime
}

func (s *ListCampaignsResponseBodyDataList) GetQueueId() *string {
	return s.QueueId
}

func (s *ListCampaignsResponseBodyDataList) GetQueueName() *string {
	return s.QueueName
}

func (s *ListCampaignsResponseBodyDataList) GetSimulation() *bool {
	return s.Simulation
}

func (s *ListCampaignsResponseBodyDataList) GetState() *string {
	return s.State
}

func (s *ListCampaignsResponseBodyDataList) GetStrategyParameters() *string {
	return s.StrategyParameters
}

func (s *ListCampaignsResponseBodyDataList) GetStrategyType() *string {
	return s.StrategyType
}

func (s *ListCampaignsResponseBodyDataList) GetTotalCases() *int64 {
	return s.TotalCases
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

func (s *ListCampaignsResponseBodyDataList) SetCasesAborted(v int64) *ListCampaignsResponseBodyDataList {
	s.CasesAborted = &v
	return s
}

func (s *ListCampaignsResponseBodyDataList) SetCasesConnected(v int64) *ListCampaignsResponseBodyDataList {
	s.CasesConnected = &v
	return s
}

func (s *ListCampaignsResponseBodyDataList) SetCasesUncompleted(v int64) *ListCampaignsResponseBodyDataList {
	s.CasesUncompleted = &v
	return s
}

func (s *ListCampaignsResponseBodyDataList) SetCompletionRate(v float32) *ListCampaignsResponseBodyDataList {
	s.CompletionRate = &v
	return s
}

func (s *ListCampaignsResponseBodyDataList) SetContactFlowId(v string) *ListCampaignsResponseBodyDataList {
	s.ContactFlowId = &v
	return s
}

func (s *ListCampaignsResponseBodyDataList) SetMaxAttemptCount(v int64) *ListCampaignsResponseBodyDataList {
	s.MaxAttemptCount = &v
	return s
}

func (s *ListCampaignsResponseBodyDataList) SetMinAttemptInterval(v int64) *ListCampaignsResponseBodyDataList {
	s.MinAttemptInterval = &v
	return s
}

func (s *ListCampaignsResponseBodyDataList) SetName(v string) *ListCampaignsResponseBodyDataList {
	s.Name = &v
	return s
}

func (s *ListCampaignsResponseBodyDataList) SetPlanedEndTime(v int64) *ListCampaignsResponseBodyDataList {
	s.PlanedEndTime = &v
	return s
}

func (s *ListCampaignsResponseBodyDataList) SetPlanedStartTime(v int64) *ListCampaignsResponseBodyDataList {
	s.PlanedStartTime = &v
	return s
}

func (s *ListCampaignsResponseBodyDataList) SetQueueId(v string) *ListCampaignsResponseBodyDataList {
	s.QueueId = &v
	return s
}

func (s *ListCampaignsResponseBodyDataList) SetQueueName(v string) *ListCampaignsResponseBodyDataList {
	s.QueueName = &v
	return s
}

func (s *ListCampaignsResponseBodyDataList) SetSimulation(v bool) *ListCampaignsResponseBodyDataList {
	s.Simulation = &v
	return s
}

func (s *ListCampaignsResponseBodyDataList) SetState(v string) *ListCampaignsResponseBodyDataList {
	s.State = &v
	return s
}

func (s *ListCampaignsResponseBodyDataList) SetStrategyParameters(v string) *ListCampaignsResponseBodyDataList {
	s.StrategyParameters = &v
	return s
}

func (s *ListCampaignsResponseBodyDataList) SetStrategyType(v string) *ListCampaignsResponseBodyDataList {
	s.StrategyType = &v
	return s
}

func (s *ListCampaignsResponseBodyDataList) SetTotalCases(v int64) *ListCampaignsResponseBodyDataList {
	s.TotalCases = &v
	return s
}

func (s *ListCampaignsResponseBodyDataList) Validate() error {
	return dara.Validate(s)
}
