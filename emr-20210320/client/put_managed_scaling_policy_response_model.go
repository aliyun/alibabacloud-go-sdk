// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPutManagedScalingPolicyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *PutManagedScalingPolicyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *PutManagedScalingPolicyResponse
	GetStatusCode() *int32
	SetBody(v *PutManagedScalingPolicyResponseBody) *PutManagedScalingPolicyResponse
	GetBody() *PutManagedScalingPolicyResponseBody
}

type PutManagedScalingPolicyResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PutManagedScalingPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PutManagedScalingPolicyResponse) String() string {
	return dara.Prettify(s)
}

func (s PutManagedScalingPolicyResponse) GoString() string {
	return s.String()
}

func (s *PutManagedScalingPolicyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *PutManagedScalingPolicyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *PutManagedScalingPolicyResponse) GetBody() *PutManagedScalingPolicyResponseBody {
	return s.Body
}

func (s *PutManagedScalingPolicyResponse) SetHeaders(v map[string]*string) *PutManagedScalingPolicyResponse {
	s.Headers = v
	return s
}

func (s *PutManagedScalingPolicyResponse) SetStatusCode(v int32) *PutManagedScalingPolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *PutManagedScalingPolicyResponse) SetBody(v *PutManagedScalingPolicyResponseBody) *PutManagedScalingPolicyResponse {
	s.Body = v
	return s
}

func (s *PutManagedScalingPolicyResponse) Validate() error {
	return dara.Validate(s)
}
