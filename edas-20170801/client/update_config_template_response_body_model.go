// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateConfigTemplateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *UpdateConfigTemplateResponseBody
	GetCode() *int32
	SetMessage(v string) *UpdateConfigTemplateResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateConfigTemplateResponseBody
	GetRequestId() *string
}

type UpdateConfigTemplateResponseBody struct {
	// The HTTP status code that is returned.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The additional information that is returned.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// D16979DC-4D42-*************
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateConfigTemplateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateConfigTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateConfigTemplateResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *UpdateConfigTemplateResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateConfigTemplateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateConfigTemplateResponseBody) SetCode(v int32) *UpdateConfigTemplateResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateConfigTemplateResponseBody) SetMessage(v string) *UpdateConfigTemplateResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateConfigTemplateResponseBody) SetRequestId(v string) *UpdateConfigTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateConfigTemplateResponseBody) Validate() error {
	return dara.Validate(s)
}
