// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSearchEventsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SearchEventsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SearchEventsResponse
	GetStatusCode() *int32
	SetBody(v *SearchEventsResponseBody) *SearchEventsResponse
	GetBody() *SearchEventsResponseBody
}

type SearchEventsResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SearchEventsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SearchEventsResponse) String() string {
	return dara.Prettify(s)
}

func (s SearchEventsResponse) GoString() string {
	return s.String()
}

func (s *SearchEventsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SearchEventsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SearchEventsResponse) GetBody() *SearchEventsResponseBody {
	return s.Body
}

func (s *SearchEventsResponse) SetHeaders(v map[string]*string) *SearchEventsResponse {
	s.Headers = v
	return s
}

func (s *SearchEventsResponse) SetStatusCode(v int32) *SearchEventsResponse {
	s.StatusCode = &v
	return s
}

func (s *SearchEventsResponse) SetBody(v *SearchEventsResponseBody) *SearchEventsResponse {
	s.Body = v
	return s
}

func (s *SearchEventsResponse) Validate() error {
	return dara.Validate(s)
}
