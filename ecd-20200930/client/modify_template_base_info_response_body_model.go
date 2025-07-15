// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyTemplateBaseInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ModifyTemplateBaseInfoResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *ModifyTemplateBaseInfoResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ModifyTemplateBaseInfoResponseBody
	GetMessage() *string
	SetRequestId(v string) *ModifyTemplateBaseInfoResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ModifyTemplateBaseInfoResponseBody
	GetSuccess() *bool
}

type ModifyTemplateBaseInfoResponseBody struct {
	// The execution result of the operation. If the request was successful, `success` is returned. If the request failed, an error message is returned.
	//
	// example:
	//
	// success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The error message. This parameter is not returned if the value of Code is `success`.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 3D99C38F-DE7D-5109-BC06-43EC5D*****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ModifyTemplateBaseInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyTemplateBaseInfoResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyTemplateBaseInfoResponseBody) GetCode() *string {
	return s.Code
}

func (s *ModifyTemplateBaseInfoResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ModifyTemplateBaseInfoResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ModifyTemplateBaseInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyTemplateBaseInfoResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ModifyTemplateBaseInfoResponseBody) SetCode(v string) *ModifyTemplateBaseInfoResponseBody {
	s.Code = &v
	return s
}

func (s *ModifyTemplateBaseInfoResponseBody) SetHttpStatusCode(v int32) *ModifyTemplateBaseInfoResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ModifyTemplateBaseInfoResponseBody) SetMessage(v string) *ModifyTemplateBaseInfoResponseBody {
	s.Message = &v
	return s
}

func (s *ModifyTemplateBaseInfoResponseBody) SetRequestId(v string) *ModifyTemplateBaseInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyTemplateBaseInfoResponseBody) SetSuccess(v bool) *ModifyTemplateBaseInfoResponseBody {
	s.Success = &v
	return s
}

func (s *ModifyTemplateBaseInfoResponseBody) Validate() error {
	return dara.Validate(s)
}
