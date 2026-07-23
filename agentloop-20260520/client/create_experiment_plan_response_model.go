// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateExperimentPlanResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateExperimentPlanResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateExperimentPlanResponse
	GetStatusCode() *int32
	SetBody(v *CreateExperimentPlanResponseBody) *CreateExperimentPlanResponse
	GetBody() *CreateExperimentPlanResponseBody
}

type CreateExperimentPlanResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateExperimentPlanResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateExperimentPlanResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateExperimentPlanResponse) GoString() string {
	return s.String()
}

func (s *CreateExperimentPlanResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateExperimentPlanResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateExperimentPlanResponse) GetBody() *CreateExperimentPlanResponseBody {
	return s.Body
}

func (s *CreateExperimentPlanResponse) SetHeaders(v map[string]*string) *CreateExperimentPlanResponse {
	s.Headers = v
	return s
}

func (s *CreateExperimentPlanResponse) SetStatusCode(v int32) *CreateExperimentPlanResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateExperimentPlanResponse) SetBody(v *CreateExperimentPlanResponseBody) *CreateExperimentPlanResponse {
	s.Body = v
	return s
}

func (s *CreateExperimentPlanResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
