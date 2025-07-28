// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitUrgentDemandPlanResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SubmitUrgentDemandPlanResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SubmitUrgentDemandPlanResponse
	GetStatusCode() *int32
	SetBody(v *SubmitUrgentDemandPlanResponseBody) *SubmitUrgentDemandPlanResponse
	GetBody() *SubmitUrgentDemandPlanResponseBody
}

type SubmitUrgentDemandPlanResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SubmitUrgentDemandPlanResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SubmitUrgentDemandPlanResponse) String() string {
	return dara.Prettify(s)
}

func (s SubmitUrgentDemandPlanResponse) GoString() string {
	return s.String()
}

func (s *SubmitUrgentDemandPlanResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SubmitUrgentDemandPlanResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SubmitUrgentDemandPlanResponse) GetBody() *SubmitUrgentDemandPlanResponseBody {
	return s.Body
}

func (s *SubmitUrgentDemandPlanResponse) SetHeaders(v map[string]*string) *SubmitUrgentDemandPlanResponse {
	s.Headers = v
	return s
}

func (s *SubmitUrgentDemandPlanResponse) SetStatusCode(v int32) *SubmitUrgentDemandPlanResponse {
	s.StatusCode = &v
	return s
}

func (s *SubmitUrgentDemandPlanResponse) SetBody(v *SubmitUrgentDemandPlanResponseBody) *SubmitUrgentDemandPlanResponse {
	s.Body = v
	return s
}

func (s *SubmitUrgentDemandPlanResponse) Validate() error {
	return dara.Validate(s)
}
