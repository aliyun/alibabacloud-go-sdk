// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteOpaStrategyNewResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteOpaStrategyNewResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteOpaStrategyNewResponse
	GetStatusCode() *int32
	SetBody(v *DeleteOpaStrategyNewResponseBody) *DeleteOpaStrategyNewResponse
	GetBody() *DeleteOpaStrategyNewResponseBody
}

type DeleteOpaStrategyNewResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteOpaStrategyNewResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteOpaStrategyNewResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteOpaStrategyNewResponse) GoString() string {
	return s.String()
}

func (s *DeleteOpaStrategyNewResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteOpaStrategyNewResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteOpaStrategyNewResponse) GetBody() *DeleteOpaStrategyNewResponseBody {
	return s.Body
}

func (s *DeleteOpaStrategyNewResponse) SetHeaders(v map[string]*string) *DeleteOpaStrategyNewResponse {
	s.Headers = v
	return s
}

func (s *DeleteOpaStrategyNewResponse) SetStatusCode(v int32) *DeleteOpaStrategyNewResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteOpaStrategyNewResponse) SetBody(v *DeleteOpaStrategyNewResponseBody) *DeleteOpaStrategyNewResponse {
	s.Body = v
	return s
}

func (s *DeleteOpaStrategyNewResponse) Validate() error {
	return dara.Validate(s)
}
