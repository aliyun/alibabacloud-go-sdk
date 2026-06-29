// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteTemplateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *DeleteTemplateResponseBody
	GetCode() *int32
	SetDetails(v string) *DeleteTemplateResponseBody
	GetDetails() *string
	SetErrorCode(v string) *DeleteTemplateResponseBody
	GetErrorCode() *string
	SetMessage(v string) *DeleteTemplateResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteTemplateResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteTemplateResponseBody
	GetSuccess() *bool
	SetTemplateId(v string) *DeleteTemplateResponseBody
	GetTemplateId() *string
}

type DeleteTemplateResponseBody struct {
	// Total amount of data under the conditions of this request. This parameter is optional and does not need to be returned by default.
	//
	// example:
	//
	// 0
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// Details
	//
	// example:
	//
	// ""
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
	// 152***348342353920
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s DeleteTemplateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteTemplateResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *DeleteTemplateResponseBody) GetDetails() *string {
	return s.Details
}

func (s *DeleteTemplateResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *DeleteTemplateResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteTemplateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteTemplateResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteTemplateResponseBody) GetTemplateId() *string {
	return s.TemplateId
}

func (s *DeleteTemplateResponseBody) SetCode(v int32) *DeleteTemplateResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteTemplateResponseBody) SetDetails(v string) *DeleteTemplateResponseBody {
	s.Details = &v
	return s
}

func (s *DeleteTemplateResponseBody) SetErrorCode(v string) *DeleteTemplateResponseBody {
	s.ErrorCode = &v
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

func (s *DeleteTemplateResponseBody) SetSuccess(v bool) *DeleteTemplateResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteTemplateResponseBody) SetTemplateId(v string) *DeleteTemplateResponseBody {
	s.TemplateId = &v
	return s
}

func (s *DeleteTemplateResponseBody) Validate() error {
	return dara.Validate(s)
}
