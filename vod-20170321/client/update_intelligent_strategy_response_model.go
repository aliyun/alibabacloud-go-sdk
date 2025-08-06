// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateIntelligentStrategyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateIntelligentStrategyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateIntelligentStrategyResponse
	GetStatusCode() *int32
	SetBody(v *UpdateIntelligentStrategyResponseBody) *UpdateIntelligentStrategyResponse
	GetBody() *UpdateIntelligentStrategyResponseBody
}

type UpdateIntelligentStrategyResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateIntelligentStrategyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateIntelligentStrategyResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateIntelligentStrategyResponse) GoString() string {
	return s.String()
}

func (s *UpdateIntelligentStrategyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateIntelligentStrategyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateIntelligentStrategyResponse) GetBody() *UpdateIntelligentStrategyResponseBody {
	return s.Body
}

func (s *UpdateIntelligentStrategyResponse) SetHeaders(v map[string]*string) *UpdateIntelligentStrategyResponse {
	s.Headers = v
	return s
}

func (s *UpdateIntelligentStrategyResponse) SetStatusCode(v int32) *UpdateIntelligentStrategyResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateIntelligentStrategyResponse) SetBody(v *UpdateIntelligentStrategyResponseBody) *UpdateIntelligentStrategyResponse {
	s.Body = v
	return s
}

func (s *UpdateIntelligentStrategyResponse) Validate() error {
	return dara.Validate(s)
}
