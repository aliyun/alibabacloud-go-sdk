// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCapacityPlanResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CapacityPlanResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CapacityPlanResponse
	GetStatusCode() *int32
	SetBody(v *CapacityPlanResponseBody) *CapacityPlanResponse
	GetBody() *CapacityPlanResponseBody
}

type CapacityPlanResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CapacityPlanResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CapacityPlanResponse) String() string {
	return dara.Prettify(s)
}

func (s CapacityPlanResponse) GoString() string {
	return s.String()
}

func (s *CapacityPlanResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CapacityPlanResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CapacityPlanResponse) GetBody() *CapacityPlanResponseBody {
	return s.Body
}

func (s *CapacityPlanResponse) SetHeaders(v map[string]*string) *CapacityPlanResponse {
	s.Headers = v
	return s
}

func (s *CapacityPlanResponse) SetStatusCode(v int32) *CapacityPlanResponse {
	s.StatusCode = &v
	return s
}

func (s *CapacityPlanResponse) SetBody(v *CapacityPlanResponseBody) *CapacityPlanResponse {
	s.Body = v
	return s
}

func (s *CapacityPlanResponse) Validate() error {
	return dara.Validate(s)
}
