// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListLegacyAgentStatusLogsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListLegacyAgentStatusLogsResponseBody
	GetCode() *string
	SetData(v *ListLegacyAgentStatusLogsResponseBodyData) *ListLegacyAgentStatusLogsResponseBody
	GetData() *ListLegacyAgentStatusLogsResponseBodyData
	SetHttpStatusCode(v int32) *ListLegacyAgentStatusLogsResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ListLegacyAgentStatusLogsResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListLegacyAgentStatusLogsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListLegacyAgentStatusLogsResponseBody
	GetSuccess() *bool
}

type ListLegacyAgentStatusLogsResponseBody struct {
	// example:
	//
	// OK
	Code *string                                    `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *ListLegacyAgentStatusLogsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 200
	HttpStatusCode *int32  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 4C9A5D93-33A6-57F3-A423-4C83BD0A8455
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListLegacyAgentStatusLogsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListLegacyAgentStatusLogsResponseBody) GoString() string {
	return s.String()
}

func (s *ListLegacyAgentStatusLogsResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListLegacyAgentStatusLogsResponseBody) GetData() *ListLegacyAgentStatusLogsResponseBodyData {
	return s.Data
}

func (s *ListLegacyAgentStatusLogsResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListLegacyAgentStatusLogsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListLegacyAgentStatusLogsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListLegacyAgentStatusLogsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListLegacyAgentStatusLogsResponseBody) SetCode(v string) *ListLegacyAgentStatusLogsResponseBody {
	s.Code = &v
	return s
}

func (s *ListLegacyAgentStatusLogsResponseBody) SetData(v *ListLegacyAgentStatusLogsResponseBodyData) *ListLegacyAgentStatusLogsResponseBody {
	s.Data = v
	return s
}

func (s *ListLegacyAgentStatusLogsResponseBody) SetHttpStatusCode(v int32) *ListLegacyAgentStatusLogsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListLegacyAgentStatusLogsResponseBody) SetMessage(v string) *ListLegacyAgentStatusLogsResponseBody {
	s.Message = &v
	return s
}

func (s *ListLegacyAgentStatusLogsResponseBody) SetRequestId(v string) *ListLegacyAgentStatusLogsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListLegacyAgentStatusLogsResponseBody) SetSuccess(v bool) *ListLegacyAgentStatusLogsResponseBody {
	s.Success = &v
	return s
}

func (s *ListLegacyAgentStatusLogsResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListLegacyAgentStatusLogsResponseBodyData struct {
	List []*ListLegacyAgentStatusLogsResponseBodyDataList `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 100
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 10
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListLegacyAgentStatusLogsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListLegacyAgentStatusLogsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListLegacyAgentStatusLogsResponseBodyData) GetList() []*ListLegacyAgentStatusLogsResponseBodyDataList {
	return s.List
}

func (s *ListLegacyAgentStatusLogsResponseBodyData) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListLegacyAgentStatusLogsResponseBodyData) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListLegacyAgentStatusLogsResponseBodyData) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListLegacyAgentStatusLogsResponseBodyData) SetList(v []*ListLegacyAgentStatusLogsResponseBodyDataList) *ListLegacyAgentStatusLogsResponseBodyData {
	s.List = v
	return s
}

