// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQuerySendDetailsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QuerySendDetailsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QuerySendDetailsResponse
	GetStatusCode() *int32
	SetBody(v *QuerySendDetailsResponseBody) *QuerySendDetailsResponse
	GetBody() *QuerySendDetailsResponseBody
}

type QuerySendDetailsResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QuerySendDetailsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QuerySendDetailsResponse) String() string {
	return dara.Prettify(s)
}

func (s QuerySendDetailsResponse) GoString() string {
	return s.String()
}

func (s *QuerySendDetailsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QuerySendDetailsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QuerySendDetailsResponse) GetBody() *QuerySendDetailsResponseBody {
	return s.Body
}

func (s *QuerySendDetailsResponse) SetHeaders(v map[string]*string) *QuerySendDetailsResponse {
	s.Headers = v
	return s
}

func (s *QuerySendDetailsResponse) SetStatusCode(v int32) *QuerySendDetailsResponse {
	s.StatusCode = &v
	return s
}

func (s *QuerySendDetailsResponse) SetBody(v *QuerySendDetailsResponseBody) *QuerySendDetailsResponse {
	s.Body = v
	return s
}

func (s *QuerySendDetailsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
