// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopChannelResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *StopChannelResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *StopChannelResponseBody
	GetSuccess() *bool
}

type StopChannelResponseBody struct {
	// **Request ID**
	//
	// example:
	//
	// xxx-xxxx-xxxxx-xxxx
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s StopChannelResponseBody) String() string {
	return dara.Prettify(s)
}

func (s StopChannelResponseBody) GoString() string {
	return s.String()
}

func (s *StopChannelResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *StopChannelResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *StopChannelResponseBody) SetRequestId(v string) *StopChannelResponseBody {
	s.RequestId = &v
	return s
}

func (s *StopChannelResponseBody) SetSuccess(v bool) *StopChannelResponseBody {
	s.Success = &v
	return s
}

func (s *StopChannelResponseBody) Validate() error {
	return dara.Validate(s)
}
