// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableKeyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DisableKeyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DisableKeyResponse
	GetStatusCode() *int32
	SetBody(v *DisableKeyResponseBody) *DisableKeyResponse
	GetBody() *DisableKeyResponseBody
}

type DisableKeyResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DisableKeyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DisableKeyResponse) String() string {
	return dara.Prettify(s)
}

func (s DisableKeyResponse) GoString() string {
	return s.String()
}

func (s *DisableKeyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DisableKeyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DisableKeyResponse) GetBody() *DisableKeyResponseBody {
	return s.Body
}

func (s *DisableKeyResponse) SetHeaders(v map[string]*string) *DisableKeyResponse {
	s.Headers = v
	return s
}

func (s *DisableKeyResponse) SetStatusCode(v int32) *DisableKeyResponse {
	s.StatusCode = &v
	return s
}

func (s *DisableKeyResponse) SetBody(v *DisableKeyResponseBody) *DisableKeyResponse {
	s.Body = v
	return s
}

func (s *DisableKeyResponse) Validate() error {
	return dara.Validate(s)
}
