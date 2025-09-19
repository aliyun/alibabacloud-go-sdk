// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteStrategyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteStrategyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteStrategyResponse
	GetStatusCode() *int32
	SetBody(v *DeleteStrategyResponseBody) *DeleteStrategyResponse
	GetBody() *DeleteStrategyResponseBody
}

type DeleteStrategyResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteStrategyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteStrategyResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteStrategyResponse) GoString() string {
	return s.String()
}

func (s *DeleteStrategyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteStrategyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteStrategyResponse) GetBody() *DeleteStrategyResponseBody {
	return s.Body
}

func (s *DeleteStrategyResponse) SetHeaders(v map[string]*string) *DeleteStrategyResponse {
	s.Headers = v
	return s
}

func (s *DeleteStrategyResponse) SetStatusCode(v int32) *DeleteStrategyResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteStrategyResponse) SetBody(v *DeleteStrategyResponseBody) *DeleteStrategyResponse {
	s.Body = v
	return s
}

func (s *DeleteStrategyResponse) Validate() error {
	return dara.Validate(s)
}
