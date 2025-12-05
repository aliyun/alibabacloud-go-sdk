// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListHostGroupsForUserResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListHostGroupsForUserResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListHostGroupsForUserResponse
	GetStatusCode() *int32
	SetBody(v *ListHostGroupsForUserResponseBody) *ListHostGroupsForUserResponse
	GetBody() *ListHostGroupsForUserResponseBody
}

type ListHostGroupsForUserResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListHostGroupsForUserResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListHostGroupsForUserResponse) String() string {
	return dara.Prettify(s)
}

func (s ListHostGroupsForUserResponse) GoString() string {
	return s.String()
}

func (s *ListHostGroupsForUserResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListHostGroupsForUserResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListHostGroupsForUserResponse) GetBody() *ListHostGroupsForUserResponseBody {
	return s.Body
}

func (s *ListHostGroupsForUserResponse) SetHeaders(v map[string]*string) *ListHostGroupsForUserResponse {
	s.Headers = v
	return s
}

func (s *ListHostGroupsForUserResponse) SetStatusCode(v int32) *ListHostGroupsForUserResponse {
	s.StatusCode = &v
	return s
}

func (s *ListHostGroupsForUserResponse) SetBody(v *ListHostGroupsForUserResponseBody) *ListHostGroupsForUserResponse {
	s.Body = v
	return s
}

func (s *ListHostGroupsForUserResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
