// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAnswerSqlSyntaxByMetaAgentResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AnswerSqlSyntaxByMetaAgentResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AnswerSqlSyntaxByMetaAgentResponse
	GetStatusCode() *int32
	SetBody(v *AnswerSqlSyntaxByMetaAgentResponseBody) *AnswerSqlSyntaxByMetaAgentResponse
	GetBody() *AnswerSqlSyntaxByMetaAgentResponseBody
}

type AnswerSqlSyntaxByMetaAgentResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AnswerSqlSyntaxByMetaAgentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AnswerSqlSyntaxByMetaAgentResponse) String() string {
	return dara.Prettify(s)
}

func (s AnswerSqlSyntaxByMetaAgentResponse) GoString() string {
	return s.String()
}

func (s *AnswerSqlSyntaxByMetaAgentResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AnswerSqlSyntaxByMetaAgentResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AnswerSqlSyntaxByMetaAgentResponse) GetBody() *AnswerSqlSyntaxByMetaAgentResponseBody {
	return s.Body
}

func (s *AnswerSqlSyntaxByMetaAgentResponse) SetHeaders(v map[string]*string) *AnswerSqlSyntaxByMetaAgentResponse {
	s.Headers = v
	return s
}

func (s *AnswerSqlSyntaxByMetaAgentResponse) SetStatusCode(v int32) *AnswerSqlSyntaxByMetaAgentResponse {
	s.StatusCode = &v
	return s
}

func (s *AnswerSqlSyntaxByMetaAgentResponse) SetBody(v *AnswerSqlSyntaxByMetaAgentResponseBody) *AnswerSqlSyntaxByMetaAgentResponse {
	s.Body = v
	return s
}

func (s *AnswerSqlSyntaxByMetaAgentResponse) Validate() error {
	return dara.Validate(s)
}
