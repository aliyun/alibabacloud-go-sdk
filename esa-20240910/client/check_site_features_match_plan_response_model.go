// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckSiteFeaturesMatchPlanResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CheckSiteFeaturesMatchPlanResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CheckSiteFeaturesMatchPlanResponse
	GetStatusCode() *int32
	SetBody(v *CheckSiteFeaturesMatchPlanResponseBody) *CheckSiteFeaturesMatchPlanResponse
	GetBody() *CheckSiteFeaturesMatchPlanResponseBody
}

type CheckSiteFeaturesMatchPlanResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CheckSiteFeaturesMatchPlanResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CheckSiteFeaturesMatchPlanResponse) String() string {
	return dara.Prettify(s)
}

func (s CheckSiteFeaturesMatchPlanResponse) GoString() string {
	return s.String()
}

func (s *CheckSiteFeaturesMatchPlanResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CheckSiteFeaturesMatchPlanResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CheckSiteFeaturesMatchPlanResponse) GetBody() *CheckSiteFeaturesMatchPlanResponseBody {
	return s.Body
}

func (s *CheckSiteFeaturesMatchPlanResponse) SetHeaders(v map[string]*string) *CheckSiteFeaturesMatchPlanResponse {
	s.Headers = v
	return s
}

func (s *CheckSiteFeaturesMatchPlanResponse) SetStatusCode(v int32) *CheckSiteFeaturesMatchPlanResponse {
	s.StatusCode = &v
	return s
}

func (s *CheckSiteFeaturesMatchPlanResponse) SetBody(v *CheckSiteFeaturesMatchPlanResponseBody) *CheckSiteFeaturesMatchPlanResponse {
	s.Body = v
	return s
}

func (s *CheckSiteFeaturesMatchPlanResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
