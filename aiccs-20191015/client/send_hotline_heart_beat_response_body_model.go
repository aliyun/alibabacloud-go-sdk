// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSendHotlineHeartBeatResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *SendHotlineHeartBeatResponseBody
	GetCode() *string
	SetMessage(v string) *SendHotlineHeartBeatResponseBody
	GetMessage() *string
	SetRequestId(v string) *SendHotlineHeartBeatResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *SendHotlineHeartBeatResponseBody
	GetSuccess() *bool
}

type SendHotlineHeartBeatResponseBody struct {
	// example:
	//
	// Success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// xxxx
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// EE338D98-9BD3-4413-B165
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SendHotlineHeartBeatResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SendHotlineHeartBeatResponseBody) GoString() string {
	return s.String()
}

func (s *SendHotlineHeartBeatResponseBody) GetCode() *string {
	return s.Code
}

func (s *SendHotlineHeartBeatResponseBody) GetMessage() *string {
	return s.Message
}

func (s *SendHotlineHeartBeatResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SendHotlineHeartBeatResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *SendHotlineHeartBeatResponseBody) SetCode(v string) *SendHotlineHeartBeatResponseBody {
	s.Code = &v
	return s
}

func (s *SendHotlineHeartBeatResponseBody) SetMessage(v string) *SendHotlineHeartBeatResponseBody {
	s.Message = &v
	return s
}

func (s *SendHotlineHeartBeatResponseBody) SetRequestId(v string) *SendHotlineHeartBeatResponseBody {
	s.RequestId = &v
	return s
}

func (s *SendHotlineHeartBeatResponseBody) SetSuccess(v bool) *SendHotlineHeartBeatResponseBody {
	s.Success = &v
	return s
}

func (s *SendHotlineHeartBeatResponseBody) Validate() error {
	return dara.Validate(s)
}
