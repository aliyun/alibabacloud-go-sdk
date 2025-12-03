// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListRepositoryGroupsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListRepositoryGroupsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListRepositoryGroupsResponse
	GetStatusCode() *int32
	SetBody(v *ListRepositoryGroupsResponseBody) *ListRepositoryGroupsResponse
	GetBody() *ListRepositoryGroupsResponseBody
}

type ListRepositoryGroupsResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListRepositoryGroupsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListRepositoryGroupsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListRepositoryGroupsResponse) GoString() string {
	return s.String()
}

func (s *ListRepositoryGroupsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListRepositoryGroupsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListRepositoryGroupsResponse) GetBody() *ListRepositoryGroupsResponseBody {
	return s.Body
}

func (s *ListRepositoryGroupsResponse) SetHeaders(v map[string]*string) *ListRepositoryGroupsResponse {
	s.Headers = v
	return s
}

func (s *ListRepositoryGroupsResponse) SetStatusCode(v int32) *ListRepositoryGroupsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListRepositoryGroupsResponse) SetBody(v *ListRepositoryGroupsResponseBody) *ListRepositoryGroupsResponse {
	s.Body = v
	return s
}

func (s *ListRepositoryGroupsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
