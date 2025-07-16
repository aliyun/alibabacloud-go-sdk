// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDeliveryPlanResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateDeliveryPlanResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateDeliveryPlanResponse
	GetStatusCode() *int32
	SetBody(v *CreateDeliveryPlanResponseBody) *CreateDeliveryPlanResponse
	GetBody() *CreateDeliveryPlanResponseBody
}

type CreateDeliveryPlanResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateDeliveryPlanResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateDeliveryPlanResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateDeliveryPlanResponse) GoString() string {
	return s.String()
}

func (s *CreateDeliveryPlanResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateDeliveryPlanResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateDeliveryPlanResponse) GetBody() *CreateDeliveryPlanResponseBody {
	return s.Body
}

func (s *CreateDeliveryPlanResponse) SetHeaders(v map[string]*string) *CreateDeliveryPlanResponse {
	s.Headers = v
	return s
}

func (s *CreateDeliveryPlanResponse) SetStatusCode(v int32) *CreateDeliveryPlanResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateDeliveryPlanResponse) SetBody(v *CreateDeliveryPlanResponseBody) *CreateDeliveryPlanResponse {
	s.Body = v
	return s
}

func (s *CreateDeliveryPlanResponse) Validate() error {
	return dara.Validate(s)
}
