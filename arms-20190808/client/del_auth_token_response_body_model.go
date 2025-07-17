// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDelAuthTokenResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v string) *DelAuthTokenResponseBody
	GetData() *string
	SetRequestId(v string) *DelAuthTokenResponseBody
	GetRequestId() *string
}

type DelAuthTokenResponseBody struct {
	// Indicates whether the call was successful.
	//
	// example:
	//
	// success
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// The request ID.
	//
	// example:
	//
	// A5EC8221-08F2-4C95-9AF1-49FD998C647A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DelAuthTokenResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DelAuthTokenResponseBody) GoString() string {
	return s.String()
}

func (s *DelAuthTokenResponseBody) GetData() *string {
	return s.Data
}

func (s *DelAuthTokenResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DelAuthTokenResponseBody) SetData(v string) *DelAuthTokenResponseBody {
	s.Data = &v
	return s
}

func (s *DelAuthTokenResponseBody) SetRequestId(v string) *DelAuthTokenResponseBody {
	s.RequestId = &v
	return s
}

func (s *DelAuthTokenResponseBody) Validate() error {
	return dara.Validate(s)
}
