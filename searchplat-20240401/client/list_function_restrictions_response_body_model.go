// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListFunctionRestrictionsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListFunctionRestrictionsResponseBody
	GetCode() *string
	SetHttpCode(v int64) *ListFunctionRestrictionsResponseBody
	GetHttpCode() *int64
	SetLatency(v float64) *ListFunctionRestrictionsResponseBody
	GetLatency() *float64
	SetMessage(v string) *ListFunctionRestrictionsResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListFunctionRestrictionsResponseBody
	GetRequestId() *string
	SetResult(v []*ListFunctionRestrictionsResponseBodyResult) *ListFunctionRestrictionsResponseBody
	GetResult() []*ListFunctionRestrictionsResponseBodyResult
	SetStatus(v string) *ListFunctionRestrictionsResponseBody
	GetStatus() *string
	SetTotalCount(v int64) *ListFunctionRestrictionsResponseBody
	GetTotalCount() *int64
}

type ListFunctionRestrictionsResponseBody struct {
	// The error code.
	//
	// example:
	//
	// not found
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpCode *int64 `json:"httpCode,omitempty" xml:"httpCode,omitempty"`
	// The execution duration.
	//
	// example:
	//
	// 10.444
	Latency *float64 `json:"latency,omitempty" xml:"latency,omitempty"`
	// The error message.
	//
	// example:
	//
	// "xx not found"
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 2423C841-91C4-5E51-B296-590D367967FC
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The returned result.
	Result []*ListFunctionRestrictionsResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
	// The request status.
	//
	// example:
	//
	// OK
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// The total number of entries.
	//
	// example:
	//
	// 1
	TotalCount *int64 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s ListFunctionRestrictionsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListFunctionRestrictionsResponseBody) GoString() string {
	return s.String()
}

func (s *ListFunctionRestrictionsResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListFunctionRestrictionsResponseBody) GetHttpCode() *int64 {
	return s.HttpCode
}

func (s *ListFunctionRestrictionsResponseBody) GetLatency() *float64 {
	return s.Latency
}

func (s *ListFunctionRestrictionsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListFunctionRestrictionsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListFunctionRestrictionsResponseBody) GetResult() []*ListFunctionRestrictionsResponseBodyResult {
	return s.Result
}

func (s *ListFunctionRestrictionsResponseBody) GetStatus() *string {
	return s.Status
}

func (s *ListFunctionRestrictionsResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListFunctionRestrictionsResponseBody) SetCode(v string) *ListFunctionRestrictionsResponseBody {
	s.Code = &v
	return s
}

func (s *ListFunctionRestrictionsResponseBody) SetHttpCode(v int64) *ListFunctionRestrictionsResponseBody {
	s.HttpCode = &v
	return s
}

func (s *ListFunctionRestrictionsResponseBody) SetLatency(v float64) *ListFunctionRestrictionsResponseBody {
	s.Latency = &v
	return s
}

func (s *ListFunctionRestrictionsResponseBody) SetMessage(v string) *ListFunctionRestrictionsResponseBody {
	s.Message = &v
	return s
}

func (s *ListFunctionRestrictionsResponseBody) SetRequestId(v string) *ListFunctionRestrictionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListFunctionRestrictionsResponseBody) SetResult(v []*ListFunctionRestrictionsResponseBodyResult) *ListFunctionRestrictionsResponseBody {
	s.Result = v
	return s
}

func (s *ListFunctionRestrictionsResponseBody) SetStatus(v string) *ListFunctionRestrictionsResponseBody {
	s.Status = &v
	return s
}

func (s *ListFunctionRestrictionsResponseBody) SetTotalCount(v int64) *ListFunctionRestrictionsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListFunctionRestrictionsResponseBody) Validate() error {
	if s.Result != nil {
		for _, item := range s.Result {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListFunctionRestrictionsResponseBodyResult struct {
	// The metadata.
	//
	// example:
	//
	// {
	//
	// 					"taskType":"text-embedding",
	//
	// 					"modelSource":[
	//
	// 						"ai_search"
	//
	// 					],
	//
	// 					"regionId":[
	//
	// 						"cn-hangzhou",
	//
	// 						"cn-zhangjiakou"
	//
	// 					],
	//
	// 					"instanceType":[
	//
	// 						"gpu.v100.16g.x1",
	//
	// 						"gpu.t4.16g.x1",
	//
	// 						"gpu.a10.24g.x1"
	//
	// 					]
	//
	// 				}
	Meta map[string]interface{} `json:"meta,omitempty" xml:"meta,omitempty"`
	// The rule name.
	//
	// example:
	//
	// test
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s ListFunctionRestrictionsResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s ListFunctionRestrictionsResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListFunctionRestrictionsResponseBodyResult) GetMeta() map[string]interface{} {
	return s.Meta
}

func (s *ListFunctionRestrictionsResponseBodyResult) GetName() *string {
	return s.Name
}

func (s *ListFunctionRestrictionsResponseBodyResult) SetMeta(v map[string]interface{}) *ListFunctionRestrictionsResponseBodyResult {
	s.Meta = v
	return s
}

func (s *ListFunctionRestrictionsResponseBodyResult) SetName(v string) *ListFunctionRestrictionsResponseBodyResult {
	s.Name = &v
	return s
}

func (s *ListFunctionRestrictionsResponseBodyResult) Validate() error {
	return dara.Validate(s)
}
