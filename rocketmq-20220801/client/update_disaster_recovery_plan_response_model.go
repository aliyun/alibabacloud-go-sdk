// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDisasterRecoveryPlanResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateDisasterRecoveryPlanResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateDisasterRecoveryPlanResponse
	GetStatusCode() *int32
	SetBody(v *UpdateDisasterRecoveryPlanResponseBody) *UpdateDisasterRecoveryPlanResponse
	GetBody() *UpdateDisasterRecoveryPlanResponseBody
}

type UpdateDisasterRecoveryPlanResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateDisasterRecoveryPlanResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateDisasterRecoveryPlanResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateDisasterRecoveryPlanResponse) GoString() string {
	return s.String()
}

func (s *UpdateDisasterRecoveryPlanResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateDisasterRecoveryPlanResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateDisasterRecoveryPlanResponse) GetBody() *UpdateDisasterRecoveryPlanResponseBody {
	return s.Body
}

func (s *UpdateDisasterRecoveryPlanResponse) SetHeaders(v map[string]*string) *UpdateDisasterRecoveryPlanResponse {
	s.Headers = v
	return s
}

func (s *UpdateDisasterRecoveryPlanResponse) SetStatusCode(v int32) *UpdateDisasterRecoveryPlanResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateDisasterRecoveryPlanResponse) SetBody(v *UpdateDisasterRecoveryPlanResponseBody) *UpdateDisasterRecoveryPlanResponse {
	s.Body = v
	return s
}

func (s *UpdateDisasterRecoveryPlanResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
