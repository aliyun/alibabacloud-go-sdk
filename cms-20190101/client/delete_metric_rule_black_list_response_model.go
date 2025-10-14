// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteMetricRuleBlackListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteMetricRuleBlackListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteMetricRuleBlackListResponse
	GetStatusCode() *int32
	SetBody(v *DeleteMetricRuleBlackListResponseBody) *DeleteMetricRuleBlackListResponse
	GetBody() *DeleteMetricRuleBlackListResponseBody
}

type DeleteMetricRuleBlackListResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteMetricRuleBlackListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteMetricRuleBlackListResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteMetricRuleBlackListResponse) GoString() string {
	return s.String()
}

func (s *DeleteMetricRuleBlackListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteMetricRuleBlackListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteMetricRuleBlackListResponse) GetBody() *DeleteMetricRuleBlackListResponseBody {
	return s.Body
}

func (s *DeleteMetricRuleBlackListResponse) SetHeaders(v map[string]*string) *DeleteMetricRuleBlackListResponse {
	s.Headers = v
	return s
}

func (s *DeleteMetricRuleBlackListResponse) SetStatusCode(v int32) *DeleteMetricRuleBlackListResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteMetricRuleBlackListResponse) SetBody(v *DeleteMetricRuleBlackListResponseBody) *DeleteMetricRuleBlackListResponse {
	s.Body = v
	return s
}

func (s *DeleteMetricRuleBlackListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
