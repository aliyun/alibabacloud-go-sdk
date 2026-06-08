// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSecurityStrategyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetSecurityStrategyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetSecurityStrategyResponse
	GetStatusCode() *int32
	SetBody(v *GetSecurityStrategyResponseBody) *GetSecurityStrategyResponse
	GetBody() *GetSecurityStrategyResponseBody
}

type GetSecurityStrategyResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetSecurityStrategyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetSecurityStrategyResponse) String() string {
	return dara.Prettify(s)
}

func (s GetSecurityStrategyResponse) GoString() string {
	return s.String()
}

func (s *GetSecurityStrategyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetSecurityStrategyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetSecurityStrategyResponse) GetBody() *GetSecurityStrategyResponseBody {
	return s.Body
}

func (s *GetSecurityStrategyResponse) SetHeaders(v map[string]*string) *GetSecurityStrategyResponse {
	s.Headers = v
	return s
}

func (s *GetSecurityStrategyResponse) SetStatusCode(v int32) *GetSecurityStrategyResponse {
	s.StatusCode = &v
	return s
}

func (s *GetSecurityStrategyResponse) SetBody(v *GetSecurityStrategyResponseBody) *GetSecurityStrategyResponse {
	s.Body = v
	return s
}

func (s *GetSecurityStrategyResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
