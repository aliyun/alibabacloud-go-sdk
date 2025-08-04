// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListIndexFileDetailsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListIndexFileDetailsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListIndexFileDetailsResponse
	GetStatusCode() *int32
	SetBody(v *ListIndexFileDetailsResponseBody) *ListIndexFileDetailsResponse
	GetBody() *ListIndexFileDetailsResponseBody
}

type ListIndexFileDetailsResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListIndexFileDetailsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListIndexFileDetailsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListIndexFileDetailsResponse) GoString() string {
	return s.String()
}

func (s *ListIndexFileDetailsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListIndexFileDetailsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListIndexFileDetailsResponse) GetBody() *ListIndexFileDetailsResponseBody {
	return s.Body
}

func (s *ListIndexFileDetailsResponse) SetHeaders(v map[string]*string) *ListIndexFileDetailsResponse {
	s.Headers = v
	return s
}

func (s *ListIndexFileDetailsResponse) SetStatusCode(v int32) *ListIndexFileDetailsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListIndexFileDetailsResponse) SetBody(v *ListIndexFileDetailsResponseBody) *ListIndexFileDetailsResponse {
	s.Body = v
	return s
}

func (s *ListIndexFileDetailsResponse) Validate() error {
	return dara.Validate(s)
}
