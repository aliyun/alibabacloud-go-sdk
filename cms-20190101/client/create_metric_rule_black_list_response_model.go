// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateMetricRuleBlackListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateMetricRuleBlackListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateMetricRuleBlackListResponse
	GetStatusCode() *int32
	SetBody(v *CreateMetricRuleBlackListResponseBody) *CreateMetricRuleBlackListResponse
	GetBody() *CreateMetricRuleBlackListResponseBody
}

type CreateMetricRuleBlackListResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateMetricRuleBlackListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateMetricRuleBlackListResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateMetricRuleBlackListResponse) GoString() string {
	return s.String()
}

func (s *CreateMetricRuleBlackListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateMetricRuleBlackListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateMetricRuleBlackListResponse) GetBody() *CreateMetricRuleBlackListResponseBody {
	return s.Body
}

func (s *CreateMetricRuleBlackListResponse) SetHeaders(v map[string]*string) *CreateMetricRuleBlackListResponse {
	s.Headers = v
	return s
}

func (s *CreateMetricRuleBlackListResponse) SetStatusCode(v int32) *CreateMetricRuleBlackListResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateMetricRuleBlackListResponse) SetBody(v *CreateMetricRuleBlackListResponseBody) *CreateMetricRuleBlackListResponse {
	s.Body = v
	return s
}

func (s *CreateMetricRuleBlackListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
