// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTongyiConversationLogsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCostTime(v string) *ListTongyiConversationLogsResponseBody
	GetCostTime() *string
	SetDatas(v []map[string]interface{}) *ListTongyiConversationLogsResponseBody
	GetDatas() []map[string]interface{}
	SetRequestId(v string) *ListTongyiConversationLogsResponseBody
	GetRequestId() *string
	SetSessionFlowDebugInfo(v *ListTongyiConversationLogsResponseBodySessionFlowDebugInfo) *ListTongyiConversationLogsResponseBody
	GetSessionFlowDebugInfo() *ListTongyiConversationLogsResponseBodySessionFlowDebugInfo
}

type ListTongyiConversationLogsResponseBody struct {
	// example:
	//
	// 66
	CostTime *string                  `json:"CostTime,omitempty" xml:"CostTime,omitempty"`
	Datas    []map[string]interface{} `json:"Datas,omitempty" xml:"Datas,omitempty" type:"Repeated"`
	// example:
	//
	// 28805A7C-D695-548C-A31B-67E52C2C274F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// if can be null:
	// true
	SessionFlowDebugInfo *ListTongyiConversationLogsResponseBodySessionFlowDebugInfo `json:"SessionFlowDebugInfo,omitempty" xml:"SessionFlowDebugInfo,omitempty" type:"Struct"`
}

func (s ListTongyiConversationLogsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListTongyiConversationLogsResponseBody) GoString() string {
	return s.String()
}

func (s *ListTongyiConversationLogsResponseBody) GetCostTime() *string {
	return s.CostTime
}

func (s *ListTongyiConversationLogsResponseBody) GetDatas() []map[string]interface{} {
	return s.Datas
}

func (s *ListTongyiConversationLogsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListTongyiConversationLogsResponseBody) GetSessionFlowDebugInfo() *ListTongyiConversationLogsResponseBodySessionFlowDebugInfo {
	return s.SessionFlowDebugInfo
}

func (s *ListTongyiConversationLogsResponseBody) SetCostTime(v string) *ListTongyiConversationLogsResponseBody {
	s.CostTime = &v
	return s
}

func (s *ListTongyiConversationLogsResponseBody) SetDatas(v []map[string]interface{}) *ListTongyiConversationLogsResponseBody {
	s.Datas = v
	return s
}

func (s *ListTongyiConversationLogsResponseBody) SetRequestId(v string) *ListTongyiConversationLogsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListTongyiConversationLogsResponseBody) SetSessionFlowDebugInfo(v *ListTongyiConversationLogsResponseBodySessionFlowDebugInfo) *ListTongyiConversationLogsResponseBody {
	s.SessionFlowDebugInfo = v
	return s
}

func (s *ListTongyiConversationLogsResponseBody) Validate() error {
	if s.SessionFlowDebugInfo != nil {
		if err := s.SessionFlowDebugInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListTongyiConversationLogsResponseBodySessionFlowDebugInfo struct {
	ApiParams  map[string]interface{} `json:"ApiParams,omitempty" xml:"ApiParams,omitempty"`
	SlotParams map[string]interface{} `json:"SlotParams,omitempty" xml:"SlotParams,omitempty"`
}

func (s ListTongyiConversationLogsResponseBodySessionFlowDebugInfo) String() string {
	return dara.Prettify(s)
}

func (s ListTongyiConversationLogsResponseBodySessionFlowDebugInfo) GoString() string {
	return s.String()
}

func (s *ListTongyiConversationLogsResponseBodySessionFlowDebugInfo) GetApiParams() map[string]interface{} {
	return s.ApiParams
}

func (s *ListTongyiConversationLogsResponseBodySessionFlowDebugInfo) GetSlotParams() map[string]interface{} {
	return s.SlotParams
}

func (s *ListTongyiConversationLogsResponseBodySessionFlowDebugInfo) SetApiParams(v map[string]interface{}) *ListTongyiConversationLogsResponseBodySessionFlowDebugInfo {
	s.ApiParams = v
	return s
}

func (s *ListTongyiConversationLogsResponseBodySessionFlowDebugInfo) SetSlotParams(v map[string]interface{}) *ListTongyiConversationLogsResponseBodySessionFlowDebugInfo {
	s.SlotParams = v
	return s
}

func (s *ListTongyiConversationLogsResponseBodySessionFlowDebugInfo) Validate() error {
	return dara.Validate(s)
}
