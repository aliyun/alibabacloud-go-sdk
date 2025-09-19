// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyClientConfStrategyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyClientConfStrategyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyClientConfStrategyResponse
	GetStatusCode() *int32
	SetBody(v *ModifyClientConfStrategyResponseBody) *ModifyClientConfStrategyResponse
	GetBody() *ModifyClientConfStrategyResponseBody
}

type ModifyClientConfStrategyResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyClientConfStrategyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyClientConfStrategyResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyClientConfStrategyResponse) GoString() string {
	return s.String()
}

func (s *ModifyClientConfStrategyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyClientConfStrategyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyClientConfStrategyResponse) GetBody() *ModifyClientConfStrategyResponseBody {
	return s.Body
}

func (s *ModifyClientConfStrategyResponse) SetHeaders(v map[string]*string) *ModifyClientConfStrategyResponse {
	s.Headers = v
	return s
}

func (s *ModifyClientConfStrategyResponse) SetStatusCode(v int32) *ModifyClientConfStrategyResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyClientConfStrategyResponse) SetBody(v *ModifyClientConfStrategyResponseBody) *ModifyClientConfStrategyResponse {
	s.Body = v
	return s
}

func (s *ModifyClientConfStrategyResponse) Validate() error {
	return dara.Validate(s)
}
