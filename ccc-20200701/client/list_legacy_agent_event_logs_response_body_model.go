// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListLegacyAgentEventLogsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListLegacyAgentEventLogsResponseBody
	GetCode() *string
	SetData(v *ListLegacyAgentEventLogsResponseBodyData) *ListLegacyAgentEventLogsResponseBody
	GetData() *ListLegacyAgentEventLogsResponseBodyData
	SetHttpStatusCode(v int32) *ListLegacyAgentEventLogsResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ListLegacyAgentEventLogsResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListLegacyAgentEventLogsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListLegacyAgentEventLogsResponseBody
	GetSuccess() *bool
}

type ListLegacyAgentEventLogsResponseBody struct {
	// example:
	//
	// OK
	Code *string                                   `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *ListLegacyAgentEventLogsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 200
	HttpStatusCode *int32  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 2802EE59-3B53-513A-A130-85E480AF689D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListLegacyAgentEventLogsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListLegacyAgentEventLogsResponseBody) GoString() string {
	return s.String()
}

func (s *ListLegacyAgentEventLogsResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListLegacyAgentEventLogsResponseBody) GetData() *ListLegacyAgentEventLogsResponseBodyData {
	return s.Data
}

func (s *ListLegacyAgentEventLogsResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListLegacyAgentEventLogsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListLegacyAgentEventLogsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListLegacyAgentEventLogsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListLegacyAgentEventLogsResponseBody) SetCode(v string) *ListLegacyAgentEventLogsResponseBody {
	s.Code = &v
	return s
}

func (s *ListLegacyAgentEventLogsResponseBody) SetData(v *ListLegacyAgentEventLogsResponseBodyData) *ListLegacyAgentEventLogsResponseBody {
	s.Data = v
	return s
}

func (s *ListLegacyAgentEventLogsResponseBody) SetHttpStatusCode(v int32) *ListLegacyAgentEventLogsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListLegacyAgentEventLogsResponseBody) SetMessage(v string) *ListLegacyAgentEventLogsResponseBody {
	s.Message = &v
	return s
}

func (s *ListLegacyAgentEventLogsResponseBody) SetRequestId(v string) *ListLegacyAgentEventLogsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListLegacyAgentEventLogsResponseBody) SetSuccess(v bool) *ListLegacyAgentEventLogsResponseBody {
	s.Success = &v
	return s
}

func (s *ListLegacyAgentEventLogsResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListLegacyAgentEventLogsResponseBodyData struct {
	List []*ListLegacyAgentEventLogsResponseBodyDataList `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
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

func (s ListLegacyAgentEventLogsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListLegacyAgentEventLogsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListLegacyAgentEventLogsResponseBodyData) GetList() []*ListLegacyAgentEventLogsResponseBodyDataList {
	return s.List
}

func (s *ListLegacyAgentEventLogsResponseBodyData) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListLegacyAgentEventLogsResponseBodyData) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListLegacyAgentEventLogsResponseBodyData) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListLegacyAgentEventLogsResponseBodyData) SetList(v []*ListLegacyAgentEventLogsResponseBodyDataList) *ListLegacyAgentEventLogsResponseBodyData {
	s.List = v
	return s
}

