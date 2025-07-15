// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyTemplateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ModifyTemplateResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *ModifyTemplateResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ModifyTemplateResponseBody
	GetMessage() *string
	SetRequestId(v string) *ModifyTemplateResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ModifyTemplateResponseBody
	GetSuccess() *bool
}

type ModifyTemplateResponseBody struct {
	// example:
	//
	// success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
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
	// 1CBAFFAB-B697-4049-A9B1-67E1FC5F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ModifyTemplateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyTemplateResponseBody) GetCode() *string {
	return s.Code
}

func (s *ModifyTemplateResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ModifyTemplateResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ModifyTemplateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyTemplateResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ModifyTemplateResponseBody) SetCode(v string) *ModifyTemplateResponseBody {
	s.Code = &v
	return s
}

func (s *ModifyTemplateResponseBody) SetHttpStatusCode(v int32) *ModifyTemplateResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ModifyTemplateResponseBody) SetMessage(v string) *ModifyTemplateResponseBody {
	s.Message = &v
	return s
}

func (s *ModifyTemplateResponseBody) SetRequestId(v string) *ModifyTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyTemplateResponseBody) SetSuccess(v bool) *ModifyTemplateResponseBody {
	s.Success = &v
	return s
}

func (s *ModifyTemplateResponseBody) Validate() error {
	return dara.Validate(s)
}
