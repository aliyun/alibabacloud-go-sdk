// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetExperimentPlanResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetExperimentPlanResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetExperimentPlanResponse
	GetStatusCode() *int32
	SetBody(v *GetExperimentPlanResponseBody) *GetExperimentPlanResponse
	GetBody() *GetExperimentPlanResponseBody
}

type GetExperimentPlanResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetExperimentPlanResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetExperimentPlanResponse) String() string {
	return dara.Prettify(s)
}

func (s GetExperimentPlanResponse) GoString() string {
	return s.String()
}

func (s *GetExperimentPlanResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetExperimentPlanResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetExperimentPlanResponse) GetBody() *GetExperimentPlanResponseBody {
	return s.Body
}

func (s *GetExperimentPlanResponse) SetHeaders(v map[string]*string) *GetExperimentPlanResponse {
	s.Headers = v
	return s
}

func (s *GetExperimentPlanResponse) SetStatusCode(v int32) *GetExperimentPlanResponse {
	s.StatusCode = &v
	return s
}

func (s *GetExperimentPlanResponse) SetBody(v *GetExperimentPlanResponseBody) *GetExperimentPlanResponse {
	s.Body = v
	return s
}

func (s *GetExperimentPlanResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
