// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetRunningPlanResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SetRunningPlanResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SetRunningPlanResponse
	GetStatusCode() *int32
	SetBody(v *SetRunningPlanResponseBody) *SetRunningPlanResponse
	GetBody() *SetRunningPlanResponseBody
}

type SetRunningPlanResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SetRunningPlanResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SetRunningPlanResponse) String() string {
	return dara.Prettify(s)
}

func (s SetRunningPlanResponse) GoString() string {
	return s.String()
}

func (s *SetRunningPlanResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SetRunningPlanResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SetRunningPlanResponse) GetBody() *SetRunningPlanResponseBody {
	return s.Body
}

func (s *SetRunningPlanResponse) SetHeaders(v map[string]*string) *SetRunningPlanResponse {
	s.Headers = v
	return s
}

func (s *SetRunningPlanResponse) SetStatusCode(v int32) *SetRunningPlanResponse {
	s.StatusCode = &v
	return s
}

func (s *SetRunningPlanResponse) SetBody(v *SetRunningPlanResponseBody) *SetRunningPlanResponse {
	s.Body = v
	return s
}

func (s *SetRunningPlanResponse) Validate() error {
	return dara.Validate(s)
}
