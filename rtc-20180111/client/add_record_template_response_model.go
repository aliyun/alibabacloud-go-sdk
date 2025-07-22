// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddRecordTemplateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddRecordTemplateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddRecordTemplateResponse
	GetStatusCode() *int32
	SetBody(v *AddRecordTemplateResponseBody) *AddRecordTemplateResponse
	GetBody() *AddRecordTemplateResponseBody
}

type AddRecordTemplateResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddRecordTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddRecordTemplateResponse) String() string {
	return dara.Prettify(s)
}

func (s AddRecordTemplateResponse) GoString() string {
	return s.String()
}

func (s *AddRecordTemplateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddRecordTemplateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddRecordTemplateResponse) GetBody() *AddRecordTemplateResponseBody {
	return s.Body
}

func (s *AddRecordTemplateResponse) SetHeaders(v map[string]*string) *AddRecordTemplateResponse {
	s.Headers = v
	return s
}

func (s *AddRecordTemplateResponse) SetStatusCode(v int32) *AddRecordTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *AddRecordTemplateResponse) SetBody(v *AddRecordTemplateResponseBody) *AddRecordTemplateResponse {
	s.Body = v
	return s
}

func (s *AddRecordTemplateResponse) Validate() error {
	return dara.Validate(s)
}