func (s *ListLegacyAgentStatusLogsResponseBodyData) SetPageNumber(v int32) *ListLegacyAgentStatusLogsResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *ListLegacyAgentStatusLogsResponseBodyData) SetPageSize(v int32) *ListLegacyAgentStatusLogsResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListLegacyAgentStatusLogsResponseBodyData) SetTotalCount(v int32) *ListLegacyAgentStatusLogsResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *ListLegacyAgentStatusLogsResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type ListLegacyAgentStatusLogsResponseBodyDataList struct {
	// example:
	//
	// false
	AgentDropCall *string `json:"AgentDropCall,omitempty" xml:"AgentDropCall,omitempty"`
	// example:
	//
	// agent@ccc-test
	AgentNo        *string `json:"AgentNo,omitempty" xml:"AgentNo,omitempty"`
	AliHangupCause *string `json:"AliHangupCause,omitempty" xml:"AliHangupCause,omitempty"`
	CallDir        *string `json:"CallDir,omitempty" xml:"CallDir,omitempty"`
	// example:
	//
	// 454326****
	CallId   *string `json:"CallId,omitempty" xml:"CallId,omitempty"`
	CallType *string `json:"CallType,omitempty" xml:"CallType,omitempty"`
	// example:
	//
	// 1312121****
	CalleeId *string `json:"CalleeId,omitempty" xml:"CalleeId,omitempty"`
	// example:
	//
	// 08337676****
	CallerId *string `json:"CallerId,omitempty" xml:"CallerId,omitempty"`
	// example:
	//
	// 487326****
	ConnId  *string `json:"ConnId,omitempty" xml:"ConnId,omitempty"`
	Extend1 *string `json:"Extend1,omitempty" xml:"Extend1,omitempty"`
	Extend2 *string `json:"Extend2,omitempty" xml:"Extend2,omitempty"`
	Extend3 *string `json:"Extend3,omitempty" xml:"Extend3,omitempty"`
	Extend4 *string `json:"Extend4,omitempty" xml:"Extend4,omitempty"`
	// example:
	//
	// skillgroup@ccc-test
	GroupNo *string `json:"GroupNo,omitempty" xml:"GroupNo,omitempty"`
	// example:
	//
	// 8032****
	MonitedAgentNo *string `json:"MonitedAgentNo,omitempty" xml:"MonitedAgentNo,omitempty"`
	// example:
	//
	// 8012****
	MonitedAgentPhoneNo *string `json:"MonitedAgentPhoneNo,omitempty" xml:"MonitedAgentPhoneNo,omitempty"`
	// example:
	//
	// false
	OutboundScenario *bool `json:"OutboundScenario,omitempty" xml:"OutboundScenario,omitempty"`
	// example:
	//
	// 8012****
	PhoneNo *string `json:"PhoneNo,omitempty" xml:"PhoneNo,omitempty"`
	// example:
	//
	// 2021-12-03T10:15:30
	StatisticDate *string `json:"StatisticDate,omitempty" xml:"StatisticDate,omitempty"`
	// example:
	//
	// 3
	StatisticTime *int32  `json:"StatisticTime,omitempty" xml:"StatisticTime,omitempty"`
	Status        *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// skillgroup@ccc-test
	TargetRequest *string `json:"TargetRequest,omitempty" xml:"TargetRequest,omitempty"`
	// example:
	//
	// skillgroup@ccc-test
	TargetSelect *string `json:"TargetSelect,omitempty" xml:"TargetSelect,omitempty"`
	// example:
	//
	// acc6736
	TenantId *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	// example:
	//
	// 07551231****
	TransferNo *string `json:"TransferNo,omitempty" xml:"TransferNo,omitempty"`
	// example:
	//
	// 05711231****
	TransferNumber *string `json:"TransferNumber,omitempty" xml:"TransferNumber,omitempty"`
}

func (s ListLegacyAgentStatusLogsResponseBodyDataList) String() string {
	return dara.Prettify(s)
}

func (s ListLegacyAgentStatusLogsResponseBodyDataList) GoString() string {
	return s.String()
}

func (s *ListLegacyAgentStatusLogsResponseBodyDataList) GetAgentDropCall() *string {
	return s.AgentDropCall
}

func (s *ListLegacyAgentStatusLogsResponseBodyDataList) GetAgentNo() *string {
	return s.AgentNo
}

func (s *ListLegacyAgentStatusLogsResponseBodyDataList) GetAliHangupCause() *string {
	return s.AliHangupCause
}

func (s *ListLegacyAgentStatusLogsResponseBodyDataList) GetCallDir() *string {
	return s.CallDir
}

func (s *ListLegacyAgentStatusLogsResponseBodyDataList) GetCallId() *string {
	return s.CallId
}

func (s *ListLegacyAgentStatusLogsResponseBodyDataList) GetCallType() *string {
	return s.CallType
}

func (s *ListLegacyAgentStatusLogsResponseBodyDataList) GetCalleeId() *string {
	return s.CalleeId
}

