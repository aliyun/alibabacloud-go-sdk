// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteConcurrencyConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteConcurrencyConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteConcurrencyConfigResponse
	GetStatusCode() *int32
}

type DeleteConcurrencyConfigResponse struct {
	Headers    map[string]*string `json:"headers" xml:"headers"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
}

func (s DeleteConcurrencyConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteConcurrencyConfigResponse) GoString() string {
	return s.String()
}

func (s *DeleteConcurrencyConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteConcurrencyConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteConcurrencyConfigResponse) SetHeaders(v map[string]*string) *DeleteConcurrencyConfigResponse {
	s.Headers = v
	return s
}

func (s *DeleteConcurrencyConfigResponse) SetStatusCode(v int32) *DeleteConcurrencyConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteConcurrencyConfigResponse) Validate() error {
	return dara.Validate(s)
}
