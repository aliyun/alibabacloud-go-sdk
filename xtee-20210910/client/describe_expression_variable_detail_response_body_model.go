// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeExpressionVariableDetailResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeExpressionVariableDetailResponseBody
	GetRequestId() *string
	SetResultObject(v map[string]interface{}) *DescribeExpressionVariableDetailResponseBody
	GetResultObject() map[string]interface{}
}

type DescribeExpressionVariableDetailResponseBody struct {
	// Request ID.
	//
	// example:
	//
	// A32FE941-35F2-5378-B37C-4B8FDB16F094
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Returned object
	//
	// example:
	//
	// true
	ResultObject map[string]interface{} `json:"resultObject,omitempty" xml:"resultObject,omitempty"`
}

func (s DescribeExpressionVariableDetailResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeExpressionVariableDetailResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeExpressionVariableDetailResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeExpressionVariableDetailResponseBody) GetResultObject() map[string]interface{} {
	return s.ResultObject
}

func (s *DescribeExpressionVariableDetailResponseBody) SetRequestId(v string) *DescribeExpressionVariableDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeExpressionVariableDetailResponseBody) SetResultObject(v map[string]interface{}) *DescribeExpressionVariableDetailResponseBody {
	s.ResultObject = v
	return s
}

func (s *DescribeExpressionVariableDetailResponseBody) Validate() error {
	return dara.Validate(s)
}
