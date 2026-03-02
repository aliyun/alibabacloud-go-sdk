// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPermissionResourcePopResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListPermissionResourcePopResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListPermissionResourcePopResponse
	GetStatusCode() *int32
	SetBody(v *PermissionAllowResource) *ListPermissionResourcePopResponse
	GetBody() *PermissionAllowResource
}

type ListPermissionResourcePopResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PermissionAllowResource `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListPermissionResourcePopResponse) String() string {
	return dara.Prettify(s)
}

func (s ListPermissionResourcePopResponse) GoString() string {
	return s.String()
}

func (s *ListPermissionResourcePopResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListPermissionResourcePopResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListPermissionResourcePopResponse) GetBody() *PermissionAllowResource {
	return s.Body
}

func (s *ListPermissionResourcePopResponse) SetHeaders(v map[string]*string) *ListPermissionResourcePopResponse {
	s.Headers = v
	return s
}

func (s *ListPermissionResourcePopResponse) SetStatusCode(v int32) *ListPermissionResourcePopResponse {
	s.StatusCode = &v
	return s
}

func (s *ListPermissionResourcePopResponse) SetBody(v *PermissionAllowResource) *ListPermissionResourcePopResponse {
	s.Body = v
	return s
}

func (s *ListPermissionResourcePopResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
