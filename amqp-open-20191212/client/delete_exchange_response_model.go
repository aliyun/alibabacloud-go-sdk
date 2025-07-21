// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteExchangeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteExchangeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteExchangeResponse
	GetStatusCode() *int32
	SetBody(v *DeleteExchangeResponseBody) *DeleteExchangeResponse
	GetBody() *DeleteExchangeResponseBody
}

type DeleteExchangeResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteExchangeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteExchangeResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteExchangeResponse) GoString() string {
	return s.String()
}

func (s *DeleteExchangeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteExchangeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteExchangeResponse) GetBody() *DeleteExchangeResponseBody {
	return s.Body
}

func (s *DeleteExchangeResponse) SetHeaders(v map[string]*string) *DeleteExchangeResponse {
	s.Headers = v
	return s
}

func (s *DeleteExchangeResponse) SetStatusCode(v int32) *DeleteExchangeResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteExchangeResponse) SetBody(v *DeleteExchangeResponseBody) *DeleteExchangeResponse {
	s.Body = v
	return s
}

func (s *DeleteExchangeResponse) Validate() error {
	return dara.Validate(s)
}
