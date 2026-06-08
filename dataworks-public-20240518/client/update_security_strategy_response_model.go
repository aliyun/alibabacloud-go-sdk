// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateSecurityStrategyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateSecurityStrategyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateSecurityStrategyResponse
	GetStatusCode() *int32
	SetBody(v *UpdateSecurityStrategyResponseBody) *UpdateSecurityStrategyResponse
	GetBody() *UpdateSecurityStrategyResponseBody
}

type UpdateSecurityStrategyResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateSecurityStrategyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateSecurityStrategyResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateSecurityStrategyResponse) GoString() string {
	return s.String()
}

func (s *UpdateSecurityStrategyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateSecurityStrategyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateSecurityStrategyResponse) GetBody() *UpdateSecurityStrategyResponseBody {
	return s.Body
}

func (s *UpdateSecurityStrategyResponse) SetHeaders(v map[string]*string) *UpdateSecurityStrategyResponse {
	s.Headers = v
	return s
}

func (s *UpdateSecurityStrategyResponse) SetStatusCode(v int32) *UpdateSecurityStrategyResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateSecurityStrategyResponse) SetBody(v *UpdateSecurityStrategyResponseBody) *UpdateSecurityStrategyResponse {
	s.Body = v
	return s
}

func (s *UpdateSecurityStrategyResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
