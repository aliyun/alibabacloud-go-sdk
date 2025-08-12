// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteMetricRuleTargetsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteMetricRuleTargetsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteMetricRuleTargetsResponse
	GetStatusCode() *int32
	SetBody(v *DeleteMetricRuleTargetsResponseBody) *DeleteMetricRuleTargetsResponse
	GetBody() *DeleteMetricRuleTargetsResponseBody
}

type DeleteMetricRuleTargetsResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteMetricRuleTargetsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteMetricRuleTargetsResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteMetricRuleTargetsResponse) GoString() string {
	return s.String()
}

func (s *DeleteMetricRuleTargetsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteMetricRuleTargetsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteMetricRuleTargetsResponse) GetBody() *DeleteMetricRuleTargetsResponseBody {
	return s.Body
}

func (s *DeleteMetricRuleTargetsResponse) SetHeaders(v map[string]*string) *DeleteMetricRuleTargetsResponse {
	s.Headers = v
	return s
}

func (s *DeleteMetricRuleTargetsResponse) SetStatusCode(v int32) *DeleteMetricRuleTargetsResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteMetricRuleTargetsResponse) SetBody(v *DeleteMetricRuleTargetsResponseBody) *DeleteMetricRuleTargetsResponse {
	s.Body = v
	return s
}

func (s *DeleteMetricRuleTargetsResponse) Validate() error {
	return dara.Validate(s)
}
