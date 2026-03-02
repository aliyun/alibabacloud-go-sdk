// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeletePrivilegeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeletePrivilegeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeletePrivilegeResponse
	GetStatusCode() *int32
}

type DeletePrivilegeResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
}

func (s DeletePrivilegeResponse) String() string {
	return dara.Prettify(s)
}

func (s DeletePrivilegeResponse) GoString() string {
	return s.String()
}

func (s *DeletePrivilegeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeletePrivilegeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeletePrivilegeResponse) SetHeaders(v map[string]*string) *DeletePrivilegeResponse {
	s.Headers = v
	return s
}

func (s *DeletePrivilegeResponse) SetStatusCode(v int32) *DeletePrivilegeResponse {
	s.StatusCode = &v
	return s
}

func (s *DeletePrivilegeResponse) Validate() error {
	return dara.Validate(s)
}
