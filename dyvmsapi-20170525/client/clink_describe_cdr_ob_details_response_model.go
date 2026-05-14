// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iClinkDescribeCdrObDetailsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ClinkDescribeCdrObDetailsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ClinkDescribeCdrObDetailsResponse
	GetStatusCode() *int32
	SetBody(v *ClinkDescribeCdrObDetailsResponseBody) *ClinkDescribeCdrObDetailsResponse
	GetBody() *ClinkDescribeCdrObDetailsResponseBody
}

type ClinkDescribeCdrObDetailsResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ClinkDescribeCdrObDetailsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ClinkDescribeCdrObDetailsResponse) String() string {
	return dara.Prettify(s)
}

func (s ClinkDescribeCdrObDetailsResponse) GoString() string {
	return s.String()
}

func (s *ClinkDescribeCdrObDetailsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ClinkDescribeCdrObDetailsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ClinkDescribeCdrObDetailsResponse) GetBody() *ClinkDescribeCdrObDetailsResponseBody {
	return s.Body
}

func (s *ClinkDescribeCdrObDetailsResponse) SetHeaders(v map[string]*string) *ClinkDescribeCdrObDetailsResponse {
	s.Headers = v
	return s
}

func (s *ClinkDescribeCdrObDetailsResponse) SetStatusCode(v int32) *ClinkDescribeCdrObDetailsResponse {
	s.StatusCode = &v
	return s
}

func (s *ClinkDescribeCdrObDetailsResponse) SetBody(v *ClinkDescribeCdrObDetailsResponseBody) *ClinkDescribeCdrObDetailsResponse {
	s.Body = v
	return s
}

func (s *ClinkDescribeCdrObDetailsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
