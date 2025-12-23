// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSearchResourcesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SearchResourcesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SearchResourcesResponse
	GetStatusCode() *int32
	SetBody(v *SearchResourcesResponseBody) *SearchResourcesResponse
	GetBody() *SearchResourcesResponseBody
}

type SearchResourcesResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SearchResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SearchResourcesResponse) String() string {
	return dara.Prettify(s)
}

func (s SearchResourcesResponse) GoString() string {
	return s.String()
}

func (s *SearchResourcesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SearchResourcesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SearchResourcesResponse) GetBody() *SearchResourcesResponseBody {
	return s.Body
}

func (s *SearchResourcesResponse) SetHeaders(v map[string]*string) *SearchResourcesResponse {
	s.Headers = v
	return s
}

func (s *SearchResourcesResponse) SetStatusCode(v int32) *SearchResourcesResponse {
	s.StatusCode = &v
	return s
}

func (s *SearchResourcesResponse) SetBody(v *SearchResourcesResponseBody) *SearchResourcesResponse {
	s.Body = v
	return s
}

func (s *SearchResourcesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
