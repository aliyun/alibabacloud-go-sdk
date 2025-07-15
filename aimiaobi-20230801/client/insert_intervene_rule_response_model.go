// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInsertInterveneRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *InsertInterveneRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *InsertInterveneRuleResponse
	GetStatusCode() *int32
	SetBody(v *InsertInterveneRuleResponseBody) *InsertInterveneRuleResponse
	GetBody() *InsertInterveneRuleResponseBody
}

type InsertInterveneRuleResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *InsertInterveneRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s InsertInterveneRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s InsertInterveneRuleResponse) GoString() string {
	return s.String()
}

func (s *InsertInterveneRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *InsertInterveneRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *InsertInterveneRuleResponse) GetBody() *InsertInterveneRuleResponseBody {
	return s.Body
}

func (s *InsertInterveneRuleResponse) SetHeaders(v map[string]*string) *InsertInterveneRuleResponse {
	s.Headers = v
	return s
}

func (s *InsertInterveneRuleResponse) SetStatusCode(v int32) *InsertInterveneRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *InsertInterveneRuleResponse) SetBody(v *InsertInterveneRuleResponseBody) *InsertInterveneRuleResponse {
	s.Body = v
	return s
}

func (s *InsertInterveneRuleResponse) Validate() error {
	return dara.Validate(s)
}
