// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateGlobalSecurityIPGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateGlobalSecurityIPGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateGlobalSecurityIPGroupResponse
	GetStatusCode() *int32
	SetBody(v *CreateGlobalSecurityIPGroupResponseBody) *CreateGlobalSecurityIPGroupResponse
	GetBody() *CreateGlobalSecurityIPGroupResponseBody
}

type CreateGlobalSecurityIPGroupResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateGlobalSecurityIPGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateGlobalSecurityIPGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateGlobalSecurityIPGroupResponse) GoString() string {
	return s.String()
}

func (s *CreateGlobalSecurityIPGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateGlobalSecurityIPGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateGlobalSecurityIPGroupResponse) GetBody() *CreateGlobalSecurityIPGroupResponseBody {
	return s.Body
}

func (s *CreateGlobalSecurityIPGroupResponse) SetHeaders(v map[string]*string) *CreateGlobalSecurityIPGroupResponse {
	s.Headers = v
	return s
}

func (s *CreateGlobalSecurityIPGroupResponse) SetStatusCode(v int32) *CreateGlobalSecurityIPGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateGlobalSecurityIPGroupResponse) SetBody(v *CreateGlobalSecurityIPGroupResponseBody) *CreateGlobalSecurityIPGroupResponse {
	s.Body = v
	return s
}

func (s *CreateGlobalSecurityIPGroupResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
