// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableElasticPlanResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DisableElasticPlanResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DisableElasticPlanResponse
	GetStatusCode() *int32
	SetBody(v *DisableElasticPlanResponseBody) *DisableElasticPlanResponse
	GetBody() *DisableElasticPlanResponseBody
}

type DisableElasticPlanResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DisableElasticPlanResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DisableElasticPlanResponse) String() string {
	return dara.Prettify(s)
}

func (s DisableElasticPlanResponse) GoString() string {
	return s.String()
}

func (s *DisableElasticPlanResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DisableElasticPlanResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DisableElasticPlanResponse) GetBody() *DisableElasticPlanResponseBody {
	return s.Body
}

func (s *DisableElasticPlanResponse) SetHeaders(v map[string]*string) *DisableElasticPlanResponse {
	s.Headers = v
	return s
}

func (s *DisableElasticPlanResponse) SetStatusCode(v int32) *DisableElasticPlanResponse {
	s.StatusCode = &v
	return s
}

func (s *DisableElasticPlanResponse) SetBody(v *DisableElasticPlanResponseBody) *DisableElasticPlanResponse {
	s.Body = v
	return s
}

func (s *DisableElasticPlanResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
