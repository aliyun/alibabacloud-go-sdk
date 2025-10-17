// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSparkTemplateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateSparkTemplateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateSparkTemplateResponse
	GetStatusCode() *int32
	SetBody(v *CreateSparkTemplateResponseBody) *CreateSparkTemplateResponse
	GetBody() *CreateSparkTemplateResponseBody
}

type CreateSparkTemplateResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateSparkTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateSparkTemplateResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateSparkTemplateResponse) GoString() string {
	return s.String()
}

func (s *CreateSparkTemplateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateSparkTemplateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateSparkTemplateResponse) GetBody() *CreateSparkTemplateResponseBody {
	return s.Body
}

func (s *CreateSparkTemplateResponse) SetHeaders(v map[string]*string) *CreateSparkTemplateResponse {
	s.Headers = v
	return s
}

func (s *CreateSparkTemplateResponse) SetStatusCode(v int32) *CreateSparkTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateSparkTemplateResponse) SetBody(v *CreateSparkTemplateResponseBody) *CreateSparkTemplateResponse {
	s.Body = v
	return s
}

func (s *CreateSparkTemplateResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
