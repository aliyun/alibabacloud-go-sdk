// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRecallRuleAuditResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RecallRuleAuditResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RecallRuleAuditResponse
	GetStatusCode() *int32
	SetBody(v *RecallRuleAuditResponseBody) *RecallRuleAuditResponse
	GetBody() *RecallRuleAuditResponseBody
}

type RecallRuleAuditResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RecallRuleAuditResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RecallRuleAuditResponse) String() string {
	return dara.Prettify(s)
}

func (s RecallRuleAuditResponse) GoString() string {
	return s.String()
}

func (s *RecallRuleAuditResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RecallRuleAuditResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RecallRuleAuditResponse) GetBody() *RecallRuleAuditResponseBody {
	return s.Body
}

func (s *RecallRuleAuditResponse) SetHeaders(v map[string]*string) *RecallRuleAuditResponse {
	s.Headers = v
	return s
}

func (s *RecallRuleAuditResponse) SetStatusCode(v int32) *RecallRuleAuditResponse {
	s.StatusCode = &v
	return s
}

func (s *RecallRuleAuditResponse) SetBody(v *RecallRuleAuditResponseBody) *RecallRuleAuditResponse {
	s.Body = v
	return s
}

func (s *RecallRuleAuditResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
