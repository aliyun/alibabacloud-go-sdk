// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListIvrTrackingDetailsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListIvrTrackingDetailsResponseBody
	GetCode() *string
	SetData(v *ListIvrTrackingDetailsResponseBodyData) *ListIvrTrackingDetailsResponseBody
	GetData() *ListIvrTrackingDetailsResponseBodyData
	SetHttpStatusCode(v int32) *ListIvrTrackingDetailsResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ListIvrTrackingDetailsResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListIvrTrackingDetailsResponseBody
	GetRequestId() *string
}

type ListIvrTrackingDetailsResponseBody struct {
	// example:
	//
	// OK
	Code *string                                 `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *ListIvrTrackingDetailsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 200
	HttpStatusCode *int32  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// D2RB671A-3E24-4A04-81E6-6C4F5B39DF75
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListIvrTrackingDetailsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListIvrTrackingDetailsResponseBody) GoString() string {
	return s.String()
}

func (s *ListIvrTrackingDetailsResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListIvrTrackingDetailsResponseBody) GetData() *ListIvrTrackingDetailsResponseBodyData {
	return s.Data
}

func (s *ListIvrTrackingDetailsResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListIvrTrackingDetailsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListIvrTrackingDetailsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListIvrTrackingDetailsResponseBody) SetCode(v string) *ListIvrTrackingDetailsResponseBody {
	s.Code = &v
	return s
}

func (s *ListIvrTrackingDetailsResponseBody) SetData(v *ListIvrTrackingDetailsResponseBodyData) *ListIvrTrackingDetailsResponseBody {
	s.Data = v
	return s
}

func (s *ListIvrTrackingDetailsResponseBody) SetHttpStatusCode(v int32) *ListIvrTrackingDetailsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListIvrTrackingDetailsResponseBody) SetMessage(v string) *ListIvrTrackingDetailsResponseBody {
	s.Message = &v
	return s
}

func (s *ListIvrTrackingDetailsResponseBody) SetRequestId(v string) *ListIvrTrackingDetailsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListIvrTrackingDetailsResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListIvrTrackingDetailsResponseBodyData struct {
	List []*ListIvrTrackingDetailsResponseBodyDataList `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 3
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListIvrTrackingDetailsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListIvrTrackingDetailsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListIvrTrackingDetailsResponseBodyData) GetList() []*ListIvrTrackingDetailsResponseBodyDataList {
	return s.List
}

func (s *ListIvrTrackingDetailsResponseBodyData) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListIvrTrackingDetailsResponseBodyData) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListIvrTrackingDetailsResponseBodyData) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListIvrTrackingDetailsResponseBodyData) SetList(v []*ListIvrTrackingDetailsResponseBodyDataList) *ListIvrTrackingDetailsResponseBodyData {
	s.List = v
	return s
}

