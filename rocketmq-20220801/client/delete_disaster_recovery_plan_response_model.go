// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDisasterRecoveryPlanResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteDisasterRecoveryPlanResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteDisasterRecoveryPlanResponse
	GetStatusCode() *int32
	SetBody(v *DeleteDisasterRecoveryPlanResponseBody) *DeleteDisasterRecoveryPlanResponse
	GetBody() *DeleteDisasterRecoveryPlanResponseBody
}

type DeleteDisasterRecoveryPlanResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteDisasterRecoveryPlanResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteDisasterRecoveryPlanResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteDisasterRecoveryPlanResponse) GoString() string {
	return s.String()
}

func (s *DeleteDisasterRecoveryPlanResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteDisasterRecoveryPlanResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteDisasterRecoveryPlanResponse) GetBody() *DeleteDisasterRecoveryPlanResponseBody {
	return s.Body
}

func (s *DeleteDisasterRecoveryPlanResponse) SetHeaders(v map[string]*string) *DeleteDisasterRecoveryPlanResponse {
	s.Headers = v
	return s
}

func (s *DeleteDisasterRecoveryPlanResponse) SetStatusCode(v int32) *DeleteDisasterRecoveryPlanResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteDisasterRecoveryPlanResponse) SetBody(v *DeleteDisasterRecoveryPlanResponseBody) *DeleteDisasterRecoveryPlanResponse {
	s.Body = v
	return s
}

func (s *DeleteDisasterRecoveryPlanResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
