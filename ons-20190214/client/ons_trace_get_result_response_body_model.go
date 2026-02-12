// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOnsTraceGetResultResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *OnsTraceGetResultResponseBody
	GetRequestId() *string
	SetTraceData(v *OnsTraceGetResultResponseBodyTraceData) *OnsTraceGetResultResponseBody
	GetTraceData() *OnsTraceGetResultResponseBodyTraceData
}

type OnsTraceGetResultResponseBody struct {
	// The ID of the request. This parameter is a common parameter. Each request has a unique ID. You can use this ID to troubleshoot issues.
	//
	// example:
	//
	// 84EE24D2-851F-40D6-B99E-4D6AB909****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The details of the message trace.
	TraceData *OnsTraceGetResultResponseBodyTraceData `json:"TraceData,omitempty" xml:"TraceData,omitempty" type:"Struct"`
}

func (s OnsTraceGetResultResponseBody) String() string {
	return dara.Prettify(s)
}

func (s OnsTraceGetResultResponseBody) GoString() string {
	return s.String()
}

func (s *OnsTraceGetResultResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *OnsTraceGetResultResponseBody) GetTraceData() *OnsTraceGetResultResponseBodyTraceData {
	return s.TraceData
}

func (s *OnsTraceGetResultResponseBody) SetRequestId(v string) *OnsTraceGetResultResponseBody {
	s.RequestId = &v
	return s
}

func (s *OnsTraceGetResultResponseBody) SetTraceData(v *OnsTraceGetResultResponseBodyTraceData) *OnsTraceGetResultResponseBody {
	s.TraceData = v
	return s
}

