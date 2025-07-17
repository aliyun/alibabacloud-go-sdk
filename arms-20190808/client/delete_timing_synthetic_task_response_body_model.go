// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteTimingSyntheticTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int64) *DeleteTimingSyntheticTaskResponseBody
	GetCode() *int64
	SetData(v bool) *DeleteTimingSyntheticTaskResponseBody
	GetData() *bool
	SetMessage(v string) *DeleteTimingSyntheticTaskResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteTimingSyntheticTaskResponseBody
	GetRequestId() *string
}

type DeleteTimingSyntheticTaskResponseBody struct {
	// The HTTP status code. The status code 200 indicates that the request was successful. Other status codes indicate that the request failed.
	//
	// example:
	//
	// 200
	Code *int64 `json:"Code,omitempty" xml:"Code,omitempty"`
	// Indicates whether the synthetic monitoring task was deleted. true: The synthetic monitoring task was deleted. false: The synthetic monitoring task failed to be deleted.
	//
	// example:
	//
	// true
	Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
	// The returned message.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// A5EC8221-08F2-4C95-9AF1-49FD998C****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteTimingSyntheticTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteTimingSyntheticTaskResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteTimingSyntheticTaskResponseBody) GetCode() *int64 {
	return s.Code
}

func (s *DeleteTimingSyntheticTaskResponseBody) GetData() *bool {
	return s.Data
}

func (s *DeleteTimingSyntheticTaskResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteTimingSyntheticTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteTimingSyntheticTaskResponseBody) SetCode(v int64) *DeleteTimingSyntheticTaskResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteTimingSyntheticTaskResponseBody) SetData(v bool) *DeleteTimingSyntheticTaskResponseBody {
	s.Data = &v
	return s
}

func (s *DeleteTimingSyntheticTaskResponseBody) SetMessage(v string) *DeleteTimingSyntheticTaskResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteTimingSyntheticTaskResponseBody) SetRequestId(v string) *DeleteTimingSyntheticTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteTimingSyntheticTaskResponseBody) Validate() error {
	return dara.Validate(s)
}
