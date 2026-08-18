// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyContextDBSecurityIpsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyContextDBSecurityIpsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyContextDBSecurityIpsResponse
	GetStatusCode() *int32
	SetBody(v *ModifyContextDBSecurityIpsResponseBody) *ModifyContextDBSecurityIpsResponse
	GetBody() *ModifyContextDBSecurityIpsResponseBody
}

type ModifyContextDBSecurityIpsResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyContextDBSecurityIpsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyContextDBSecurityIpsResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyContextDBSecurityIpsResponse) GoString() string {
	return s.String()
}

func (s *ModifyContextDBSecurityIpsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyContextDBSecurityIpsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyContextDBSecurityIpsResponse) GetBody() *ModifyContextDBSecurityIpsResponseBody {
	return s.Body
}

func (s *ModifyContextDBSecurityIpsResponse) SetHeaders(v map[string]*string) *ModifyContextDBSecurityIpsResponse {
	s.Headers = v
	return s
}

func (s *ModifyContextDBSecurityIpsResponse) SetStatusCode(v int32) *ModifyContextDBSecurityIpsResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyContextDBSecurityIpsResponse) SetBody(v *ModifyContextDBSecurityIpsResponseBody) *ModifyContextDBSecurityIpsResponse {
	s.Body = v
	return s
}

func (s *ModifyContextDBSecurityIpsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
