// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPermissionResourceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListPermissionResourceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListPermissionResourceResponse
	GetStatusCode() *int32
	SetBody(v *PermissionAllowResource) *ListPermissionResourceResponse
	GetBody() *PermissionAllowResource
}

type ListPermissionResourceResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PermissionAllowResource `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListPermissionResourceResponse) String() string {
	return dara.Prettify(s)
}

func (s ListPermissionResourceResponse) GoString() string {
	return s.String()
}

func (s *ListPermissionResourceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListPermissionResourceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListPermissionResourceResponse) GetBody() *PermissionAllowResource {
	return s.Body
}

func (s *ListPermissionResourceResponse) SetHeaders(v map[string]*string) *ListPermissionResourceResponse {
	s.Headers = v
	return s
}

func (s *ListPermissionResourceResponse) SetStatusCode(v int32) *ListPermissionResourceResponse {
	s.StatusCode = &v
	return s
}

func (s *ListPermissionResourceResponse) SetBody(v *PermissionAllowResource) *ListPermissionResourceResponse {
	s.Body = v
	return s
}

func (s *ListPermissionResourceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
