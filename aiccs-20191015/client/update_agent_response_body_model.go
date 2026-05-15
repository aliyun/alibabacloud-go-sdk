// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAgentResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UpdateAgentResponseBody
	GetCode() *string
	SetHttpStatusCode(v int64) *UpdateAgentResponseBody
	GetHttpStatusCode() *int64
	SetMessage(v string) *UpdateAgentResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateAgentResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateAgentResponseBody
	GetSuccess() *bool
}

type UpdateAgentResponseBody struct {
	// example:
	//
	// Success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int64 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// OK
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

func (s UpdateAgentResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateAgentResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateAgentResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpdateAgentResponseBody) GetHttpStatusCode() *int64 {
	return s.HttpStatusCode
}

func (s *UpdateAgentResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateAgentResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateAgentResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateAgentResponseBody) SetCode(v string) *UpdateAgentResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateAgentResponseBody) SetHttpStatusCode(v int64) *UpdateAgentResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *UpdateAgentResponseBody) SetMessage(v string) *UpdateAgentResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateAgentResponseBody) SetRequestId(v string) *UpdateAgentResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateAgentResponseBody) SetSuccess(v bool) *UpdateAgentResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateAgentResponseBody) Validate() error {
	return dara.Validate(s)
}
