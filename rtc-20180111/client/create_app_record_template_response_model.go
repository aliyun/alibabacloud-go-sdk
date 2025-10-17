// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAppRecordTemplateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateAppRecordTemplateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateAppRecordTemplateResponse
	GetStatusCode() *int32
	SetBody(v *CreateAppRecordTemplateResponseBody) *CreateAppRecordTemplateResponse
	GetBody() *CreateAppRecordTemplateResponseBody
}

type CreateAppRecordTemplateResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateAppRecordTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateAppRecordTemplateResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateAppRecordTemplateResponse) GoString() string {
	return s.String()
}

func (s *CreateAppRecordTemplateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateAppRecordTemplateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateAppRecordTemplateResponse) GetBody() *CreateAppRecordTemplateResponseBody {
	return s.Body
}

func (s *CreateAppRecordTemplateResponse) SetHeaders(v map[string]*string) *CreateAppRecordTemplateResponse {
	s.Headers = v
	return s
}

func (s *CreateAppRecordTemplateResponse) SetStatusCode(v int32) *CreateAppRecordTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateAppRecordTemplateResponse) SetBody(v *CreateAppRecordTemplateResponseBody) *CreateAppRecordTemplateResponse {
	s.Body = v
	return s
}

func (s *CreateAppRecordTemplateResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
