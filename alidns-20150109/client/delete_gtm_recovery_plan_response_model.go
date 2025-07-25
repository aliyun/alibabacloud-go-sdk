// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteGtmRecoveryPlanResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteGtmRecoveryPlanResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteGtmRecoveryPlanResponse
	GetStatusCode() *int32
	SetBody(v *DeleteGtmRecoveryPlanResponseBody) *DeleteGtmRecoveryPlanResponse
	GetBody() *DeleteGtmRecoveryPlanResponseBody
}

type DeleteGtmRecoveryPlanResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteGtmRecoveryPlanResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteGtmRecoveryPlanResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteGtmRecoveryPlanResponse) GoString() string {
	return s.String()
}

func (s *DeleteGtmRecoveryPlanResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteGtmRecoveryPlanResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteGtmRecoveryPlanResponse) GetBody() *DeleteGtmRecoveryPlanResponseBody {
	return s.Body
}

func (s *DeleteGtmRecoveryPlanResponse) SetHeaders(v map[string]*string) *DeleteGtmRecoveryPlanResponse {
	s.Headers = v
	return s
}

func (s *DeleteGtmRecoveryPlanResponse) SetStatusCode(v int32) *DeleteGtmRecoveryPlanResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteGtmRecoveryPlanResponse) SetBody(v *DeleteGtmRecoveryPlanResponseBody) *DeleteGtmRecoveryPlanResponse {
	s.Body = v
	return s
}

func (s *DeleteGtmRecoveryPlanResponse) Validate() error {
	return dara.Validate(s)
}
