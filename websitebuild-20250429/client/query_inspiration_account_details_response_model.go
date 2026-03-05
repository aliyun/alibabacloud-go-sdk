// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryInspirationAccountDetailsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryInspirationAccountDetailsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryInspirationAccountDetailsResponse
	GetStatusCode() *int32
	SetBody(v *QueryInspirationAccountDetailsResponseBody) *QueryInspirationAccountDetailsResponse
	GetBody() *QueryInspirationAccountDetailsResponseBody
}

type QueryInspirationAccountDetailsResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryInspirationAccountDetailsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryInspirationAccountDetailsResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryInspirationAccountDetailsResponse) GoString() string {
	return s.String()
}

func (s *QueryInspirationAccountDetailsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryInspirationAccountDetailsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryInspirationAccountDetailsResponse) GetBody() *QueryInspirationAccountDetailsResponseBody {
	return s.Body
}

func (s *QueryInspirationAccountDetailsResponse) SetHeaders(v map[string]*string) *QueryInspirationAccountDetailsResponse {
	s.Headers = v
	return s
}

func (s *QueryInspirationAccountDetailsResponse) SetStatusCode(v int32) *QueryInspirationAccountDetailsResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryInspirationAccountDetailsResponse) SetBody(v *QueryInspirationAccountDetailsResponseBody) *QueryInspirationAccountDetailsResponse {
	s.Body = v
	return s
}

func (s *QueryInspirationAccountDetailsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
