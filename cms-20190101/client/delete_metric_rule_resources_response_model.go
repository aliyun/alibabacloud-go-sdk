// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteMetricRuleResourcesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteMetricRuleResourcesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteMetricRuleResourcesResponse
	GetStatusCode() *int32
	SetBody(v *DeleteMetricRuleResourcesResponseBody) *DeleteMetricRuleResourcesResponse
	GetBody() *DeleteMetricRuleResourcesResponseBody
}

type DeleteMetricRuleResourcesResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteMetricRuleResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteMetricRuleResourcesResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteMetricRuleResourcesResponse) GoString() string {
	return s.String()
}

func (s *DeleteMetricRuleResourcesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteMetricRuleResourcesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteMetricRuleResourcesResponse) GetBody() *DeleteMetricRuleResourcesResponseBody {
	return s.Body
}

func (s *DeleteMetricRuleResourcesResponse) SetHeaders(v map[string]*string) *DeleteMetricRuleResourcesResponse {
	s.Headers = v
	return s
}

func (s *DeleteMetricRuleResourcesResponse) SetStatusCode(v int32) *DeleteMetricRuleResourcesResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteMetricRuleResourcesResponse) SetBody(v *DeleteMetricRuleResourcesResponseBody) *DeleteMetricRuleResourcesResponse {
	s.Body = v
	return s
}

func (s *DeleteMetricRuleResourcesResponse) Validate() error {
	return dara.Validate(s)
}
