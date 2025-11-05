// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGrantDatabasePermissionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GrantDatabasePermissionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GrantDatabasePermissionResponse
	GetStatusCode() *int32
	SetBody(v *GrantDatabasePermissionResponseBody) *GrantDatabasePermissionResponse
	GetBody() *GrantDatabasePermissionResponseBody
}

type GrantDatabasePermissionResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GrantDatabasePermissionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GrantDatabasePermissionResponse) String() string {
	return dara.Prettify(s)
}

func (s GrantDatabasePermissionResponse) GoString() string {
	return s.String()
}

func (s *GrantDatabasePermissionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GrantDatabasePermissionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GrantDatabasePermissionResponse) GetBody() *GrantDatabasePermissionResponseBody {
	return s.Body
}

func (s *GrantDatabasePermissionResponse) SetHeaders(v map[string]*string) *GrantDatabasePermissionResponse {
	s.Headers = v
	return s
}

func (s *GrantDatabasePermissionResponse) SetStatusCode(v int32) *GrantDatabasePermissionResponse {
	s.StatusCode = &v
	return s
}

func (s *GrantDatabasePermissionResponse) SetBody(v *GrantDatabasePermissionResponseBody) *GrantDatabasePermissionResponse {
	s.Body = v
	return s
}

func (s *GrantDatabasePermissionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
