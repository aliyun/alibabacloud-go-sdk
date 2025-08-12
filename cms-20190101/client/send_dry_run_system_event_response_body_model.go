// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSendDryRunSystemEventResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *SendDryRunSystemEventResponseBody
	GetCode() *string
	SetMessage(v string) *SendDryRunSystemEventResponseBody
	GetMessage() *string
	SetRequestId(v string) *SendDryRunSystemEventResponseBody
	GetRequestId() *string
	SetSuccess(v string) *SendDryRunSystemEventResponseBody
	GetSuccess() *string
}

type SendDryRunSystemEventResponseBody struct {
	// The status code.
	//
	// >  The status code 200 indicates that the request was successful.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
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
	// 486029C9-53E1-44B4-85A8-16A571A043FD
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
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SendDryRunSystemEventResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SendDryRunSystemEventResponseBody) GoString() string {
	return s.String()
}

func (s *SendDryRunSystemEventResponseBody) GetCode() *string {
	return s.Code
}

func (s *SendDryRunSystemEventResponseBody) GetMessage() *string {
	return s.Message
}

func (s *SendDryRunSystemEventResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SendDryRunSystemEventResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *SendDryRunSystemEventResponseBody) SetCode(v string) *SendDryRunSystemEventResponseBody {
	s.Code = &v
	return s
}

func (s *SendDryRunSystemEventResponseBody) SetMessage(v string) *SendDryRunSystemEventResponseBody {
	s.Message = &v
	return s
}

func (s *SendDryRunSystemEventResponseBody) SetRequestId(v string) *SendDryRunSystemEventResponseBody {
	s.RequestId = &v
	return s
}

func (s *SendDryRunSystemEventResponseBody) SetSuccess(v string) *SendDryRunSystemEventResponseBody {
	s.Success = &v
	return s
}

func (s *SendDryRunSystemEventResponseBody) Validate() error {
	return dara.Validate(s)
}
