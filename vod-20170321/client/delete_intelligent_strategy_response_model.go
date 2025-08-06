// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteIntelligentStrategyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteIntelligentStrategyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteIntelligentStrategyResponse
	GetStatusCode() *int32
	SetBody(v *DeleteIntelligentStrategyResponseBody) *DeleteIntelligentStrategyResponse
	GetBody() *DeleteIntelligentStrategyResponseBody
}

type DeleteIntelligentStrategyResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteIntelligentStrategyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteIntelligentStrategyResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteIntelligentStrategyResponse) GoString() string {
	return s.String()
}

func (s *DeleteIntelligentStrategyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteIntelligentStrategyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteIntelligentStrategyResponse) GetBody() *DeleteIntelligentStrategyResponseBody {
	return s.Body
}

func (s *DeleteIntelligentStrategyResponse) SetHeaders(v map[string]*string) *DeleteIntelligentStrategyResponse {
	s.Headers = v
	return s
}

func (s *DeleteIntelligentStrategyResponse) SetStatusCode(v int32) *DeleteIntelligentStrategyResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteIntelligentStrategyResponse) SetBody(v *DeleteIntelligentStrategyResponseBody) *DeleteIntelligentStrategyResponse {
	s.Body = v
	return s
}

func (s *DeleteIntelligentStrategyResponse) Validate() error {
	return dara.Validate(s)
}
