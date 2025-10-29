// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateIndexTemplateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateIndexTemplateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateIndexTemplateResponse
	GetStatusCode() *int32
	SetBody(v *CreateIndexTemplateResponseBody) *CreateIndexTemplateResponse
	GetBody() *CreateIndexTemplateResponseBody
}

type CreateIndexTemplateResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateIndexTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateIndexTemplateResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateIndexTemplateResponse) GoString() string {
	return s.String()
}

func (s *CreateIndexTemplateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateIndexTemplateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateIndexTemplateResponse) GetBody() *CreateIndexTemplateResponseBody {
	return s.Body
}

func (s *CreateIndexTemplateResponse) SetHeaders(v map[string]*string) *CreateIndexTemplateResponse {
	s.Headers = v
	return s
}

func (s *CreateIndexTemplateResponse) SetStatusCode(v int32) *CreateIndexTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateIndexTemplateResponse) SetBody(v *CreateIndexTemplateResponseBody) *CreateIndexTemplateResponse {
	s.Body = v
	return s
}

func (s *CreateIndexTemplateResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
