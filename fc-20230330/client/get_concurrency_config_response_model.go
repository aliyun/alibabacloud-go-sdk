// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetConcurrencyConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetConcurrencyConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetConcurrencyConfigResponse
	GetStatusCode() *int32
	SetBody(v *ConcurrencyConfig) *GetConcurrencyConfigResponse
	GetBody() *ConcurrencyConfig
}

type GetConcurrencyConfigResponse struct {
	Headers    map[string]*string `json:"headers" xml:"headers"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ConcurrencyConfig `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetConcurrencyConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s GetConcurrencyConfigResponse) GoString() string {
	return s.String()
}

func (s *GetConcurrencyConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetConcurrencyConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetConcurrencyConfigResponse) GetBody() *ConcurrencyConfig {
	return s.Body
}

func (s *GetConcurrencyConfigResponse) SetHeaders(v map[string]*string) *GetConcurrencyConfigResponse {
	s.Headers = v
	return s
}

func (s *GetConcurrencyConfigResponse) SetStatusCode(v int32) *GetConcurrencyConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *GetConcurrencyConfigResponse) SetBody(v *ConcurrencyConfig) *GetConcurrencyConfigResponse {
	s.Body = v
	return s
}

func (s *GetConcurrencyConfigResponse) Validate() error {
	return dara.Validate(s)
}
