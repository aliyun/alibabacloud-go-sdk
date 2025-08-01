// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetLocalityRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *GetLocalityRuleResponseBody
	GetCode() *int32
	SetData(v string) *GetLocalityRuleResponseBody
	GetData() *string
	SetHttpStatusCode(v int32) *GetLocalityRuleResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetLocalityRuleResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetLocalityRuleResponseBody
	GetRequestId() *string
	SetSuccess(v string) *GetLocalityRuleResponseBody
	GetSuccess() *string
}

type GetLocalityRuleResponseBody struct {
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// {
	//
	//   "pageNumber":0,
	//
	//   "namespaceId":"cn-hangzhou",
	//
	//   "enable":false,
	//
	//   "appId":"hkhon1po62@3aa3582********",
	//
	//   "pageSize":0,
	//
	//   "region":"cn-hangzhou",
	//
	//   "routeRules":[
	//
	//     {
	//
	//       "threshold":0.2,
	//
	//       "tags":[]
	//
	//     }
	//
	//   ]
	//
	// }
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// D3971C60-3F07-58B0-8EA0-A194********
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetLocalityRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetLocalityRuleResponseBody) GoString() string {
	return s.String()
}

func (s *GetLocalityRuleResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *GetLocalityRuleResponseBody) GetData() *string {
	return s.Data
}

func (s *GetLocalityRuleResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetLocalityRuleResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetLocalityRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetLocalityRuleResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *GetLocalityRuleResponseBody) SetCode(v int32) *GetLocalityRuleResponseBody {
	s.Code = &v
	return s
}

func (s *GetLocalityRuleResponseBody) SetData(v string) *GetLocalityRuleResponseBody {
	s.Data = &v
	return s
}

func (s *GetLocalityRuleResponseBody) SetHttpStatusCode(v int32) *GetLocalityRuleResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetLocalityRuleResponseBody) SetMessage(v string) *GetLocalityRuleResponseBody {
	s.Message = &v
	return s
}

func (s *GetLocalityRuleResponseBody) SetRequestId(v string) *GetLocalityRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetLocalityRuleResponseBody) SetSuccess(v string) *GetLocalityRuleResponseBody {
	s.Success = &v
	return s
}

func (s *GetLocalityRuleResponseBody) Validate() error {
	return dara.Validate(s)
}
