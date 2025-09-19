// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFixSqlByMetaAgentResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *FixSqlByMetaAgentResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *FixSqlByMetaAgentResponse
	GetStatusCode() *int32
	SetBody(v *FixSqlByMetaAgentResponseBody) *FixSqlByMetaAgentResponse
	GetBody() *FixSqlByMetaAgentResponseBody
}

type FixSqlByMetaAgentResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *FixSqlByMetaAgentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s FixSqlByMetaAgentResponse) String() string {
	return dara.Prettify(s)
}

func (s FixSqlByMetaAgentResponse) GoString() string {
	return s.String()
}

func (s *FixSqlByMetaAgentResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *FixSqlByMetaAgentResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *FixSqlByMetaAgentResponse) GetBody() *FixSqlByMetaAgentResponseBody {
	return s.Body
}

func (s *FixSqlByMetaAgentResponse) SetHeaders(v map[string]*string) *FixSqlByMetaAgentResponse {
	s.Headers = v
	return s
}

func (s *FixSqlByMetaAgentResponse) SetStatusCode(v int32) *FixSqlByMetaAgentResponse {
	s.StatusCode = &v
	return s
}

func (s *FixSqlByMetaAgentResponse) SetBody(v *FixSqlByMetaAgentResponseBody) *FixSqlByMetaAgentResponse {
	s.Body = v
	return s
}

func (s *FixSqlByMetaAgentResponse) Validate() error {
	return dara.Validate(s)
}
