// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateRecordTemplateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateRecordTemplateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateRecordTemplateResponse
	GetStatusCode() *int32
	SetBody(v *UpdateRecordTemplateResponseBody) *UpdateRecordTemplateResponse
	GetBody() *UpdateRecordTemplateResponseBody
}

type UpdateRecordTemplateResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateRecordTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateRecordTemplateResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateRecordTemplateResponse) GoString() string {
	return s.String()
}

func (s *UpdateRecordTemplateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateRecordTemplateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateRecordTemplateResponse) GetBody() *UpdateRecordTemplateResponseBody {
	return s.Body
}

func (s *UpdateRecordTemplateResponse) SetHeaders(v map[string]*string) *UpdateRecordTemplateResponse {
	s.Headers = v
	return s
}

func (s *UpdateRecordTemplateResponse) SetStatusCode(v int32) *UpdateRecordTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateRecordTemplateResponse) SetBody(v *UpdateRecordTemplateResponseBody) *UpdateRecordTemplateResponse {
	s.Body = v
	return s
}

func (s *UpdateRecordTemplateResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
