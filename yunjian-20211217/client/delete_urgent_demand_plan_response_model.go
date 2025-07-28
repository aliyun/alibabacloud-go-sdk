// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteUrgentDemandPlanResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteUrgentDemandPlanResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteUrgentDemandPlanResponse
	GetStatusCode() *int32
	SetBody(v *DeleteUrgentDemandPlanResponseBody) *DeleteUrgentDemandPlanResponse
	GetBody() *DeleteUrgentDemandPlanResponseBody
}

type DeleteUrgentDemandPlanResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteUrgentDemandPlanResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteUrgentDemandPlanResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteUrgentDemandPlanResponse) GoString() string {
	return s.String()
}

func (s *DeleteUrgentDemandPlanResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteUrgentDemandPlanResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteUrgentDemandPlanResponse) GetBody() *DeleteUrgentDemandPlanResponseBody {
	return s.Body
}

func (s *DeleteUrgentDemandPlanResponse) SetHeaders(v map[string]*string) *DeleteUrgentDemandPlanResponse {
	s.Headers = v
	return s
}

func (s *DeleteUrgentDemandPlanResponse) SetStatusCode(v int32) *DeleteUrgentDemandPlanResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteUrgentDemandPlanResponse) SetBody(v *DeleteUrgentDemandPlanResponseBody) *DeleteUrgentDemandPlanResponse {
	s.Body = v
	return s
}

func (s *DeleteUrgentDemandPlanResponse) Validate() error {
	return dara.Validate(s)
}
