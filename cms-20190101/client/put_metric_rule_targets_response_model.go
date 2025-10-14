// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPutMetricRuleTargetsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *PutMetricRuleTargetsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *PutMetricRuleTargetsResponse
	GetStatusCode() *int32
	SetBody(v *PutMetricRuleTargetsResponseBody) *PutMetricRuleTargetsResponse
	GetBody() *PutMetricRuleTargetsResponseBody
}

type PutMetricRuleTargetsResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PutMetricRuleTargetsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PutMetricRuleTargetsResponse) String() string {
	return dara.Prettify(s)
}

func (s PutMetricRuleTargetsResponse) GoString() string {
	return s.String()
}

func (s *PutMetricRuleTargetsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *PutMetricRuleTargetsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *PutMetricRuleTargetsResponse) GetBody() *PutMetricRuleTargetsResponseBody {
	return s.Body
}

func (s *PutMetricRuleTargetsResponse) SetHeaders(v map[string]*string) *PutMetricRuleTargetsResponse {
	s.Headers = v
	return s
}

func (s *PutMetricRuleTargetsResponse) SetStatusCode(v int32) *PutMetricRuleTargetsResponse {
	s.StatusCode = &v
	return s
}

func (s *PutMetricRuleTargetsResponse) SetBody(v *PutMetricRuleTargetsResponseBody) *PutMetricRuleTargetsResponse {
	s.Body = v
	return s
}

func (s *PutMetricRuleTargetsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
