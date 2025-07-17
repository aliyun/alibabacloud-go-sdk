// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDataQualityRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetDataQualityRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetDataQualityRuleResponse
	GetStatusCode() *int32
	SetBody(v *GetDataQualityRuleResponseBody) *GetDataQualityRuleResponse
	GetBody() *GetDataQualityRuleResponseBody
}

type GetDataQualityRuleResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetDataQualityRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetDataQualityRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s GetDataQualityRuleResponse) GoString() string {
	return s.String()
}

func (s *GetDataQualityRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetDataQualityRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetDataQualityRuleResponse) GetBody() *GetDataQualityRuleResponseBody {
	return s.Body
}

func (s *GetDataQualityRuleResponse) SetHeaders(v map[string]*string) *GetDataQualityRuleResponse {
	s.Headers = v
	return s
}

func (s *GetDataQualityRuleResponse) SetStatusCode(v int32) *GetDataQualityRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDataQualityRuleResponse) SetBody(v *GetDataQualityRuleResponseBody) *GetDataQualityRuleResponse {
	s.Body = v
	return s
}

func (s *GetDataQualityRuleResponse) Validate() error {
	return dara.Validate(s)
}
