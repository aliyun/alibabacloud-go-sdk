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

func (s *ListTongyiConversationLogsResponseBody) Validate() error {
	return dara.Validate(s)
}
