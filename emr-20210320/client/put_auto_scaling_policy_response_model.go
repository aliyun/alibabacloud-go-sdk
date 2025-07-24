// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPutAutoScalingPolicyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *PutAutoScalingPolicyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *PutAutoScalingPolicyResponse
	GetStatusCode() *int32
	SetBody(v *PutAutoScalingPolicyResponseBody) *PutAutoScalingPolicyResponse
	GetBody() *PutAutoScalingPolicyResponseBody
}

type PutAutoScalingPolicyResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PutAutoScalingPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PutAutoScalingPolicyResponse) String() string {
	return dara.Prettify(s)
}

func (s PutAutoScalingPolicyResponse) GoString() string {
	return s.String()
}

func (s *PutAutoScalingPolicyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *PutAutoScalingPolicyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *PutAutoScalingPolicyResponse) GetBody() *PutAutoScalingPolicyResponseBody {
	return s.Body
}

func (s *PutAutoScalingPolicyResponse) SetHeaders(v map[string]*string) *PutAutoScalingPolicyResponse {
	s.Headers = v
	return s
}

func (s *PutAutoScalingPolicyResponse) SetStatusCode(v int32) *PutAutoScalingPolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *PutAutoScalingPolicyResponse) SetBody(v *PutAutoScalingPolicyResponseBody) *PutAutoScalingPolicyResponse {
	s.Body = v
	return s
}

func (s *PutAutoScalingPolicyResponse) Validate() error {
	return dara.Validate(s)
}
