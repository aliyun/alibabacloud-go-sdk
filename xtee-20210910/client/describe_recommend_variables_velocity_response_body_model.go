// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRecommendVariablesVelocityResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeRecommendVariablesVelocityResponseBody
	GetRequestId() *string
	SetResultObject(v bool) *DescribeRecommendVariablesVelocityResponseBody
	GetResultObject() *bool
}

type DescribeRecommendVariablesVelocityResponseBody struct {
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

func (s DescribeRecommendVariablesVelocityResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeRecommendVariablesVelocityResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRecommendVariablesVelocityResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeRecommendVariablesVelocityResponseBody) GetResultObject() *bool {
	return s.ResultObject
}

func (s *DescribeRecommendVariablesVelocityResponseBody) SetRequestId(v string) *DescribeRecommendVariablesVelocityResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeRecommendVariablesVelocityResponseBody) SetResultObject(v bool) *DescribeRecommendVariablesVelocityResponseBody {
	s.ResultObject = &v
	return s
}

func (s *DescribeRecommendVariablesVelocityResponseBody) Validate() error {
	return dara.Validate(s)
}
