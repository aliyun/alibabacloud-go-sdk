// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListIntelligentStrategyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListIntelligentStrategyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListIntelligentStrategyResponse
	GetStatusCode() *int32
	SetBody(v *ListIntelligentStrategyResponseBody) *ListIntelligentStrategyResponse
	GetBody() *ListIntelligentStrategyResponseBody
}

type ListIntelligentStrategyResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListIntelligentStrategyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListIntelligentStrategyResponse) String() string {
	return dara.Prettify(s)
}

func (s ListIntelligentStrategyResponse) GoString() string {
	return s.String()
}

func (s *ListIntelligentStrategyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListIntelligentStrategyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListIntelligentStrategyResponse) GetBody() *ListIntelligentStrategyResponseBody {
	return s.Body
}

func (s *ListIntelligentStrategyResponse) SetHeaders(v map[string]*string) *ListIntelligentStrategyResponse {
	s.Headers = v
	return s
}

func (s *ListIntelligentStrategyResponse) SetStatusCode(v int32) *ListIntelligentStrategyResponse {
	s.StatusCode = &v
	return s
}

func (s *ListIntelligentStrategyResponse) SetBody(v *ListIntelligentStrategyResponseBody) *ListIntelligentStrategyResponse {
	s.Body = v
	return s
}

func (s *ListIntelligentStrategyResponse) Validate() error {
	return dara.Validate(s)
}
