// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateJobPlanResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateJobPlanResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateJobPlanResponse
	GetStatusCode() *int32
	SetBody(v *UpdateJobPlanResponseBody) *UpdateJobPlanResponse
	GetBody() *UpdateJobPlanResponseBody
}

type UpdateJobPlanResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateJobPlanResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateJobPlanResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateJobPlanResponse) GoString() string {
	return s.String()
}

func (s *UpdateJobPlanResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateJobPlanResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateJobPlanResponse) GetBody() *UpdateJobPlanResponseBody {
	return s.Body
}

func (s *UpdateJobPlanResponse) SetHeaders(v map[string]*string) *UpdateJobPlanResponse {
	s.Headers = v
	return s
}

func (s *UpdateJobPlanResponse) SetStatusCode(v int32) *UpdateJobPlanResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateJobPlanResponse) SetBody(v *UpdateJobPlanResponseBody) *UpdateJobPlanResponse {
	s.Body = v
	return s
}

func (s *UpdateJobPlanResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
