// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRecommendSceneVariablesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeRecommendSceneVariablesResponseBody
	GetRequestId() *string
	SetResultObject(v bool) *DescribeRecommendSceneVariablesResponseBody
	GetResultObject() *bool
}

type DescribeRecommendSceneVariablesResponseBody struct {
	// Request ID.
	//
	// example:
	//
	// A32FE941-35F2-5378-B37C-4B8FDB16F094
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Return Object
	//
	// example:
	//
	// true
	ResultObject *bool `json:"resultObject,omitempty" xml:"resultObject,omitempty"`
}

func (s DescribeRecommendSceneVariablesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeRecommendSceneVariablesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRecommendSceneVariablesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeRecommendSceneVariablesResponseBody) GetResultObject() *bool {
	return s.ResultObject
}

func (s *DescribeRecommendSceneVariablesResponseBody) SetRequestId(v string) *DescribeRecommendSceneVariablesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeRecommendSceneVariablesResponseBody) SetResultObject(v bool) *DescribeRecommendSceneVariablesResponseBody {
	s.ResultObject = &v
	return s
}

func (s *DescribeRecommendSceneVariablesResponseBody) Validate() error {
	return dara.Validate(s)
}
