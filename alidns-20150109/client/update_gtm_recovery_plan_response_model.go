// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateGtmRecoveryPlanResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateGtmRecoveryPlanResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateGtmRecoveryPlanResponse
	GetStatusCode() *int32
	SetBody(v *UpdateGtmRecoveryPlanResponseBody) *UpdateGtmRecoveryPlanResponse
	GetBody() *UpdateGtmRecoveryPlanResponseBody
}

type UpdateGtmRecoveryPlanResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateGtmRecoveryPlanResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateGtmRecoveryPlanResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateGtmRecoveryPlanResponse) GoString() string {
	return s.String()
}

func (s *UpdateGtmRecoveryPlanResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateGtmRecoveryPlanResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateGtmRecoveryPlanResponse) GetBody() *UpdateGtmRecoveryPlanResponseBody {
	return s.Body
}

func (s *UpdateGtmRecoveryPlanResponse) SetHeaders(v map[string]*string) *UpdateGtmRecoveryPlanResponse {
	s.Headers = v
	return s
}

func (s *UpdateGtmRecoveryPlanResponse) SetStatusCode(v int32) *UpdateGtmRecoveryPlanResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateGtmRecoveryPlanResponse) SetBody(v *UpdateGtmRecoveryPlanResponseBody) *UpdateGtmRecoveryPlanResponse {
	s.Body = v
	return s
}

func (s *UpdateGtmRecoveryPlanResponse) Validate() error {
	return dara.Validate(s)
}
