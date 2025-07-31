// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyGlobalSecurityIPGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyGlobalSecurityIPGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyGlobalSecurityIPGroupResponse
	GetStatusCode() *int32
	SetBody(v *ModifyGlobalSecurityIPGroupResponseBody) *ModifyGlobalSecurityIPGroupResponse
	GetBody() *ModifyGlobalSecurityIPGroupResponseBody
}

type ModifyGlobalSecurityIPGroupResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyGlobalSecurityIPGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyGlobalSecurityIPGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyGlobalSecurityIPGroupResponse) GoString() string {
	return s.String()
}

func (s *ModifyGlobalSecurityIPGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyGlobalSecurityIPGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyGlobalSecurityIPGroupResponse) GetBody() *ModifyGlobalSecurityIPGroupResponseBody {
	return s.Body
}

func (s *ModifyGlobalSecurityIPGroupResponse) SetHeaders(v map[string]*string) *ModifyGlobalSecurityIPGroupResponse {
	s.Headers = v
	return s
}

func (s *ModifyGlobalSecurityIPGroupResponse) SetStatusCode(v int32) *ModifyGlobalSecurityIPGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyGlobalSecurityIPGroupResponse) SetBody(v *ModifyGlobalSecurityIPGroupResponseBody) *ModifyGlobalSecurityIPGroupResponse {
	s.Body = v
	return s
}

func (s *ModifyGlobalSecurityIPGroupResponse) Validate() error {
	return dara.Validate(s)
}
