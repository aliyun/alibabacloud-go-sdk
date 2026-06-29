// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateTemplateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *CreateTemplateResponseBody
	GetCode() *int32
	SetDetails(v string) *CreateTemplateResponseBody
	GetDetails() *string
	SetErrorCode(v string) *CreateTemplateResponseBody
	GetErrorCode() *string
	SetMessage(v string) *CreateTemplateResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateTemplateResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateTemplateResponseBody
	GetSuccess() *bool
	SetTemplateId(v string) *CreateTemplateResponseBody
	GetTemplateId() *string
}

type CreateTemplateResponseBody struct {
	// Return code. The default value is 0, indicating normal execution.
	//
	// example:
	//
	// 0
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// Details.
	//
	// example:
	//
	// -
	Details *string `json:"Details,omitempty" xml:"Details,omitempty"`
	// Error code.
	//
	// example:
	//
	// -
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// Response message of the request.
	//
	// This parameter is required.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Request ID.
	//
	// example:
	//
	// 90ABA848-AD74-1F6E-84BC-4182A7F1****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the operation succeeded. Valid values:
	//
	// - true: The request succeeded.
	//
	// - false: The request failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// Template ID.
	//
	// example:
	//
	// 152***0348342353920
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s CreateTemplateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *CreateTemplateResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *CreateTemplateResponseBody) GetDetails() *string {
	return s.Details
}

func (s *CreateTemplateResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *CreateTemplateResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateTemplateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateTemplateResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateTemplateResponseBody) GetTemplateId() *string {
	return s.TemplateId
}

func (s *CreateTemplateResponseBody) SetCode(v int32) *CreateTemplateResponseBody {
	s.Code = &v
	return s
}

func (s *CreateTemplateResponseBody) SetDetails(v string) *CreateTemplateResponseBody {
	s.Details = &v
	return s
}

func (s *CreateTemplateResponseBody) SetErrorCode(v string) *CreateTemplateResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *CreateTemplateResponseBody) SetMessage(v string) *CreateTemplateResponseBody {
	s.Message = &v
	return s
}

func (s *CreateTemplateResponseBody) SetRequestId(v string) *CreateTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateTemplateResponseBody) SetSuccess(v bool) *CreateTemplateResponseBody {
	s.Success = &v
	return s
}

func (s *CreateTemplateResponseBody) SetTemplateId(v string) *CreateTemplateResponseBody {
	s.TemplateId = &v
	return s
}

func (s *CreateTemplateResponseBody) Validate() error {
	return dara.Validate(s)
}
