// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifySecurityIpsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifySecurityIpsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifySecurityIpsResponse
	GetStatusCode() *int32
	SetBody(v *ModifySecurityIpsResponseBody) *ModifySecurityIpsResponse
	GetBody() *ModifySecurityIpsResponseBody
}

type ModifySecurityIpsResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifySecurityIpsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifySecurityIpsResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifySecurityIpsResponse) GoString() string {
	return s.String()
}

func (s *ModifySecurityIpsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifySecurityIpsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifySecurityIpsResponse) GetBody() *ModifySecurityIpsResponseBody {
	return s.Body
}

func (s *ModifySecurityIpsResponse) SetHeaders(v map[string]*string) *ModifySecurityIpsResponse {
	s.Headers = v
	return s
}

func (s *ModifySecurityIpsResponse) SetStatusCode(v int32) *ModifySecurityIpsResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifySecurityIpsResponse) SetBody(v *ModifySecurityIpsResponseBody) *ModifySecurityIpsResponse {
	s.Body = v
	return s
}

func (s *ModifySecurityIpsResponse) Validate() error {
	return dara.Validate(s)
}
