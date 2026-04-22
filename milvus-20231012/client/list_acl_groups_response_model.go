// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAclGroupsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListAclGroupsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListAclGroupsResponse
	GetStatusCode() *int32
	SetBody(v *ListAclGroupsResponseBody) *ListAclGroupsResponse
	GetBody() *ListAclGroupsResponseBody
}

type ListAclGroupsResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAclGroupsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAclGroupsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListAclGroupsResponse) GoString() string {
	return s.String()
}

func (s *ListAclGroupsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListAclGroupsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListAclGroupsResponse) GetBody() *ListAclGroupsResponseBody {
	return s.Body
}

func (s *ListAclGroupsResponse) SetHeaders(v map[string]*string) *ListAclGroupsResponse {
	s.Headers = v
	return s
}

func (s *ListAclGroupsResponse) SetStatusCode(v int32) *ListAclGroupsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAclGroupsResponse) SetBody(v *ListAclGroupsResponseBody) *ListAclGroupsResponse {
	s.Body = v
	return s
}

func (s *ListAclGroupsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
