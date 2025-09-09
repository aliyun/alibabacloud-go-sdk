// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyHostShareKeyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyHostShareKeyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyHostShareKeyResponse
	GetStatusCode() *int32
	SetBody(v *ModifyHostShareKeyResponseBody) *ModifyHostShareKeyResponse
	GetBody() *ModifyHostShareKeyResponseBody
}

type ModifyHostShareKeyResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyHostShareKeyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyHostShareKeyResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyHostShareKeyResponse) GoString() string {
	return s.String()
}

func (s *ModifyHostShareKeyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyHostShareKeyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyHostShareKeyResponse) GetBody() *ModifyHostShareKeyResponseBody {
	return s.Body
}

func (s *ModifyHostShareKeyResponse) SetHeaders(v map[string]*string) *ModifyHostShareKeyResponse {
	s.Headers = v
	return s
}

func (s *ModifyHostShareKeyResponse) SetStatusCode(v int32) *ModifyHostShareKeyResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyHostShareKeyResponse) SetBody(v *ModifyHostShareKeyResponseBody) *ModifyHostShareKeyResponse {
	s.Body = v
	return s
}

func (s *ModifyHostShareKeyResponse) Validate() error {
	return dara.Validate(s)
}
