// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchDeleteRouteStrategyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *BatchDeleteRouteStrategyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *BatchDeleteRouteStrategyResponse
	GetStatusCode() *int32
	SetBody(v *BatchDeleteRouteStrategyResponseBody) *BatchDeleteRouteStrategyResponse
	GetBody() *BatchDeleteRouteStrategyResponseBody
}

type BatchDeleteRouteStrategyResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *BatchDeleteRouteStrategyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s BatchDeleteRouteStrategyResponse) String() string {
	return dara.Prettify(s)
}

func (s BatchDeleteRouteStrategyResponse) GoString() string {
	return s.String()
}

func (s *BatchDeleteRouteStrategyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *BatchDeleteRouteStrategyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *BatchDeleteRouteStrategyResponse) GetBody() *BatchDeleteRouteStrategyResponseBody {
	return s.Body
}

func (s *BatchDeleteRouteStrategyResponse) SetHeaders(v map[string]*string) *BatchDeleteRouteStrategyResponse {
	s.Headers = v
	return s
}

func (s *BatchDeleteRouteStrategyResponse) SetStatusCode(v int32) *BatchDeleteRouteStrategyResponse {
	s.StatusCode = &v
	return s
}

func (s *BatchDeleteRouteStrategyResponse) SetBody(v *BatchDeleteRouteStrategyResponseBody) *BatchDeleteRouteStrategyResponse {
	s.Body = v
	return s
}

func (s *BatchDeleteRouteStrategyResponse) Validate() error {
	return dara.Validate(s)
}
