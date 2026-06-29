// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTaskTemplateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetTaskTemplateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetTaskTemplateResponse
	GetStatusCode() *int32
	SetBody(v *GetTaskTemplateResponseBody) *GetTaskTemplateResponse
	GetBody() *GetTaskTemplateResponseBody
}

type GetTaskTemplateResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetTaskTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetTaskTemplateResponse) String() string {
	return dara.Prettify(s)
}

func (s GetTaskTemplateResponse) GoString() string {
	return s.String()
}

func (s *GetTaskTemplateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetTaskTemplateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetTaskTemplateResponse) GetBody() *GetTaskTemplateResponseBody {
	return s.Body
}

func (s *GetTaskTemplateResponse) SetHeaders(v map[string]*string) *GetTaskTemplateResponse {
	s.Headers = v
	return s
}

func (s *GetTaskTemplateResponse) SetStatusCode(v int32) *GetTaskTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *GetTaskTemplateResponse) SetBody(v *GetTaskTemplateResponseBody) *GetTaskTemplateResponse {
	s.Body = v
	return s
}

func (s *GetTaskTemplateResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
