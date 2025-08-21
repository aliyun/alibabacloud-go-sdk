// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTongyiChatHistorysResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCostTime(v string) *ListTongyiChatHistorysResponseBody
	GetCostTime() *string
	SetDatas(v []map[string]interface{}) *ListTongyiChatHistorysResponseBody
	GetDatas() []map[string]interface{}
	SetRequestId(v string) *ListTongyiChatHistorysResponseBody
	GetRequestId() *string
}

type ListTongyiChatHistorysResponseBody struct {
	// example:
	//
	// 116
	CostTime *string                  `json:"CostTime,omitempty" xml:"CostTime,omitempty"`
	Datas    []map[string]interface{} `json:"Datas,omitempty" xml:"Datas,omitempty" type:"Repeated"`
	// example:
	//
	// D0DDFC4C-D66D-4787-9AE4-4D757481EDEE
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListTongyiChatHistorysResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListTongyiChatHistorysResponseBody) GoString() string {
	return s.String()
}

func (s *ListTongyiChatHistorysResponseBody) GetCostTime() *string {
	return s.CostTime
}

func (s *ListTongyiChatHistorysResponseBody) GetDatas() []map[string]interface{} {
	return s.Datas
}

func (s *ListTongyiChatHistorysResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListTongyiChatHistorysResponseBody) SetCostTime(v string) *ListTongyiChatHistorysResponseBody {
	s.CostTime = &v
	return s
}

func (s *ListTongyiChatHistorysResponseBody) SetDatas(v []map[string]interface{}) *ListTongyiChatHistorysResponseBody {
	s.Datas = v
	return s
}

func (s *ListTongyiChatHistorysResponseBody) SetRequestId(v string) *ListTongyiChatHistorysResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListTongyiChatHistorysResponseBody) Validate() error {
	return dara.Validate(s)
}
