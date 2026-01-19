// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCloudAccountRolesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListCloudAccountRolesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListCloudAccountRolesResponse
	GetStatusCode() *int32
	SetBody(v *ListCloudAccountRolesResponseBody) *ListCloudAccountRolesResponse
	GetBody() *ListCloudAccountRolesResponseBody
}

type ListCloudAccountRolesResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListCloudAccountRolesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListCloudAccountRolesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListCloudAccountRolesResponse) GoString() string {
	return s.String()
}

func (s *ListCloudAccountRolesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListCloudAccountRolesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListCloudAccountRolesResponse) GetBody() *ListCloudAccountRolesResponseBody {
	return s.Body
}

func (s *ListCloudAccountRolesResponse) SetHeaders(v map[string]*string) *ListCloudAccountRolesResponse {
	s.Headers = v
	return s
}

func (s *ListCloudAccountRolesResponse) SetStatusCode(v int32) *ListCloudAccountRolesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListCloudAccountRolesResponse) SetBody(v *ListCloudAccountRolesResponseBody) *ListCloudAccountRolesResponse {
	s.Body = v
	return s
}

func (s *ListCloudAccountRolesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
