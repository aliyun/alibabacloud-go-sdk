// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListResourcesByAdvancedSearchResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListResourcesByAdvancedSearchResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListResourcesByAdvancedSearchResponse
	GetStatusCode() *int32
	SetBody(v *ListResourcesByAdvancedSearchResponseBody) *ListResourcesByAdvancedSearchResponse
	GetBody() *ListResourcesByAdvancedSearchResponseBody
}

type ListResourcesByAdvancedSearchResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListResourcesByAdvancedSearchResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListResourcesByAdvancedSearchResponse) String() string {
	return dara.Prettify(s)
}

func (s ListResourcesByAdvancedSearchResponse) GoString() string {
	return s.String()
}

func (s *ListResourcesByAdvancedSearchResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListResourcesByAdvancedSearchResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListResourcesByAdvancedSearchResponse) GetBody() *ListResourcesByAdvancedSearchResponseBody {
	return s.Body
}

func (s *ListResourcesByAdvancedSearchResponse) SetHeaders(v map[string]*string) *ListResourcesByAdvancedSearchResponse {
	s.Headers = v
	return s
}

func (s *ListResourcesByAdvancedSearchResponse) SetStatusCode(v int32) *ListResourcesByAdvancedSearchResponse {
	s.StatusCode = &v
	return s
}

func (s *ListResourcesByAdvancedSearchResponse) SetBody(v *ListResourcesByAdvancedSearchResponseBody) *ListResourcesByAdvancedSearchResponse {
	s.Body = v
	return s
}

func (s *ListResourcesByAdvancedSearchResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
