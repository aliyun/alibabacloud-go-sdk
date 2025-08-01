// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGenerateCopilotResponseResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GenerateCopilotResponseResponseBody
	GetCode() *string
	SetData(v string) *GenerateCopilotResponseResponseBody
	GetData() *string
	SetMassage(v string) *GenerateCopilotResponseResponseBody
	GetMassage() *string
	SetRequestId(v string) *GenerateCopilotResponseResponseBody
	GetRequestId() *string
}

type GenerateCopilotResponseResponseBody struct {
	// example:
	//
	// SysomOpenAPI.ServerError
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	Data *string `json:"data,omitempty" xml:"data,omitempty"`
	// example:
	//
	// Requests for llm service failed
	Massage *string `json:"massage,omitempty" xml:"massage,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 2D693121-C925-5154-8DF6-C09A8B369822
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s GenerateCopilotResponseResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GenerateCopilotResponseResponseBody) GoString() string {
	return s.String()
}

func (s *GenerateCopilotResponseResponseBody) GetCode() *string {
	return s.Code
}

func (s *GenerateCopilotResponseResponseBody) GetData() *string {
	return s.Data
}

func (s *GenerateCopilotResponseResponseBody) GetMassage() *string {
	return s.Massage
}

func (s *GenerateCopilotResponseResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GenerateCopilotResponseResponseBody) SetCode(v string) *GenerateCopilotResponseResponseBody {
	s.Code = &v
	return s
}

func (s *GenerateCopilotResponseResponseBody) SetData(v string) *GenerateCopilotResponseResponseBody {
	s.Data = &v
	return s
}

func (s *GenerateCopilotResponseResponseBody) SetMassage(v string) *GenerateCopilotResponseResponseBody {
	s.Massage = &v
	return s
}

func (s *GenerateCopilotResponseResponseBody) SetRequestId(v string) *GenerateCopilotResponseResponseBody {
	s.RequestId = &v
	return s
}

func (s *GenerateCopilotResponseResponseBody) Validate() error {
	return dara.Validate(s)
}
