// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveAutoScalingPolicyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RemoveAutoScalingPolicyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RemoveAutoScalingPolicyResponse
	GetStatusCode() *int32
	SetBody(v *RemoveAutoScalingPolicyResponseBody) *RemoveAutoScalingPolicyResponse
	GetBody() *RemoveAutoScalingPolicyResponseBody
}

type RemoveAutoScalingPolicyResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RemoveAutoScalingPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RemoveAutoScalingPolicyResponse) String() string {
	return dara.Prettify(s)
}

func (s RemoveAutoScalingPolicyResponse) GoString() string {
	return s.String()
}

func (s *RemoveAutoScalingPolicyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RemoveAutoScalingPolicyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RemoveAutoScalingPolicyResponse) GetBody() *RemoveAutoScalingPolicyResponseBody {
	return s.Body
}

func (s *RemoveAutoScalingPolicyResponse) SetHeaders(v map[string]*string) *RemoveAutoScalingPolicyResponse {
	s.Headers = v
	return s
}

func (s *RemoveAutoScalingPolicyResponse) SetStatusCode(v int32) *RemoveAutoScalingPolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *RemoveAutoScalingPolicyResponse) SetBody(v *RemoveAutoScalingPolicyResponseBody) *RemoveAutoScalingPolicyResponse {
	s.Body = v
	return s
}

func (s *RemoveAutoScalingPolicyResponse) Validate() error {
	return dara.Validate(s)
}
