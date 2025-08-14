// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteChannelResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteChannelResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteChannelResponseBody
	GetSuccess() *bool
}

type DeleteChannelResponseBody struct {
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

func (s DeleteChannelResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteChannelResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteChannelResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteChannelResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteChannelResponseBody) SetRequestId(v string) *DeleteChannelResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteChannelResponseBody) SetSuccess(v bool) *DeleteChannelResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteChannelResponseBody) Validate() error {
	return dara.Validate(s)
}
