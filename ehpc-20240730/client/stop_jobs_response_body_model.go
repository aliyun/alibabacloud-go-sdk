// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopJobsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *StopJobsResponseBody
	GetRequestId() *string
	SetSuccess(v string) *StopJobsResponseBody
	GetSuccess() *string
}

type StopJobsResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// F8868A00-6757-5542-BDD6-E1040D94****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- true: The request was successful.
	//
	// 	- false: The request failed.
	//
	// example:
	//
	// true
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s StopJobsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s StopJobsResponseBody) GoString() string {
	return s.String()
}

func (s *StopJobsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *StopJobsResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *StopJobsResponseBody) SetRequestId(v string) *StopJobsResponseBody {
	s.RequestId = &v
	return s
}

func (s *StopJobsResponseBody) SetSuccess(v string) *StopJobsResponseBody {
	s.Success = &v
	return s
}

func (s *StopJobsResponseBody) Validate() error {
	return dara.Validate(s)
}
