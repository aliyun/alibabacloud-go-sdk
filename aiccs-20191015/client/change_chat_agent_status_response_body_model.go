// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iChangeChatAgentStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ChangeChatAgentStatusResponseBody
	GetCode() *string
	SetData(v string) *ChangeChatAgentStatusResponseBody
	GetData() *string
	SetHttpStatusCode(v int32) *ChangeChatAgentStatusResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ChangeChatAgentStatusResponseBody
	GetMessage() *string
	SetRequestId(v string) *ChangeChatAgentStatusResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ChangeChatAgentStatusResponseBody
	GetSuccess() *bool
}

type ChangeChatAgentStatusResponseBody struct {
	// example:
	//
	// Success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 0
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// DF6A3FB7-A5AA-43BE-A65B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ChangeChatAgentStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ChangeChatAgentStatusResponseBody) GoString() string {
	return s.String()
}

func (s *ChangeChatAgentStatusResponseBody) GetCode() *string {
	return s.Code
}

func (s *ChangeChatAgentStatusResponseBody) GetData() *string {
	return s.Data
}

func (s *ChangeChatAgentStatusResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ChangeChatAgentStatusResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ChangeChatAgentStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ChangeChatAgentStatusResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ChangeChatAgentStatusResponseBody) SetCode(v string) *ChangeChatAgentStatusResponseBody {
	s.Code = &v
	return s
}

func (s *ChangeChatAgentStatusResponseBody) SetData(v string) *ChangeChatAgentStatusResponseBody {
	s.Data = &v
	return s
}

func (s *ChangeChatAgentStatusResponseBody) SetHttpStatusCode(v int32) *ChangeChatAgentStatusResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ChangeChatAgentStatusResponseBody) SetMessage(v string) *ChangeChatAgentStatusResponseBody {
	s.Message = &v
	return s
}

func (s *ChangeChatAgentStatusResponseBody) SetRequestId(v string) *ChangeChatAgentStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *ChangeChatAgentStatusResponseBody) SetSuccess(v bool) *ChangeChatAgentStatusResponseBody {
	s.Success = &v
	return s
}

func (s *ChangeChatAgentStatusResponseBody) Validate() error {
	return dara.Validate(s)
}
