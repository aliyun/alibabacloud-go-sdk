// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAutoScalingPolicyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetAutoScalingPolicyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetAutoScalingPolicyResponse
	GetStatusCode() *int32
	SetBody(v *GetAutoScalingPolicyResponseBody) *GetAutoScalingPolicyResponse
	GetBody() *GetAutoScalingPolicyResponseBody
}

type GetAutoScalingPolicyResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAutoScalingPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAutoScalingPolicyResponse) String() string {
	return dara.Prettify(s)
}

func (s GetAutoScalingPolicyResponse) GoString() string {
	return s.String()
}

func (s *GetAutoScalingPolicyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetAutoScalingPolicyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetAutoScalingPolicyResponse) GetBody() *GetAutoScalingPolicyResponseBody {
	return s.Body
}

func (s *GetAutoScalingPolicyResponse) SetHeaders(v map[string]*string) *GetAutoScalingPolicyResponse {
	s.Headers = v
	return s
}

func (s *GetAutoScalingPolicyResponse) SetStatusCode(v int32) *GetAutoScalingPolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAutoScalingPolicyResponse) SetBody(v *GetAutoScalingPolicyResponseBody) *GetAutoScalingPolicyResponse {
	s.Body = v
	return s
}

func (s *GetAutoScalingPolicyResponse) Validate() error {
	return dara.Validate(s)
}
