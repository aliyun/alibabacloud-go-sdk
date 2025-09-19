// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOptimizeSqlByMetaAgentResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *OptimizeSqlByMetaAgentResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *OptimizeSqlByMetaAgentResponse
	GetStatusCode() *int32
	SetBody(v *OptimizeSqlByMetaAgentResponseBody) *OptimizeSqlByMetaAgentResponse
	GetBody() *OptimizeSqlByMetaAgentResponseBody
}

type OptimizeSqlByMetaAgentResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *OptimizeSqlByMetaAgentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s OptimizeSqlByMetaAgentResponse) String() string {
	return dara.Prettify(s)
}

func (s OptimizeSqlByMetaAgentResponse) GoString() string {
	return s.String()
}

func (s *OptimizeSqlByMetaAgentResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *OptimizeSqlByMetaAgentResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *OptimizeSqlByMetaAgentResponse) GetBody() *OptimizeSqlByMetaAgentResponseBody {
	return s.Body
}

func (s *OptimizeSqlByMetaAgentResponse) SetHeaders(v map[string]*string) *OptimizeSqlByMetaAgentResponse {
	s.Headers = v
	return s
}

func (s *OptimizeSqlByMetaAgentResponse) SetStatusCode(v int32) *OptimizeSqlByMetaAgentResponse {
	s.StatusCode = &v
	return s
}

func (s *OptimizeSqlByMetaAgentResponse) SetBody(v *OptimizeSqlByMetaAgentResponseBody) *OptimizeSqlByMetaAgentResponse {
	s.Body = v
	return s
}

func (s *OptimizeSqlByMetaAgentResponse) Validate() error {
	return dara.Validate(s)
}
