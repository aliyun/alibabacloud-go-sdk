// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateTemplateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *UpdateTemplateResponseBody
	GetCode() *int32
	SetDetails(v string) *UpdateTemplateResponseBody
	GetDetails() *string
	SetErrorCode(v string) *UpdateTemplateResponseBody
	GetErrorCode() *string
	SetMessage(v string) *UpdateTemplateResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateTemplateResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateTemplateResponseBody
	GetSuccess() *bool
	SetTemplateId(v string) *UpdateTemplateResponseBody
	GetTemplateId() *string
}

type UpdateTemplateResponseBody struct {
	// Total amount of data under the current request conditions. This parameter is optional and does not need to be returned by default.
	//
	// example:
	//
	// 0
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// Details
	//
	// example:
	//
	// null
	Details *string `json:"Details,omitempty" xml:"Details,omitempty"`
	// error code
	//
	// example:
	//
	// ""
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// Return message of the request
	//
	// This parameter is required.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Request ID
	//
	// example:
	//
	// 90ABA848-AD74-1F6E-84BC-4182A7F1F29E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// is succeeded
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// Template ID
	//
	// example:
	//
	// 1529360348342353920
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s UpdateTemplateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateTemplateResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *UpdateTemplateResponseBody) GetDetails() *string {
	return s.Details
}

func (s *UpdateTemplateResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *UpdateTemplateResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateTemplateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateTemplateResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateTemplateResponseBody) GetTemplateId() *string {
	return s.TemplateId
}

func (s *UpdateTemplateResponseBody) SetCode(v int32) *UpdateTemplateResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateTemplateResponseBody) SetDetails(v string) *UpdateTemplateResponseBody {
	s.Details = &v
	return s
}

func (s *UpdateTemplateResponseBody) SetErrorCode(v string) *UpdateTemplateResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *UpdateTemplateResponseBody) SetMessage(v string) *UpdateTemplateResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateTemplateResponseBody) SetRequestId(v string) *UpdateTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateTemplateResponseBody) SetSuccess(v bool) *UpdateTemplateResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateTemplateResponseBody) SetTemplateId(v string) *UpdateTemplateResponseBody {
	s.TemplateId = &v
	return s
}

func (s *UpdateTemplateResponseBody) Validate() error {
	return dara.Validate(s)
}
