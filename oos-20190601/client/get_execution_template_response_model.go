// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetExecutionTemplateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetExecutionTemplateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetExecutionTemplateResponse
	GetStatusCode() *int32
	SetBody(v *GetExecutionTemplateResponseBody) *GetExecutionTemplateResponse
	GetBody() *GetExecutionTemplateResponseBody
}

type GetExecutionTemplateResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetExecutionTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetExecutionTemplateResponse) String() string {
	return dara.Prettify(s)
}

func (s GetExecutionTemplateResponse) GoString() string {
	return s.String()
}

func (s *GetExecutionTemplateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetExecutionTemplateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetExecutionTemplateResponse) GetBody() *GetExecutionTemplateResponseBody {
	return s.Body
}

func (s *GetExecutionTemplateResponse) SetHeaders(v map[string]*string) *GetExecutionTemplateResponse {
	s.Headers = v
	return s
}

func (s *GetExecutionTemplateResponse) SetStatusCode(v int32) *GetExecutionTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *GetExecutionTemplateResponse) SetBody(v *GetExecutionTemplateResponseBody) *GetExecutionTemplateResponse {
	s.Body = v
	return s
}

func (s *GetExecutionTemplateResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
