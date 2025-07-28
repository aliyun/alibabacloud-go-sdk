// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetUrgentDemandPlanDetailResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetUrgentDemandPlanDetailResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetUrgentDemandPlanDetailResponse
	GetStatusCode() *int32
	SetBody(v *GetUrgentDemandPlanDetailResponseBody) *GetUrgentDemandPlanDetailResponse
	GetBody() *GetUrgentDemandPlanDetailResponseBody
}

type GetUrgentDemandPlanDetailResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetUrgentDemandPlanDetailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetUrgentDemandPlanDetailResponse) String() string {
	return dara.Prettify(s)
}

func (s GetUrgentDemandPlanDetailResponse) GoString() string {
	return s.String()
}

func (s *GetUrgentDemandPlanDetailResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetUrgentDemandPlanDetailResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetUrgentDemandPlanDetailResponse) GetBody() *GetUrgentDemandPlanDetailResponseBody {
	return s.Body
}

func (s *GetUrgentDemandPlanDetailResponse) SetHeaders(v map[string]*string) *GetUrgentDemandPlanDetailResponse {
	s.Headers = v
	return s
}

func (s *GetUrgentDemandPlanDetailResponse) SetStatusCode(v int32) *GetUrgentDemandPlanDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *GetUrgentDemandPlanDetailResponse) SetBody(v *GetUrgentDemandPlanDetailResponseBody) *GetUrgentDemandPlanDetailResponse {
	s.Body = v
	return s
}

func (s *GetUrgentDemandPlanDetailResponse) Validate() error {
	return dara.Validate(s)
}
