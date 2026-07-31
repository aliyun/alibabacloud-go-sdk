// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSearchSemanticViewsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SearchSemanticViewsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SearchSemanticViewsResponse
	GetStatusCode() *int32
	SetBody(v *SearchSemanticViewsResponseBody) *SearchSemanticViewsResponse
	GetBody() *SearchSemanticViewsResponseBody
}

type SearchSemanticViewsResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SearchSemanticViewsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SearchSemanticViewsResponse) String() string {
	return dara.Prettify(s)
}

func (s SearchSemanticViewsResponse) GoString() string {
	return s.String()
}

func (s *SearchSemanticViewsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SearchSemanticViewsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SearchSemanticViewsResponse) GetBody() *SearchSemanticViewsResponseBody {
	return s.Body
}

func (s *SearchSemanticViewsResponse) SetHeaders(v map[string]*string) *SearchSemanticViewsResponse {
	s.Headers = v
	return s
}

func (s *SearchSemanticViewsResponse) SetStatusCode(v int32) *SearchSemanticViewsResponse {
	s.StatusCode = &v
	return s
}

func (s *SearchSemanticViewsResponse) SetBody(v *SearchSemanticViewsResponseBody) *SearchSemanticViewsResponse {
	s.Body = v
	return s
}

func (s *SearchSemanticViewsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
