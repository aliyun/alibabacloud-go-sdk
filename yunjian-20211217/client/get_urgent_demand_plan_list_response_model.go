// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetUrgentDemandPlanListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetUrgentDemandPlanListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetUrgentDemandPlanListResponse
	GetStatusCode() *int32
	SetBody(v *GetUrgentDemandPlanListResponseBody) *GetUrgentDemandPlanListResponse
	GetBody() *GetUrgentDemandPlanListResponseBody
}

type GetUrgentDemandPlanListResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetUrgentDemandPlanListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetUrgentDemandPlanListResponse) String() string {
	return dara.Prettify(s)
}

func (s GetUrgentDemandPlanListResponse) GoString() string {
	return s.String()
}

func (s *GetUrgentDemandPlanListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetUrgentDemandPlanListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetUrgentDemandPlanListResponse) GetBody() *GetUrgentDemandPlanListResponseBody {
	return s.Body
}

func (s *GetUrgentDemandPlanListResponse) SetHeaders(v map[string]*string) *GetUrgentDemandPlanListResponse {
	s.Headers = v
	return s
}

func (s *GetUrgentDemandPlanListResponse) SetStatusCode(v int32) *GetUrgentDemandPlanListResponse {
	s.StatusCode = &v
	return s
}

func (s *GetUrgentDemandPlanListResponse) SetBody(v *GetUrgentDemandPlanListResponseBody) *GetUrgentDemandPlanListResponse {
	s.Body = v
	return s
}

func (s *GetUrgentDemandPlanListResponse) Validate() error {
	return dara.Validate(s)
}
