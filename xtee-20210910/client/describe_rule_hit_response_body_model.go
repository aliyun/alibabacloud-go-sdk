// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRuleHitResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeRuleHitResponseBody
	GetRequestId() *string
	SetResultObject(v map[string]interface{}) *DescribeRuleHitResponseBody
	GetResultObject() map[string]interface{}
}

type DescribeRuleHitResponseBody struct {
	// Request ID.
	//
	// example:
	//
	// A32FE941-35F2-5378-B37C-4B8FDB16F094
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Return object
	//
	// example:
	//
	// true
	ResultObject map[string]interface{} `json:"resultObject,omitempty" xml:"resultObject,omitempty"`
}

func (s DescribeRuleHitResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeRuleHitResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRuleHitResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeRuleHitResponseBody) GetResultObject() map[string]interface{} {
	return s.ResultObject
}

func (s *DescribeRuleHitResponseBody) SetRequestId(v string) *DescribeRuleHitResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeRuleHitResponseBody) SetResultObject(v map[string]interface{}) *DescribeRuleHitResponseBody {
	s.ResultObject = v
	return s
}

func (s *DescribeRuleHitResponseBody) Validate() error {
	return dara.Validate(s)
}
