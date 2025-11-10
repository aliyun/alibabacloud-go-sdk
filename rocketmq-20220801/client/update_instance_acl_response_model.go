// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateInstanceAclResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateInstanceAclResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateInstanceAclResponse
	GetStatusCode() *int32
	SetBody(v *UpdateInstanceAclResponseBody) *UpdateInstanceAclResponse
	GetBody() *UpdateInstanceAclResponseBody
}

type UpdateInstanceAclResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateInstanceAclResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateInstanceAclResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateInstanceAclResponse) GoString() string {
	return s.String()
}

func (s *UpdateInstanceAclResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateInstanceAclResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateInstanceAclResponse) GetBody() *UpdateInstanceAclResponseBody {
	return s.Body
}

func (s *UpdateInstanceAclResponse) SetHeaders(v map[string]*string) *UpdateInstanceAclResponse {
	s.Headers = v
	return s
}

func (s *UpdateInstanceAclResponse) SetStatusCode(v int32) *UpdateInstanceAclResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateInstanceAclResponse) SetBody(v *UpdateInstanceAclResponseBody) *UpdateInstanceAclResponse {
	s.Body = v
	return s
}

func (s *UpdateInstanceAclResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
