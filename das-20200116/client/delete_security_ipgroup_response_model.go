// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteSecurityIPGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteSecurityIPGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteSecurityIPGroupResponse
	GetStatusCode() *int32
	SetBody(v *DeleteSecurityIPGroupResponseBody) *DeleteSecurityIPGroupResponse
	GetBody() *DeleteSecurityIPGroupResponseBody
}

type DeleteSecurityIPGroupResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteSecurityIPGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteSecurityIPGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteSecurityIPGroupResponse) GoString() string {
	return s.String()
}

func (s *DeleteSecurityIPGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteSecurityIPGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteSecurityIPGroupResponse) GetBody() *DeleteSecurityIPGroupResponseBody {
	return s.Body
}

func (s *DeleteSecurityIPGroupResponse) SetHeaders(v map[string]*string) *DeleteSecurityIPGroupResponse {
	s.Headers = v
	return s
}

func (s *DeleteSecurityIPGroupResponse) SetStatusCode(v int32) *DeleteSecurityIPGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteSecurityIPGroupResponse) SetBody(v *DeleteSecurityIPGroupResponseBody) *DeleteSecurityIPGroupResponse {
	s.Body = v
	return s
}

func (s *DeleteSecurityIPGroupResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