func (s *ListLegacyAgentEventLogsResponseBodyData) SetPageNumber(v int32) *ListLegacyAgentEventLogsResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *ListLegacyAgentEventLogsResponseBodyData) SetPageSize(v int32) *ListLegacyAgentEventLogsResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListLegacyAgentEventLogsResponseBodyData) SetTotalCount(v int32) *ListLegacyAgentEventLogsResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *ListLegacyAgentEventLogsResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type ListLegacyAgentEventLogsResponseBodyDataList struct {
	// example:
	//
	// false
	AgentDropCall *string `json:"AgentDropCall,omitempty" xml:"AgentDropCall,omitempty"`
	// example:
	//
	// agent@ccc-test
	AgentNo *string `json:"AgentNo,omitempty" xml:"AgentNo,omitempty"`
	CallDir *string `json:"CallDir,omitempty" xml:"CallDir,omitempty"`
	// example:
	//
	// 378654****
	CallId   *string `json:"CallId,omitempty" xml:"CallId,omitempty"`
	CallMode *string `json:"CallMode,omitempty" xml:"CallMode,omitempty"`
	CallType *string `json:"CallType,omitempty" xml:"CallType,omitempty"`
	// example:
	//
	// 1312343****
	CalleeId *string `json:"CalleeId,omitempty" xml:"CalleeId,omitempty"`
	// example:
	//
	// 08331111****
	CallerId *string `json:"CallerId,omitempty" xml:"CallerId,omitempty"`
	// example:
	//
	// 345467****
	ConnId *string `json:"ConnId,omitempty" xml:"ConnId,omitempty"`
	Event  *string `json:"Event,omitempty" xml:"Event,omitempty"`
	// example:
	//
	// skillgroup1@ccc-test,skillgroup2@ccc-test
	GroupNo *string `json:"GroupNo,omitempty" xml:"GroupNo,omitempty"`
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
	// skillgroup@ccc-test
	TargetRequest *string `json:"TargetRequest,omitempty" xml:"TargetRequest,omitempty"`
	// example:
	//
	// skillgroup@ccc-test
	TargetSelect *string `json:"TargetSelect,omitempty" xml:"TargetSelect,omitempty"`
	// example:
	//
	// acc101
	TenantId *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	// example:
	//
	// 07518888****
	TransferNumber *string `json:"TransferNumber,omitempty" xml:"TransferNumber,omitempty"`
}

func (s ListLegacyAgentEventLogsResponseBodyDataList) String() string {
	return dara.Prettify(s)
}

func (s ListLegacyAgentEventLogsResponseBodyDataList) GoString() string {
	return s.String()
}

func (s *ListLegacyAgentEventLogsResponseBodyDataList) GetAgentDropCall() *string {
	return s.AgentDropCall
}

func (s *ListLegacyAgentEventLogsResponseBodyDataList) GetAgentNo() *string {
	return s.AgentNo
}

func (s *ListLegacyAgentEventLogsResponseBodyDataList) GetCallDir() *string {
	return s.CallDir
}

func (s *ListLegacyAgentEventLogsResponseBodyDataList) GetCallId() *string {
	return s.CallId
}

func (s *ListLegacyAgentEventLogsResponseBodyDataList) GetCallMode() *string {
	return s.CallMode
}

func (s *ListLegacyAgentEventLogsResponseBodyDataList) GetCallType() *string {
	return s.CallType
}

func (s *ListLegacyAgentEventLogsResponseBodyDataList) GetCalleeId() *string {
	return s.CalleeId
}

func (s *ListLegacyAgentEventLogsResponseBodyDataList) GetCallerId() *string {
	return s.CallerId
}

func (s *ListLegacyAgentEventLogsResponseBodyDataList) GetConnId() *string {
	return s.ConnId
}

func (s *ListLegacyAgentEventLogsResponseBodyDataList) GetEvent() *string {
	return s.Event
}

func (s *ListLegacyAgentEventLogsResponseBodyDataList) GetGroupNo() *string {
	return s.GroupNo
}

func (s *ListLegacyAgentEventLogsResponseBodyDataList) GetOutboundScenario() *bool {
	return s.OutboundScenario
}

func (s *ListLegacyAgentEventLogsResponseBodyDataList) GetPhoneNo() *string {
	return s.PhoneNo
}

func (s *ListLegacyAgentEventLogsResponseBodyDataList) GetStatisticDate() *string {
	return s.StatisticDate
}

func (s *ListLegacyAgentEventLogsResponseBodyDataList) GetTargetRequest() *string {
	return s.TargetRequest
}

func (s *ListLegacyAgentEventLogsResponseBodyDataList) GetTargetSelect() *string {
	return s.TargetSelect
}

func (s *ListLegacyAgentEventLogsResponseBodyDataList) GetTenantId() *string {
	return s.TenantId
}

