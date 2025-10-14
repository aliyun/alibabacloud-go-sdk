// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRollbackGtmRecoveryPlanResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RollbackGtmRecoveryPlanResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RollbackGtmRecoveryPlanResponse
	GetStatusCode() *int32
	SetBody(v *RollbackGtmRecoveryPlanResponseBody) *RollbackGtmRecoveryPlanResponse
	GetBody() *RollbackGtmRecoveryPlanResponseBody
}

type RollbackGtmRecoveryPlanResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RollbackGtmRecoveryPlanResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RollbackGtmRecoveryPlanResponse) String() string {
	return dara.Prettify(s)
}

func (s RollbackGtmRecoveryPlanResponse) GoString() string {
	return s.String()
}

func (s *RollbackGtmRecoveryPlanResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RollbackGtmRecoveryPlanResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RollbackGtmRecoveryPlanResponse) GetBody() *RollbackGtmRecoveryPlanResponseBody {
	return s.Body
}

func (s *RollbackGtmRecoveryPlanResponse) SetHeaders(v map[string]*string) *RollbackGtmRecoveryPlanResponse {
	s.Headers = v
	return s
}

func (s *RollbackGtmRecoveryPlanResponse) SetStatusCode(v int32) *RollbackGtmRecoveryPlanResponse {
	s.StatusCode = &v
	return s
}

func (s *RollbackGtmRecoveryPlanResponse) SetBody(v *RollbackGtmRecoveryPlanResponseBody) *RollbackGtmRecoveryPlanResponse {
	s.Body = v
	return s
}

func (s *RollbackGtmRecoveryPlanResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
