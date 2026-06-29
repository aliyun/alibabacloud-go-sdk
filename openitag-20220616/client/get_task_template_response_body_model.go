// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTaskTemplateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *GetTaskTemplateResponseBody
	GetCode() *int32
	SetDetails(v string) *GetTaskTemplateResponseBody
	GetDetails() *string
	SetErrorCode(v string) *GetTaskTemplateResponseBody
	GetErrorCode() *string
	SetMessage(v string) *GetTaskTemplateResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetTaskTemplateResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetTaskTemplateResponseBody
	GetSuccess() *bool
	SetTemplate(v *TemplateDetail) *GetTaskTemplateResponseBody
	GetTemplate() *TemplateDetail
}

type GetTaskTemplateResponseBody struct {
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
	// ""
	Details *string `json:"Details,omitempty" xml:"Details,omitempty"`
	// Error code.
	//
	// example:
	//
	// ""
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The response message of the request.
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
	// Indicates whether the operation succeeded.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// Template details.
	Template *TemplateDetail `json:"Template,omitempty" xml:"Template,omitempty"`
}

func (s GetTaskTemplateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetTaskTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *GetTaskTemplateResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *GetTaskTemplateResponseBody) GetDetails() *string {
	return s.Details
}

func (s *GetTaskTemplateResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *GetTaskTemplateResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetTaskTemplateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetTaskTemplateResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetTaskTemplateResponseBody) GetTemplate() *TemplateDetail {
	return s.Template
}

func (s *GetTaskTemplateResponseBody) SetCode(v int32) *GetTaskTemplateResponseBody {
	s.Code = &v
	return s
}

func (s *GetTaskTemplateResponseBody) SetDetails(v string) *GetTaskTemplateResponseBody {
	s.Details = &v
	return s
}

func (s *GetTaskTemplateResponseBody) SetErrorCode(v string) *GetTaskTemplateResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetTaskTemplateResponseBody) SetMessage(v string) *GetTaskTemplateResponseBody {
	s.Message = &v
	return s
}

func (s *GetTaskTemplateResponseBody) SetRequestId(v string) *GetTaskTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetTaskTemplateResponseBody) SetSuccess(v bool) *GetTaskTemplateResponseBody {
	s.Success = &v
	return s
}

func (s *GetTaskTemplateResponseBody) SetTemplate(v *TemplateDetail) *GetTaskTemplateResponseBody {
	s.Template = v
	return s
}

func (s *GetTaskTemplateResponseBody) Validate() error {
	if s.Template != nil {
		if err := s.Template.Validate(); err != nil {
			return err
		}
	}
	return nil
}
