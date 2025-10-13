// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableHiveAccessResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DisableHiveAccessResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DisableHiveAccessResponse
	GetStatusCode() *int32
	SetBody(v *DisableHiveAccessResponseBody) *DisableHiveAccessResponse
	GetBody() *DisableHiveAccessResponseBody
}

type DisableHiveAccessResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DisableHiveAccessResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DisableHiveAccessResponse) String() string {
	return dara.Prettify(s)
}

func (s DisableHiveAccessResponse) GoString() string {
	return s.String()
}

func (s *DisableHiveAccessResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DisableHiveAccessResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DisableHiveAccessResponse) GetBody() *DisableHiveAccessResponseBody {
	return s.Body
}

func (s *DisableHiveAccessResponse) SetHeaders(v map[string]*string) *DisableHiveAccessResponse {
	s.Headers = v
	return s
}

func (s *DisableHiveAccessResponse) SetStatusCode(v int32) *DisableHiveAccessResponse {
	s.StatusCode = &v
	return s
}

func (s *DisableHiveAccessResponse) SetBody(v *DisableHiveAccessResponseBody) *DisableHiveAccessResponse {
	s.Body = v
	return s
}

func (s *DisableHiveAccessResponse) Validate() error {
	return dara.Validate(s)
}
