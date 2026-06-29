// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTemplateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *GetTemplateResponseBody
	GetCode() *int32
	SetDetails(v string) *GetTemplateResponseBody
	GetDetails() *string
	SetErrorCode(v string) *GetTemplateResponseBody
	GetErrorCode() *string
	SetMessage(v string) *GetTemplateResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetTemplateResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetTemplateResponseBody
	GetSuccess() *bool
	SetTemplate(v *TemplateDetail) *GetTemplateResponseBody
	GetTemplate() *TemplateDetail
}

type GetTemplateResponseBody struct {
	// Total amount of data under the conditions of this request. This parameter is optional and does not need to be returned by default.
	//
	// example:
	//
	// 0
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// Detail ID
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
	// template
	Template *TemplateDetail `json:"Template,omitempty" xml:"Template,omitempty"`
}

func (s GetTemplateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *GetTemplateResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *GetTemplateResponseBody) GetDetails() *string {
	return s.Details
}

func (s *GetTemplateResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *GetTemplateResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetTemplateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetTemplateResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetTemplateResponseBody) GetTemplate() *TemplateDetail {
	return s.Template
}

func (s *GetTemplateResponseBody) SetCode(v int32) *GetTemplateResponseBody {
	s.Code = &v
	return s
}

func (s *GetTemplateResponseBody) SetDetails(v string) *GetTemplateResponseBody {
	s.Details = &v
	return s
}

func (s *GetTemplateResponseBody) SetErrorCode(v string) *GetTemplateResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetTemplateResponseBody) SetMessage(v string) *GetTemplateResponseBody {
	s.Message = &v
	return s
}

func (s *GetTemplateResponseBody) SetRequestId(v string) *GetTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetTemplateResponseBody) SetSuccess(v bool) *GetTemplateResponseBody {
	s.Success = &v
	return s
}

func (s *GetTemplateResponseBody) SetTemplate(v *TemplateDetail) *GetTemplateResponseBody {
	s.Template = v
	return s
}

func (s *GetTemplateResponseBody) Validate() error {
	if s.Template != nil {
		if err := s.Template.Validate(); err != nil {
			return err
		}
	}
	return nil
}
