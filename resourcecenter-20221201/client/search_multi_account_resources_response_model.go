// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSearchMultiAccountResourcesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SearchMultiAccountResourcesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SearchMultiAccountResourcesResponse
	GetStatusCode() *int32
	SetBody(v *SearchMultiAccountResourcesResponseBody) *SearchMultiAccountResourcesResponse
	GetBody() *SearchMultiAccountResourcesResponseBody
}

type SearchMultiAccountResourcesResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SearchMultiAccountResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SearchMultiAccountResourcesResponse) String() string {
	return dara.Prettify(s)
}

func (s SearchMultiAccountResourcesResponse) GoString() string {
	return s.String()
}

func (s *SearchMultiAccountResourcesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SearchMultiAccountResourcesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SearchMultiAccountResourcesResponse) GetBody() *SearchMultiAccountResourcesResponseBody {
	return s.Body
}

func (s *SearchMultiAccountResourcesResponse) SetHeaders(v map[string]*string) *SearchMultiAccountResourcesResponse {
	s.Headers = v
	return s
}

func (s *SearchMultiAccountResourcesResponse) SetStatusCode(v int32) *SearchMultiAccountResourcesResponse {
	s.StatusCode = &v
	return s
}

func (s *SearchMultiAccountResourcesResponse) SetBody(v *SearchMultiAccountResourcesResponseBody) *SearchMultiAccountResourcesResponse {
	s.Body = v
	return s
}

func (s *SearchMultiAccountResourcesResponse) Validate() error {
	return dara.Validate(s)
}
