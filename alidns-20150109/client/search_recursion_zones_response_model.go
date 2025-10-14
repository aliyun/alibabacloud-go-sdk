// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSearchRecursionZonesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SearchRecursionZonesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SearchRecursionZonesResponse
	GetStatusCode() *int32
	SetBody(v *SearchRecursionZonesResponseBody) *SearchRecursionZonesResponse
	GetBody() *SearchRecursionZonesResponseBody
}

type SearchRecursionZonesResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SearchRecursionZonesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SearchRecursionZonesResponse) String() string {
	return dara.Prettify(s)
}

func (s SearchRecursionZonesResponse) GoString() string {
	return s.String()
}

func (s *SearchRecursionZonesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SearchRecursionZonesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SearchRecursionZonesResponse) GetBody() *SearchRecursionZonesResponseBody {
	return s.Body
}

func (s *SearchRecursionZonesResponse) SetHeaders(v map[string]*string) *SearchRecursionZonesResponse {
	s.Headers = v
	return s
}

func (s *SearchRecursionZonesResponse) SetStatusCode(v int32) *SearchRecursionZonesResponse {
	s.StatusCode = &v
	return s
}

func (s *SearchRecursionZonesResponse) SetBody(v *SearchRecursionZonesResponseBody) *SearchRecursionZonesResponse {
	s.Body = v
	return s
}

func (s *SearchRecursionZonesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