func (s *ListLegacyAgentStatusLogsResponseBodyDataList) GetCallerId() *string {
	return s.CallerId
}

func (s *ListLegacyAgentStatusLogsResponseBodyDataList) GetConnId() *string {
	return s.ConnId
}

func (s *ListLegacyAgentStatusLogsResponseBodyDataList) GetExtend1() *string {
	return s.Extend1
}

func (s *ListLegacyAgentStatusLogsResponseBodyDataList) GetExtend2() *string {
	return s.Extend2
}

func (s *ListLegacyAgentStatusLogsResponseBodyDataList) GetExtend3() *string {
	return s.Extend3
}

func (s *ListLegacyAgentStatusLogsResponseBodyDataList) GetExtend4() *string {
	return s.Extend4
}

func (s *ListLegacyAgentStatusLogsResponseBodyDataList) GetGroupNo() *string {
	return s.GroupNo
}

func (s *ListLegacyAgentStatusLogsResponseBodyDataList) GetMonitedAgentNo() *string {
	return s.MonitedAgentNo
}

func (s *ListLegacyAgentStatusLogsResponseBodyDataList) GetMonitedAgentPhoneNo() *string {
	return s.MonitedAgentPhoneNo
}

func (s *ListLegacyAgentStatusLogsResponseBodyDataList) GetOutboundScenario() *bool {
	return s.OutboundScenario
}

func (s *ListLegacyAgentStatusLogsResponseBodyDataList) GetPhoneNo() *string {
	return s.PhoneNo
}

func (s *ListLegacyAgentStatusLogsResponseBodyDataList) GetStatisticDate() *string {
	return s.StatisticDate
}

func (s *ListLegacyAgentStatusLogsResponseBodyDataList) GetStatisticTime() *int32 {
	return s.StatisticTime
}

func (s *ListLegacyAgentStatusLogsResponseBodyDataList) GetStatus() *string {
	return s.Status
}

func (s *ListLegacyAgentStatusLogsResponseBodyDataList) GetTargetRequest() *string {
	return s.TargetRequest
}

func (s *ListLegacyAgentStatusLogsResponseBodyDataList) GetTargetSelect() *string {
	return s.TargetSelect
}

func (s *ListLegacyAgentStatusLogsResponseBodyDataList) GetTenantId() *string {
	return s.TenantId
}

func (s *ListLegacyAgentStatusLogsResponseBodyDataList) GetTransferNo() *string {
	return s.TransferNo
}

func (s *ListLegacyAgentStatusLogsResponseBodyDataList) GetTransferNumber() *string {
	return s.TransferNumber
}

func (s *ListLegacyAgentStatusLogsResponseBodyDataList) SetAgentDropCall(v string) *ListLegacyAgentStatusLogsResponseBodyDataList {
	s.AgentDropCall = &v
	return s
}

func (s *ListLegacyAgentStatusLogsResponseBodyDataList) SetAgentNo(v string) *ListLegacyAgentStatusLogsResponseBodyDataList {
	s.AgentNo = &v
	return s
}

func (s *ListLegacyAgentStatusLogsResponseBodyDataList) SetAliHangupCause(v string) *ListLegacyAgentStatusLogsResponseBodyDataList {
	s.AliHangupCause = &v
	return s
}

func (s *ListLegacyAgentStatusLogsResponseBodyDataList) SetCallDir(v string) *ListLegacyAgentStatusLogsResponseBodyDataList {
	s.CallDir = &v
	return s
}

func (s *ListLegacyAgentStatusLogsResponseBodyDataList) SetCallId(v string) *ListLegacyAgentStatusLogsResponseBodyDataList {
	s.CallId = &v
	return s
}

func (s *ListLegacyAgentStatusLogsResponseBodyDataList) SetCallType(v string) *ListLegacyAgentStatusLogsResponseBodyDataList {
	s.CallType = &v
	return s
}

func (s *ListLegacyAgentStatusLogsResponseBodyDataList) SetCalleeId(v string) *ListLegacyAgentStatusLogsResponseBodyDataList {
	s.CalleeId = &v
	return s
}

