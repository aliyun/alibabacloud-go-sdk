// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateMetricRuleResourcesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateMetricRuleResourcesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateMetricRuleResourcesResponse
	GetStatusCode() *int32
	SetBody(v *CreateMetricRuleResourcesResponseBody) *CreateMetricRuleResourcesResponse
	GetBody() *CreateMetricRuleResourcesResponseBody
}

type CreateMetricRuleResourcesResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateMetricRuleResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateMetricRuleResourcesResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateMetricRuleResourcesResponse) GoString() string {
	return s.String()
}

func (s *CreateMetricRuleResourcesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateMetricRuleResourcesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateMetricRuleResourcesResponse) GetBody() *CreateMetricRuleResourcesResponseBody {
	return s.Body
}

func (s *CreateMetricRuleResourcesResponse) SetHeaders(v map[string]*string) *CreateMetricRuleResourcesResponse {
	s.Headers = v
	return s
}

func (s *CreateMetricRuleResourcesResponse) SetStatusCode(v int32) *CreateMetricRuleResourcesResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateMetricRuleResourcesResponse) SetBody(v *CreateMetricRuleResourcesResponseBody) *CreateMetricRuleResourcesResponse {
	s.Body = v
	return s
}

func (s *CreateMetricRuleResourcesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
