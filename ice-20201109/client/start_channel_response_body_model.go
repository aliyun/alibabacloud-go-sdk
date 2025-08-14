// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartChannelResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *StartChannelResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *StartChannelResponseBody
	GetSuccess() *bool
}

type StartChannelResponseBody struct {
	// **Request ID**
	//
	// example:
	//
	// xxx-xxxx-xxxxx-xxxx
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s StartChannelResponseBody) String() string {
	return dara.Prettify(s)
}

func (s StartChannelResponseBody) GoString() string {
	return s.String()
}

func (s *StartChannelResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *StartChannelResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *StartChannelResponseBody) SetRequestId(v string) *StartChannelResponseBody {
	s.RequestId = &v
	return s
}

func (s *StartChannelResponseBody) SetSuccess(v bool) *StartChannelResponseBody {
	s.Success = &v
	return s
}

func (s *StartChannelResponseBody) Validate() error {
	return dara.Validate(s)
}
