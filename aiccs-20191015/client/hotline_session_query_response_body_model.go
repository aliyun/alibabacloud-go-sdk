// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iHotlineSessionQueryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *HotlineSessionQueryResponseBody
	GetCode() *string
	SetData(v *HotlineSessionQueryResponseBodyData) *HotlineSessionQueryResponseBody
	GetData() *HotlineSessionQueryResponseBodyData
	SetMessage(v string) *HotlineSessionQueryResponseBody
	GetMessage() *string
	SetRequestId(v string) *HotlineSessionQueryResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *HotlineSessionQueryResponseBody
	GetSuccess() *bool
}

type HotlineSessionQueryResponseBody struct {
	// Status code. A value of "Success" indicates that the request succeeded.
	//
	// example:
	//
	// Success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Call data.
	Data *HotlineSessionQueryResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
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
	// EE339D98-9BD3-4413-B165
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the API call succeeded.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s HotlineSessionQueryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s HotlineSessionQueryResponseBody) GoString() string {
	return s.String()
}

func (s *HotlineSessionQueryResponseBody) GetCode() *string {
	return s.Code
}

func (s *HotlineSessionQueryResponseBody) GetData() *HotlineSessionQueryResponseBodyData {
	return s.Data
}

func (s *HotlineSessionQueryResponseBody) GetMessage() *string {
	return s.Message
}

func (s *HotlineSessionQueryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *HotlineSessionQueryResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *HotlineSessionQueryResponseBody) SetCode(v string) *HotlineSessionQueryResponseBody {
	s.Code = &v
	return s
}

func (s *HotlineSessionQueryResponseBody) SetData(v *HotlineSessionQueryResponseBodyData) *HotlineSessionQueryResponseBody {
	s.Data = v
	return s
}

func (s *HotlineSessionQueryResponseBody) SetMessage(v string) *HotlineSessionQueryResponseBody {
	s.Message = &v
	return s
}

func (s *HotlineSessionQueryResponseBody) SetRequestId(v string) *HotlineSessionQueryResponseBody {
	s.RequestId = &v
	return s
}

func (s *HotlineSessionQueryResponseBody) SetSuccess(v bool) *HotlineSessionQueryResponseBody {
	s.Success = &v
	return s
}

func (s *HotlineSessionQueryResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type HotlineSessionQueryResponseBodyData struct {
	// Call detail records.
	CallDetailRecord []*HotlineSessionQueryResponseBodyDataCallDetailRecord `json:"CallDetailRecord,omitempty" xml:"CallDetailRecord,omitempty" type:"Repeated"`
	// Current page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// Number of items per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Total number of records.
	//
	// example:
	//
	// 26
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s HotlineSessionQueryResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s HotlineSessionQueryResponseBodyData) GoString() string {
	return s.String()
}

func (s *HotlineSessionQueryResponseBodyData) GetCallDetailRecord() []*HotlineSessionQueryResponseBodyDataCallDetailRecord {
	return s.CallDetailRecord
}

func (s *HotlineSessionQueryResponseBodyData) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *HotlineSessionQueryResponseBodyData) GetPageSize() *int32 {
	return s.PageSize
}

func (s *HotlineSessionQueryResponseBodyData) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *HotlineSessionQueryResponseBodyData) SetCallDetailRecord(v []*HotlineSessionQueryResponseBodyDataCallDetailRecord) *HotlineSessionQueryResponseBodyData {
	s.CallDetailRecord = v
	return s
}

