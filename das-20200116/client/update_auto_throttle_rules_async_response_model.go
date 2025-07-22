// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAutoThrottleRulesAsyncResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateAutoThrottleRulesAsyncResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateAutoThrottleRulesAsyncResponse
	GetStatusCode() *int32
	SetBody(v *UpdateAutoThrottleRulesAsyncResponseBody) *UpdateAutoThrottleRulesAsyncResponse
	GetBody() *UpdateAutoThrottleRulesAsyncResponseBody
}

type UpdateAutoThrottleRulesAsyncResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateAutoThrottleRulesAsyncResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateAutoThrottleRulesAsyncResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateAutoThrottleRulesAsyncResponse) GoString() string {
	return s.String()
}

func (s *UpdateAutoThrottleRulesAsyncResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateAutoThrottleRulesAsyncResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateAutoThrottleRulesAsyncResponse) GetBody() *UpdateAutoThrottleRulesAsyncResponseBody {
	return s.Body
}

func (s *UpdateAutoThrottleRulesAsyncResponse) SetHeaders(v map[string]*string) *UpdateAutoThrottleRulesAsyncResponse {
	s.Headers = v
	return s
}

func (s *UpdateAutoThrottleRulesAsyncResponse) SetStatusCode(v int32) *UpdateAutoThrottleRulesAsyncResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateAutoThrottleRulesAsyncResponse) SetBody(v *UpdateAutoThrottleRulesAsyncResponseBody) *UpdateAutoThrottleRulesAsyncResponse {
	s.Body = v
	return s
}

func (s *UpdateAutoThrottleRulesAsyncResponse) Validate() error {
	return dara.Validate(s)
}
