// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateUserPermissionsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateUserPermissionsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateUserPermissionsResponse
	GetStatusCode() *int32
}

type UpdateUserPermissionsResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
}

func (s UpdateUserPermissionsResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateUserPermissionsResponse) GoString() string {
	return s.String()
}

func (s *UpdateUserPermissionsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateUserPermissionsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateUserPermissionsResponse) SetHeaders(v map[string]*string) *UpdateUserPermissionsResponse {
	s.Headers = v
	return s
}

func (s *UpdateUserPermissionsResponse) SetStatusCode(v int32) *UpdateUserPermissionsResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateUserPermissionsResponse) Validate() error {
	return dara.Validate(s)
}