func (s *OnsTraceGetResultResponseBody) Validate() error {
	if s.TraceData != nil {
		if err := s.TraceData.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type OnsTraceGetResultResponseBodyTraceData struct {
	// The point in time when the task was created.
	//
	// example:
	//
	// 1570966857000
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The ID of the instance
	//
	// example:
	//
	// MQ_INST_111111111111_DOxxxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The ID of the message that is queried.
	//
	// example:
	//
	// 1E05791C117818B4AAC23B1BB0CE****
	MsgId *string `json:"MsgId,omitempty" xml:"MsgId,omitempty"`
	// The key of the message that is queried.
	//
	// example:
	//
	// ORDERID_100
	MsgKey *string `json:"MsgKey,omitempty" xml:"MsgKey,omitempty"`
	// The ID of the task.
	//
	// example:
	//
	// 272967562652883649157096685****
	QueryId *string `json:"QueryId,omitempty" xml:"QueryId,omitempty"`
	// The status of the task. Valid values:
	//
	// 	- **finish**: The task is complete.
	//
	// 	- **working**: The task is in progress.
	//
	// 	- **removed**: The task is deleted.
	//
	// example:
	//
	// finish
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The topic in which the message is stored.
	//
	// example:
	//
	// test
	Topic     *string                                          `json:"Topic,omitempty" xml:"Topic,omitempty"`
	TraceList *OnsTraceGetResultResponseBodyTraceDataTraceList `json:"TraceList,omitempty" xml:"TraceList,omitempty" type:"Struct"`
	// The most recent point in time when the task was updated.
	//
	// example:
	//
	// 1570966877000
	UpdateTime *int64 `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	// The ID of the user who created the task.
	//
	// example:
	//
	// 27296756265288****
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s OnsTraceGetResultResponseBodyTraceData) String() string {
	return dara.Prettify(s)
}

func (s OnsTraceGetResultResponseBodyTraceData) GoString() string {
	return s.String()
}

func (s *OnsTraceGetResultResponseBodyTraceData) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *OnsTraceGetResultResponseBodyTraceData) GetInstanceId() *string {
	return s.InstanceId
}

func (s *OnsTraceGetResultResponseBodyTraceData) GetMsgId() *string {
	return s.MsgId
}

func (s *OnsTraceGetResultResponseBodyTraceData) GetMsgKey() *string {
	return s.MsgKey
}

func (s *OnsTraceGetResultResponseBodyTraceData) GetQueryId() *string {
	return s.QueryId
}

func (s *OnsTraceGetResultResponseBodyTraceData) GetStatus() *string {
	return s.Status
}

func (s *OnsTraceGetResultResponseBodyTraceData) GetTopic() *string {
	return s.Topic
}

func (s *OnsTraceGetResultResponseBodyTraceData) GetTraceList() *OnsTraceGetResultResponseBodyTraceDataTraceList {
	return s.TraceList
}

func (s *OnsTraceGetResultResponseBodyTraceData) GetUpdateTime() *int64 {
	return s.UpdateTime
}

func (s *OnsTraceGetResultResponseBodyTraceData) GetUserId() *string {
	return s.UserId
}

func (s *OnsTraceGetResultResponseBodyTraceData) SetCreateTime(v int64) *OnsTraceGetResultResponseBodyTraceData {
	s.CreateTime = &v
	return s
}

func (s *OnsTraceGetResultResponseBodyTraceData) SetInstanceId(v string) *OnsTraceGetResultResponseBodyTraceData {
	s.InstanceId = &v
	return s
}

func (s *OnsTraceGetResultResponseBodyTraceData) SetMsgId(v string) *OnsTraceGetResultResponseBodyTraceData {
	s.MsgId = &v
	return s
}

func (s *OnsTraceGetResultResponseBodyTraceData) SetMsgKey(v string) *OnsTraceGetResultResponseBodyTraceData {
	s.MsgKey = &v
	return s
}

func (s *OnsTraceGetResultResponseBodyTraceData) SetQueryId(v string) *OnsTraceGetResultResponseBodyTraceData {
	s.QueryId = &v
	return s
}

func (s *OnsTraceGetResultResponseBodyTraceData) SetStatus(v string) *OnsTraceGetResultResponseBodyTraceData {
	s.Status = &v
	return s
}

func (s *OnsTraceGetResultResponseBodyTraceData) SetTopic(v string) *OnsTraceGetResultResponseBodyTraceData {
	s.Topic = &v
	return s
}

func (s *OnsTraceGetResultResponseBodyTraceData) SetTraceList(v *OnsTraceGetResultResponseBodyTraceDataTraceList) *OnsTraceGetResultResponseBodyTraceData {
	s.TraceList = v
	return s
}

func (s *OnsTraceGetResultResponseBodyTraceData) SetUpdateTime(v int64) *OnsTraceGetResultResponseBodyTraceData {
	s.UpdateTime = &v
	return s
}

func (s *OnsTraceGetResultResponseBodyTraceData) SetUserId(v string) *OnsTraceGetResultResponseBodyTraceData {
	s.UserId = &v
	return s
}

func (s *OnsTraceGetResultResponseBodyTraceData) Validate() error {
	if s.TraceList != nil {
		if err := s.TraceList.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type OnsTraceGetResultResponseBodyTraceDataTraceList struct {
	TraceMapDo []*OnsTraceGetResultResponseBodyTraceDataTraceListTraceMapDo `json:"TraceMapDo,omitempty" xml:"TraceMapDo,omitempty" type:"Repeated"`
}

func (s OnsTraceGetResultResponseBodyTraceDataTraceList) String() string {
	return dara.Prettify(s)
}

func (s OnsTraceGetResultResponseBodyTraceDataTraceList) GoString() string {
	return s.String()
}

func (s *OnsTraceGetResultResponseBodyTraceDataTraceList) GetTraceMapDo() []*OnsTraceGetResultResponseBodyTraceDataTraceListTraceMapDo {
	return s.TraceMapDo
}

func (s *OnsTraceGetResultResponseBodyTraceDataTraceList) SetTraceMapDo(v []*OnsTraceGetResultResponseBodyTraceDataTraceListTraceMapDo) *OnsTraceGetResultResponseBodyTraceDataTraceList {
	s.TraceMapDo = v
	return s
}

func (s *OnsTraceGetResultResponseBodyTraceDataTraceList) Validate() error {
	if s.TraceMapDo != nil {
		for _, item := range s.TraceMapDo {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type OnsTraceGetResultResponseBodyTraceDataTraceListTraceMapDo struct {
	BornHost     *string                                                           `json:"BornHost,omitempty" xml:"BornHost,omitempty"`
	CostTime     *int32                                                            `json:"CostTime,omitempty" xml:"CostTime,omitempty"`
	MsgId        *string                                                           `json:"MsgId,omitempty" xml:"MsgId,omitempty"`
	MsgKey       *string                                                           `json:"MsgKey,omitempty" xml:"MsgKey,omitempty"`
	PubGroupName *string                                                           `json:"PubGroupName,omitempty" xml:"PubGroupName,omitempty"`
	PubTime      *int64                                                            `json:"PubTime,omitempty" xml:"PubTime,omitempty"`
	Status       *string                                                           `json:"Status,omitempty" xml:"Status,omitempty"`
	SubList      *OnsTraceGetResultResponseBodyTraceDataTraceListTraceMapDoSubList `json:"SubList,omitempty" xml:"SubList,omitempty" type:"Struct"`
	Tag          *string                                                           `json:"Tag,omitempty" xml:"Tag,omitempty"`
	Topic        *string                                                           `json:"Topic,omitempty" xml:"Topic,omitempty"`
}

func (s OnsTraceGetResultResponseBodyTraceDataTraceListTraceMapDo) String() string {
	return dara.Prettify(s)
}

func (s OnsTraceGetResultResponseBodyTraceDataTraceListTraceMapDo) GoString() string {
	return s.String()
}

func (s *OnsTraceGetResultResponseBodyTraceDataTraceListTraceMapDo) GetBornHost() *string {
	return s.BornHost
}

func (s *OnsTraceGetResultResponseBodyTraceDataTraceListTraceMapDo) GetCostTime() *int32 {
	return s.CostTime
}

func (s *OnsTraceGetResultResponseBodyTraceDataTraceListTraceMapDo) GetMsgId() *string {
	return s.MsgId
}

func (s *OnsTraceGetResultResponseBodyTraceDataTraceListTraceMapDo) GetMsgKey() *string {
	return s.MsgKey
}

func (s *OnsTraceGetResultResponseBodyTraceDataTraceListTraceMapDo) GetPubGroupName() *string {
	return s.PubGroupName
}

func (s *OnsTraceGetResultResponseBodyTraceDataTraceListTraceMapDo) GetPubTime() *int64 {
	return s.PubTime
}

func (s *OnsTraceGetResultResponseBodyTraceDataTraceListTraceMapDo) GetStatus() *string {
	return s.Status
}

func (s *OnsTraceGetResultResponseBodyTraceDataTraceListTraceMapDo) GetSubList() *OnsTraceGetResultResponseBodyTraceDataTraceListTraceMapDoSubList {
	return s.SubList
}

func (s *OnsTraceGetResultResponseBodyTraceDataTraceListTraceMapDo) GetTag() *string {
	return s.Tag
}

func (s *OnsTraceGetResultResponseBodyTraceDataTraceListTraceMapDo) GetTopic() *string {
	return s.Topic
}

func (s *OnsTraceGetResultResponseBodyTraceDataTraceListTraceMapDo) SetBornHost(v string) *OnsTraceGetResultResponseBodyTraceDataTraceListTraceMapDo {
	s.BornHost = &v
	return s
}

func (s *OnsTraceGetResultResponseBodyTraceDataTraceListTraceMapDo) SetCostTime(v int32) *OnsTraceGetResultResponseBodyTraceDataTraceListTraceMapDo {
	s.CostTime = &v
	return s
}

func (s *OnsTraceGetResultResponseBodyTraceDataTraceListTraceMapDo) SetMsgId(v string) *OnsTraceGetResultResponseBodyTraceDataTraceListTraceMapDo {
	s.MsgId = &v
	return s
}

func (s *OnsTraceGetResultResponseBodyTraceDataTraceListTraceMapDo) SetMsgKey(v string) *OnsTraceGetResultResponseBodyTraceDataTraceListTraceMapDo {
	s.MsgKey = &v
	return s
}

func (s *OnsTraceGetResultResponseBodyTraceDataTraceListTraceMapDo) SetPubGroupName(v string) *OnsTraceGetResultResponseBodyTraceDataTraceListTraceMapDo {
	s.PubGroupName = &v
	return s
}

func (s *OnsTraceGetResultResponseBodyTraceDataTraceListTraceMapDo) SetPubTime(v int64) *OnsTraceGetResultResponseBodyTraceDataTraceListTraceMapDo {
	s.PubTime = &v
	return s
}

func (s *OnsTraceGetResultResponseBodyTraceDataTraceListTraceMapDo) SetStatus(v string) *OnsTraceGetResultResponseBodyTraceDataTraceListTraceMapDo {
	s.Status = &v
	return s
}

func (s *OnsTraceGetResultResponseBodyTraceDataTraceListTraceMapDo) SetSubList(v *OnsTraceGetResultResponseBodyTraceDataTraceListTraceMapDoSubList) *OnsTraceGetResultResponseBodyTraceDataTraceListTraceMapDo {
	s.SubList = v
	return s
}

func (s *OnsTraceGetResultResponseBodyTraceDataTraceListTraceMapDo) SetTag(v string) *OnsTraceGetResultResponseBodyTraceDataTraceListTraceMapDo {
	s.Tag = &v
	return s
}

func (s *OnsTraceGetResultResponseBodyTraceDataTraceListTraceMapDo) SetTopic(v string) *OnsTraceGetResultResponseBodyTraceDataTraceListTraceMapDo {
	s.Topic = &v
	return s
}

func (s *OnsTraceGetResultResponseBodyTraceDataTraceListTraceMapDo) Validate() error {
	if s.SubList != nil {
		if err := s.SubList.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type OnsTraceGetResultResponseBodyTraceDataTraceListTraceMapDoSubList struct {
	SubMapDo []*OnsTraceGetResultResponseBodyTraceDataTraceListTraceMapDoSubListSubMapDo `json:"SubMapDo,omitempty" xml:"SubMapDo,omitempty" type:"Repeated"`
}

func (s OnsTraceGetResultResponseBodyTraceDataTraceListTraceMapDoSubList) String() string {
	return dara.Prettify(s)
}

func (s OnsTraceGetResultResponseBodyTraceDataTraceListTraceMapDoSubList) GoString() string {
	return s.String()
}

func (s *OnsTraceGetResultResponseBodyTraceDataTraceListTraceMapDoSubList) GetSubMapDo() []*OnsTraceGetResultResponseBodyTraceDataTraceListTraceMapDoSubListSubMapDo {
	return s.SubMapDo
}

func (s *OnsTraceGetResultResponseBodyTraceDataTraceListTraceMapDoSubList) SetSubMapDo(v []*OnsTraceGetResultResponseBodyTraceDataTraceListTraceMapDoSubListSubMapDo) *OnsTraceGetResultResponseBodyTraceDataTraceListTraceMapDoSubList {
	s.SubMapDo = v
	return s
}

func (s *OnsTraceGetResultResponseBodyTraceDataTraceListTraceMapDoSubList) Validate() error {
	if s.SubMapDo != nil {
		for _, item := range s.SubMapDo {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type OnsTraceGetResultResponseBodyTraceDataTraceListTraceMapDoSubListSubMapDo struct {
	ClientList   *OnsTraceGetResultResponseBodyTraceDataTraceListTraceMapDoSubListSubMapDoClientList `json:"ClientList,omitempty" xml:"ClientList,omitempty" type:"Struct"`
	FailCount    *int32                                                                              `json:"FailCount,omitempty" xml:"FailCount,omitempty"`
	SubGroupName *string                                                                             `json:"SubGroupName,omitempty" xml:"SubGroupName,omitempty"`
	SuccessCount *int32                                                                              `json:"SuccessCount,omitempty" xml:"SuccessCount,omitempty"`
}

func (s OnsTraceGetResultResponseBodyTraceDataTraceListTraceMapDoSubListSubMapDo) String() string {
	return dara.Prettify(s)
}

func (s OnsTraceGetResultResponseBodyTraceDataTraceListTraceMapDoSubListSubMapDo) GoString() string {
	return s.String()
}

func (s *OnsTraceGetResultResponseBodyTraceDataTraceListTraceMapDoSubListSubMapDo) GetClientList() *OnsTraceGetResultResponseBodyTraceDataTraceListTraceMapDoSubListSubMapDoClientList {
	return s.ClientList
}

func (s *OnsTraceGetResultResponseBodyTraceDataTraceListTraceMapDoSubListSubMapDo) GetFailCount() *int32 {
	return s.FailCount
}

func (s *OnsTraceGetResultResponseBodyTraceDataTraceListTraceMapDoSubListSubMapDo) GetSubGroupName() *string {
	return s.SubGroupName
}

func (s *OnsTraceGetResultResponseBodyTraceDataTraceListTraceMapDoSubListSubMapDo) GetSuccessCount() *int32 {
	return s.SuccessCount
}

func (s *OnsTraceGetResultResponseBodyTraceDataTraceListTraceMapDoSubListSubMapDo) SetClientList(v *OnsTraceGetResultResponseBodyTraceDataTraceListTraceMapDoSubListSubMapDoClientList) *OnsTraceGetResultResponseBodyTraceDataTraceListTraceMapDoSubListSubMapDo {
	s.ClientList = v
	return s
}

func (s *OnsTraceGetResultResponseBodyTraceDataTraceListTraceMapDoSubListSubMapDo) SetFailCount(v int32) *OnsTraceGetResultResponseBodyTraceDataTraceListTraceMapDoSubListSubMapDo {
	s.FailCount = &v
	return s
}

func (s *OnsTraceGetResultResponseBodyTraceDataTraceListTraceMapDoSubListSubMapDo) SetSubGroupName(v string) *OnsTraceGetResultResponseBodyTraceDataTraceListTraceMapDoSubListSubMapDo {
	s.SubGroupName = &v
	return s
}

func (s *OnsTraceGetResultResponseBodyTraceDataTraceListTraceMapDoSubListSubMapDo) SetSuccessCount(v int32) *OnsTraceGetResultResponseBodyTraceDataTraceListTraceMapDoSubListSubMapDo {
	s.SuccessCount = &v
	return s
}

func (s *OnsTraceGetResultResponseBodyTraceDataTraceListTraceMapDoSubListSubMapDo) Validate() error {
	if s.ClientList != nil {
		if err := s.ClientList.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type OnsTraceGetResultResponseBodyTraceDataTraceListTraceMapDoSubListSubMapDoClientList struct {
	SubClientInfoDo []*OnsTraceGetResultResponseBodyTraceDataTraceListTraceMapDoSubListSubMapDoClientListSubClientInfoDo `json:"SubClientInfoDo,omitempty" xml:"SubClientInfoDo,omitempty" type:"Repeated"`
}

func (s OnsTraceGetResultResponseBodyTraceDataTraceListTraceMapDoSubListSubMapDoClientList) String() string {
	return dara.Prettify(s)
}

func (s OnsTraceGetResultResponseBodyTraceDataTraceListTraceMapDoSubListSubMapDoClientList) GoString() string {
	return s.String()
}

func (s *OnsTraceGetResultResponseBodyTraceDataTraceListTraceMapDoSubListSubMapDoClientList) GetSubClientInfoDo() []*OnsTraceGetResultResponseBodyTraceDataTraceListTraceMapDoSubListSubMapDoClientListSubClientInfoDo {
	return s.SubClientInfoDo
}

func (s *OnsTraceGetResultResponseBodyTraceDataTraceListTraceMapDoSubListSubMapDoClientList) SetSubClientInfoDo(v []*OnsTraceGetResultResponseBodyTraceDataTraceListTraceMapDoSubListSubMapDoClientListSubClientInfoDo) *OnsTraceGetResultResponseBodyTraceDataTraceListTraceMapDoSubListSubMapDoClientList {
	s.SubClientInfoDo = v
	return s
}

func (s *OnsTraceGetResultResponseBodyTraceDataTraceListTraceMapDoSubListSubMapDoClientList) Validate() error {
	if s.SubClientInfoDo != nil {
		for _, item := range s.SubClientInfoDo {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type OnsTraceGetResultResponseBodyTraceDataTraceListTraceMapDoSubListSubMapDoClientListSubClientInfoDo struct {
	ClientHost     *string `json:"ClientHost,omitempty" xml:"ClientHost,omitempty"`
	CostTime       *int32  `json:"CostTime,omitempty" xml:"CostTime,omitempty"`
	ReconsumeTimes *int32  `json:"ReconsumeTimes,omitempty" xml:"ReconsumeTimes,omitempty"`
	Status         *string `json:"Status,omitempty" xml:"Status,omitempty"`
	SubGroupName   *string `json:"SubGroupName,omitempty" xml:"SubGroupName,omitempty"`
	SubTime        *int64  `json:"SubTime,omitempty" xml:"SubTime,omitempty"`
}

func (s OnsTraceGetResultResponseBodyTraceDataTraceListTraceMapDoSubListSubMapDoClientListSubClientInfoDo) String() string {
	return dara.Prettify(s)
}

func (s OnsTraceGetResultResponseBodyTraceDataTraceListTraceMapDoSubListSubMapDoClientListSubClientInfoDo) GoString() string {
	return s.String()
}

func (s *OnsTraceGetResultResponseBodyTraceDataTraceListTraceMapDoSubListSubMapDoClientListSubClientInfoDo) GetClientHost() *string {
	return s.ClientHost
}

func (s *OnsTraceGetResultResponseBodyTraceDataTraceListTraceMapDoSubListSubMapDoClientListSubClientInfoDo) GetCostTime() *int32 {
	return s.CostTime
}

func (s *OnsTraceGetResultResponseBodyTraceDataTraceListTraceMapDoSubListSubMapDoClientListSubClientInfoDo) GetReconsumeTimes() *int32 {
	return s.ReconsumeTimes
}

func (s *OnsTraceGetResultResponseBodyTraceDataTraceListTraceMapDoSubListSubMapDoClientListSubClientInfoDo) GetStatus() *string {
	return s.Status
}

func (s *OnsTraceGetResultResponseBodyTraceDataTraceListTraceMapDoSubListSubMapDoClientListSubClientInfoDo) GetSubGroupName() *string {
	return s.SubGroupName
}

func (s *OnsTraceGetResultResponseBodyTraceDataTraceListTraceMapDoSubListSubMapDoClientListSubClientInfoDo) GetSubTime() *int64 {
	return s.SubTime
}

func (s *OnsTraceGetResultResponseBodyTraceDataTraceListTraceMapDoSubListSubMapDoClientListSubClientInfoDo) SetClientHost(v string) *OnsTraceGetResultResponseBodyTraceDataTraceListTraceMapDoSubListSubMapDoClientListSubClientInfoDo {
	s.ClientHost = &v
	return s
}

func (s *OnsTraceGetResultResponseBodyTraceDataTraceListTraceMapDoSubListSubMapDoClientListSubClientInfoDo) SetCostTime(v int32) *OnsTraceGetResultResponseBodyTraceDataTraceListTraceMapDoSubListSubMapDoClientListSubClientInfoDo {
	s.CostTime = &v
	return s
}

func (s *OnsTraceGetResultResponseBodyTraceDataTraceListTraceMapDoSubListSubMapDoClientListSubClientInfoDo) SetReconsumeTimes(v int32) *OnsTraceGetResultResponseBodyTraceDataTraceListTraceMapDoSubListSubMapDoClientListSubClientInfoDo {
	s.ReconsumeTimes = &v
	return s
}

func (s *OnsTraceGetResultResponseBodyTraceDataTraceListTraceMapDoSubListSubMapDoClientListSubClientInfoDo) SetStatus(v string) *OnsTraceGetResultResponseBodyTraceDataTraceListTraceMapDoSubListSubMapDoClientListSubClientInfoDo {
	s.Status = &v
	return s
}

func (s *OnsTraceGetResultResponseBodyTraceDataTraceListTraceMapDoSubListSubMapDoClientListSubClientInfoDo) SetSubGroupName(v string) *OnsTraceGetResultResponseBodyTraceDataTraceListTraceMapDoSubListSubMapDoClientListSubClientInfoDo {
	s.SubGroupName = &v
	return s
}

func (s *OnsTraceGetResultResponseBodyTraceDataTraceListTraceMapDoSubListSubMapDoClientListSubClientInfoDo) SetSubTime(v int64) *OnsTraceGetResultResponseBodyTraceDataTraceListTraceMapDoSubListSubMapDoClientListSubClientInfoDo {
	s.SubTime = &v
	return s
}

func (s *OnsTraceGetResultResponseBodyTraceDataTraceListTraceMapDoSubListSubMapDoClientListSubClientInfoDo) Validate() error {
	return dara.Validate(s)
}
