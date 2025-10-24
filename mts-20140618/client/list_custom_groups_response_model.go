// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCustomGroupsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListCustomGroupsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListCustomGroupsResponse
	GetStatusCode() *int32
	SetBody(v *ListCustomGroupsResponseBody) *ListCustomGroupsResponse
	GetBody() *ListCustomGroupsResponseBody
}

type ListCustomGroupsResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListCustomGroupsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListCustomGroupsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListCustomGroupsResponse) GoString() string {
	return s.String()
}

func (s *ListCustomGroupsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListCustomGroupsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListCustomGroupsResponse) GetBody() *ListCustomGroupsResponseBody {
	return s.Body
}

func (s *ListCustomGroupsResponse) SetHeaders(v map[string]*string) *ListCustomGroupsResponse {
	s.Headers = v
	return s
}

func (s *ListCustomGroupsResponse) SetStatusCode(v int32) *ListCustomGroupsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListCustomGroupsResponse) SetBody(v *ListCustomGroupsResponseBody) *ListCustomGroupsResponse {
	s.Body = v
	return s
}

func (s *ListCustomGroupsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
