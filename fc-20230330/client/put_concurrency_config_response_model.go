// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPutConcurrencyConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *PutConcurrencyConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *PutConcurrencyConfigResponse
	GetStatusCode() *int32
	SetBody(v *ConcurrencyConfig) *PutConcurrencyConfigResponse
	GetBody() *ConcurrencyConfig
}

type PutConcurrencyConfigResponse struct {
	Headers    map[string]*string `json:"headers" xml:"headers"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ConcurrencyConfig `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PutConcurrencyConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s PutConcurrencyConfigResponse) GoString() string {
	return s.String()
}

func (s *PutConcurrencyConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *PutConcurrencyConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *PutConcurrencyConfigResponse) GetBody() *ConcurrencyConfig {
	return s.Body
}

func (s *PutConcurrencyConfigResponse) SetHeaders(v map[string]*string) *PutConcurrencyConfigResponse {
	s.Headers = v
	return s
}

func (s *PutConcurrencyConfigResponse) SetStatusCode(v int32) *PutConcurrencyConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *PutConcurrencyConfigResponse) SetBody(v *ConcurrencyConfig) *PutConcurrencyConfigResponse {
	s.Body = v
	return s
}

func (s *PutConcurrencyConfigResponse) Validate() error {
	return dara.Validate(s)
}
