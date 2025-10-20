// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTableDetailsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListTableDetailsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListTableDetailsResponse
	GetStatusCode() *int32
	SetBody(v *ListTableDetailsResponseBody) *ListTableDetailsResponse
	GetBody() *ListTableDetailsResponseBody
}

type ListTableDetailsResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListTableDetailsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListTableDetailsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListTableDetailsResponse) GoString() string {
	return s.String()
}

func (s *ListTableDetailsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListTableDetailsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListTableDetailsResponse) GetBody() *ListTableDetailsResponseBody {
	return s.Body
}

func (s *ListTableDetailsResponse) SetHeaders(v map[string]*string) *ListTableDetailsResponse {
	s.Headers = v
	return s
}

func (s *ListTableDetailsResponse) SetStatusCode(v int32) *ListTableDetailsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListTableDetailsResponse) SetBody(v *ListTableDetailsResponseBody) *ListTableDetailsResponse {
	s.Body = v
	return s
}

func (s *ListTableDetailsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
