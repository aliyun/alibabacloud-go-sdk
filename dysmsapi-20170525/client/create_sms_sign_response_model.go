// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSmsSignResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateSmsSignResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateSmsSignResponse
	GetStatusCode() *int32
	SetBody(v *CreateSmsSignResponseBody) *CreateSmsSignResponse
	GetBody() *CreateSmsSignResponseBody
}

type CreateSmsSignResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateSmsSignResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateSmsSignResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateSmsSignResponse) GoString() string {
	return s.String()
}

func (s *CreateSmsSignResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateSmsSignResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateSmsSignResponse) GetBody() *CreateSmsSignResponseBody {
	return s.Body
}

func (s *CreateSmsSignResponse) SetHeaders(v map[string]*string) *CreateSmsSignResponse {
	s.Headers = v
	return s
}

func (s *CreateSmsSignResponse) SetStatusCode(v int32) *CreateSmsSignResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateSmsSignResponse) SetBody(v *CreateSmsSignResponseBody) *CreateSmsSignResponse {
	s.Body = v
	return s
}

func (s *CreateSmsSignResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
