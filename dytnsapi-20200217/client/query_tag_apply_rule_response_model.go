// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryTagApplyRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryTagApplyRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryTagApplyRuleResponse
	GetStatusCode() *int32
	SetBody(v *QueryTagApplyRuleResponseBody) *QueryTagApplyRuleResponse
	GetBody() *QueryTagApplyRuleResponseBody
}

type QueryTagApplyRuleResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryTagApplyRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryTagApplyRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryTagApplyRuleResponse) GoString() string {
	return s.String()
}

func (s *QueryTagApplyRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryTagApplyRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryTagApplyRuleResponse) GetBody() *QueryTagApplyRuleResponseBody {
	return s.Body
}

func (s *QueryTagApplyRuleResponse) SetHeaders(v map[string]*string) *QueryTagApplyRuleResponse {
	s.Headers = v
	return s
}

func (s *QueryTagApplyRuleResponse) SetStatusCode(v int32) *QueryTagApplyRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryTagApplyRuleResponse) SetBody(v *QueryTagApplyRuleResponseBody) *QueryTagApplyRuleResponse {
	s.Body = v
	return s
}

func (s *QueryTagApplyRuleResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
