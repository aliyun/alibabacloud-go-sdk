// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAutoScalingRulesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListAutoScalingRulesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListAutoScalingRulesResponse
	GetStatusCode() *int32
	SetBody(v *ListAutoScalingRulesResponseBody) *ListAutoScalingRulesResponse
	GetBody() *ListAutoScalingRulesResponseBody
}

type ListAutoScalingRulesResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAutoScalingRulesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAutoScalingRulesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListAutoScalingRulesResponse) GoString() string {
	return s.String()
}

func (s *ListAutoScalingRulesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListAutoScalingRulesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListAutoScalingRulesResponse) GetBody() *ListAutoScalingRulesResponseBody {
	return s.Body
}

func (s *ListAutoScalingRulesResponse) SetHeaders(v map[string]*string) *ListAutoScalingRulesResponse {
	s.Headers = v
	return s
}

func (s *ListAutoScalingRulesResponse) SetStatusCode(v int32) *ListAutoScalingRulesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAutoScalingRulesResponse) SetBody(v *ListAutoScalingRulesResponseBody) *ListAutoScalingRulesResponse {
	s.Body = v
	return s
}

func (s *ListAutoScalingRulesResponse) Validate() error {
	return dara.Validate(s)
}
