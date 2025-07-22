// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *StartInstanceResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *StartInstanceResponseBody
	GetHttpStatusCode() *int32
	SetInstanceId(v string) *StartInstanceResponseBody
	GetInstanceId() *string
	SetMessage(v string) *StartInstanceResponseBody
	GetMessage() *string
	SetRequestId(v string) *StartInstanceResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *StartInstanceResponseBody
	GetSuccess() *bool
}

type StartInstanceResponseBody struct {
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

func (s StartInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s StartInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *StartInstanceResponseBody) GetCode() *string {
	return s.Code
}

func (s *StartInstanceResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *StartInstanceResponseBody) GetInstanceId() *string {
	return s.InstanceId
}

func (s *StartInstanceResponseBody) GetMessage() *string {
	return s.Message
}

func (s *StartInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *StartInstanceResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *StartInstanceResponseBody) SetCode(v string) *StartInstanceResponseBody {
	s.Code = &v
	return s
}

func (s *StartInstanceResponseBody) SetHttpStatusCode(v int32) *StartInstanceResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *StartInstanceResponseBody) SetInstanceId(v string) *StartInstanceResponseBody {
	s.InstanceId = &v
	return s
}

func (s *StartInstanceResponseBody) SetMessage(v string) *StartInstanceResponseBody {
	s.Message = &v
	return s
}

func (s *StartInstanceResponseBody) SetRequestId(v string) *StartInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *StartInstanceResponseBody) SetSuccess(v bool) *StartInstanceResponseBody {
	s.Success = &v
	return s
}

func (s *StartInstanceResponseBody) Validate() error {
	return dara.Validate(s)
}
