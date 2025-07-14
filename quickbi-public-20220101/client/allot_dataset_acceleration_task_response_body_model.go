// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAllotDatasetAccelerationTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *AllotDatasetAccelerationTaskResponseBody
	GetRequestId() *string
	SetResult(v bool) *AllotDatasetAccelerationTaskResponseBody
	GetResult() *bool
	SetSuccess(v bool) *AllotDatasetAccelerationTaskResponseBody
	GetSuccess() *bool
}

type AllotDatasetAccelerationTaskResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// D8749D65-E80A-433C-AF1B-CE9C180FF3B4
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the IP address whitelist is updated. Valid values:
	//
	// 	- true: The task is triggered.
	//
	// 	- false: The task fails to be triggered.
	//
	// example:
	//
	// true
	Result *bool `json:"Result,omitempty" xml:"Result,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- true: The request was successful.
	//
	// 	- false: The request failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s AllotDatasetAccelerationTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AllotDatasetAccelerationTaskResponseBody) GoString() string {
	return s.String()
}

func (s *AllotDatasetAccelerationTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AllotDatasetAccelerationTaskResponseBody) GetResult() *bool {
	return s.Result
}

func (s *AllotDatasetAccelerationTaskResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *AllotDatasetAccelerationTaskResponseBody) SetRequestId(v string) *AllotDatasetAccelerationTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *AllotDatasetAccelerationTaskResponseBody) SetResult(v bool) *AllotDatasetAccelerationTaskResponseBody {
	s.Result = &v
	return s
}

func (s *AllotDatasetAccelerationTaskResponseBody) SetSuccess(v bool) *AllotDatasetAccelerationTaskResponseBody {
	s.Success = &v
	return s
}

func (s *AllotDatasetAccelerationTaskResponseBody) Validate() error {
	return dara.Validate(s)
}
