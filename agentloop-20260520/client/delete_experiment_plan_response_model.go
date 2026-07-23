// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteExperimentPlanResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteExperimentPlanResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteExperimentPlanResponse
	GetStatusCode() *int32
	SetBody(v *DeleteExperimentPlanResponseBody) *DeleteExperimentPlanResponse
	GetBody() *DeleteExperimentPlanResponseBody
}

type DeleteExperimentPlanResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteExperimentPlanResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteExperimentPlanResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteExperimentPlanResponse) GoString() string {
	return s.String()
}

func (s *DeleteExperimentPlanResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteExperimentPlanResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteExperimentPlanResponse) GetBody() *DeleteExperimentPlanResponseBody {
	return s.Body
}

func (s *DeleteExperimentPlanResponse) SetHeaders(v map[string]*string) *DeleteExperimentPlanResponse {
	s.Headers = v
	return s
}

func (s *DeleteExperimentPlanResponse) SetStatusCode(v int32) *DeleteExperimentPlanResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteExperimentPlanResponse) SetBody(v *DeleteExperimentPlanResponseBody) *DeleteExperimentPlanResponse {
	s.Body = v
	return s
}

func (s *DeleteExperimentPlanResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
