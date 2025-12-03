// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAccessGroupsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListAccessGroupsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListAccessGroupsResponse
	GetStatusCode() *int32
	SetBody(v *ListAccessGroupsResponseBody) *ListAccessGroupsResponse
	GetBody() *ListAccessGroupsResponseBody
}

type ListAccessGroupsResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAccessGroupsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAccessGroupsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListAccessGroupsResponse) GoString() string {
	return s.String()
}

func (s *ListAccessGroupsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListAccessGroupsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListAccessGroupsResponse) GetBody() *ListAccessGroupsResponseBody {
	return s.Body
}

func (s *ListAccessGroupsResponse) SetHeaders(v map[string]*string) *ListAccessGroupsResponse {
	s.Headers = v
	return s
}

func (s *ListAccessGroupsResponse) SetStatusCode(v int32) *ListAccessGroupsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAccessGroupsResponse) SetBody(v *ListAccessGroupsResponseBody) *ListAccessGroupsResponse {
	s.Body = v
	return s
}

func (s *ListAccessGroupsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
