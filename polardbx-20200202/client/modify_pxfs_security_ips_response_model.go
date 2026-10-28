// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyPxfsSecurityIpsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyPxfsSecurityIpsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyPxfsSecurityIpsResponse
	GetStatusCode() *int32
	SetBody(v *ModifyPxfsSecurityIpsResponseBody) *ModifyPxfsSecurityIpsResponse
	GetBody() *ModifyPxfsSecurityIpsResponseBody
}

type ModifyPxfsSecurityIpsResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyPxfsSecurityIpsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyPxfsSecurityIpsResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyPxfsSecurityIpsResponse) GoString() string {
	return s.String()
}

func (s *ModifyPxfsSecurityIpsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyPxfsSecurityIpsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyPxfsSecurityIpsResponse) GetBody() *ModifyPxfsSecurityIpsResponseBody {
	return s.Body
}

func (s *ModifyPxfsSecurityIpsResponse) SetHeaders(v map[string]*string) *ModifyPxfsSecurityIpsResponse {
	s.Headers = v
	return s
}

func (s *ModifyPxfsSecurityIpsResponse) SetStatusCode(v int32) *ModifyPxfsSecurityIpsResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyPxfsSecurityIpsResponse) SetBody(v *ModifyPxfsSecurityIpsResponseBody) *ModifyPxfsSecurityIpsResponse {
	s.Body = v
	return s
}

func (s *ModifyPxfsSecurityIpsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
