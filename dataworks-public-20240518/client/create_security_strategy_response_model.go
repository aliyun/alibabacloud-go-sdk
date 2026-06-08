// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSecurityStrategyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateSecurityStrategyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateSecurityStrategyResponse
	GetStatusCode() *int32
	SetBody(v *CreateSecurityStrategyResponseBody) *CreateSecurityStrategyResponse
	GetBody() *CreateSecurityStrategyResponseBody
}

type CreateSecurityStrategyResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateSecurityStrategyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateSecurityStrategyResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateSecurityStrategyResponse) GoString() string {
	return s.String()
}

func (s *CreateSecurityStrategyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateSecurityStrategyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateSecurityStrategyResponse) GetBody() *CreateSecurityStrategyResponseBody {
	return s.Body
}

func (s *CreateSecurityStrategyResponse) SetHeaders(v map[string]*string) *CreateSecurityStrategyResponse {
	s.Headers = v
	return s
}

func (s *CreateSecurityStrategyResponse) SetStatusCode(v int32) *CreateSecurityStrategyResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateSecurityStrategyResponse) SetBody(v *CreateSecurityStrategyResponseBody) *CreateSecurityStrategyResponse {
	s.Body = v
	return s
}

func (s *CreateSecurityStrategyResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
