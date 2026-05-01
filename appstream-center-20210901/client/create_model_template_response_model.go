// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateModelTemplateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateModelTemplateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateModelTemplateResponse
	GetStatusCode() *int32
	SetBody(v *CreateModelTemplateResponseBody) *CreateModelTemplateResponse
	GetBody() *CreateModelTemplateResponseBody
}

type CreateModelTemplateResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateModelTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateModelTemplateResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateModelTemplateResponse) GoString() string {
	return s.String()
}

func (s *CreateModelTemplateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateModelTemplateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateModelTemplateResponse) GetBody() *CreateModelTemplateResponseBody {
	return s.Body
}

func (s *CreateModelTemplateResponse) SetHeaders(v map[string]*string) *CreateModelTemplateResponse {
	s.Headers = v
	return s
}

func (s *CreateModelTemplateResponse) SetStatusCode(v int32) *CreateModelTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateModelTemplateResponse) SetBody(v *CreateModelTemplateResponseBody) *CreateModelTemplateResponse {
	s.Body = v
	return s
}

func (s *CreateModelTemplateResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