func (s *HotlineSessionQueryResponseBodyData) SetPageNumber(v int32) *HotlineSessionQueryResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *HotlineSessionQueryResponseBodyData) SetPageSize(v int32) *HotlineSessionQueryResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *HotlineSessionQueryResponseBodyData) SetTotalCount(v int32) *HotlineSessionQueryResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *HotlineSessionQueryResponseBodyData) Validate() error {
	if s.CallDetailRecord != nil {
		for _, item := range s.CallDetailRecord {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type HotlineSessionQueryResponseBodyDataCallDetailRecord struct {
	// Session ID. The acid in WebSocket after an incoming call.
	//
	// example:
	//
	// 7719786
	Acid *string `json:"Acid,omitempty" xml:"Acid,omitempty"`
	// Agent ID.
	//
	// > This field is null in non–change owner scenarios.
	//
	// example:
	//
	// 12
	ActiveTransferId *string `json:"ActiveTransferId,omitempty" xml:"ActiveTransferId,omitempty"`
	// Call duration, in seconds.
	//
	// > Unconnected calls do not include call duration.
	//
	// example:
	//
	// 37
	CallContinueTime *int32 `json:"CallContinueTime,omitempty" xml:"CallContinueTime,omitempty"`
	// Call result. Valid values:
	//
	// - **normal**: Normal hang-up.
	//
	// - **touchRouteError**: Queue hang-up.
	//
	// - **touchInQueue**: Queue hang-up.
	//
	// - **touchInLoss**: Queue hang-up.
	//
	// - **userHangup**: User hang-up or IVR hang-up.
	//
	// - **sysHangup**: System hang-up or IVR hang-up.
	//
	// - **transferAgent**: User hang-up or IVR hang-up.
	//
	// - **dailing**: Agent hang-up or ring-off hang-up.
	//
	// - **TouchRingCallLoss**: Queue hang-up or ring-off hang-up.
	//
	// example:
	//
	// normal
	CallResult *string `json:"CallResult,omitempty" xml:"CallResult,omitempty"`
	// Call type. Valid values:
	//
	// - **1**: Outbound call
	//
	// - **2**: Inbound call
	//
	// - **3**: Change owner
	//
	// example:
	//
	// 1
	CallType *int32 `json:"CallType,omitempty" xml:"CallType,omitempty"`
	// Called number.
	//
	// example:
	//
	// 135615*****
	CalledNumber *string `json:"CalledNumber,omitempty" xml:"CalledNumber,omitempty"`
	// Calling party number, such as a user\\"s phone number, agent number, or machine number.
	//
	// example:
	//
	// 0571773
	CallingNumber *string `json:"CallingNumber,omitempty" xml:"CallingNumber,omitempty"`
	// Call creation time.
	//
	// > - In outbound scenarios, this is the time when the outbound call was initiated.
	//
	// > - In inbound scenarios, this is the time when the call entered the ACC system.
	//
	// example:
	//
	// 2020-10-02 22:32:55
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// Satisfaction rating, indicated by star level. Valid values:
	//
	// - **2**: Two-star satisfaction
	//
	// - **3**: Three-star satisfaction
	//
	// - **4**: Four-star satisfaction
	//
	// - **5**: Five-star satisfaction
	//
	// > This field has no data in outbound scenarios or scenarios where the call was not answered.
	//
	// example:
	//
	// 4
	EvaluationLevel *int32 `json:"EvaluationLevel,omitempty" xml:"EvaluationLevel,omitempty"`
	// Satisfaction score. Valid values:
	//
	// - **1**: Very dissatisfied
	//
	// - **2**: Dissatisfied
	//
	// - **3**: Neutral
	//
	// - **4**: Satisfied
	//
	// - **5**: Very satisfied
	//
	// > This field has no data in outbound scenarios or scenarios where the call was not answered.
	//
	// example:
	//
	// 4
	EvaluationScore *int32 `json:"EvaluationScore,omitempty" xml:"EvaluationScore,omitempty"`
	// Skill group ID.
	//
	// > When CallType is **1**, outbound call scenarios do not include skill group information.
	//
	// example:
	//
	// 123456
	GroupId *int64 `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// Skill group name.
	//
	// > When CallType is **1**, outbound call scenarios do not include skill group information.
	//
	// example:
	//
	// 自动化技能组
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// Party that hung up. Valid values:
	//
	// - **1**: System hung up
	//
	// - **2**: Customer hung up
	//
	// - **3**: Agent hung up
	//
	// - **null**: Unknown
	//
	// example:
	//
	// 2
	HangUpRole *string `json:"HangUpRole,omitempty" xml:"HangUpRole,omitempty"`
	// Hang-up time.
	//
	// example:
	//
	// 2020-10-02 22:33:46
	HangUpTime *string `json:"HangUpTime,omitempty" xml:"HangUpTime,omitempty"`
	// The GUID of the call detail record.
	//
	// example:
	//
	// acc1c58dab4a4dd280e3813c66
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// Time when the call entered the queue for assignment.
	//
	// > Outbound call scenarios do not include queue entry time.
	//
	// example:
	//
	// 2020-10-02 22:32:55
	InQueueTime *string `json:"InQueueTime,omitempty" xml:"InQueueTime,omitempty"`
	// Membership ID.
	//
	// example:
	//
	// 7856876
	MemberId *string `json:"MemberId,omitempty" xml:"MemberId,omitempty"`
	// Membership name.
	//
	// example:
	//
	// 匿名会员
	MemberName *string `json:"MemberName,omitempty" xml:"MemberName,omitempty"`
	// The time when the hotline call is assigned and dequeued.
	//
	// > Outbound scenarios do not have a dequeue time.
	//
	// example:
	//
	// 2020-10-02 22:32:59
	OutQueueTime *string `json:"OutQueueTime,omitempty" xml:"OutQueueTime,omitempty"`
	// Agent ID. The phone number to which the call is transferred.
	//
	// > This field is null in non-transfer scenarios.
	//
	// example:
	//
	// 12
	PassiveTransferId *string `json:"PassiveTransferId,omitempty" xml:"PassiveTransferId,omitempty"`
	// The recipient of the transferred session. Valid values:
	//
	// - **1**: Agent ID.
	//
	// - **2**: Transferred phone number.
	//
	// > This field is null in non-transfer scenarios.
	//
	// example:
	//
	// 1
	PassiveTransferIdType *string `json:"PassiveTransferIdType,omitempty" xml:"PassiveTransferIdType,omitempty"`
	// The time when the call is answered.
	//
	// example:
	//
	// 2020-10-02 22:33:09
	PickUpTime *string `json:"PickUpTime,omitempty" xml:"PickUpTime,omitempty"`
	// Queue duration.
	//
	// example:
	//
	// 4
	QueueUpContinueTime *int32 `json:"QueueUpContinueTime,omitempty" xml:"QueueUpContinueTime,omitempty"`
	// Ringing duration, in seconds.
	//
	// > Outbound scenarios do not have ringing duration.
	//
	// example:
	//
	// 10
	RingContinueTime *int32 `json:"RingContinueTime,omitempty" xml:"RingContinueTime,omitempty"`
	// The time when ringing ends.
	//
	// > Outbound scenarios do not have a ring end time.
	//
	// example:
	//
	// 2020-10-02 22:33:09
	RingEndTime *string `json:"RingEndTime,omitempty" xml:"RingEndTime,omitempty"`
	// Ringing start time.
	//
	// > Outbound call scenarios do not include ringing start time.
	//
	// example:
	//
	// 2020-10-02 22:32:59
	RingStartTime *string `json:"RingStartTime,omitempty" xml:"RingStartTime,omitempty"`
	// Agent ID.
	//
	// > In inbound scenarios, agent information is unavailable until the call is assigned to an agent.
	//
	// example:
	//
	// 555555
	ServicerId *string `json:"ServicerId,omitempty" xml:"ServicerId,omitempty"`
	// Agent name.
	//
	// > Agent information is unavailable before the call is assigned to an agent in inbound scenarios.
	//
	// example:
	//
	// 刘测试
	ServicerName *string `json:"ServicerName,omitempty" xml:"ServicerName,omitempty"`
	// Long-distance call.
	//
	// example:
	//
	// 1861111****
	TrunkCall *string `json:"TrunkCall,omitempty" xml:"TrunkCall,omitempty"`
}

func (s HotlineSessionQueryResponseBodyDataCallDetailRecord) String() string {
	return dara.Prettify(s)
}

func (s HotlineSessionQueryResponseBodyDataCallDetailRecord) GoString() string {
	return s.String()
}

func (s *HotlineSessionQueryResponseBodyDataCallDetailRecord) GetAcid() *string {
	return s.Acid
}

func (s *HotlineSessionQueryResponseBodyDataCallDetailRecord) GetActiveTransferId() *string {
	return s.ActiveTransferId
}

func (s *HotlineSessionQueryResponseBodyDataCallDetailRecord) GetCallContinueTime() *int32 {
	return s.CallContinueTime
}

func (s *HotlineSessionQueryResponseBodyDataCallDetailRecord) GetCallResult() *string {
	return s.CallResult
}

func (s *HotlineSessionQueryResponseBodyDataCallDetailRecord) GetCallType() *int32 {
	return s.CallType
}

func (s *HotlineSessionQueryResponseBodyDataCallDetailRecord) GetCalledNumber() *string {
	return s.CalledNumber
}

func (s *HotlineSessionQueryResponseBodyDataCallDetailRecord) GetCallingNumber() *string {
	return s.CallingNumber
}

func (s *HotlineSessionQueryResponseBodyDataCallDetailRecord) GetCreateTime() *string {
	return s.CreateTime
}

func (s *HotlineSessionQueryResponseBodyDataCallDetailRecord) GetEvaluationLevel() *int32 {
	return s.EvaluationLevel
}

func (s *HotlineSessionQueryResponseBodyDataCallDetailRecord) GetEvaluationScore() *int32 {
	return s.EvaluationScore
}

func (s *HotlineSessionQueryResponseBodyDataCallDetailRecord) GetGroupId() *int64 {
	return s.GroupId
}

func (s *HotlineSessionQueryResponseBodyDataCallDetailRecord) GetGroupName() *string {
	return s.GroupName
}

func (s *HotlineSessionQueryResponseBodyDataCallDetailRecord) GetHangUpRole() *string {
	return s.HangUpRole
}

func (s *HotlineSessionQueryResponseBodyDataCallDetailRecord) GetHangUpTime() *string {
	return s.HangUpTime
}

func (s *HotlineSessionQueryResponseBodyDataCallDetailRecord) GetId() *string {
	return s.Id
}

func (s *HotlineSessionQueryResponseBodyDataCallDetailRecord) GetInQueueTime() *string {
	return s.InQueueTime
}

func (s *HotlineSessionQueryResponseBodyDataCallDetailRecord) GetMemberId() *string {
	return s.MemberId
}

func (s *HotlineSessionQueryResponseBodyDataCallDetailRecord) GetMemberName() *string {
	return s.MemberName
}

func (s *HotlineSessionQueryResponseBodyDataCallDetailRecord) GetOutQueueTime() *string {
	return s.OutQueueTime
}

func (s *HotlineSessionQueryResponseBodyDataCallDetailRecord) GetPassiveTransferId() *string {
	return s.PassiveTransferId
}

func (s *HotlineSessionQueryResponseBodyDataCallDetailRecord) GetPassiveTransferIdType() *string {
	return s.PassiveTransferIdType
}

func (s *HotlineSessionQueryResponseBodyDataCallDetailRecord) GetPickUpTime() *string {
	return s.PickUpTime
}

func (s *HotlineSessionQueryResponseBodyDataCallDetailRecord) GetQueueUpContinueTime() *int32 {
	return s.QueueUpContinueTime
}

func (s *HotlineSessionQueryResponseBodyDataCallDetailRecord) GetRingContinueTime() *int32 {
	return s.RingContinueTime
}

func (s *HotlineSessionQueryResponseBodyDataCallDetailRecord) GetRingEndTime() *string {
	return s.RingEndTime
}

func (s *HotlineSessionQueryResponseBodyDataCallDetailRecord) GetRingStartTime() *string {
	return s.RingStartTime
}

func (s *HotlineSessionQueryResponseBodyDataCallDetailRecord) GetServicerId() *string {
	return s.ServicerId
}

func (s *HotlineSessionQueryResponseBodyDataCallDetailRecord) GetServicerName() *string {
	return s.ServicerName
}

func (s *HotlineSessionQueryResponseBodyDataCallDetailRecord) GetTrunkCall() *string {
	return s.TrunkCall
}

func (s *HotlineSessionQueryResponseBodyDataCallDetailRecord) SetAcid(v string) *HotlineSessionQueryResponseBodyDataCallDetailRecord {
	s.Acid = &v
	return s
}

func (s *HotlineSessionQueryResponseBodyDataCallDetailRecord) SetActiveTransferId(v string) *HotlineSessionQueryResponseBodyDataCallDetailRecord {
	s.ActiveTransferId = &v
	return s
}

func (s *HotlineSessionQueryResponseBodyDataCallDetailRecord) SetCallContinueTime(v int32) *HotlineSessionQueryResponseBodyDataCallDetailRecord {
	s.CallContinueTime = &v
	return s
}

func (s *HotlineSessionQueryResponseBodyDataCallDetailRecord) SetCallResult(v string) *HotlineSessionQueryResponseBodyDataCallDetailRecord {
	s.CallResult = &v
	return s
}

func (s *HotlineSessionQueryResponseBodyDataCallDetailRecord) SetCallType(v int32) *HotlineSessionQueryResponseBodyDataCallDetailRecord {
	s.CallType = &v
	return s
}

func (s *HotlineSessionQueryResponseBodyDataCallDetailRecord) SetCalledNumber(v string) *HotlineSessionQueryResponseBodyDataCallDetailRecord {
	s.CalledNumber = &v
	return s
}

func (s *HotlineSessionQueryResponseBodyDataCallDetailRecord) SetCallingNumber(v string) *HotlineSessionQueryResponseBodyDataCallDetailRecord {
	s.CallingNumber = &v
	return s
}

func (s *HotlineSessionQueryResponseBodyDataCallDetailRecord) SetCreateTime(v string) *HotlineSessionQueryResponseBodyDataCallDetailRecord {
	s.CreateTime = &v
	return s
}

func (s *HotlineSessionQueryResponseBodyDataCallDetailRecord) SetEvaluationLevel(v int32) *HotlineSessionQueryResponseBodyDataCallDetailRecord {
	s.EvaluationLevel = &v
	return s
}

func (s *HotlineSessionQueryResponseBodyDataCallDetailRecord) SetEvaluationScore(v int32) *HotlineSessionQueryResponseBodyDataCallDetailRecord {
	s.EvaluationScore = &v
	return s
}

func (s *HotlineSessionQueryResponseBodyDataCallDetailRecord) SetGroupId(v int64) *HotlineSessionQueryResponseBodyDataCallDetailRecord {
	s.GroupId = &v
	return s
}

func (s *HotlineSessionQueryResponseBodyDataCallDetailRecord) SetGroupName(v string) *HotlineSessionQueryResponseBodyDataCallDetailRecord {
	s.GroupName = &v
	return s
}

func (s *HotlineSessionQueryResponseBodyDataCallDetailRecord) SetHangUpRole(v string) *HotlineSessionQueryResponseBodyDataCallDetailRecord {
	s.HangUpRole = &v
	return s
}

func (s *HotlineSessionQueryResponseBodyDataCallDetailRecord) SetHangUpTime(v string) *HotlineSessionQueryResponseBodyDataCallDetailRecord {
	s.HangUpTime = &v
	return s
}

func (s *HotlineSessionQueryResponseBodyDataCallDetailRecord) SetId(v string) *HotlineSessionQueryResponseBodyDataCallDetailRecord {
	s.Id = &v
	return s
}

func (s *HotlineSessionQueryResponseBodyDataCallDetailRecord) SetInQueueTime(v string) *HotlineSessionQueryResponseBodyDataCallDetailRecord {
	s.InQueueTime = &v
	return s
}

func (s *HotlineSessionQueryResponseBodyDataCallDetailRecord) SetMemberId(v string) *HotlineSessionQueryResponseBodyDataCallDetailRecord {
	s.MemberId = &v
	return s
}

func (s *HotlineSessionQueryResponseBodyDataCallDetailRecord) SetMemberName(v string) *HotlineSessionQueryResponseBodyDataCallDetailRecord {
	s.MemberName = &v
	return s
}

func (s *HotlineSessionQueryResponseBodyDataCallDetailRecord) SetOutQueueTime(v string) *HotlineSessionQueryResponseBodyDataCallDetailRecord {
	s.OutQueueTime = &v
	return s
}

func (s *HotlineSessionQueryResponseBodyDataCallDetailRecord) SetPassiveTransferId(v string) *HotlineSessionQueryResponseBodyDataCallDetailRecord {
	s.PassiveTransferId = &v
	return s
}

func (s *HotlineSessionQueryResponseBodyDataCallDetailRecord) SetPassiveTransferIdType(v string) *HotlineSessionQueryResponseBodyDataCallDetailRecord {
	s.PassiveTransferIdType = &v
	return s
}

func (s *HotlineSessionQueryResponseBodyDataCallDetailRecord) SetPickUpTime(v string) *HotlineSessionQueryResponseBodyDataCallDetailRecord {
	s.PickUpTime = &v
	return s
}

func (s *HotlineSessionQueryResponseBodyDataCallDetailRecord) SetQueueUpContinueTime(v int32) *HotlineSessionQueryResponseBodyDataCallDetailRecord {
	s.QueueUpContinueTime = &v
	return s
}

func (s *HotlineSessionQueryResponseBodyDataCallDetailRecord) SetRingContinueTime(v int32) *HotlineSessionQueryResponseBodyDataCallDetailRecord {
	s.RingContinueTime = &v
	return s
}

func (s *HotlineSessionQueryResponseBodyDataCallDetailRecord) SetRingEndTime(v string) *HotlineSessionQueryResponseBodyDataCallDetailRecord {
	s.RingEndTime = &v
	return s
}

func (s *HotlineSessionQueryResponseBodyDataCallDetailRecord) SetRingStartTime(v string) *HotlineSessionQueryResponseBodyDataCallDetailRecord {
	s.RingStartTime = &v
	return s
}

func (s *HotlineSessionQueryResponseBodyDataCallDetailRecord) SetServicerId(v string) *HotlineSessionQueryResponseBodyDataCallDetailRecord {
	s.ServicerId = &v
	return s
}

func (s *HotlineSessionQueryResponseBodyDataCallDetailRecord) SetServicerName(v string) *HotlineSessionQueryResponseBodyDataCallDetailRecord {
	s.ServicerName = &v
	return s
}

func (s *HotlineSessionQueryResponseBodyDataCallDetailRecord) SetTrunkCall(v string) *HotlineSessionQueryResponseBodyDataCallDetailRecord {
	s.TrunkCall = &v
	return s
}

func (s *HotlineSessionQueryResponseBodyDataCallDetailRecord) Validate() error {
	return dara.Validate(s)
}
