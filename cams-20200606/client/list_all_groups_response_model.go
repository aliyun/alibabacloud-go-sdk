// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAllGroupsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListAllGroupsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListAllGroupsResponse
	GetStatusCode() *int32
	SetBody(v *ListAllGroupsResponseBody) *ListAllGroupsResponse
	GetBody() *ListAllGroupsResponseBody
}

type ListAllGroupsResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAllGroupsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAllGroupsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListAllGroupsResponse) GoString() string {
	return s.String()
}

func (s *ListAllGroupsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListAllGroupsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListAllGroupsResponse) GetBody() *ListAllGroupsResponseBody {
	return s.Body
}

func (s *ListAllGroupsResponse) SetHeaders(v map[string]*string) *ListAllGroupsResponse {
	s.Headers = v
	return s
}

func (s *ListAllGroupsResponse) SetStatusCode(v int32) *ListAllGroupsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAllGroupsResponse) SetBody(v *ListAllGroupsResponseBody) *ListAllGroupsResponse {
	s.Body = v
	return s
}

func (s *ListAllGroupsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
