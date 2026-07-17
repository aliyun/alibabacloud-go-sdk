// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveElasticPlanResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *RemoveElasticPlanResponseBody
	GetCode() *string
	SetHttpCode(v int64) *RemoveElasticPlanResponseBody
	GetHttpCode() *int64
	SetMessage(v string) *RemoveElasticPlanResponseBody
	GetMessage() *string
	SetRequestId(v string) *RemoveElasticPlanResponseBody
	GetRequestId() *string
	SetResult(v map[string]interface{}) *RemoveElasticPlanResponseBody
	GetResult() map[string]interface{}
}

type RemoveElasticPlanResponseBody struct {
	// example:
	//
	// ElasticPlan.NotFound
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// example:
	//
	// 200
	HttpCode *int64 `json:"httpCode,omitempty" xml:"httpCode,omitempty"`
	// example:
	//
	// Elastic plan not found
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// example:
	//
	// 0A6EB64B-B4C8-CF02-810F-E660812972FF
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// -
	Result map[string]interface{} `json:"result,omitempty" xml:"result,omitempty"`
}

func (s RemoveElasticPlanResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RemoveElasticPlanResponseBody) GoString() string {
	return s.String()
}

func (s *RemoveElasticPlanResponseBody) GetCode() *string {
	return s.Code
}

func (s *RemoveElasticPlanResponseBody) GetHttpCode() *int64 {
	return s.HttpCode
}

func (s *RemoveElasticPlanResponseBody) GetMessage() *string {
	return s.Message
}

func (s *RemoveElasticPlanResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RemoveElasticPlanResponseBody) GetResult() map[string]interface{} {
	return s.Result
}

func (s *RemoveElasticPlanResponseBody) SetCode(v string) *RemoveElasticPlanResponseBody {
	s.Code = &v
	return s
}

func (s *RemoveElasticPlanResponseBody) SetHttpCode(v int64) *RemoveElasticPlanResponseBody {
	s.HttpCode = &v
	return s
}

func (s *RemoveElasticPlanResponseBody) SetMessage(v string) *RemoveElasticPlanResponseBody {
	s.Message = &v
	return s
}

func (s *RemoveElasticPlanResponseBody) SetRequestId(v string) *RemoveElasticPlanResponseBody {
	s.RequestId = &v
	return s
}

func (s *RemoveElasticPlanResponseBody) SetResult(v map[string]interface{}) *RemoveElasticPlanResponseBody {
	s.Result = v
	return s
}

func (s *RemoveElasticPlanResponseBody) Validate() error {
	return dara.Validate(s)
}
