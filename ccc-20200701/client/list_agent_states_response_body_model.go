// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAgentStatesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListAgentStatesResponseBody
	GetCode() *string
	SetData(v *ListAgentStatesResponseBodyData) *ListAgentStatesResponseBody
	GetData() *ListAgentStatesResponseBodyData
	SetHttpStatusCode(v int32) *ListAgentStatesResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ListAgentStatesResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListAgentStatesResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListAgentStatesResponseBody
	GetSuccess() *bool
}

type ListAgentStatesResponseBody struct {
	// example:
	//
	// OK
	Code *string                          `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *ListAgentStatesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 200
	HttpStatusCode *int32  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 03C67DAD-EB26-41D8-949D-9B0C470FB716
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListAgentStatesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListAgentStatesResponseBody) GoString() string {
	return s.String()
}

func (s *ListAgentStatesResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListAgentStatesResponseBody) GetData() *ListAgentStatesResponseBodyData {
	return s.Data
}

func (s *ListAgentStatesResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListAgentStatesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListAgentStatesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListAgentStatesResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListAgentStatesResponseBody) SetCode(v string) *ListAgentStatesResponseBody {
	s.Code = &v
	return s
}

func (s *ListAgentStatesResponseBody) SetData(v *ListAgentStatesResponseBodyData) *ListAgentStatesResponseBody {
	s.Data = v
	return s
}

func (s *ListAgentStatesResponseBody) SetHttpStatusCode(v int32) *ListAgentStatesResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListAgentStatesResponseBody) SetMessage(v string) *ListAgentStatesResponseBody {
	s.Message = &v
	return s
}

func (s *ListAgentStatesResponseBody) SetRequestId(v string) *ListAgentStatesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAgentStatesResponseBody) SetSuccess(v bool) *ListAgentStatesResponseBody {
	s.Success = &v
	return s
}

func (s *ListAgentStatesResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListAgentStatesResponseBodyData struct {
	List []*ListAgentStatesResponseBodyDataList `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
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
	// 11
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListAgentStatesResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListAgentStatesResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListAgentStatesResponseBodyData) GetList() []*ListAgentStatesResponseBodyDataList {
	return s.List
}

func (s *ListAgentStatesResponseBodyData) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListAgentStatesResponseBodyData) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListAgentStatesResponseBodyData) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListAgentStatesResponseBodyData) SetList(v []*ListAgentStatesResponseBodyDataList) *ListAgentStatesResponseBodyData {
	s.List = v
	return s
}

func (s *ListAgentStatesResponseBodyData) SetPageNumber(v int32) *ListAgentStatesResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *ListAgentStatesResponseBodyData) SetPageSize(v int32) *ListAgentStatesResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListAgentStatesResponseBodyData) SetTotalCount(v int32) *ListAgentStatesResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *ListAgentStatesResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type ListAgentStatesResponseBodyDataList struct {
	// example:
	//
	// agent@ccc-test
	AgentId   *string `json:"AgentId,omitempty" xml:"AgentId,omitempty"`
	AgentName *string `json:"AgentName,omitempty" xml:"AgentName,omitempty"`
	// example:
	//
	// Warm-up
	BreakCode *string `json:"BreakCode,omitempty" xml:"BreakCode,omitempty"`
	// example:
	//
	// 8030****
	Dn *string `json:"Dn,omitempty" xml:"Dn,omitempty"`
	// example:
	//
	// ccc-test
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// agent
	LoginName *string `json:"LoginName,omitempty" xml:"LoginName,omitempty"`
	// example:
	//
	// false
	OutboundScenario *bool `json:"OutboundScenario,omitempty" xml:"OutboundScenario,omitempty"`
	// example:
	//
	// Ready
	State *string `json:"State,omitempty" xml:"State,omitempty"`
	// example:
	//
	// 10
	StateDuration *string `json:"StateDuration,omitempty" xml:"StateDuration,omitempty"`
}

func (s ListAgentStatesResponseBodyDataList) String() string {
	return dara.Prettify(s)
}

func (s ListAgentStatesResponseBodyDataList) GoString() string {
	return s.String()
}

func (s *ListAgentStatesResponseBodyDataList) GetAgentId() *string {
	return s.AgentId
}

func (s *ListAgentStatesResponseBodyDataList) GetAgentName() *string {
	return s.AgentName
}

func (s *ListAgentStatesResponseBodyDataList) GetBreakCode() *string {
	return s.BreakCode
}

func (s *ListAgentStatesResponseBodyDataList) GetDn() *string {
	return s.Dn
}

func (s *ListAgentStatesResponseBodyDataList) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListAgentStatesResponseBodyDataList) GetLoginName() *string {
	return s.LoginName
}

func (s *ListAgentStatesResponseBodyDataList) GetOutboundScenario() *bool {
	return s.OutboundScenario
}

func (s *ListAgentStatesResponseBodyDataList) GetState() *string {
	return s.State
}

func (s *ListAgentStatesResponseBodyDataList) GetStateDuration() *string {
	return s.StateDuration
}

func (s *ListAgentStatesResponseBodyDataList) SetAgentId(v string) *ListAgentStatesResponseBodyDataList {
	s.AgentId = &v
	return s
}

func (s *ListAgentStatesResponseBodyDataList) SetAgentName(v string) *ListAgentStatesResponseBodyDataList {
	s.AgentName = &v
	return s
}

func (s *ListAgentStatesResponseBodyDataList) SetBreakCode(v string) *ListAgentStatesResponseBodyDataList {
	s.BreakCode = &v
	return s
}

func (s *ListAgentStatesResponseBodyDataList) SetDn(v string) *ListAgentStatesResponseBodyDataList {
	s.Dn = &v
	return s
}

func (s *ListAgentStatesResponseBodyDataList) SetInstanceId(v string) *ListAgentStatesResponseBodyDataList {
	s.InstanceId = &v
	return s
}

func (s *ListAgentStatesResponseBodyDataList) SetLoginName(v string) *ListAgentStatesResponseBodyDataList {
	s.LoginName = &v
	return s
}

func (s *ListAgentStatesResponseBodyDataList) SetOutboundScenario(v bool) *ListAgentStatesResponseBodyDataList {
	s.OutboundScenario = &v
	return s
}

func (s *ListAgentStatesResponseBodyDataList) SetState(v string) *ListAgentStatesResponseBodyDataList {
	s.State = &v
	return s
}

func (s *ListAgentStatesResponseBodyDataList) SetStateDuration(v string) *ListAgentStatesResponseBodyDataList {
	s.StateDuration = &v
	return s
}

func (s *ListAgentStatesResponseBodyDataList) Validate() error {
	return dara.Validate(s)
}
