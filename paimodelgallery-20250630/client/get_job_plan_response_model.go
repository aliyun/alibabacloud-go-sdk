// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetJobPlanResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetJobPlanResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetJobPlanResponse
	GetStatusCode() *int32
	SetBody(v *GetJobPlanResponseBody) *GetJobPlanResponse
	GetBody() *GetJobPlanResponseBody
}

type GetJobPlanResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetJobPlanResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetJobPlanResponse) String() string {
	return dara.Prettify(s)
}

func (s GetJobPlanResponse) GoString() string {
	return s.String()
}

func (s *GetJobPlanResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetJobPlanResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetJobPlanResponse) GetBody() *GetJobPlanResponseBody {
	return s.Body
}

func (s *GetJobPlanResponse) SetHeaders(v map[string]*string) *GetJobPlanResponse {
	s.Headers = v
	return s
}

func (s *GetJobPlanResponse) SetStatusCode(v int32) *GetJobPlanResponse {
	s.StatusCode = &v
	return s
}

func (s *GetJobPlanResponse) SetBody(v *GetJobPlanResponseBody) *GetJobPlanResponse {
	s.Body = v
	return s
}

func (s *GetJobPlanResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
