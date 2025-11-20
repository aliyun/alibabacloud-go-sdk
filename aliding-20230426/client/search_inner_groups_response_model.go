// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSearchInnerGroupsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SearchInnerGroupsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SearchInnerGroupsResponse
	GetStatusCode() *int32
	SetBody(v *SearchInnerGroupsResponseBody) *SearchInnerGroupsResponse
	GetBody() *SearchInnerGroupsResponseBody
}

type SearchInnerGroupsResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SearchInnerGroupsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SearchInnerGroupsResponse) String() string {
	return dara.Prettify(s)
}

func (s SearchInnerGroupsResponse) GoString() string {
	return s.String()
}

func (s *SearchInnerGroupsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SearchInnerGroupsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SearchInnerGroupsResponse) GetBody() *SearchInnerGroupsResponseBody {
	return s.Body
}

func (s *SearchInnerGroupsResponse) SetHeaders(v map[string]*string) *SearchInnerGroupsResponse {
	s.Headers = v
	return s
}

func (s *SearchInnerGroupsResponse) SetStatusCode(v int32) *SearchInnerGroupsResponse {
	s.StatusCode = &v
	return s
}

func (s *SearchInnerGroupsResponse) SetBody(v *SearchInnerGroupsResponseBody) *SearchInnerGroupsResponse {
	s.Body = v
	return s
}

func (s *SearchInnerGroupsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
