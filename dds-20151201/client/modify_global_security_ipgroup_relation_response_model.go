// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyGlobalSecurityIPGroupRelationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyGlobalSecurityIPGroupRelationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyGlobalSecurityIPGroupRelationResponse
	GetStatusCode() *int32
	SetBody(v *ModifyGlobalSecurityIPGroupRelationResponseBody) *ModifyGlobalSecurityIPGroupRelationResponse
	GetBody() *ModifyGlobalSecurityIPGroupRelationResponseBody
}

type ModifyGlobalSecurityIPGroupRelationResponse struct {
	Headers    map[string]*string                               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyGlobalSecurityIPGroupRelationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyGlobalSecurityIPGroupRelationResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyGlobalSecurityIPGroupRelationResponse) GoString() string {
	return s.String()
}

func (s *ModifyGlobalSecurityIPGroupRelationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyGlobalSecurityIPGroupRelationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyGlobalSecurityIPGroupRelationResponse) GetBody() *ModifyGlobalSecurityIPGroupRelationResponseBody {
	return s.Body
}

func (s *ModifyGlobalSecurityIPGroupRelationResponse) SetHeaders(v map[string]*string) *ModifyGlobalSecurityIPGroupRelationResponse {
	s.Headers = v
	return s
}

func (s *ModifyGlobalSecurityIPGroupRelationResponse) SetStatusCode(v int32) *ModifyGlobalSecurityIPGroupRelationResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyGlobalSecurityIPGroupRelationResponse) SetBody(v *ModifyGlobalSecurityIPGroupRelationResponseBody) *ModifyGlobalSecurityIPGroupRelationResponse {
	s.Body = v
	return s
}

func (s *ModifyGlobalSecurityIPGroupRelationResponse) Validate() error {
	return dara.Validate(s)
}
