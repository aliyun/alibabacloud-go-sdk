// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteTaskResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteTaskResponseBody
	GetSuccess() *bool
}

type DeleteTaskResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 22C97E95-F023-56B5-8852-B1A77A17XXXX
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteTaskResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteTaskResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteTaskResponseBody) SetRequestId(v string) *DeleteTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteTaskResponseBody) SetSuccess(v bool) *DeleteTaskResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteTaskResponseBody) Validate() error {
	return dara.Validate(s)
}
