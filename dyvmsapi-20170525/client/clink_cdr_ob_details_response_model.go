// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iClinkCdrObDetailsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ClinkCdrObDetailsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ClinkCdrObDetailsResponse
	GetStatusCode() *int32
	SetBody(v *ClinkCdrObDetailsResponseBody) *ClinkCdrObDetailsResponse
	GetBody() *ClinkCdrObDetailsResponseBody
}

type ClinkCdrObDetailsResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ClinkCdrObDetailsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ClinkCdrObDetailsResponse) String() string {
	return dara.Prettify(s)
}

func (s ClinkCdrObDetailsResponse) GoString() string {
	return s.String()
}

func (s *ClinkCdrObDetailsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ClinkCdrObDetailsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ClinkCdrObDetailsResponse) GetBody() *ClinkCdrObDetailsResponseBody {
	return s.Body
}

func (s *ClinkCdrObDetailsResponse) SetHeaders(v map[string]*string) *ClinkCdrObDetailsResponse {
	s.Headers = v
	return s
}

func (s *ClinkCdrObDetailsResponse) SetStatusCode(v int32) *ClinkCdrObDetailsResponse {
	s.StatusCode = &v
	return s
}

func (s *ClinkCdrObDetailsResponse) SetBody(v *ClinkCdrObDetailsResponseBody) *ClinkCdrObDetailsResponse {
	s.Body = v
	return s
}

func (s *ClinkCdrObDetailsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
