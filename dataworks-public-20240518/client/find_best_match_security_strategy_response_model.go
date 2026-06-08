// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFindBestMatchSecurityStrategyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *FindBestMatchSecurityStrategyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *FindBestMatchSecurityStrategyResponse
	GetStatusCode() *int32
	SetBody(v *FindBestMatchSecurityStrategyResponseBody) *FindBestMatchSecurityStrategyResponse
	GetBody() *FindBestMatchSecurityStrategyResponseBody
}

type FindBestMatchSecurityStrategyResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *FindBestMatchSecurityStrategyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s FindBestMatchSecurityStrategyResponse) String() string {
	return dara.Prettify(s)
}

func (s FindBestMatchSecurityStrategyResponse) GoString() string {
	return s.String()
}

func (s *FindBestMatchSecurityStrategyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *FindBestMatchSecurityStrategyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *FindBestMatchSecurityStrategyResponse) GetBody() *FindBestMatchSecurityStrategyResponseBody {
	return s.Body
}

func (s *FindBestMatchSecurityStrategyResponse) SetHeaders(v map[string]*string) *FindBestMatchSecurityStrategyResponse {
	s.Headers = v
	return s
}

func (s *FindBestMatchSecurityStrategyResponse) SetStatusCode(v int32) *FindBestMatchSecurityStrategyResponse {
	s.StatusCode = &v
	return s
}

func (s *FindBestMatchSecurityStrategyResponse) SetBody(v *FindBestMatchSecurityStrategyResponseBody) *FindBestMatchSecurityStrategyResponse {
	s.Body = v
	return s
}

func (s *FindBestMatchSecurityStrategyResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
