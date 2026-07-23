// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateExperimentPlanResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateExperimentPlanResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateExperimentPlanResponse
	GetStatusCode() *int32
	SetBody(v *UpdateExperimentPlanResponseBody) *UpdateExperimentPlanResponse
	GetBody() *UpdateExperimentPlanResponseBody
}

type UpdateExperimentPlanResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateExperimentPlanResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateExperimentPlanResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateExperimentPlanResponse) GoString() string {
	return s.String()
}

func (s *UpdateExperimentPlanResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateExperimentPlanResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateExperimentPlanResponse) GetBody() *UpdateExperimentPlanResponseBody {
	return s.Body
}

func (s *UpdateExperimentPlanResponse) SetHeaders(v map[string]*string) *UpdateExperimentPlanResponse {
	s.Headers = v
	return s
}

func (s *UpdateExperimentPlanResponse) SetStatusCode(v int32) *UpdateExperimentPlanResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateExperimentPlanResponse) SetBody(v *UpdateExperimentPlanResponseBody) *UpdateExperimentPlanResponse {
	s.Body = v
	return s
}

func (s *UpdateExperimentPlanResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
