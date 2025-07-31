// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyGlobalSecurityIPGroupNameResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyGlobalSecurityIPGroupNameResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyGlobalSecurityIPGroupNameResponse
	GetStatusCode() *int32
	SetBody(v *ModifyGlobalSecurityIPGroupNameResponseBody) *ModifyGlobalSecurityIPGroupNameResponse
	GetBody() *ModifyGlobalSecurityIPGroupNameResponseBody
}

type ModifyGlobalSecurityIPGroupNameResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyGlobalSecurityIPGroupNameResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyGlobalSecurityIPGroupNameResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyGlobalSecurityIPGroupNameResponse) GoString() string {
	return s.String()
}

func (s *ModifyGlobalSecurityIPGroupNameResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyGlobalSecurityIPGroupNameResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyGlobalSecurityIPGroupNameResponse) GetBody() *ModifyGlobalSecurityIPGroupNameResponseBody {
	return s.Body
}

func (s *ModifyGlobalSecurityIPGroupNameResponse) SetHeaders(v map[string]*string) *ModifyGlobalSecurityIPGroupNameResponse {
	s.Headers = v
	return s
}

func (s *ModifyGlobalSecurityIPGroupNameResponse) SetStatusCode(v int32) *ModifyGlobalSecurityIPGroupNameResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyGlobalSecurityIPGroupNameResponse) SetBody(v *ModifyGlobalSecurityIPGroupNameResponseBody) *ModifyGlobalSecurityIPGroupNameResponse {
	s.Body = v
	return s
}

func (s *ModifyGlobalSecurityIPGroupNameResponse) Validate() error {
	return dara.Validate(s)
}
