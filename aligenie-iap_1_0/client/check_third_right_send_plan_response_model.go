// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckThirdRightSendPlanResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CheckThirdRightSendPlanResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CheckThirdRightSendPlanResponse
	GetStatusCode() *int32
	SetBody(v *CheckThirdRightSendPlanResponseBody) *CheckThirdRightSendPlanResponse
	GetBody() *CheckThirdRightSendPlanResponseBody
}

type CheckThirdRightSendPlanResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CheckThirdRightSendPlanResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CheckThirdRightSendPlanResponse) String() string {
	return dara.Prettify(s)
}

func (s CheckThirdRightSendPlanResponse) GoString() string {
	return s.String()
}

func (s *CheckThirdRightSendPlanResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CheckThirdRightSendPlanResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CheckThirdRightSendPlanResponse) GetBody() *CheckThirdRightSendPlanResponseBody {
	return s.Body
}

func (s *CheckThirdRightSendPlanResponse) SetHeaders(v map[string]*string) *CheckThirdRightSendPlanResponse {
	s.Headers = v
	return s
}

func (s *CheckThirdRightSendPlanResponse) SetStatusCode(v int32) *CheckThirdRightSendPlanResponse {
	s.StatusCode = &v
	return s
}

func (s *CheckThirdRightSendPlanResponse) SetBody(v *CheckThirdRightSendPlanResponseBody) *CheckThirdRightSendPlanResponse {
	s.Body = v
	return s
}

func (s *CheckThirdRightSendPlanResponse) Validate() error {
	return dara.Validate(s)
}
