// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDisposeStrategyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListDisposeStrategyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListDisposeStrategyResponse
	GetStatusCode() *int32
	SetBody(v *ListDisposeStrategyResponseBody) *ListDisposeStrategyResponse
	GetBody() *ListDisposeStrategyResponseBody
}

type ListDisposeStrategyResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListDisposeStrategyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListDisposeStrategyResponse) String() string {
	return dara.Prettify(s)
}

func (s ListDisposeStrategyResponse) GoString() string {
	return s.String()
}

func (s *ListDisposeStrategyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListDisposeStrategyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListDisposeStrategyResponse) GetBody() *ListDisposeStrategyResponseBody {
	return s.Body
}

func (s *ListDisposeStrategyResponse) SetHeaders(v map[string]*string) *ListDisposeStrategyResponse {
	s.Headers = v
	return s
}

func (s *ListDisposeStrategyResponse) SetStatusCode(v int32) *ListDisposeStrategyResponse {
	s.StatusCode = &v
	return s
}

func (s *ListDisposeStrategyResponse) SetBody(v *ListDisposeStrategyResponseBody) *ListDisposeStrategyResponse {
	s.Body = v
	return s
}

func (s *ListDisposeStrategyResponse) Validate() error {
	return dara.Validate(s)
}