func (s *ListLegacyAgentEventLogsResponseBodyDataList) GetTransferNumber() *string {
	return s.TransferNumber
}

func (s *ListLegacyAgentEventLogsResponseBodyDataList) SetAgentDropCall(v string) *ListLegacyAgentEventLogsResponseBodyDataList {
	s.AgentDropCall = &v
	return s
}

func (s *ListLegacyAgentEventLogsResponseBodyDataList) SetAgentNo(v string) *ListLegacyAgentEventLogsResponseBodyDataList {
	s.AgentNo = &v
	return s
}

func (s *ListLegacyAgentEventLogsResponseBodyDataList) SetCallDir(v string) *ListLegacyAgentEventLogsResponseBodyDataList {
	s.CallDir = &v
	return s
}

func (s *ListLegacyAgentEventLogsResponseBodyDataList) SetCallId(v string) *ListLegacyAgentEventLogsResponseBodyDataList {
	s.CallId = &v
	return s
}

func (s *ListLegacyAgentEventLogsResponseBodyDataList) SetCallMode(v string) *ListLegacyAgentEventLogsResponseBodyDataList {
	s.CallMode = &v
	return s
}

func (s *ListLegacyAgentEventLogsResponseBodyDataList) SetCallType(v string) *ListLegacyAgentEventLogsResponseBodyDataList {
	s.CallType = &v
	return s
}

func (s *ListLegacyAgentEventLogsResponseBodyDataList) SetCalleeId(v string) *ListLegacyAgentEventLogsResponseBodyDataList {
	s.CalleeId = &v
	return s
}

func (s *ListLegacyAgentEventLogsResponseBodyDataList) SetCallerId(v string) *ListLegacyAgentEventLogsResponseBodyDataList {
	s.CallerId = &v
	return s
}

func (s *ListLegacyAgentEventLogsResponseBodyDataList) SetConnId(v string) *ListLegacyAgentEventLogsResponseBodyDataList {
	s.ConnId = &v
	return s
}

func (s *ListLegacyAgentEventLogsResponseBodyDataList) SetEvent(v string) *ListLegacyAgentEventLogsResponseBodyDataList {
	s.Event = &v
	return s
}

func (s *ListLegacyAgentEventLogsResponseBodyDataList) SetGroupNo(v string) *ListLegacyAgentEventLogsResponseBodyDataList {
	s.GroupNo = &v
	return s
}

func (s *ListLegacyAgentEventLogsResponseBodyDataList) SetOutboundScenario(v bool) *ListLegacyAgentEventLogsResponseBodyDataList {
	s.OutboundScenario = &v
	return s
}

func (s *ListLegacyAgentEventLogsResponseBodyDataList) SetPhoneNo(v string) *ListLegacyAgentEventLogsResponseBodyDataList {
	s.PhoneNo = &v
	return s
}

func (s *ListLegacyAgentEventLogsResponseBodyDataList) SetStatisticDate(v string) *ListLegacyAgentEventLogsResponseBodyDataList {
	s.StatisticDate = &v
	return s
}

func (s *ListLegacyAgentEventLogsResponseBodyDataList) SetTargetRequest(v string) *ListLegacyAgentEventLogsResponseBodyDataList {
	s.TargetRequest = &v
	return s
}

func (s *ListLegacyAgentEventLogsResponseBodyDataList) SetTargetSelect(v string) *ListLegacyAgentEventLogsResponseBodyDataList {
	s.TargetSelect = &v
	return s
}

func (s *ListLegacyAgentEventLogsResponseBodyDataList) SetTenantId(v string) *ListLegacyAgentEventLogsResponseBodyDataList {
	s.TenantId = &v
	return s
}

func (s *ListLegacyAgentEventLogsResponseBodyDataList) SetTransferNumber(v string) *ListLegacyAgentEventLogsResponseBodyDataList {
	s.TransferNumber = &v
	return s
}

func (s *ListLegacyAgentEventLogsResponseBodyDataList) Validate() error {
	return dara.Validate(s)
}
