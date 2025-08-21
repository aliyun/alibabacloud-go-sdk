// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCreatePrePaidInstanceResultResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceCreateResult(v *DescribeCreatePrePaidInstanceResultResponseBodyInstanceCreateResult) *DescribeCreatePrePaidInstanceResultResponseBody
	GetInstanceCreateResult() *DescribeCreatePrePaidInstanceResultResponseBodyInstanceCreateResult
	SetRequestId(v string) *DescribeCreatePrePaidInstanceResultResponseBody
	GetRequestId() *string
}

type DescribeCreatePrePaidInstanceResultResponseBody struct {
	// Returned results of creating an instance.
	InstanceCreateResult *DescribeCreatePrePaidInstanceResultResponseBodyInstanceCreateResult `json:"InstanceCreateResult,omitempty" xml:"InstanceCreateResult,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// AAE90880-4970-4D81-A534-A6C0F3631F74
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeCreatePrePaidInstanceResultResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeCreatePrePaidInstanceResultResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCreatePrePaidInstanceResultResponseBody) GetInstanceCreateResult() *DescribeCreatePrePaidInstanceResultResponseBodyInstanceCreateResult {
	return s.InstanceCreateResult
}

func (s *DescribeCreatePrePaidInstanceResultResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeCreatePrePaidInstanceResultResponseBody) SetInstanceCreateResult(v *DescribeCreatePrePaidInstanceResultResponseBodyInstanceCreateResult) *DescribeCreatePrePaidInstanceResultResponseBody {
	s.InstanceCreateResult = v
	return s
}

func (s *DescribeCreatePrePaidInstanceResultResponseBody) SetRequestId(v string) *DescribeCreatePrePaidInstanceResultResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeCreatePrePaidInstanceResultResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeCreatePrePaidInstanceResultResponseBodyInstanceCreateResult struct {
	// The status of the instance creation.
	//
	// 	- Accepted
	//
	// 	- Creating
	//
	// 	- Failed
	//
	// 	- Successed
	//
	// example:
	//
	// Successed
	InstanceCreateStatus *string `json:"InstanceCreateStatus,omitempty" xml:"InstanceCreateStatus,omitempty"`
	// The ID of the instance.
	//
	// example:
	//
	// i-6ecpqvkicnchxccozrpxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s DescribeCreatePrePaidInstanceResultResponseBodyInstanceCreateResult) String() string {
	return dara.Prettify(s)
}

func (s DescribeCreatePrePaidInstanceResultResponseBodyInstanceCreateResult) GoString() string {
	return s.String()
}

func (s *DescribeCreatePrePaidInstanceResultResponseBodyInstanceCreateResult) GetInstanceCreateStatus() *string {
	return s.InstanceCreateStatus
}

func (s *DescribeCreatePrePaidInstanceResultResponseBodyInstanceCreateResult) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeCreatePrePaidInstanceResultResponseBodyInstanceCreateResult) SetInstanceCreateStatus(v string) *DescribeCreatePrePaidInstanceResultResponseBodyInstanceCreateResult {
	s.InstanceCreateStatus = &v
	return s
}

func (s *DescribeCreatePrePaidInstanceResultResponseBodyInstanceCreateResult) SetInstanceId(v string) *DescribeCreatePrePaidInstanceResultResponseBodyInstanceCreateResult {
	s.InstanceId = &v
	return s
}

func (s *DescribeCreatePrePaidInstanceResultResponseBodyInstanceCreateResult) Validate() error {
	return dara.Validate(s)
}
