// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteGlobalSecurityIPGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteGlobalSecurityIPGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteGlobalSecurityIPGroupResponse
	GetStatusCode() *int32
	SetBody(v *DeleteGlobalSecurityIPGroupResponseBody) *DeleteGlobalSecurityIPGroupResponse
	GetBody() *DeleteGlobalSecurityIPGroupResponseBody
}

type DeleteGlobalSecurityIPGroupResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteGlobalSecurityIPGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteGlobalSecurityIPGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteGlobalSecurityIPGroupResponse) GoString() string {
	return s.String()
}

func (s *DeleteGlobalSecurityIPGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteGlobalSecurityIPGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteGlobalSecurityIPGroupResponse) GetBody() *DeleteGlobalSecurityIPGroupResponseBody {
	return s.Body
}

func (s *DeleteGlobalSecurityIPGroupResponse) SetHeaders(v map[string]*string) *DeleteGlobalSecurityIPGroupResponse {
	s.Headers = v
	return s
}

func (s *DeleteGlobalSecurityIPGroupResponse) SetStatusCode(v int32) *DeleteGlobalSecurityIPGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteGlobalSecurityIPGroupResponse) SetBody(v *DeleteGlobalSecurityIPGroupResponseBody) *DeleteGlobalSecurityIPGroupResponse {
	s.Body = v
	return s
}

func (s *DeleteGlobalSecurityIPGroupResponse) Validate() error {
	return dara.Validate(s)
}
