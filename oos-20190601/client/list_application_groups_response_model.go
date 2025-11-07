// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListApplicationGroupsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListApplicationGroupsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListApplicationGroupsResponse
	GetStatusCode() *int32
	SetBody(v *ListApplicationGroupsResponseBody) *ListApplicationGroupsResponse
	GetBody() *ListApplicationGroupsResponseBody
}

type ListApplicationGroupsResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListApplicationGroupsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListApplicationGroupsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListApplicationGroupsResponse) GoString() string {
	return s.String()
}

func (s *ListApplicationGroupsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListApplicationGroupsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListApplicationGroupsResponse) GetBody() *ListApplicationGroupsResponseBody {
	return s.Body
}

func (s *ListApplicationGroupsResponse) SetHeaders(v map[string]*string) *ListApplicationGroupsResponse {
	s.Headers = v
	return s
}

func (s *ListApplicationGroupsResponse) SetStatusCode(v int32) *ListApplicationGroupsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListApplicationGroupsResponse) SetBody(v *ListApplicationGroupsResponseBody) *ListApplicationGroupsResponse {
	s.Body = v
	return s
}

func (s *ListApplicationGroupsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
