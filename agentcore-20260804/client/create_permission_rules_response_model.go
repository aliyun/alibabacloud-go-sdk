// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePermissionRulesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreatePermissionRulesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreatePermissionRulesResponse
	GetStatusCode() *int32
	SetBody(v *CreatePermissionRulesResponseBody) *CreatePermissionRulesResponse
	GetBody() *CreatePermissionRulesResponseBody
}

type CreatePermissionRulesResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreatePermissionRulesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreatePermissionRulesResponse) String() string {
	return dara.Prettify(s)
}

func (s CreatePermissionRulesResponse) GoString() string {
	return s.String()
}

func (s *CreatePermissionRulesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreatePermissionRulesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreatePermissionRulesResponse) GetBody() *CreatePermissionRulesResponseBody {
	return s.Body
}

func (s *CreatePermissionRulesResponse) SetHeaders(v map[string]*string) *CreatePermissionRulesResponse {
	s.Headers = v
	return s
}

func (s *CreatePermissionRulesResponse) SetStatusCode(v int32) *CreatePermissionRulesResponse {
	s.StatusCode = &v
	return s
}

func (s *CreatePermissionRulesResponse) SetBody(v *CreatePermissionRulesResponseBody) *CreatePermissionRulesResponse {
	s.Body = v
	return s
}

func (s *CreatePermissionRulesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
