// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGenerateCopilotStreamResponseResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GenerateCopilotStreamResponseResponseBody
	GetCode() *string
	SetData(v string) *GenerateCopilotStreamResponseResponseBody
	GetData() *string
	SetMessage(v string) *GenerateCopilotStreamResponseResponseBody
	GetMessage() *string
	SetRequestId(v string) *GenerateCopilotStreamResponseResponseBody
	GetRequestId() *string
}

type GenerateCopilotStreamResponseResponseBody struct {
	// example:
	//
	// SysomOpenAPI.ServerError
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	Data *string `json:"data,omitempty" xml:"data,omitempty"`
	// example:
	//
	// Requests for llm service failed
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 2D693121-C925-5154-8DF6-C09A8B369822
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s GenerateCopilotStreamResponseResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GenerateCopilotStreamResponseResponseBody) GoString() string {
	return s.String()
}

func (s *GenerateCopilotStreamResponseResponseBody) GetCode() *string {
	return s.Code
}

func (s *GenerateCopilotStreamResponseResponseBody) GetData() *string {
	return s.Data
}

func (s *GenerateCopilotStreamResponseResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GenerateCopilotStreamResponseResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GenerateCopilotStreamResponseResponseBody) SetCode(v string) *GenerateCopilotStreamResponseResponseBody {
	s.Code = &v
	return s
}

func (s *GenerateCopilotStreamResponseResponseBody) SetData(v string) *GenerateCopilotStreamResponseResponseBody {
	s.Data = &v
	return s
}

func (s *GenerateCopilotStreamResponseResponseBody) SetMessage(v string) *GenerateCopilotStreamResponseResponseBody {
	s.Message = &v
	return s
}

func (s *GenerateCopilotStreamResponseResponseBody) SetRequestId(v string) *GenerateCopilotStreamResponseResponseBody {
	s.RequestId = &v
	return s
}

func (s *GenerateCopilotStreamResponseResponseBody) Validate() error {
	return dara.Validate(s)
}
