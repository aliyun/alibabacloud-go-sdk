// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateSecurityLevelResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateSecurityLevelResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateSecurityLevelResponse
	GetStatusCode() *int32
	SetBody(v *UpdateSecurityLevelResponseBody) *UpdateSecurityLevelResponse
	GetBody() *UpdateSecurityLevelResponseBody
}

type UpdateSecurityLevelResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateSecurityLevelResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateSecurityLevelResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateSecurityLevelResponse) GoString() string {
	return s.String()
}

func (s *UpdateSecurityLevelResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateSecurityLevelResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateSecurityLevelResponse) GetBody() *UpdateSecurityLevelResponseBody {
	return s.Body
}

func (s *UpdateSecurityLevelResponse) SetHeaders(v map[string]*string) *UpdateSecurityLevelResponse {
	s.Headers = v
	return s
}

func (s *UpdateSecurityLevelResponse) SetStatusCode(v int32) *UpdateSecurityLevelResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateSecurityLevelResponse) SetBody(v *UpdateSecurityLevelResponseBody) *UpdateSecurityLevelResponse {
	s.Body = v
	return s
}

func (s *UpdateSecurityLevelResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
