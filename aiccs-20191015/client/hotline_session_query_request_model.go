// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iHotlineSessionQueryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcid(v string) *HotlineSessionQueryRequest
	GetAcid() *string
	SetAcidList(v []*string) *HotlineSessionQueryRequest
	GetAcidList() []*string
	SetCallResult(v string) *HotlineSessionQueryRequest
	GetCallResult() *string
	SetCallResultList(v []*string) *HotlineSessionQueryRequest
	GetCallResultList() []*string
	SetCallType(v int32) *HotlineSessionQueryRequest
	GetCallType() *int32
	SetCallTypeList(v []*int32) *HotlineSessionQueryRequest
	GetCallTypeList() []*int32
	SetCalledNumber(v string) *HotlineSessionQueryRequest
	GetCalledNumber() *string
	SetCalledNumberList(v []*string) *HotlineSessionQueryRequest
	GetCalledNumberList() []*string
	SetCallingNumber(v string) *HotlineSessionQueryRequest
	GetCallingNumber() *string
	SetCallingNumberList(v []*string) *HotlineSessionQueryRequest
	GetCallingNumberList() []*string
	SetGroupId(v int64) *HotlineSessionQueryRequest
	GetGroupId() *int64
	SetGroupIdList(v []*int64) *HotlineSessionQueryRequest
	GetGroupIdList() []*int64
	SetGroupName(v string) *HotlineSessionQueryRequest
	GetGroupName() *string
	SetId(v string) *HotlineSessionQueryRequest
	GetId() *string
	SetInstanceId(v string) *HotlineSessionQueryRequest
	GetInstanceId() *string
	SetMemberId(v string) *HotlineSessionQueryRequest
	GetMemberId() *string
	SetMemberIdList(v []*string) *HotlineSessionQueryRequest
	GetMemberIdList() []*string
	SetMemberName(v string) *HotlineSessionQueryRequest
	GetMemberName() *string
	SetPageNo(v int32) *HotlineSessionQueryRequest
	GetPageNo() *int32
	SetPageSize(v int32) *HotlineSessionQueryRequest
	GetPageSize() *int32
	SetParams(v string) *HotlineSessionQueryRequest
	GetParams() *string
	SetQueryEndTime(v int64) *HotlineSessionQueryRequest
	GetQueryEndTime() *int64
	SetQueryStartTime(v int64) *HotlineSessionQueryRequest
	GetQueryStartTime() *int64
	SetRequestId(v string) *HotlineSessionQueryRequest
	GetRequestId() *string
	SetServicerId(v string) *HotlineSessionQueryRequest
	GetServicerId() *string
	SetServicerIdList(v []*string) *HotlineSessionQueryRequest
	GetServicerIdList() []*string
	SetServicerName(v string) *HotlineSessionQueryRequest
	GetServicerName() *string
}

