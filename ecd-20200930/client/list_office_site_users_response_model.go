// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListOfficeSiteUsersResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListOfficeSiteUsersResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListOfficeSiteUsersResponse
	GetStatusCode() *int32
	SetBody(v *ListOfficeSiteUsersResponseBody) *ListOfficeSiteUsersResponse
	GetBody() *ListOfficeSiteUsersResponseBody
}

type ListOfficeSiteUsersResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListOfficeSiteUsersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListOfficeSiteUsersResponse) String() string {
	return dara.Prettify(s)
}

func (s ListOfficeSiteUsersResponse) GoString() string {
	return s.String()
}

func (s *ListOfficeSiteUsersResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListOfficeSiteUsersResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListOfficeSiteUsersResponse) GetBody() *ListOfficeSiteUsersResponseBody {
	return s.Body
}

func (s *ListOfficeSiteUsersResponse) SetHeaders(v map[string]*string) *ListOfficeSiteUsersResponse {
	s.Headers = v
	return s
}

func (s *ListOfficeSiteUsersResponse) SetStatusCode(v int32) *ListOfficeSiteUsersResponse {
	s.StatusCode = &v
	return s
}

func (s *ListOfficeSiteUsersResponse) SetBody(v *ListOfficeSiteUsersResponseBody) *ListOfficeSiteUsersResponse {
	s.Body = v
	return s
}

func (s *ListOfficeSiteUsersResponse) Validate() error {
	return dara.Validate(s)
}
