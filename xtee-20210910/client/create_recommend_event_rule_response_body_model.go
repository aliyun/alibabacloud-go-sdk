// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateRecommendEventRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateRecommendEventRuleResponseBody
	GetRequestId() *string
	SetResultObject(v bool) *CreateRecommendEventRuleResponseBody
	GetResultObject() *bool
}

type CreateRecommendEventRuleResponseBody struct {
	// Request ID.
	//
	// example:
	//
	// A32FE941-35F2-5378-B37C-4B8FDB16F094
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Return object.
	//
	// example:
	//
	// true
	ResultObject *bool `json:"resultObject,omitempty" xml:"resultObject,omitempty"`
}

func (s CreateRecommendEventRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateRecommendEventRuleResponseBody) GoString() string {
	return s.String()
}

func (s *CreateRecommendEventRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateRecommendEventRuleResponseBody) GetResultObject() *bool {
	return s.ResultObject
}

func (s *CreateRecommendEventRuleResponseBody) SetRequestId(v string) *CreateRecommendEventRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateRecommendEventRuleResponseBody) SetResultObject(v bool) *CreateRecommendEventRuleResponseBody {
	s.ResultObject = &v
	return s
}

func (s *CreateRecommendEventRuleResponseBody) Validate() error {
	return dara.Validate(s)
}
