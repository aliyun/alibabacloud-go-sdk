// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iClinkAgentStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *ClinkAgentStatusResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *ClinkAgentStatusResponseBody
	GetCode() *string
	SetData(v *ClinkAgentStatusResponseBodyData) *ClinkAgentStatusResponseBody
	GetData() *ClinkAgentStatusResponseBodyData
	SetMessage(v string) *ClinkAgentStatusResponseBody
	GetMessage() *string
	SetRequestId(v string) *ClinkAgentStatusResponseBody
	GetRequestId() *string
}

type ClinkAgentStatusResponseBody struct {
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// example:
	//
	// OK
	Code *string                           `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *ClinkAgentStatusResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// A90E4451-FED7-49D2-87C8-00700A8C4D0D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ClinkAgentStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ClinkAgentStatusResponseBody) GoString() string {
	return s.String()
}

func (s *ClinkAgentStatusResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *ClinkAgentStatusResponseBody) GetCode() *string {
	return s.Code
}

func (s *ClinkAgentStatusResponseBody) GetData() *ClinkAgentStatusResponseBodyData {
	return s.Data
}

func (s *ClinkAgentStatusResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ClinkAgentStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ClinkAgentStatusResponseBody) SetAccessDeniedDetail(v string) *ClinkAgentStatusResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *ClinkAgentStatusResponseBody) SetCode(v string) *ClinkAgentStatusResponseBody {
	s.Code = &v
	return s
}

func (s *ClinkAgentStatusResponseBody) SetData(v *ClinkAgentStatusResponseBodyData) *ClinkAgentStatusResponseBody {
	s.Data = v
	return s
}

func (s *ClinkAgentStatusResponseBody) SetMessage(v string) *ClinkAgentStatusResponseBody {
	s.Message = &v
	return s
}

func (s *ClinkAgentStatusResponseBody) SetRequestId(v string) *ClinkAgentStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *ClinkAgentStatusResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ClinkAgentStatusResponseBodyData struct {
	// [座席状态详情] #座席状态详情
	AgentStatus []*ClinkAgentStatusResponseBodyDataAgentStatus `json:"AgentStatus,omitempty" xml:"AgentStatus,omitempty" type:"Repeated"`
	// 请求 id
	//
	// example:
	//
	// null
	ClinkRequestId *string `json:"ClinkRequestId,omitempty" xml:"ClinkRequestId,omitempty"`
	// 当前页码
	//
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// 一页展示条数
	//
	// example:
	//
	// 10
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// 总条数
	//
	// example:
	//
	// 63
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ClinkAgentStatusResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ClinkAgentStatusResponseBodyData) GoString() string {
	return s.String()
}

func (s *ClinkAgentStatusResponseBodyData) GetAgentStatus() []*ClinkAgentStatusResponseBodyDataAgentStatus {
	return s.AgentStatus
}

func (s *ClinkAgentStatusResponseBodyData) GetClinkRequestId() *string {
	return s.ClinkRequestId
}

func (s *ClinkAgentStatusResponseBodyData) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *ClinkAgentStatusResponseBodyData) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ClinkAgentStatusResponseBodyData) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ClinkAgentStatusResponseBodyData) SetAgentStatus(v []*ClinkAgentStatusResponseBodyDataAgentStatus) *ClinkAgentStatusResponseBodyData {
	s.AgentStatus = v
	return s
}

func (s *ClinkAgentStatusResponseBodyData) SetClinkRequestId(v string) *ClinkAgentStatusResponseBodyData {
	s.ClinkRequestId = &v
	return s
}

func (s *ClinkAgentStatusResponseBodyData) SetPageNumber(v int64) *ClinkAgentStatusResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *ClinkAgentStatusResponseBodyData) SetPageSize(v int64) *ClinkAgentStatusResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ClinkAgentStatusResponseBodyData) SetTotalCount(v int64) *ClinkAgentStatusResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *ClinkAgentStatusResponseBodyData) Validate() error {
	if s.AgentStatus != nil {
		for _, item := range s.AgentStatus {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ClinkAgentStatusResponseBodyDataAgentStatus struct {
	// 座席状态
	//
	// example:
	//
	// 示例值
	AgentStatus *string `json:"AgentStatus,omitempty" xml:"AgentStatus,omitempty"`
	// 座席状态详情
	//
	// example:
	//
	// 示例值
	AgentStatusDetail *string `json:"AgentStatusDetail,omitempty" xml:"AgentStatusDetail,omitempty"`
	// 绑定号码
	//
	// example:
	//
	// 1327770
	BindTel *string `json:"BindTel,omitempty" xml:"BindTel,omitempty"`
	// 电话类型，1:电话；2:IP话机；3:软电话
	//
	// example:
	//
	// 19
	BindType *int64 `json:"BindType,omitempty" xml:"BindType,omitempty"`
	// 座席来电接听次数
	//
	// example:
	//
	// 59
	BridgeCallCount *int64 `json:"BridgeCallCount,omitempty" xml:"BridgeCallCount,omitempty"`
	// 座席名称
	//
	// example:
	//
	// name1
	ClientName *string `json:"ClientName,omitempty" xml:"ClientName,omitempty"`
	// 座席号
	//
	// example:
	//
	// 1111
	Cno *string `json:"Cno,omitempty" xml:"Cno,omitempty"`
	// 座席code
	//
	// example:
	//
	// 示例值示例值
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// 客户号码
	//
	// example:
	//
	// 177xxxx7750
	CustomerNumber *string `json:"CustomerNumber,omitempty" xml:"CustomerNumber,omitempty"`
	// 客户来电号码加密串
	//
	// example:
	//
	// ******
	CustomerNumberEncrypt *string `json:"CustomerNumberEncrypt,omitempty" xml:"CustomerNumberEncrypt,omitempty"`
	// 客户状态，空闲，呼叫中，响铃，通话
	//
	// example:
	//
	// 示例值示例值
	CustomerStatus *string `json:"CustomerStatus,omitempty" xml:"CustomerStatus,omitempty"`
	// 呼入接听总响铃时长
	//
	// example:
	//
	// 15
	IbAnsweredRingingDuration *int64 `json:"IbAnsweredRingingDuration,omitempty" xml:"IbAnsweredRingingDuration,omitempty"`
	// 客户呼入接听数
	//
	// example:
	//
	// 66
	IbClientAnsweredCount *int64 `json:"IbClientAnsweredCount,omitempty" xml:"IbClientAnsweredCount,omitempty"`
	// 客户呼入接听率
	//
	// example:
	//
	// 0.0
	IbClientAnsweredRate *float64 `json:"IbClientAnsweredRate,omitempty" xml:"IbClientAnsweredRate,omitempty"`
	// 客户呼入数
	//
	// example:
	//
	// 15
	IbClientTotalCount *int64 `json:"IbClientTotalCount,omitempty" xml:"IbClientTotalCount,omitempty"`
	// 客户呼入未接听数
	//
	// example:
	//
	// 71
	IbClientUnansweredCount *int64 `json:"IbClientUnansweredCount,omitempty" xml:"IbClientUnansweredCount,omitempty"`
	// 空闲超时时间（单位分钟）
	//
	// example:
	//
	// 83
	IdleTimeOut *int64 `json:"IdleTimeOut,omitempty" xml:"IdleTimeOut,omitempty"`
	// 座席来电次数
	//
	// example:
	//
	// 29
	IncomingCallCount *int64 `json:"IncomingCallCount,omitempty" xml:"IncomingCallCount,omitempty"`
	// 登录时长
	//
	// example:
	//
	// 97
	LoginDuration *int64 `json:"LoginDuration,omitempty" xml:"LoginDuration,omitempty"`
	// 外呼平均通话时长
	//
	// example:
	//
	// 25
	ObAvgBridgeDuration *int64 `json:"ObAvgBridgeDuration,omitempty" xml:"ObAvgBridgeDuration,omitempty"`
	// 外呼通话时长
	//
	// example:
	//
	// 59
	ObBridgeDuration *int64 `json:"ObBridgeDuration,omitempty" xml:"ObBridgeDuration,omitempty"`
	// 外呼接听率
	//
	// example:
	//
	// 0.0
	ObBridgeRate *float64 `json:"ObBridgeRate,omitempty" xml:"ObBridgeRate,omitempty"`
	// 座席外呼数
	//
	// example:
	//
	// 96
	ObCallCount *int64 `json:"ObCallCount,omitempty" xml:"ObCallCount,omitempty"`
	// 外呼座席未接听数
	//
	// example:
	//
	// 75
	ObClientUnbridgeCount *int64 `json:"ObClientUnbridgeCount,omitempty" xml:"ObClientUnbridgeCount,omitempty"`
	// 外呼客户接听数
	//
	// example:
	//
	// 55
	ObCustomerBridgeCount *int64 `json:"ObCustomerBridgeCount,omitempty" xml:"ObCustomerBridgeCount,omitempty"`
	// 外呼客户未接听数
	//
	// example:
	//
	// 15
	ObCustomerUnbridgeCount *int64 `json:"ObCustomerUnbridgeCount,omitempty" xml:"ObCustomerUnbridgeCount,omitempty"`
	// 置忙状态描述
	//
	// example:
	//
	// 示例值示例值示例值
	PauseDescription *string `json:"PauseDescription,omitempty" xml:"PauseDescription,omitempty"`
	// 置忙类型
	//
	// example:
	//
	// 87
	PauseType *int64 `json:"PauseType,omitempty" xml:"PauseType,omitempty"`
	// 座席是否处于预测外呼状态 1：是 0：否
	//
	// example:
	//
	// 0
	PredictToCall *int64 `json:"PredictToCall,omitempty" xml:"PredictToCall,omitempty"`
	// 队列来电接听数
	//
	// example:
	//
	// 20
	QueueIncomingCallCount *int64 `json:"QueueIncomingCallCount,omitempty" xml:"QueueIncomingCallCount,omitempty"`
	// 队列信息
	//
	// example:
	//
	// null
	QueueInfo *string `json:"QueueInfo,omitempty" xml:"QueueInfo,omitempty"`
	// 队列编号
	//
	// example:
	//
	// 0035
	Queues *string `json:"Queues,omitempty" xml:"Queues,omitempty"`
	// 客户速挂数
	//
	// example:
	//
	// 63
	QuickUnlinkCount *int64 `json:"QuickUnlinkCount,omitempty" xml:"QuickUnlinkCount,omitempty"`
	// 状态时长
	//
	// example:
	//
	// 89
	StateDuration *int64 `json:"StateDuration,omitempty" xml:"StateDuration,omitempty"`
	// 座席标签
	TagNames []*string `json:"TagNames,omitempty" xml:"TagNames,omitempty" type:"Repeated"`
}

func (s ClinkAgentStatusResponseBodyDataAgentStatus) String() string {
	return dara.Prettify(s)
}

func (s ClinkAgentStatusResponseBodyDataAgentStatus) GoString() string {
	return s.String()
}

func (s *ClinkAgentStatusResponseBodyDataAgentStatus) GetAgentStatus() *string {
	return s.AgentStatus
}

func (s *ClinkAgentStatusResponseBodyDataAgentStatus) GetAgentStatusDetail() *string {
	return s.AgentStatusDetail
}

func (s *ClinkAgentStatusResponseBodyDataAgentStatus) GetBindTel() *string {
	return s.BindTel
}

func (s *ClinkAgentStatusResponseBodyDataAgentStatus) GetBindType() *int64 {
	return s.BindType
}

func (s *ClinkAgentStatusResponseBodyDataAgentStatus) GetBridgeCallCount() *int64 {
	return s.BridgeCallCount
}

func (s *ClinkAgentStatusResponseBodyDataAgentStatus) GetClientName() *string {
	return s.ClientName
}

func (s *ClinkAgentStatusResponseBodyDataAgentStatus) GetCno() *string {
	return s.Cno
}

func (s *ClinkAgentStatusResponseBodyDataAgentStatus) GetCode() *string {
	return s.Code
}

func (s *ClinkAgentStatusResponseBodyDataAgentStatus) GetCustomerNumber() *string {
	return s.CustomerNumber
}

func (s *ClinkAgentStatusResponseBodyDataAgentStatus) GetCustomerNumberEncrypt() *string {
	return s.CustomerNumberEncrypt
}

func (s *ClinkAgentStatusResponseBodyDataAgentStatus) GetCustomerStatus() *string {
	return s.CustomerStatus
}

func (s *ClinkAgentStatusResponseBodyDataAgentStatus) GetIbAnsweredRingingDuration() *int64 {
	return s.IbAnsweredRingingDuration
}

func (s *ClinkAgentStatusResponseBodyDataAgentStatus) GetIbClientAnsweredCount() *int64 {
	return s.IbClientAnsweredCount
}

func (s *ClinkAgentStatusResponseBodyDataAgentStatus) GetIbClientAnsweredRate() *float64 {
	return s.IbClientAnsweredRate
}

func (s *ClinkAgentStatusResponseBodyDataAgentStatus) GetIbClientTotalCount() *int64 {
	return s.IbClientTotalCount
}

func (s *ClinkAgentStatusResponseBodyDataAgentStatus) GetIbClientUnansweredCount() *int64 {
	return s.IbClientUnansweredCount
}

func (s *ClinkAgentStatusResponseBodyDataAgentStatus) GetIdleTimeOut() *int64 {
	return s.IdleTimeOut
}

func (s *ClinkAgentStatusResponseBodyDataAgentStatus) GetIncomingCallCount() *int64 {
	return s.IncomingCallCount
}

func (s *ClinkAgentStatusResponseBodyDataAgentStatus) GetLoginDuration() *int64 {
	return s.LoginDuration
}

func (s *ClinkAgentStatusResponseBodyDataAgentStatus) GetObAvgBridgeDuration() *int64 {
	return s.ObAvgBridgeDuration
}

func (s *ClinkAgentStatusResponseBodyDataAgentStatus) GetObBridgeDuration() *int64 {
	return s.ObBridgeDuration
}

func (s *ClinkAgentStatusResponseBodyDataAgentStatus) GetObBridgeRate() *float64 {
	return s.ObBridgeRate
}

func (s *ClinkAgentStatusResponseBodyDataAgentStatus) GetObCallCount() *int64 {
	return s.ObCallCount
}

func (s *ClinkAgentStatusResponseBodyDataAgentStatus) GetObClientUnbridgeCount() *int64 {
	return s.ObClientUnbridgeCount
}

func (s *ClinkAgentStatusResponseBodyDataAgentStatus) GetObCustomerBridgeCount() *int64 {
	return s.ObCustomerBridgeCount
}

func (s *ClinkAgentStatusResponseBodyDataAgentStatus) GetObCustomerUnbridgeCount() *int64 {
	return s.ObCustomerUnbridgeCount
}

func (s *ClinkAgentStatusResponseBodyDataAgentStatus) GetPauseDescription() *string {
	return s.PauseDescription
}

func (s *ClinkAgentStatusResponseBodyDataAgentStatus) GetPauseType() *int64 {
	return s.PauseType
}

func (s *ClinkAgentStatusResponseBodyDataAgentStatus) GetPredictToCall() *int64 {
	return s.PredictToCall
}

func (s *ClinkAgentStatusResponseBodyDataAgentStatus) GetQueueIncomingCallCount() *int64 {
	return s.QueueIncomingCallCount
}

func (s *ClinkAgentStatusResponseBodyDataAgentStatus) GetQueueInfo() *string {
	return s.QueueInfo
}

func (s *ClinkAgentStatusResponseBodyDataAgentStatus) GetQueues() *string {
	return s.Queues
}

func (s *ClinkAgentStatusResponseBodyDataAgentStatus) GetQuickUnlinkCount() *int64 {
	return s.QuickUnlinkCount
}

func (s *ClinkAgentStatusResponseBodyDataAgentStatus) GetStateDuration() *int64 {
	return s.StateDuration
}

func (s *ClinkAgentStatusResponseBodyDataAgentStatus) GetTagNames() []*string {
	return s.TagNames
}

func (s *ClinkAgentStatusResponseBodyDataAgentStatus) SetAgentStatus(v string) *ClinkAgentStatusResponseBodyDataAgentStatus {
	s.AgentStatus = &v
	return s
}

func (s *ClinkAgentStatusResponseBodyDataAgentStatus) SetAgentStatusDetail(v string) *ClinkAgentStatusResponseBodyDataAgentStatus {
	s.AgentStatusDetail = &v
	return s
}

func (s *ClinkAgentStatusResponseBodyDataAgentStatus) SetBindTel(v string) *ClinkAgentStatusResponseBodyDataAgentStatus {
	s.BindTel = &v
	return s
}

func (s *ClinkAgentStatusResponseBodyDataAgentStatus) SetBindType(v int64) *ClinkAgentStatusResponseBodyDataAgentStatus {
	s.BindType = &v
	return s
}

func (s *ClinkAgentStatusResponseBodyDataAgentStatus) SetBridgeCallCount(v int64) *ClinkAgentStatusResponseBodyDataAgentStatus {
	s.BridgeCallCount = &v
	return s
}

func (s *ClinkAgentStatusResponseBodyDataAgentStatus) SetClientName(v string) *ClinkAgentStatusResponseBodyDataAgentStatus {
	s.ClientName = &v
	return s
}

func (s *ClinkAgentStatusResponseBodyDataAgentStatus) SetCno(v string) *ClinkAgentStatusResponseBodyDataAgentStatus {
	s.Cno = &v
	return s
}

func (s *ClinkAgentStatusResponseBodyDataAgentStatus) SetCode(v string) *ClinkAgentStatusResponseBodyDataAgentStatus {
	s.Code = &v
	return s
}

func (s *ClinkAgentStatusResponseBodyDataAgentStatus) SetCustomerNumber(v string) *ClinkAgentStatusResponseBodyDataAgentStatus {
	s.CustomerNumber = &v
	return s
}

func (s *ClinkAgentStatusResponseBodyDataAgentStatus) SetCustomerNumberEncrypt(v string) *ClinkAgentStatusResponseBodyDataAgentStatus {
	s.CustomerNumberEncrypt = &v
	return s
}

func (s *ClinkAgentStatusResponseBodyDataAgentStatus) SetCustomerStatus(v string) *ClinkAgentStatusResponseBodyDataAgentStatus {
	s.CustomerStatus = &v
	return s
}

func (s *ClinkAgentStatusResponseBodyDataAgentStatus) SetIbAnsweredRingingDuration(v int64) *ClinkAgentStatusResponseBodyDataAgentStatus {
	s.IbAnsweredRingingDuration = &v
	return s
}

func (s *ClinkAgentStatusResponseBodyDataAgentStatus) SetIbClientAnsweredCount(v int64) *ClinkAgentStatusResponseBodyDataAgentStatus {
	s.IbClientAnsweredCount = &v
	return s
}

func (s *ClinkAgentStatusResponseBodyDataAgentStatus) SetIbClientAnsweredRate(v float64) *ClinkAgentStatusResponseBodyDataAgentStatus {
	s.IbClientAnsweredRate = &v
	return s
}

func (s *ClinkAgentStatusResponseBodyDataAgentStatus) SetIbClientTotalCount(v int64) *ClinkAgentStatusResponseBodyDataAgentStatus {
	s.IbClientTotalCount = &v
	return s
}

func (s *ClinkAgentStatusResponseBodyDataAgentStatus) SetIbClientUnansweredCount(v int64) *ClinkAgentStatusResponseBodyDataAgentStatus {
	s.IbClientUnansweredCount = &v
	return s
}

func (s *ClinkAgentStatusResponseBodyDataAgentStatus) SetIdleTimeOut(v int64) *ClinkAgentStatusResponseBodyDataAgentStatus {
	s.IdleTimeOut = &v
	return s
}

func (s *ClinkAgentStatusResponseBodyDataAgentStatus) SetIncomingCallCount(v int64) *ClinkAgentStatusResponseBodyDataAgentStatus {
	s.IncomingCallCount = &v
	return s
}

func (s *ClinkAgentStatusResponseBodyDataAgentStatus) SetLoginDuration(v int64) *ClinkAgentStatusResponseBodyDataAgentStatus {
	s.LoginDuration = &v
	return s
}

func (s *ClinkAgentStatusResponseBodyDataAgentStatus) SetObAvgBridgeDuration(v int64) *ClinkAgentStatusResponseBodyDataAgentStatus {
	s.ObAvgBridgeDuration = &v
	return s
}

func (s *ClinkAgentStatusResponseBodyDataAgentStatus) SetObBridgeDuration(v int64) *ClinkAgentStatusResponseBodyDataAgentStatus {
	s.ObBridgeDuration = &v
	return s
}

func (s *ClinkAgentStatusResponseBodyDataAgentStatus) SetObBridgeRate(v float64) *ClinkAgentStatusResponseBodyDataAgentStatus {
	s.ObBridgeRate = &v
	return s
}

func (s *ClinkAgentStatusResponseBodyDataAgentStatus) SetObCallCount(v int64) *ClinkAgentStatusResponseBodyDataAgentStatus {
	s.ObCallCount = &v
	return s
}

func (s *ClinkAgentStatusResponseBodyDataAgentStatus) SetObClientUnbridgeCount(v int64) *ClinkAgentStatusResponseBodyDataAgentStatus {
	s.ObClientUnbridgeCount = &v
	return s
}

func (s *ClinkAgentStatusResponseBodyDataAgentStatus) SetObCustomerBridgeCount(v int64) *ClinkAgentStatusResponseBodyDataAgentStatus {
	s.ObCustomerBridgeCount = &v
	return s
}

func (s *ClinkAgentStatusResponseBodyDataAgentStatus) SetObCustomerUnbridgeCount(v int64) *ClinkAgentStatusResponseBodyDataAgentStatus {
	s.ObCustomerUnbridgeCount = &v
	return s
}

func (s *ClinkAgentStatusResponseBodyDataAgentStatus) SetPauseDescription(v string) *ClinkAgentStatusResponseBodyDataAgentStatus {
	s.PauseDescription = &v
	return s
}

func (s *ClinkAgentStatusResponseBodyDataAgentStatus) SetPauseType(v int64) *ClinkAgentStatusResponseBodyDataAgentStatus {
	s.PauseType = &v
	return s
}

func (s *ClinkAgentStatusResponseBodyDataAgentStatus) SetPredictToCall(v int64) *ClinkAgentStatusResponseBodyDataAgentStatus {
	s.PredictToCall = &v
	return s
}

func (s *ClinkAgentStatusResponseBodyDataAgentStatus) SetQueueIncomingCallCount(v int64) *ClinkAgentStatusResponseBodyDataAgentStatus {
	s.QueueIncomingCallCount = &v
	return s
}

func (s *ClinkAgentStatusResponseBodyDataAgentStatus) SetQueueInfo(v string) *ClinkAgentStatusResponseBodyDataAgentStatus {
	s.QueueInfo = &v
	return s
}

func (s *ClinkAgentStatusResponseBodyDataAgentStatus) SetQueues(v string) *ClinkAgentStatusResponseBodyDataAgentStatus {
	s.Queues = &v
	return s
}

func (s *ClinkAgentStatusResponseBodyDataAgentStatus) SetQuickUnlinkCount(v int64) *ClinkAgentStatusResponseBodyDataAgentStatus {
	s.QuickUnlinkCount = &v
	return s
}

func (s *ClinkAgentStatusResponseBodyDataAgentStatus) SetStateDuration(v int64) *ClinkAgentStatusResponseBodyDataAgentStatus {
	s.StateDuration = &v
	return s
}

func (s *ClinkAgentStatusResponseBodyDataAgentStatus) SetTagNames(v []*string) *ClinkAgentStatusResponseBodyDataAgentStatus {
	s.TagNames = v
	return s
}

func (s *ClinkAgentStatusResponseBodyDataAgentStatus) Validate() error {
	return dara.Validate(s)
}
