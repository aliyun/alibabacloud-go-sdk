// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteSecurityStrategyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteSecurityStrategyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteSecurityStrategyResponse
	GetStatusCode() *int32
	SetBody(v *DeleteSecurityStrategyResponseBody) *DeleteSecurityStrategyResponse
	GetBody() *DeleteSecurityStrategyResponseBody
}

type DeleteSecurityStrategyResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteSecurityStrategyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteSecurityStrategyResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteSecurityStrategyResponse) GoString() string {
	return s.String()
}

func (s *DeleteSecurityStrategyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteSecurityStrategyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteSecurityStrategyResponse) GetBody() *DeleteSecurityStrategyResponseBody {
	return s.Body
}

func (s *DeleteSecurityStrategyResponse) SetHeaders(v map[string]*string) *DeleteSecurityStrategyResponse {
	s.Headers = v
	return s
}

func (s *DeleteSecurityStrategyResponse) SetStatusCode(v int32) *DeleteSecurityStrategyResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteSecurityStrategyResponse) SetBody(v *DeleteSecurityStrategyResponseBody) *DeleteSecurityStrategyResponse {
	s.Body = v
	return s
}

func (s *DeleteSecurityStrategyResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
