// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListClusterMembersResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListClusterMembersResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListClusterMembersResponse
	GetStatusCode() *int32
	SetBody(v *ListClusterMembersResponseBody) *ListClusterMembersResponse
	GetBody() *ListClusterMembersResponseBody
}

type ListClusterMembersResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListClusterMembersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListClusterMembersResponse) String() string {
	return dara.Prettify(s)
}

func (s ListClusterMembersResponse) GoString() string {
	return s.String()
}

func (s *ListClusterMembersResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListClusterMembersResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListClusterMembersResponse) GetBody() *ListClusterMembersResponseBody {
	return s.Body
}

func (s *ListClusterMembersResponse) SetHeaders(v map[string]*string) *ListClusterMembersResponse {
	s.Headers = v
	return s
}

func (s *ListClusterMembersResponse) SetStatusCode(v int32) *ListClusterMembersResponse {
	s.StatusCode = &v
	return s
}

func (s *ListClusterMembersResponse) SetBody(v *ListClusterMembersResponseBody) *ListClusterMembersResponse {
	s.Body = v
	return s
}

func (s *ListClusterMembersResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
