// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateJobPlanResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateJobPlanResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateJobPlanResponse
	GetStatusCode() *int32
	SetBody(v *CreateJobPlanResponseBody) *CreateJobPlanResponse
	GetBody() *CreateJobPlanResponseBody
}

type CreateJobPlanResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateJobPlanResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateJobPlanResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateJobPlanResponse) GoString() string {
	return s.String()
}

func (s *CreateJobPlanResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateJobPlanResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateJobPlanResponse) GetBody() *CreateJobPlanResponseBody {
	return s.Body
}

func (s *CreateJobPlanResponse) SetHeaders(v map[string]*string) *CreateJobPlanResponse {
	s.Headers = v
	return s
}

func (s *CreateJobPlanResponse) SetStatusCode(v int32) *CreateJobPlanResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateJobPlanResponse) SetBody(v *CreateJobPlanResponseBody) *CreateJobPlanResponse {
	s.Body = v
	return s
}

func (s *CreateJobPlanResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
