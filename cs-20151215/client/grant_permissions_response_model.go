// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGrantPermissionsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GrantPermissionsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GrantPermissionsResponse
	GetStatusCode() *int32
}

type GrantPermissionsResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
}

func (s GrantPermissionsResponse) String() string {
	return dara.Prettify(s)
}

func (s GrantPermissionsResponse) GoString() string {
	return s.String()
}

func (s *GrantPermissionsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GrantPermissionsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GrantPermissionsResponse) SetHeaders(v map[string]*string) *GrantPermissionsResponse {
	s.Headers = v
	return s
}

func (s *GrantPermissionsResponse) SetStatusCode(v int32) *GrantPermissionsResponse {
	s.StatusCode = &v
	return s
}

func (s *GrantPermissionsResponse) Validate() error {
	return dara.Validate(s)
}