func (s *ListIvrTrackingDetailsResponseBodyData) SetPageNumber(v int32) *ListIvrTrackingDetailsResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *ListIvrTrackingDetailsResponseBodyData) SetPageSize(v int32) *ListIvrTrackingDetailsResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListIvrTrackingDetailsResponseBodyData) SetTotalCount(v int32) *ListIvrTrackingDetailsResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *ListIvrTrackingDetailsResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type ListIvrTrackingDetailsResponseBodyDataList struct {
	// example:
	//
	// 0533128****
	Callee *string `json:"Callee,omitempty" xml:"Callee,omitempty"`
	// example:
	//
	// 1332315****
	Caller *string `json:"Caller,omitempty" xml:"Caller,omitempty"`
	// example:
	//
	// 65cp2c15-92ac-4e67-98b2-073a3c541c5d
	ChannelId *string `json:"ChannelId,omitempty" xml:"ChannelId,omitempty"`
	// example:
	//
	// A=B;C=D
	ChannelVariables *string `json:"ChannelVariables,omitempty" xml:"ChannelVariables,omitempty"`
	// example:
	//
	// job-10963442671187****
	ContactId *string `json:"ContactId,omitempty" xml:"ContactId,omitempty"`
	// example:
	//
	// 1621910542876
	EnterTime *int64 `json:"EnterTime,omitempty" xml:"EnterTime,omitempty"`
	// example:
	//
	// abc99462-1058-47d0-a114-f145ea7444ff
	FlowId   *string `json:"FlowId,omitempty" xml:"FlowId,omitempty"`
	FlowName *string `json:"FlowName,omitempty" xml:"FlowName,omitempty"`
	// example:
	//
	// ccc-test
	Instance *string `json:"Instance,omitempty" xml:"Instance,omitempty"`
	// example:
	//
	// 1621910545105
	LeaveTime *int64 `json:"LeaveTime,omitempty" xml:"LeaveTime,omitempty"`
	// example:
	//
	// Success
	NodeExitCode *string `json:"NodeExitCode,omitempty" xml:"NodeExitCode,omitempty"`
	// example:
	//
	// e0bc19a3
	NodeId         *string                `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	NodeName       *string                `json:"NodeName,omitempty" xml:"NodeName,omitempty"`
	NodeProperties map[string]interface{} `json:"NodeProperties,omitempty" xml:"NodeProperties,omitempty"`
	// example:
	//
	// PLAY_OR_SAY
	NodeType *string `json:"NodeType,omitempty" xml:"NodeType,omitempty"`
	// example:
	//
	// {"digits":"2"}
	NodeVariables map[string]interface{} `json:"NodeVariables,omitempty" xml:"NodeVariables,omitempty"`
}

func (s ListIvrTrackingDetailsResponseBodyDataList) String() string {
	return dara.Prettify(s)
}

func (s ListIvrTrackingDetailsResponseBodyDataList) GoString() string {
	return s.String()
}

func (s *ListIvrTrackingDetailsResponseBodyDataList) GetCallee() *string {
	return s.Callee
}

func (s *ListIvrTrackingDetailsResponseBodyDataList) GetCaller() *string {
	return s.Caller
}

func (s *ListIvrTrackingDetailsResponseBodyDataList) GetChannelId() *string {
	return s.ChannelId
}

func (s *ListIvrTrackingDetailsResponseBodyDataList) GetChannelVariables() *string {
	return s.ChannelVariables
}

func (s *ListIvrTrackingDetailsResponseBodyDataList) GetContactId() *string {
	return s.ContactId
}

func (s *ListIvrTrackingDetailsResponseBodyDataList) GetEnterTime() *int64 {
	return s.EnterTime
}

func (s *ListIvrTrackingDetailsResponseBodyDataList) GetFlowId() *string {
	return s.FlowId
}

func (s *ListIvrTrackingDetailsResponseBodyDataList) GetFlowName() *string {
	return s.FlowName
}

func (s *ListIvrTrackingDetailsResponseBodyDataList) GetInstance() *string {
	return s.Instance
}

func (s *ListIvrTrackingDetailsResponseBodyDataList) GetLeaveTime() *int64 {
	return s.LeaveTime
}

func (s *ListIvrTrackingDetailsResponseBodyDataList) GetNodeExitCode() *string {
	return s.NodeExitCode
}

func (s *ListIvrTrackingDetailsResponseBodyDataList) GetNodeId() *string {
	return s.NodeId
}

func (s *ListIvrTrackingDetailsResponseBodyDataList) GetNodeName() *string {
	return s.NodeName
}

func (s *ListIvrTrackingDetailsResponseBodyDataList) GetNodeProperties() map[string]interface{} {
	return s.NodeProperties
}

func (s *ListIvrTrackingDetailsResponseBodyDataList) GetNodeType() *string {
	return s.NodeType
}

func (s *ListIvrTrackingDetailsResponseBodyDataList) GetNodeVariables() map[string]interface{} {
	return s.NodeVariables
}

func (s *ListIvrTrackingDetailsResponseBodyDataList) SetCallee(v string) *ListIvrTrackingDetailsResponseBodyDataList {
	s.Callee = &v
	return s
}

func (s *ListIvrTrackingDetailsResponseBodyDataList) SetCaller(v string) *ListIvrTrackingDetailsResponseBodyDataList {
	s.Caller = &v
	return s
}

func (s *ListIvrTrackingDetailsResponseBodyDataList) SetChannelId(v string) *ListIvrTrackingDetailsResponseBodyDataList {
	s.ChannelId = &v
	return s
}

func (s *ListIvrTrackingDetailsResponseBodyDataList) SetChannelVariables(v string) *ListIvrTrackingDetailsResponseBodyDataList {
	s.ChannelVariables = &v
	return s
}

func (s *ListIvrTrackingDetailsResponseBodyDataList) SetContactId(v string) *ListIvrTrackingDetailsResponseBodyDataList {
	s.ContactId = &v
	return s
}

func (s *ListIvrTrackingDetailsResponseBodyDataList) SetEnterTime(v int64) *ListIvrTrackingDetailsResponseBodyDataList {
	s.EnterTime = &v
	return s
}

func (s *ListIvrTrackingDetailsResponseBodyDataList) SetFlowId(v string) *ListIvrTrackingDetailsResponseBodyDataList {
	s.FlowId = &v
	return s
}

func (s *ListIvrTrackingDetailsResponseBodyDataList) SetFlowName(v string) *ListIvrTrackingDetailsResponseBodyDataList {
	s.FlowName = &v
	return s
}

func (s *ListIvrTrackingDetailsResponseBodyDataList) SetInstance(v string) *ListIvrTrackingDetailsResponseBodyDataList {
	s.Instance = &v
	return s
}

func (s *ListIvrTrackingDetailsResponseBodyDataList) SetLeaveTime(v int64) *ListIvrTrackingDetailsResponseBodyDataList {
	s.LeaveTime = &v
	return s
}

func (s *ListIvrTrackingDetailsResponseBodyDataList) SetNodeExitCode(v string) *ListIvrTrackingDetailsResponseBodyDataList {
	s.NodeExitCode = &v
	return s
}

func (s *ListIvrTrackingDetailsResponseBodyDataList) SetNodeId(v string) *ListIvrTrackingDetailsResponseBodyDataList {
	s.NodeId = &v
	return s
}

func (s *ListIvrTrackingDetailsResponseBodyDataList) SetNodeName(v string) *ListIvrTrackingDetailsResponseBodyDataList {
	s.NodeName = &v
	return s
}

func (s *ListIvrTrackingDetailsResponseBodyDataList) SetNodeProperties(v map[string]interface{}) *ListIvrTrackingDetailsResponseBodyDataList {
	s.NodeProperties = v
	return s
}

func (s *ListIvrTrackingDetailsResponseBodyDataList) SetNodeType(v string) *ListIvrTrackingDetailsResponseBodyDataList {
	s.NodeType = &v
	return s
}

func (s *ListIvrTrackingDetailsResponseBodyDataList) SetNodeVariables(v map[string]interface{}) *ListIvrTrackingDetailsResponseBodyDataList {
	s.NodeVariables = v
	return s
}

func (s *ListIvrTrackingDetailsResponseBodyDataList) Validate() error {
	return dara.Validate(s)
}
