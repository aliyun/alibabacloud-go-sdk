// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetServiceLastAccessedDetailsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetServiceLastAccessedDetailsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetServiceLastAccessedDetailsResponse
	GetStatusCode() *int32
	SetBody(v *GetServiceLastAccessedDetailsResponseBody) *GetServiceLastAccessedDetailsResponse
	GetBody() *GetServiceLastAccessedDetailsResponseBody
}

type GetServiceLastAccessedDetailsResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetServiceLastAccessedDetailsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetServiceLastAccessedDetailsResponse) String() string {
	return dara.Prettify(s)
}

func (s GetServiceLastAccessedDetailsResponse) GoString() string {
	return s.String()
}

func (s *GetServiceLastAccessedDetailsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetServiceLastAccessedDetailsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetServiceLastAccessedDetailsResponse) GetBody() *GetServiceLastAccessedDetailsResponseBody {
	return s.Body
}

func (s *GetServiceLastAccessedDetailsResponse) SetHeaders(v map[string]*string) *GetServiceLastAccessedDetailsResponse {
	s.Headers = v
	return s
}

func (s *GetServiceLastAccessedDetailsResponse) SetStatusCode(v int32) *GetServiceLastAccessedDetailsResponse {
	s.StatusCode = &v
	return s
}

func (s *GetServiceLastAccessedDetailsResponse) SetBody(v *GetServiceLastAccessedDetailsResponseBody) *GetServiceLastAccessedDetailsResponse {
	s.Body = v
	return s
}

func (s *GetServiceLastAccessedDetailsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
