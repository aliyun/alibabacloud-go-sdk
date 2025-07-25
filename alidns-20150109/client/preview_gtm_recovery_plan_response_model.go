// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPreviewGtmRecoveryPlanResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *PreviewGtmRecoveryPlanResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *PreviewGtmRecoveryPlanResponse
	GetStatusCode() *int32
	SetBody(v *PreviewGtmRecoveryPlanResponseBody) *PreviewGtmRecoveryPlanResponse
	GetBody() *PreviewGtmRecoveryPlanResponseBody
}

type PreviewGtmRecoveryPlanResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PreviewGtmRecoveryPlanResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PreviewGtmRecoveryPlanResponse) String() string {
	return dara.Prettify(s)
}

func (s PreviewGtmRecoveryPlanResponse) GoString() string {
	return s.String()
}

func (s *PreviewGtmRecoveryPlanResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *PreviewGtmRecoveryPlanResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *PreviewGtmRecoveryPlanResponse) GetBody() *PreviewGtmRecoveryPlanResponseBody {
	return s.Body
}

func (s *PreviewGtmRecoveryPlanResponse) SetHeaders(v map[string]*string) *PreviewGtmRecoveryPlanResponse {
	s.Headers = v
	return s
}

func (s *PreviewGtmRecoveryPlanResponse) SetStatusCode(v int32) *PreviewGtmRecoveryPlanResponse {
	s.StatusCode = &v
	return s
}

func (s *PreviewGtmRecoveryPlanResponse) SetBody(v *PreviewGtmRecoveryPlanResponseBody) *PreviewGtmRecoveryPlanResponse {
	s.Body = v
	return s
}

func (s *PreviewGtmRecoveryPlanResponse) Validate() error {
	return dara.Validate(s)
}
