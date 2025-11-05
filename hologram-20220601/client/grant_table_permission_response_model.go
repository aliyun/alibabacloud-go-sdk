// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGrantTablePermissionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GrantTablePermissionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GrantTablePermissionResponse
	GetStatusCode() *int32
	SetBody(v *GrantTablePermissionResponseBody) *GrantTablePermissionResponse
	GetBody() *GrantTablePermissionResponseBody
}

type GrantTablePermissionResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GrantTablePermissionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GrantTablePermissionResponse) String() string {
	return dara.Prettify(s)
}

func (s GrantTablePermissionResponse) GoString() string {
	return s.String()
}

func (s *GrantTablePermissionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GrantTablePermissionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GrantTablePermissionResponse) GetBody() *GrantTablePermissionResponseBody {
	return s.Body
}

func (s *GrantTablePermissionResponse) SetHeaders(v map[string]*string) *GrantTablePermissionResponse {
	s.Headers = v
	return s
}

func (s *GrantTablePermissionResponse) SetStatusCode(v int32) *GrantTablePermissionResponse {
	s.StatusCode = &v
	return s
}

func (s *GrantTablePermissionResponse) SetBody(v *GrantTablePermissionResponseBody) *GrantTablePermissionResponse {
	s.Body = v
	return s
}

func (s *GrantTablePermissionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
