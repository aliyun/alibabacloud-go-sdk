// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateKnowledgeBaseFileShardingStrategyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateKnowledgeBaseFileShardingStrategyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateKnowledgeBaseFileShardingStrategyResponse
	GetStatusCode() *int32
	SetBody(v *UpdateKnowledgeBaseFileShardingStrategyResponseBody) *UpdateKnowledgeBaseFileShardingStrategyResponse
	GetBody() *UpdateKnowledgeBaseFileShardingStrategyResponseBody
}

type UpdateKnowledgeBaseFileShardingStrategyResponse struct {
	Headers    map[string]*string                                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateKnowledgeBaseFileShardingStrategyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateKnowledgeBaseFileShardingStrategyResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateKnowledgeBaseFileShardingStrategyResponse) GoString() string {
	return s.String()
}

func (s *UpdateKnowledgeBaseFileShardingStrategyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateKnowledgeBaseFileShardingStrategyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateKnowledgeBaseFileShardingStrategyResponse) GetBody() *UpdateKnowledgeBaseFileShardingStrategyResponseBody {
	return s.Body
}

func (s *UpdateKnowledgeBaseFileShardingStrategyResponse) SetHeaders(v map[string]*string) *UpdateKnowledgeBaseFileShardingStrategyResponse {
	s.Headers = v
	return s
}

func (s *UpdateKnowledgeBaseFileShardingStrategyResponse) SetStatusCode(v int32) *UpdateKnowledgeBaseFileShardingStrategyResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateKnowledgeBaseFileShardingStrategyResponse) SetBody(v *UpdateKnowledgeBaseFileShardingStrategyResponseBody) *UpdateKnowledgeBaseFileShardingStrategyResponse {
	s.Body = v
	return s
}

func (s *UpdateKnowledgeBaseFileShardingStrategyResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
