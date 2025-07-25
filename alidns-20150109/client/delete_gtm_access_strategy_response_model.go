// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteGtmAccessStrategyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteGtmAccessStrategyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteGtmAccessStrategyResponse
	GetStatusCode() *int32
	SetBody(v *DeleteGtmAccessStrategyResponseBody) *DeleteGtmAccessStrategyResponse
	GetBody() *DeleteGtmAccessStrategyResponseBody
}

type DeleteGtmAccessStrategyResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteGtmAccessStrategyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteGtmAccessStrategyResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteGtmAccessStrategyResponse) GoString() string {
	return s.String()
}

func (s *DeleteGtmAccessStrategyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteGtmAccessStrategyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteGtmAccessStrategyResponse) GetBody() *DeleteGtmAccessStrategyResponseBody {
	return s.Body
}

func (s *DeleteGtmAccessStrategyResponse) SetHeaders(v map[string]*string) *DeleteGtmAccessStrategyResponse {
	s.Headers = v
	return s
}

func (s *DeleteGtmAccessStrategyResponse) SetStatusCode(v int32) *DeleteGtmAccessStrategyResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteGtmAccessStrategyResponse) SetBody(v *DeleteGtmAccessStrategyResponseBody) *DeleteGtmAccessStrategyResponse {
	s.Body = v
	return s
}

func (s *DeleteGtmAccessStrategyResponse) Validate() error {
	return dara.Validate(s)
}