type HotlineSessionQueryRequest struct {
	// The session ID. The acid in the websocket after an inbound call.
	//
	// example:
	//
	// 7719786****
	Acid *string `json:"Acid,omitempty" xml:"Acid,omitempty"`
	// The list of session IDs.
	AcidList []*string `json:"AcidList,omitempty" xml:"AcidList,omitempty" type:"Repeated"`
	// The call result. Valid values:
	//
	// - **normal**: normal hangup.
	//
	// - **touchRouteError**: queue hangup.
	//
	// - **touchInQueue**: queue hangup.
	//
	// - **touchInLoss**: queue hangup.
	//
	// - **userHangup**: user hangup or IVR hangup.
	//
	// - **sysHangup**: system hangup or IVR hangup.
	//
	// - **transferAgent**: user hangup or IVR hangup.
	//
	// - **dailing**: agent hangup or ringing hangup.
	//
	// - **TouchRingCallLoss**: queue hangup or ringing hangup.
	//
	// example:
	//
	// normal
	CallResult *string `json:"CallResult,omitempty" xml:"CallResult,omitempty"`
	// The list of call results.
	CallResultList []*string `json:"CallResultList,omitempty" xml:"CallResultList,omitempty" type:"Repeated"`
	// The call type. Valid values:
	//
	// - **1**: outbound call.
	//
	// - **2**: inbound call.
	//
	// - **3**: transferred call.
	//
	// example:
	//
	// 1
	CallType *int32 `json:"CallType,omitempty" xml:"CallType,omitempty"`
	// The list of call types.
	CallTypeList []*int32 `json:"CallTypeList,omitempty" xml:"CallTypeList,omitempty" type:"Repeated"`
	// The number of the caller. For example, a mobile phone number of a user, an agent number, or a robot number.
	//
	// example:
	//
	// 135615****
	CalledNumber *string `json:"CalledNumber,omitempty" xml:"CalledNumber,omitempty"`
	// The list of called numbers.
	CalledNumberList []*string `json:"CalledNumberList,omitempty" xml:"CalledNumberList,omitempty" type:"Repeated"`
	// The number of the callee. For example, a mobile phone number of a user, an agent number, or a robot number.
	//
	// example:
	//
	// 057177****
	CallingNumber *string `json:"CallingNumber,omitempty" xml:"CallingNumber,omitempty"`
	// The list of calling numbers.
	CallingNumberList []*string `json:"CallingNumberList,omitempty" xml:"CallingNumberList,omitempty" type:"Repeated"`
	// The ID of the skill group.
	//
	// example:
	//
	// 123456
	GroupId *int64 `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The list of skill group IDs.
	GroupIdList []*int64 `json:"GroupIdList,omitempty" xml:"GroupIdList,omitempty" type:"Repeated"`
	// The name of the skill group.
	//
	// example:
	//
	// AutomationSkillGroup
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// The globally unique ID of the call details.
	//
	// example:
	//
	// acc1c58dab4a4****0e3813c66
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The ID of the Artificial Intelligence Cloud Call Service (AICCS) instance.
	//
	// You can obtain the instance ID from <b>Instance Management</b> in the left-side navigation pane of the [Artificial Intelligence Cloud Call Service console](https://aiccs.console.aliyun.com/overview).
	//
	// This parameter is required.
	//
	// example:
	//
	// ccc_xp_pre-cn-***
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The member ID.
	//
	// example:
	//
	// 7856****
	MemberId *string `json:"MemberId,omitempty" xml:"MemberId,omitempty"`
	// The list of member IDs.
	MemberIdList []*string `json:"MemberIdList,omitempty" xml:"MemberIdList,omitempty" type:"Repeated"`
	// The member name.
	//
	// example:
	//
	// AnonymousMember
	MemberName *string `json:"MemberName,omitempty" xml:"MemberName,omitempty"`
	// The current page number. The value must be greater than **0**. Default value: **1**.
	//
	// example:
	//
	// 1
	PageNo *int32 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// The number of entries per page. The value must be greater than **0**. Default value: **20**.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The extended parameters.
	//
	// example:
	//
	// xxxx
	Params *string `json:"Params,omitempty" xml:"Params,omitempty"`
	// The end timestamp. Unit: milliseconds.
	//
	// example:
	//
	// 1614829721
	QueryEndTime *int64 `json:"QueryEndTime,omitempty" xml:"QueryEndTime,omitempty"`
	// The start timestamp. Unit: milliseconds.
	//
	// example:
	//
	// 1614828721
	QueryStartTime *int64 `json:"QueryStartTime,omitempty" xml:"QueryStartTime,omitempty"`
	// The request ID.
	//
	// example:
	//
	// EE338D98-9BD3-4413-B165
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The agent ID.
	//
	// example:
	//
	// 555555
	ServicerId *string `json:"ServicerId,omitempty" xml:"ServicerId,omitempty"`
	// The list of agent IDs.
	ServicerIdList []*string `json:"ServicerIdList,omitempty" xml:"ServicerIdList,omitempty" type:"Repeated"`
	// The agent name.
	//
	// example:
	//
	// TestAgent
	ServicerName *string `json:"ServicerName,omitempty" xml:"ServicerName,omitempty"`
}

func (s HotlineSessionQueryRequest) String() string {
	return dara.Prettify(s)
}

func (s HotlineSessionQueryRequest) GoString() string {
	return s.String()
}

func (s *HotlineSessionQueryRequest) GetAcid() *string {
	return s.Acid
}

func (s *HotlineSessionQueryRequest) GetAcidList() []*string {
	return s.AcidList
}

func (s *HotlineSessionQueryRequest) GetCallResult() *string {
	return s.CallResult
}

func (s *HotlineSessionQueryRequest) GetCallResultList() []*string {
	return s.CallResultList
}

func (s *HotlineSessionQueryRequest) GetCallType() *int32 {
	return s.CallType
}

func (s *HotlineSessionQueryRequest) GetCallTypeList() []*int32 {
	return s.CallTypeList
}

func (s *HotlineSessionQueryRequest) GetCalledNumber() *string {
	return s.CalledNumber
}

func (s *HotlineSessionQueryRequest) GetCalledNumberList() []*string {
	return s.CalledNumberList
}

func (s *HotlineSessionQueryRequest) GetCallingNumber() *string {
	return s.CallingNumber
}

func (s *HotlineSessionQueryRequest) GetCallingNumberList() []*string {
	return s.CallingNumberList
}

func (s *HotlineSessionQueryRequest) GetGroupId() *int64 {
	return s.GroupId
}

func (s *HotlineSessionQueryRequest) GetGroupIdList() []*int64 {
	return s.GroupIdList
}

func (s *HotlineSessionQueryRequest) GetGroupName() *string {
	return s.GroupName
}

func (s *HotlineSessionQueryRequest) GetId() *string {
	return s.Id
}

func (s *HotlineSessionQueryRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *HotlineSessionQueryRequest) GetMemberId() *string {
	return s.MemberId
}

func (s *HotlineSessionQueryRequest) GetMemberIdList() []*string {
	return s.MemberIdList
}

func (s *HotlineSessionQueryRequest) GetMemberName() *string {
	return s.MemberName
}

func (s *HotlineSessionQueryRequest) GetPageNo() *int32 {
	return s.PageNo
}

func (s *HotlineSessionQueryRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *HotlineSessionQueryRequest) GetParams() *string {
	return s.Params
}

func (s *HotlineSessionQueryRequest) GetQueryEndTime() *int64 {
	return s.QueryEndTime
}

func (s *HotlineSessionQueryRequest) GetQueryStartTime() *int64 {
	return s.QueryStartTime
}

func (s *HotlineSessionQueryRequest) GetRequestId() *string {
	return s.RequestId
}

func (s *HotlineSessionQueryRequest) GetServicerId() *string {
	return s.ServicerId
}

func (s *HotlineSessionQueryRequest) GetServicerIdList() []*string {
	return s.ServicerIdList
}

func (s *HotlineSessionQueryRequest) GetServicerName() *string {
	return s.ServicerName
}

func (s *HotlineSessionQueryRequest) SetAcid(v string) *HotlineSessionQueryRequest {
	s.Acid = &v
	return s
}

func (s *HotlineSessionQueryRequest) SetAcidList(v []*string) *HotlineSessionQueryRequest {
	s.AcidList = v
	return s
}

func (s *HotlineSessionQueryRequest) SetCallResult(v string) *HotlineSessionQueryRequest {
	s.CallResult = &v
	return s
}

func (s *HotlineSessionQueryRequest) SetCallResultList(v []*string) *HotlineSessionQueryRequest {
	s.CallResultList = v
	return s
}

func (s *HotlineSessionQueryRequest) SetCallType(v int32) *HotlineSessionQueryRequest {
	s.CallType = &v
	return s
}

func (s *HotlineSessionQueryRequest) SetCallTypeList(v []*int32) *HotlineSessionQueryRequest {
	s.CallTypeList = v
	return s
}

func (s *HotlineSessionQueryRequest) SetCalledNumber(v string) *HotlineSessionQueryRequest {
	s.CalledNumber = &v
	return s
}

func (s *HotlineSessionQueryRequest) SetCalledNumberList(v []*string) *HotlineSessionQueryRequest {
	s.CalledNumberList = v
	return s
}

func (s *HotlineSessionQueryRequest) SetCallingNumber(v string) *HotlineSessionQueryRequest {
	s.CallingNumber = &v
	return s
}

func (s *HotlineSessionQueryRequest) SetCallingNumberList(v []*string) *HotlineSessionQueryRequest {
	s.CallingNumberList = v
	return s
}

func (s *HotlineSessionQueryRequest) SetGroupId(v int64) *HotlineSessionQueryRequest {
	s.GroupId = &v
	return s
}

func (s *HotlineSessionQueryRequest) SetGroupIdList(v []*int64) *HotlineSessionQueryRequest {
	s.GroupIdList = v
	return s
}

func (s *HotlineSessionQueryRequest) SetGroupName(v string) *HotlineSessionQueryRequest {
	s.GroupName = &v
	return s
}

func (s *HotlineSessionQueryRequest) SetId(v string) *HotlineSessionQueryRequest {
	s.Id = &v
	return s
}

func (s *HotlineSessionQueryRequest) SetInstanceId(v string) *HotlineSessionQueryRequest {
	s.InstanceId = &v
	return s
}

func (s *HotlineSessionQueryRequest) SetMemberId(v string) *HotlineSessionQueryRequest {
	s.MemberId = &v
	return s
}

func (s *HotlineSessionQueryRequest) SetMemberIdList(v []*string) *HotlineSessionQueryRequest {
	s.MemberIdList = v
	return s
}

func (s *HotlineSessionQueryRequest) SetMemberName(v string) *HotlineSessionQueryRequest {
	s.MemberName = &v
	return s
}

func (s *HotlineSessionQueryRequest) SetPageNo(v int32) *HotlineSessionQueryRequest {
	s.PageNo = &v
	return s
}

func (s *HotlineSessionQueryRequest) SetPageSize(v int32) *HotlineSessionQueryRequest {
	s.PageSize = &v
	return s
}

func (s *HotlineSessionQueryRequest) SetParams(v string) *HotlineSessionQueryRequest {
	s.Params = &v
	return s
}

func (s *HotlineSessionQueryRequest) SetQueryEndTime(v int64) *HotlineSessionQueryRequest {
	s.QueryEndTime = &v
	return s
}

func (s *HotlineSessionQueryRequest) SetQueryStartTime(v int64) *HotlineSessionQueryRequest {
	s.QueryStartTime = &v
	return s
}

func (s *HotlineSessionQueryRequest) SetRequestId(v string) *HotlineSessionQueryRequest {
	s.RequestId = &v
	return s
}

func (s *HotlineSessionQueryRequest) SetServicerId(v string) *HotlineSessionQueryRequest {
	s.ServicerId = &v
	return s
}

func (s *HotlineSessionQueryRequest) SetServicerIdList(v []*string) *HotlineSessionQueryRequest {
	s.ServicerIdList = v
	return s
}

func (s *HotlineSessionQueryRequest) SetServicerName(v string) *HotlineSessionQueryRequest {
	s.ServicerName = &v
	return s
}

func (s *HotlineSessionQueryRequest) Validate() error {
	return dara.Validate(s)
}
