// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyAIServiceSecurityIpsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyAIServiceSecurityIpsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyAIServiceSecurityIpsResponse
	GetStatusCode() *int32
	SetBody(v *ModifyAIServiceSecurityIpsResponseBody) *ModifyAIServiceSecurityIpsResponse
	GetBody() *ModifyAIServiceSecurityIpsResponseBody
}

type ModifyAIServiceSecurityIpsResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyAIServiceSecurityIpsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyAIServiceSecurityIpsResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyAIServiceSecurityIpsResponse) GoString() string {
	return s.String()
}

func (s *ModifyAIServiceSecurityIpsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyAIServiceSecurityIpsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyAIServiceSecurityIpsResponse) GetBody() *ModifyAIServiceSecurityIpsResponseBody {
	return s.Body
}

func (s *ModifyAIServiceSecurityIpsResponse) SetHeaders(v map[string]*string) *ModifyAIServiceSecurityIpsResponse {
	s.Headers = v
	return s
}

func (s *ModifyAIServiceSecurityIpsResponse) SetStatusCode(v int32) *ModifyAIServiceSecurityIpsResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyAIServiceSecurityIpsResponse) SetBody(v *ModifyAIServiceSecurityIpsResponseBody) *ModifyAIServiceSecurityIpsResponse {
	s.Body = v
	return s
}

func (s *ModifyAIServiceSecurityIpsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
