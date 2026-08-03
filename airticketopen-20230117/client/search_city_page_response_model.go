// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSearchCityPageResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SearchCityPageResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SearchCityPageResponse
	GetStatusCode() *int32
	SetBody(v *SearchCityPageResponseBody) *SearchCityPageResponse
	GetBody() *SearchCityPageResponseBody
}

type SearchCityPageResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SearchCityPageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SearchCityPageResponse) String() string {
	return dara.Prettify(s)
}

func (s SearchCityPageResponse) GoString() string {
	return s.String()
}

func (s *SearchCityPageResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SearchCityPageResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SearchCityPageResponse) GetBody() *SearchCityPageResponseBody {
	return s.Body
}

func (s *SearchCityPageResponse) SetHeaders(v map[string]*string) *SearchCityPageResponse {
	s.Headers = v
	return s
}

func (s *SearchCityPageResponse) SetStatusCode(v int32) *SearchCityPageResponse {
	s.StatusCode = &v
	return s
}

func (s *SearchCityPageResponse) SetBody(v *SearchCityPageResponseBody) *SearchCityPageResponse {
	s.Body = v
	return s
}

func (s *SearchCityPageResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
