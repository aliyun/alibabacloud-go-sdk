// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDpiGroupsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListDpiGroupsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListDpiGroupsResponse
	GetStatusCode() *int32
	SetBody(v *ListDpiGroupsResponseBody) *ListDpiGroupsResponse
	GetBody() *ListDpiGroupsResponseBody
}

type ListDpiGroupsResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListDpiGroupsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListDpiGroupsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListDpiGroupsResponse) GoString() string {
	return s.String()
}

func (s *ListDpiGroupsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListDpiGroupsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListDpiGroupsResponse) GetBody() *ListDpiGroupsResponseBody {
	return s.Body
}

func (s *ListDpiGroupsResponse) SetHeaders(v map[string]*string) *ListDpiGroupsResponse {
	s.Headers = v
	return s
}

func (s *ListDpiGroupsResponse) SetStatusCode(v int32) *ListDpiGroupsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListDpiGroupsResponse) SetBody(v *ListDpiGroupsResponseBody) *ListDpiGroupsResponse {
	s.Body = v
	return s
}

func (s *ListDpiGroupsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
