// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteTemplateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeleteTemplateResponseBody
	GetCode() *string
	SetMessage(v string) *DeleteTemplateResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteTemplateResponseBody
	GetRequestId() *string
}

type DeleteTemplateResponseBody struct {
	// The error code.
	//
	// example:
	//
	// 200
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The response message.
	//
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// B5AD8B54-4358-5F5B-ACAA-52F2016459C6
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s DeleteTemplateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteTemplateResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteTemplateResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteTemplateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteTemplateResponseBody) SetCode(v string) *DeleteTemplateResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteTemplateResponseBody) SetMessage(v string) *DeleteTemplateResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteTemplateResponseBody) SetRequestId(v string) *DeleteTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteTemplateResponseBody) Validate() error {
	return dara.Validate(s)
}
