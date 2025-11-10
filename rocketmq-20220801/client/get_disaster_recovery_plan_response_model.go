// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDisasterRecoveryPlanResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetDisasterRecoveryPlanResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetDisasterRecoveryPlanResponse
	GetStatusCode() *int32
	SetBody(v *GetDisasterRecoveryPlanResponseBody) *GetDisasterRecoveryPlanResponse
	GetBody() *GetDisasterRecoveryPlanResponseBody
}

type GetDisasterRecoveryPlanResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetDisasterRecoveryPlanResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetDisasterRecoveryPlanResponse) String() string {
	return dara.Prettify(s)
}

func (s GetDisasterRecoveryPlanResponse) GoString() string {
	return s.String()
}

func (s *GetDisasterRecoveryPlanResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetDisasterRecoveryPlanResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetDisasterRecoveryPlanResponse) GetBody() *GetDisasterRecoveryPlanResponseBody {
	return s.Body
}

func (s *GetDisasterRecoveryPlanResponse) SetHeaders(v map[string]*string) *GetDisasterRecoveryPlanResponse {
	s.Headers = v
	return s
}

func (s *GetDisasterRecoveryPlanResponse) SetStatusCode(v int32) *GetDisasterRecoveryPlanResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDisasterRecoveryPlanResponse) SetBody(v *GetDisasterRecoveryPlanResponseBody) *GetDisasterRecoveryPlanResponse {
	s.Body = v
	return s
}

func (s *GetDisasterRecoveryPlanResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
