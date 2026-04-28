// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSearchAddressGroupsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SearchAddressGroupsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SearchAddressGroupsResponse
	GetStatusCode() *int32
	SetBody(v *SearchAddressGroupsResponseBody) *SearchAddressGroupsResponse
	GetBody() *SearchAddressGroupsResponseBody
}

type SearchAddressGroupsResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SearchAddressGroupsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SearchAddressGroupsResponse) String() string {
	return dara.Prettify(s)
}

func (s SearchAddressGroupsResponse) GoString() string {
	return s.String()
}

func (s *SearchAddressGroupsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SearchAddressGroupsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SearchAddressGroupsResponse) GetBody() *SearchAddressGroupsResponseBody {
	return s.Body
}

func (s *SearchAddressGroupsResponse) SetHeaders(v map[string]*string) *SearchAddressGroupsResponse {
	s.Headers = v
	return s
}

func (s *SearchAddressGroupsResponse) SetStatusCode(v int32) *SearchAddressGroupsResponse {
	s.StatusCode = &v
	return s
}

func (s *SearchAddressGroupsResponse) SetBody(v *SearchAddressGroupsResponseBody) *SearchAddressGroupsResponse {
	s.Body = v
	return s
}

func (s *SearchAddressGroupsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
