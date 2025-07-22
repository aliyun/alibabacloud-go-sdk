// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *StopInstanceResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *StopInstanceResponseBody
	GetHttpStatusCode() *int32
	SetInstanceId(v string) *StopInstanceResponseBody
	GetInstanceId() *string
	SetMessage(v string) *StopInstanceResponseBody
	GetMessage() *string
	SetRequestId(v string) *StopInstanceResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *StopInstanceResponseBody
	GetSuccess() *bool
}

type StopInstanceResponseBody struct {
	// example:
	//
	// null
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// null
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// dsw-730xxxxxxxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// "XXX"
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// E7D55162-4489-1619-AAF5-3F97D5FCA948
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s StopInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s StopInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *StopInstanceResponseBody) GetCode() *string {
	return s.Code
}

func (s *StopInstanceResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *StopInstanceResponseBody) GetInstanceId() *string {
	return s.InstanceId
}

func (s *StopInstanceResponseBody) GetMessage() *string {
	return s.Message
}

func (s *StopInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *StopInstanceResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *StopInstanceResponseBody) SetCode(v string) *StopInstanceResponseBody {
	s.Code = &v
	return s
}

func (s *StopInstanceResponseBody) SetHttpStatusCode(v int32) *StopInstanceResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *StopInstanceResponseBody) SetInstanceId(v string) *StopInstanceResponseBody {
	s.InstanceId = &v
	return s
}

func (s *StopInstanceResponseBody) SetMessage(v string) *StopInstanceResponseBody {
	s.Message = &v
	return s
}

func (s *StopInstanceResponseBody) SetRequestId(v string) *StopInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *StopInstanceResponseBody) SetSuccess(v bool) *StopInstanceResponseBody {
	s.Success = &v
	return s
}

func (s *StopInstanceResponseBody) Validate() error {
	return dara.Validate(s)
}
