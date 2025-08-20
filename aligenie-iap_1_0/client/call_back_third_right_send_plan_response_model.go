// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCallBackThirdRightSendPlanResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CallBackThirdRightSendPlanResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CallBackThirdRightSendPlanResponse
	GetStatusCode() *int32
	SetBody(v *CallBackThirdRightSendPlanResponseBody) *CallBackThirdRightSendPlanResponse
	GetBody() *CallBackThirdRightSendPlanResponseBody
}

type CallBackThirdRightSendPlanResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CallBackThirdRightSendPlanResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CallBackThirdRightSendPlanResponse) String() string {
	return dara.Prettify(s)
}

func (s CallBackThirdRightSendPlanResponse) GoString() string {
	return s.String()
}

func (s *CallBackThirdRightSendPlanResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CallBackThirdRightSendPlanResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CallBackThirdRightSendPlanResponse) GetBody() *CallBackThirdRightSendPlanResponseBody {
	return s.Body
}

func (s *CallBackThirdRightSendPlanResponse) SetHeaders(v map[string]*string) *CallBackThirdRightSendPlanResponse {
	s.Headers = v
	return s
}

func (s *CallBackThirdRightSendPlanResponse) SetStatusCode(v int32) *CallBackThirdRightSendPlanResponse {
	s.StatusCode = &v
	return s
}

func (s *CallBackThirdRightSendPlanResponse) SetBody(v *CallBackThirdRightSendPlanResponseBody) *CallBackThirdRightSendPlanResponse {
	s.Body = v
	return s
}

func (s *CallBackThirdRightSendPlanResponse) Validate() error {
	return dara.Validate(s)
}
