// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateElasticPlanResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateElasticPlanResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateElasticPlanResponse
	GetStatusCode() *int32
	SetBody(v *CreateElasticPlanResponseBody) *CreateElasticPlanResponse
	GetBody() *CreateElasticPlanResponseBody
}

type CreateElasticPlanResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateElasticPlanResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateElasticPlanResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateElasticPlanResponse) GoString() string {
	return s.String()
}

func (s *CreateElasticPlanResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateElasticPlanResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateElasticPlanResponse) GetBody() *CreateElasticPlanResponseBody {
	return s.Body
}

func (s *CreateElasticPlanResponse) SetHeaders(v map[string]*string) *CreateElasticPlanResponse {
	s.Headers = v
	return s
}

func (s *CreateElasticPlanResponse) SetStatusCode(v int32) *CreateElasticPlanResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateElasticPlanResponse) SetBody(v *CreateElasticPlanResponseBody) *CreateElasticPlanResponse {
	s.Body = v
	return s
}

func (s *CreateElasticPlanResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
