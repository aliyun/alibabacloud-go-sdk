// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGenerateServiceLastAccessedDetailsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GenerateServiceLastAccessedDetailsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GenerateServiceLastAccessedDetailsResponse
	GetStatusCode() *int32
	SetBody(v *GenerateServiceLastAccessedDetailsResponseBody) *GenerateServiceLastAccessedDetailsResponse
	GetBody() *GenerateServiceLastAccessedDetailsResponseBody
}

type GenerateServiceLastAccessedDetailsResponse struct {
	Headers    map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GenerateServiceLastAccessedDetailsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GenerateServiceLastAccessedDetailsResponse) String() string {
	return dara.Prettify(s)
}

func (s GenerateServiceLastAccessedDetailsResponse) GoString() string {
	return s.String()
}

func (s *GenerateServiceLastAccessedDetailsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GenerateServiceLastAccessedDetailsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GenerateServiceLastAccessedDetailsResponse) GetBody() *GenerateServiceLastAccessedDetailsResponseBody {
	return s.Body
}

func (s *GenerateServiceLastAccessedDetailsResponse) SetHeaders(v map[string]*string) *GenerateServiceLastAccessedDetailsResponse {
	s.Headers = v
	return s
}

func (s *GenerateServiceLastAccessedDetailsResponse) SetStatusCode(v int32) *GenerateServiceLastAccessedDetailsResponse {
	s.StatusCode = &v
	return s
}

func (s *GenerateServiceLastAccessedDetailsResponse) SetBody(v *GenerateServiceLastAccessedDetailsResponseBody) *GenerateServiceLastAccessedDetailsResponse {
	s.Body = v
	return s
}

func (s *GenerateServiceLastAccessedDetailsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
