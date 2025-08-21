// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetBotSessionDataResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCostTime(v string) *GetBotSessionDataResponseBody
	GetCostTime() *string
	SetDatas(v []map[string]interface{}) *GetBotSessionDataResponseBody
	GetDatas() []map[string]interface{}
	SetRequestId(v string) *GetBotSessionDataResponseBody
	GetRequestId() *string
}

type GetBotSessionDataResponseBody struct {
	// example:
	//
	// 116
	CostTime *string                  `json:"CostTime,omitempty" xml:"CostTime,omitempty"`
	Datas    []map[string]interface{} `json:"Datas,omitempty" xml:"Datas,omitempty" type:"Repeated"`
	// example:
	//
	// 15CD94CC-CBEB-4189-806C-A132D1F45D51
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetBotSessionDataResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetBotSessionDataResponseBody) GoString() string {
	return s.String()
}

func (s *GetBotSessionDataResponseBody) GetCostTime() *string {
	return s.CostTime
}

func (s *GetBotSessionDataResponseBody) GetDatas() []map[string]interface{} {
	return s.Datas
}

func (s *GetBotSessionDataResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetBotSessionDataResponseBody) SetCostTime(v string) *GetBotSessionDataResponseBody {
	s.CostTime = &v
	return s
}

func (s *GetBotSessionDataResponseBody) SetDatas(v []map[string]interface{}) *GetBotSessionDataResponseBody {
	s.Datas = v
	return s
}

func (s *GetBotSessionDataResponseBody) SetRequestId(v string) *GetBotSessionDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetBotSessionDataResponseBody) Validate() error {
	return dara.Validate(s)
}
