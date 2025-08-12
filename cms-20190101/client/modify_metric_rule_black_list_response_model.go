// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyMetricRuleBlackListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyMetricRuleBlackListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyMetricRuleBlackListResponse
	GetStatusCode() *int32
	SetBody(v *ModifyMetricRuleBlackListResponseBody) *ModifyMetricRuleBlackListResponse
	GetBody() *ModifyMetricRuleBlackListResponseBody
}

type ModifyMetricRuleBlackListResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyMetricRuleBlackListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyMetricRuleBlackListResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyMetricRuleBlackListResponse) GoString() string {
	return s.String()
}

func (s *ModifyMetricRuleBlackListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyMetricRuleBlackListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyMetricRuleBlackListResponse) GetBody() *ModifyMetricRuleBlackListResponseBody {
	return s.Body
}

func (s *ModifyMetricRuleBlackListResponse) SetHeaders(v map[string]*string) *ModifyMetricRuleBlackListResponse {
	s.Headers = v
	return s
}

func (s *ModifyMetricRuleBlackListResponse) SetStatusCode(v int32) *ModifyMetricRuleBlackListResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyMetricRuleBlackListResponse) SetBody(v *ModifyMetricRuleBlackListResponseBody) *ModifyMetricRuleBlackListResponse {
	s.Body = v
	return s
}

func (s *ModifyMetricRuleBlackListResponse) Validate() error {
	return dara.Validate(s)
}