func (s *ListLegacyAgentStatusLogsResponseBodyDataList) SetCallerId(v string) *ListLegacyAgentStatusLogsResponseBodyDataList {
	s.CallerId = &v
	return s
}

func (s *ListLegacyAgentStatusLogsResponseBodyDataList) SetConnId(v string) *ListLegacyAgentStatusLogsResponseBodyDataList {
	s.ConnId = &v
	return s
}

func (s *ListLegacyAgentStatusLogsResponseBodyDataList) SetExtend1(v string) *ListLegacyAgentStatusLogsResponseBodyDataList {
	s.Extend1 = &v
	return s
}

func (s *ListLegacyAgentStatusLogsResponseBodyDataList) SetExtend2(v string) *ListLegacyAgentStatusLogsResponseBodyDataList {
	s.Extend2 = &v
	return s
}

func (s *ListLegacyAgentStatusLogsResponseBodyDataList) SetExtend3(v string) *ListLegacyAgentStatusLogsResponseBodyDataList {
	s.Extend3 = &v
	return s
}

func (s *ListLegacyAgentStatusLogsResponseBodyDataList) SetExtend4(v string) *ListLegacyAgentStatusLogsResponseBodyDataList {
	s.Extend4 = &v
	return s
}

func (s *ListLegacyAgentStatusLogsResponseBodyDataList) SetGroupNo(v string) *ListLegacyAgentStatusLogsResponseBodyDataList {
	s.GroupNo = &v
	return s
}

func (s *ListLegacyAgentStatusLogsResponseBodyDataList) SetMonitedAgentNo(v string) *ListLegacyAgentStatusLogsResponseBodyDataList {
	s.MonitedAgentNo = &v
	return s
}

func (s *ListLegacyAgentStatusLogsResponseBodyDataList) SetMonitedAgentPhoneNo(v string) *ListLegacyAgentStatusLogsResponseBodyDataList {
	s.MonitedAgentPhoneNo = &v
	return s
}

func (s *ListLegacyAgentStatusLogsResponseBodyDataList) SetOutboundScenario(v bool) *ListLegacyAgentStatusLogsResponseBodyDataList {
	s.OutboundScenario = &v
	return s
}

func (s *ListLegacyAgentStatusLogsResponseBodyDataList) SetPhoneNo(v string) *ListLegacyAgentStatusLogsResponseBodyDataList {
	s.PhoneNo = &v
	return s
}

func (s *ListLegacyAgentStatusLogsResponseBodyDataList) SetStatisticDate(v string) *ListLegacyAgentStatusLogsResponseBodyDataList {
	s.StatisticDate = &v
	return s
}

func (s *ListLegacyAgentStatusLogsResponseBodyDataList) SetStatisticTime(v int32) *ListLegacyAgentStatusLogsResponseBodyDataList {
	s.StatisticTime = &v
	return s
}

func (s *ListLegacyAgentStatusLogsResponseBodyDataList) SetStatus(v string) *ListLegacyAgentStatusLogsResponseBodyDataList {
	s.Status = &v
	return s
}

func (s *ListLegacyAgentStatusLogsResponseBodyDataList) SetTargetRequest(v string) *ListLegacyAgentStatusLogsResponseBodyDataList {
	s.TargetRequest = &v
	return s
}

func (s *ListLegacyAgentStatusLogsResponseBodyDataList) SetTargetSelect(v string) *ListLegacyAgentStatusLogsResponseBodyDataList {
	s.TargetSelect = &v
	return s
}

func (s *ListLegacyAgentStatusLogsResponseBodyDataList) SetTenantId(v string) *ListLegacyAgentStatusLogsResponseBodyDataList {
	s.TenantId = &v
	return s
}

func (s *ListLegacyAgentStatusLogsResponseBodyDataList) SetTransferNo(v string) *ListLegacyAgentStatusLogsResponseBodyDataList {
	s.TransferNo = &v
	return s
}

func (s *ListLegacyAgentStatusLogsResponseBodyDataList) SetTransferNumber(v string) *ListLegacyAgentStatusLogsResponseBodyDataList {
	s.TransferNumber = &v
	return s
}

func (s *ListLegacyAgentStatusLogsResponseBodyDataList) Validate() error {
	return dara.Validate(s)
}
