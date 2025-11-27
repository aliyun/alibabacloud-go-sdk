// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyAccountMaskingPrivilegeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyAccountMaskingPrivilegeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyAccountMaskingPrivilegeResponse
	GetStatusCode() *int32
	SetBody(v *ModifyAccountMaskingPrivilegeResponseBody) *ModifyAccountMaskingPrivilegeResponse
	GetBody() *ModifyAccountMaskingPrivilegeResponseBody
}

type ModifyAccountMaskingPrivilegeResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyAccountMaskingPrivilegeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyAccountMaskingPrivilegeResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyAccountMaskingPrivilegeResponse) GoString() string {
	return s.String()
}

func (s *ModifyAccountMaskingPrivilegeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyAccountMaskingPrivilegeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyAccountMaskingPrivilegeResponse) GetBody() *ModifyAccountMaskingPrivilegeResponseBody {
	return s.Body
}

func (s *ModifyAccountMaskingPrivilegeResponse) SetHeaders(v map[string]*string) *ModifyAccountMaskingPrivilegeResponse {
	s.Headers = v
	return s
}

func (s *ModifyAccountMaskingPrivilegeResponse) SetStatusCode(v int32) *ModifyAccountMaskingPrivilegeResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyAccountMaskingPrivilegeResponse) SetBody(v *ModifyAccountMaskingPrivilegeResponseBody) *ModifyAccountMaskingPrivilegeResponse {
	s.Body = v
	return s
}

func (s *ModifyAccountMaskingPrivilegeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
