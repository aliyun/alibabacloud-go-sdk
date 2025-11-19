// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSecurityIPGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateSecurityIPGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateSecurityIPGroupResponse
	GetStatusCode() *int32
	SetBody(v *CreateSecurityIPGroupResponseBody) *CreateSecurityIPGroupResponse
	GetBody() *CreateSecurityIPGroupResponseBody
}

type CreateSecurityIPGroupResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateSecurityIPGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateSecurityIPGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateSecurityIPGroupResponse) GoString() string {
	return s.String()
}

func (s *CreateSecurityIPGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateSecurityIPGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateSecurityIPGroupResponse) GetBody() *CreateSecurityIPGroupResponseBody {
	return s.Body
}

func (s *CreateSecurityIPGroupResponse) SetHeaders(v map[string]*string) *CreateSecurityIPGroupResponse {
	s.Headers = v
	return s
}

func (s *CreateSecurityIPGroupResponse) SetStatusCode(v int32) *CreateSecurityIPGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateSecurityIPGroupResponse) SetBody(v *CreateSecurityIPGroupResponseBody) *CreateSecurityIPGroupResponse {
	s.Body = v
	return s
}

func (s *CreateSecurityIPGroupResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
