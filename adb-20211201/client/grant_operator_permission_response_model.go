// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGrantOperatorPermissionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GrantOperatorPermissionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GrantOperatorPermissionResponse
	GetStatusCode() *int32
	SetBody(v *GrantOperatorPermissionResponseBody) *GrantOperatorPermissionResponse
	GetBody() *GrantOperatorPermissionResponseBody
}

type GrantOperatorPermissionResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GrantOperatorPermissionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GrantOperatorPermissionResponse) String() string {
	return dara.Prettify(s)
}

func (s GrantOperatorPermissionResponse) GoString() string {
	return s.String()
}

func (s *GrantOperatorPermissionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GrantOperatorPermissionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GrantOperatorPermissionResponse) GetBody() *GrantOperatorPermissionResponseBody {
	return s.Body
}

func (s *GrantOperatorPermissionResponse) SetHeaders(v map[string]*string) *GrantOperatorPermissionResponse {
	s.Headers = v
	return s
}

func (s *GrantOperatorPermissionResponse) SetStatusCode(v int32) *GrantOperatorPermissionResponse {
	s.StatusCode = &v
	return s
}

func (s *GrantOperatorPermissionResponse) SetBody(v *GrantOperatorPermissionResponseBody) *GrantOperatorPermissionResponse {
	s.Body = v
	return s
}

func (s *GrantOperatorPermissionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
