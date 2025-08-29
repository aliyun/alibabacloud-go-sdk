// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCrowdUsersResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListCrowdUsersResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListCrowdUsersResponse
	GetStatusCode() *int32
	SetBody(v *ListCrowdUsersResponseBody) *ListCrowdUsersResponse
	GetBody() *ListCrowdUsersResponseBody
}

type ListCrowdUsersResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListCrowdUsersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListCrowdUsersResponse) String() string {
	return dara.Prettify(s)
}

func (s ListCrowdUsersResponse) GoString() string {
	return s.String()
}

func (s *ListCrowdUsersResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListCrowdUsersResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListCrowdUsersResponse) GetBody() *ListCrowdUsersResponseBody {
	return s.Body
}

func (s *ListCrowdUsersResponse) SetHeaders(v map[string]*string) *ListCrowdUsersResponse {
	s.Headers = v
	return s
}

func (s *ListCrowdUsersResponse) SetStatusCode(v int32) *ListCrowdUsersResponse {
	s.StatusCode = &v
	return s
}

func (s *ListCrowdUsersResponse) SetBody(v *ListCrowdUsersResponseBody) *ListCrowdUsersResponse {
	s.Body = v
	return s
}

func (s *ListCrowdUsersResponse) Validate() error {
	return dara.Validate(s)
}
