// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteJobPlanResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteJobPlanResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteJobPlanResponse
	GetStatusCode() *int32
	SetBody(v *DeleteJobPlanResponseBody) *DeleteJobPlanResponse
	GetBody() *DeleteJobPlanResponseBody
}

type DeleteJobPlanResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteJobPlanResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteJobPlanResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteJobPlanResponse) GoString() string {
	return s.String()
}

func (s *DeleteJobPlanResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteJobPlanResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteJobPlanResponse) GetBody() *DeleteJobPlanResponseBody {
	return s.Body
}

func (s *DeleteJobPlanResponse) SetHeaders(v map[string]*string) *DeleteJobPlanResponse {
	s.Headers = v
	return s
}

func (s *DeleteJobPlanResponse) SetStatusCode(v int32) *DeleteJobPlanResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteJobPlanResponse) SetBody(v *DeleteJobPlanResponseBody) *DeleteJobPlanResponse {
	s.Body = v
	return s
}

func (s *DeleteJobPlanResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
