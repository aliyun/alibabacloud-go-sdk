// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyLangfuseSecurityIPListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyLangfuseSecurityIPListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyLangfuseSecurityIPListResponse
	GetStatusCode() *int32
	SetBody(v *ModifyLangfuseSecurityIPListResponseBody) *ModifyLangfuseSecurityIPListResponse
	GetBody() *ModifyLangfuseSecurityIPListResponseBody
}

type ModifyLangfuseSecurityIPListResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyLangfuseSecurityIPListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyLangfuseSecurityIPListResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyLangfuseSecurityIPListResponse) GoString() string {
	return s.String()
}

func (s *ModifyLangfuseSecurityIPListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyLangfuseSecurityIPListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyLangfuseSecurityIPListResponse) GetBody() *ModifyLangfuseSecurityIPListResponseBody {
	return s.Body
}

func (s *ModifyLangfuseSecurityIPListResponse) SetHeaders(v map[string]*string) *ModifyLangfuseSecurityIPListResponse {
	s.Headers = v
	return s
}

func (s *ModifyLangfuseSecurityIPListResponse) SetStatusCode(v int32) *ModifyLangfuseSecurityIPListResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyLangfuseSecurityIPListResponse) SetBody(v *ModifyLangfuseSecurityIPListResponseBody) *ModifyLangfuseSecurityIPListResponse {
	s.Body = v
	return s
}

func (s *ModifyLangfuseSecurityIPListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
