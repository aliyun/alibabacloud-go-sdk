// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddGtmRecoveryPlanResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddGtmRecoveryPlanResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddGtmRecoveryPlanResponse
	GetStatusCode() *int32
	SetBody(v *AddGtmRecoveryPlanResponseBody) *AddGtmRecoveryPlanResponse
	GetBody() *AddGtmRecoveryPlanResponseBody
}

type AddGtmRecoveryPlanResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddGtmRecoveryPlanResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddGtmRecoveryPlanResponse) String() string {
	return dara.Prettify(s)
}

func (s AddGtmRecoveryPlanResponse) GoString() string {
	return s.String()
}

func (s *AddGtmRecoveryPlanResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddGtmRecoveryPlanResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddGtmRecoveryPlanResponse) GetBody() *AddGtmRecoveryPlanResponseBody {
	return s.Body
}

func (s *AddGtmRecoveryPlanResponse) SetHeaders(v map[string]*string) *AddGtmRecoveryPlanResponse {
	s.Headers = v
	return s
}

func (s *AddGtmRecoveryPlanResponse) SetStatusCode(v int32) *AddGtmRecoveryPlanResponse {
	s.StatusCode = &v
	return s
}

func (s *AddGtmRecoveryPlanResponse) SetBody(v *AddGtmRecoveryPlanResponseBody) *AddGtmRecoveryPlanResponse {
	s.Body = v
	return s
}

func (s *AddGtmRecoveryPlanResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
