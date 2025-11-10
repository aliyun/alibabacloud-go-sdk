// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDisasterRecoveryPlanResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateDisasterRecoveryPlanResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateDisasterRecoveryPlanResponse
	GetStatusCode() *int32
	SetBody(v *CreateDisasterRecoveryPlanResponseBody) *CreateDisasterRecoveryPlanResponse
	GetBody() *CreateDisasterRecoveryPlanResponseBody
}

type CreateDisasterRecoveryPlanResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateDisasterRecoveryPlanResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateDisasterRecoveryPlanResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateDisasterRecoveryPlanResponse) GoString() string {
	return s.String()
}

func (s *CreateDisasterRecoveryPlanResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateDisasterRecoveryPlanResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateDisasterRecoveryPlanResponse) GetBody() *CreateDisasterRecoveryPlanResponseBody {
	return s.Body
}

func (s *CreateDisasterRecoveryPlanResponse) SetHeaders(v map[string]*string) *CreateDisasterRecoveryPlanResponse {
	s.Headers = v
	return s
}

func (s *CreateDisasterRecoveryPlanResponse) SetStatusCode(v int32) *CreateDisasterRecoveryPlanResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateDisasterRecoveryPlanResponse) SetBody(v *CreateDisasterRecoveryPlanResponseBody) *CreateDisasterRecoveryPlanResponse {
	s.Body = v
	return s
}

func (s *CreateDisasterRecoveryPlanResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
