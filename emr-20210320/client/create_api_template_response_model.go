// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateApiTemplateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateApiTemplateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateApiTemplateResponse
	GetStatusCode() *int32
	SetBody(v *CreateApiTemplateResponseBody) *CreateApiTemplateResponse
	GetBody() *CreateApiTemplateResponseBody
}

type CreateApiTemplateResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateApiTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateApiTemplateResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateApiTemplateResponse) GoString() string {
	return s.String()
}

func (s *CreateApiTemplateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateApiTemplateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateApiTemplateResponse) GetBody() *CreateApiTemplateResponseBody {
	return s.Body
}

func (s *CreateApiTemplateResponse) SetHeaders(v map[string]*string) *CreateApiTemplateResponse {
	s.Headers = v
	return s
}

func (s *CreateApiTemplateResponse) SetStatusCode(v int32) *CreateApiTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateApiTemplateResponse) SetBody(v *CreateApiTemplateResponseBody) *CreateApiTemplateResponse {
	s.Body = v
	return s
}

func (s *CreateApiTemplateResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
