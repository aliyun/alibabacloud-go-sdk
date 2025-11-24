// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGrantUserPermissionsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GrantUserPermissionsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GrantUserPermissionsResponse
	GetStatusCode() *int32
	SetBody(v *GrantUserPermissionsResponseBody) *GrantUserPermissionsResponse
	GetBody() *GrantUserPermissionsResponseBody
}

type GrantUserPermissionsResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GrantUserPermissionsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GrantUserPermissionsResponse) String() string {
	return dara.Prettify(s)
}

func (s GrantUserPermissionsResponse) GoString() string {
	return s.String()
}

func (s *GrantUserPermissionsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GrantUserPermissionsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GrantUserPermissionsResponse) GetBody() *GrantUserPermissionsResponseBody {
	return s.Body
}

func (s *GrantUserPermissionsResponse) SetHeaders(v map[string]*string) *GrantUserPermissionsResponse {
	s.Headers = v
	return s
}

func (s *GrantUserPermissionsResponse) SetStatusCode(v int32) *GrantUserPermissionsResponse {
	s.StatusCode = &v
	return s
}

func (s *GrantUserPermissionsResponse) SetBody(v *GrantUserPermissionsResponseBody) *GrantUserPermissionsResponse {
	s.Body = v
	return s
}

func (s *GrantUserPermissionsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
