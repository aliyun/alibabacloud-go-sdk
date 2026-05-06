// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateModelProviderTemplateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateModelProviderTemplateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateModelProviderTemplateResponse
	GetStatusCode() *int32
	SetBody(v *CreateModelProviderTemplateResponseBody) *CreateModelProviderTemplateResponse
	GetBody() *CreateModelProviderTemplateResponseBody
}

type CreateModelProviderTemplateResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateModelProviderTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateModelProviderTemplateResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateModelProviderTemplateResponse) GoString() string {
	return s.String()
}

func (s *CreateModelProviderTemplateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateModelProviderTemplateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateModelProviderTemplateResponse) GetBody() *CreateModelProviderTemplateResponseBody {
	return s.Body
}

func (s *CreateModelProviderTemplateResponse) SetHeaders(v map[string]*string) *CreateModelProviderTemplateResponse {
	s.Headers = v
	return s
}

func (s *CreateModelProviderTemplateResponse) SetStatusCode(v int32) *CreateModelProviderTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateModelProviderTemplateResponse) SetBody(v *CreateModelProviderTemplateResponseBody) *CreateModelProviderTemplateResponse {
	s.Body = v
	return s
}

func (s *CreateModelProviderTemplateResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
