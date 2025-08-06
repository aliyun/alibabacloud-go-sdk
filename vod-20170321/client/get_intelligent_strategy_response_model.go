// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetIntelligentStrategyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetIntelligentStrategyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetIntelligentStrategyResponse
	GetStatusCode() *int32
	SetBody(v *GetIntelligentStrategyResponseBody) *GetIntelligentStrategyResponse
	GetBody() *GetIntelligentStrategyResponseBody
}

type GetIntelligentStrategyResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetIntelligentStrategyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetIntelligentStrategyResponse) String() string {
	return dara.Prettify(s)
}

func (s GetIntelligentStrategyResponse) GoString() string {
	return s.String()
}

func (s *GetIntelligentStrategyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetIntelligentStrategyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetIntelligentStrategyResponse) GetBody() *GetIntelligentStrategyResponseBody {
	return s.Body
}

func (s *GetIntelligentStrategyResponse) SetHeaders(v map[string]*string) *GetIntelligentStrategyResponse {
	s.Headers = v
	return s
}

func (s *GetIntelligentStrategyResponse) SetStatusCode(v int32) *GetIntelligentStrategyResponse {
	s.StatusCode = &v
	return s
}

func (s *GetIntelligentStrategyResponse) SetBody(v *GetIntelligentStrategyResponseBody) *GetIntelligentStrategyResponse {
	s.Body = v
	return s
}

func (s *GetIntelligentStrategyResponse) Validate() error {
	return dara.Validate(s)
}
