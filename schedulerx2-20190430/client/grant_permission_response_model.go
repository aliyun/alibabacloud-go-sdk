// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGrantPermissionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GrantPermissionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GrantPermissionResponse
	GetStatusCode() *int32
	SetBody(v *GrantPermissionResponseBody) *GrantPermissionResponse
	GetBody() *GrantPermissionResponseBody
}

type GrantPermissionResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GrantPermissionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GrantPermissionResponse) String() string {
	return dara.Prettify(s)
}

func (s GrantPermissionResponse) GoString() string {
	return s.String()
}

func (s *GrantPermissionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GrantPermissionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GrantPermissionResponse) GetBody() *GrantPermissionResponseBody {
	return s.Body
}

func (s *GrantPermissionResponse) SetHeaders(v map[string]*string) *GrantPermissionResponse {
	s.Headers = v
	return s
}

func (s *GrantPermissionResponse) SetStatusCode(v int32) *GrantPermissionResponse {
	s.StatusCode = &v
	return s
}

func (s *GrantPermissionResponse) SetBody(v *GrantPermissionResponseBody) *GrantPermissionResponse {
	s.Body = v
	return s
}

func (s *GrantPermissionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
