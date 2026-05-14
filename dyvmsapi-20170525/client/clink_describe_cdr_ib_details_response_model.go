// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iClinkDescribeCdrIbDetailsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ClinkDescribeCdrIbDetailsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ClinkDescribeCdrIbDetailsResponse
	GetStatusCode() *int32
	SetBody(v *ClinkDescribeCdrIbDetailsResponseBody) *ClinkDescribeCdrIbDetailsResponse
	GetBody() *ClinkDescribeCdrIbDetailsResponseBody
}

type ClinkDescribeCdrIbDetailsResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ClinkDescribeCdrIbDetailsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ClinkDescribeCdrIbDetailsResponse) String() string {
	return dara.Prettify(s)
}

func (s ClinkDescribeCdrIbDetailsResponse) GoString() string {
	return s.String()
}

func (s *ClinkDescribeCdrIbDetailsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ClinkDescribeCdrIbDetailsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ClinkDescribeCdrIbDetailsResponse) GetBody() *ClinkDescribeCdrIbDetailsResponseBody {
	return s.Body
}

func (s *ClinkDescribeCdrIbDetailsResponse) SetHeaders(v map[string]*string) *ClinkDescribeCdrIbDetailsResponse {
	s.Headers = v
	return s
}

func (s *ClinkDescribeCdrIbDetailsResponse) SetStatusCode(v int32) *ClinkDescribeCdrIbDetailsResponse {
	s.StatusCode = &v
	return s
}

func (s *ClinkDescribeCdrIbDetailsResponse) SetBody(v *ClinkDescribeCdrIbDetailsResponseBody) *ClinkDescribeCdrIbDetailsResponse {
	s.Body = v
	return s
}

func (s *ClinkDescribeCdrIbDetailsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
