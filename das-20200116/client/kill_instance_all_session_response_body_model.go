// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iKillInstanceAllSessionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *KillInstanceAllSessionResponseBody
	GetCode() *string
	SetData(v string) *KillInstanceAllSessionResponseBody
	GetData() *string
	SetMessage(v string) *KillInstanceAllSessionResponseBody
	GetMessage() *string
	SetRequestId(v string) *KillInstanceAllSessionResponseBody
	GetRequestId() *string
	SetSuccess(v string) *KillInstanceAllSessionResponseBody
	GetSuccess() *string
}

type KillInstanceAllSessionResponseBody struct {
	// The HTTP status code returned.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The reserved parameter.
	//
	// example:
	//
	// None
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// The returned message.
	//
	// >  If the request was successful, **Successful*	- is returned. If the request failed, an error message such as an error code is returned.
	//
	// example:
	//
	// Successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// B6D17591-B48B-4D31-9CD6-9B9796B2****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s KillInstanceAllSessionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s KillInstanceAllSessionResponseBody) GoString() string {
	return s.String()
}

func (s *KillInstanceAllSessionResponseBody) GetCode() *string {
	return s.Code
}

func (s *KillInstanceAllSessionResponseBody) GetData() *string {
	return s.Data
}

func (s *KillInstanceAllSessionResponseBody) GetMessage() *string {
	return s.Message
}

func (s *KillInstanceAllSessionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *KillInstanceAllSessionResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *KillInstanceAllSessionResponseBody) SetCode(v string) *KillInstanceAllSessionResponseBody {
	s.Code = &v
	return s
}

func (s *KillInstanceAllSessionResponseBody) SetData(v string) *KillInstanceAllSessionResponseBody {
	s.Data = &v
	return s
}

func (s *KillInstanceAllSessionResponseBody) SetMessage(v string) *KillInstanceAllSessionResponseBody {
	s.Message = &v
	return s
}

func (s *KillInstanceAllSessionResponseBody) SetRequestId(v string) *KillInstanceAllSessionResponseBody {
	s.RequestId = &v
	return s
}

func (s *KillInstanceAllSessionResponseBody) SetSuccess(v string) *KillInstanceAllSessionResponseBody {
	s.Success = &v
	return s
}

func (s *KillInstanceAllSessionResponseBody) Validate() error {
	return dara.Validate(s)
}
