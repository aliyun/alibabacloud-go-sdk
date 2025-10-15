// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateApplicationTokenResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateApplicationTokenResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateApplicationTokenResponse
	GetStatusCode() *int32
	SetBody(v *CreateApplicationTokenResponseBody) *CreateApplicationTokenResponse
	GetBody() *CreateApplicationTokenResponseBody
}

type CreateApplicationTokenResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateApplicationTokenResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateApplicationTokenResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateApplicationTokenResponse) GoString() string {
	return s.String()
}

func (s *CreateApplicationTokenResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateApplicationTokenResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateApplicationTokenResponse) GetBody() *CreateApplicationTokenResponseBody {
	return s.Body
}

func (s *CreateApplicationTokenResponse) SetHeaders(v map[string]*string) *CreateApplicationTokenResponse {
	s.Headers = v
	return s
}

func (s *CreateApplicationTokenResponse) SetStatusCode(v int32) *CreateApplicationTokenResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateApplicationTokenResponse) SetBody(v *CreateApplicationTokenResponseBody) *CreateApplicationTokenResponse {
	s.Body = v
	return s
}

func (s *CreateApplicationTokenResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
