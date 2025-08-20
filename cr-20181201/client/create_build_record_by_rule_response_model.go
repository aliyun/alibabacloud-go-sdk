// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateBuildRecordByRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateBuildRecordByRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateBuildRecordByRuleResponse
	GetStatusCode() *int32
	SetBody(v *CreateBuildRecordByRuleResponseBody) *CreateBuildRecordByRuleResponse
	GetBody() *CreateBuildRecordByRuleResponseBody
}

type CreateBuildRecordByRuleResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateBuildRecordByRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateBuildRecordByRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateBuildRecordByRuleResponse) GoString() string {
	return s.String()
}

func (s *CreateBuildRecordByRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateBuildRecordByRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateBuildRecordByRuleResponse) GetBody() *CreateBuildRecordByRuleResponseBody {
	return s.Body
}

func (s *CreateBuildRecordByRuleResponse) SetHeaders(v map[string]*string) *CreateBuildRecordByRuleResponse {
	s.Headers = v
	return s
}

func (s *CreateBuildRecordByRuleResponse) SetStatusCode(v int32) *CreateBuildRecordByRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateBuildRecordByRuleResponse) SetBody(v *CreateBuildRecordByRuleResponseBody) *CreateBuildRecordByRuleResponse {
	s.Body = v
	return s
}

func (s *CreateBuildRecordByRuleResponse) Validate() error {
	return dara.Validate(s)
}
