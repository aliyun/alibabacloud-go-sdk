// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInvokeEsRequestResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *InvokeEsRequestResponseBody
	GetRequestId() *string
	SetResult(v interface{}) *InvokeEsRequestResponseBody
	GetResult() interface{}
}

type InvokeEsRequestResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// 8E5A2C41-D96B-4308-AF72-5C0B14E6****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The response content returned as-is from ES. The structure is determined by the ES API being called.
	//
	// example:
	//
	// {"esResult":{"took":5,"timed_out":false,"hits":{"total":{"value":1,"relation":"eq"},"max_score":1.0,"hits":[{"_index":"my-index","_id":"1","_score":1.0,"_source":{"title":"Wireless Bluetooth Headphones"}}]}}}
	Result interface{} `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s InvokeEsRequestResponseBody) String() string {
	return dara.Prettify(s)
}

func (s InvokeEsRequestResponseBody) GoString() string {
	return s.String()
}

func (s *InvokeEsRequestResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *InvokeEsRequestResponseBody) GetResult() interface{} {
	return s.Result
}

func (s *InvokeEsRequestResponseBody) SetRequestId(v string) *InvokeEsRequestResponseBody {
	s.RequestId = &v
	return s
}

func (s *InvokeEsRequestResponseBody) SetResult(v interface{}) *InvokeEsRequestResponseBody {
	s.Result = v
	return s
}

func (s *InvokeEsRequestResponseBody) Validate() error {
	return dara.Validate(s)
}
