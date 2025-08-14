// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRecommendVelocitiesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeRecommendVelocitiesResponseBody
	GetRequestId() *string
	SetResultObject(v bool) *DescribeRecommendVelocitiesResponseBody
	GetResultObject() *bool
}

type DescribeRecommendVelocitiesResponseBody struct {
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
	ResultObject *bool `json:"resultObject,omitempty" xml:"resultObject,omitempty"`
}

func (s DescribeRecommendVelocitiesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeRecommendVelocitiesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRecommendVelocitiesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeRecommendVelocitiesResponseBody) GetResultObject() *bool {
	return s.ResultObject
}

func (s *DescribeRecommendVelocitiesResponseBody) SetRequestId(v string) *DescribeRecommendVelocitiesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeRecommendVelocitiesResponseBody) SetResultObject(v bool) *DescribeRecommendVelocitiesResponseBody {
	s.ResultObject = &v
	return s
}

func (s *DescribeRecommendVelocitiesResponseBody) Validate() error {
	return dara.Validate(s)
}
