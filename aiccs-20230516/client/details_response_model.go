// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDetailsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DetailsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DetailsResponse
	GetStatusCode() *int32
	SetBody(v *DetailsResponseBody) *DetailsResponse
	GetBody() *DetailsResponseBody
}

type DetailsResponse struct {
	Headers    map[string]*string   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DetailsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DetailsResponse) String() string {
	return dara.Prettify(s)
}

func (s DetailsResponse) GoString() string {
	return s.String()
}

func (s *DetailsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DetailsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DetailsResponse) GetBody() *DetailsResponseBody {
	return s.Body
}

func (s *DetailsResponse) SetHeaders(v map[string]*string) *DetailsResponse {
	s.Headers = v
	return s
}

func (s *DetailsResponse) SetStatusCode(v int32) *DetailsResponse {
	s.StatusCode = &v
	return s
}

func (s *DetailsResponse) SetBody(v *DetailsResponseBody) *DetailsResponse {
	s.Body = v
	return s
}

func (s *DetailsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
