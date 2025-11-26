// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetManagedScalingPolicyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetManagedScalingPolicyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetManagedScalingPolicyResponse
	GetStatusCode() *int32
	SetBody(v *GetManagedScalingPolicyResponseBody) *GetManagedScalingPolicyResponse
	GetBody() *GetManagedScalingPolicyResponseBody
}

type GetManagedScalingPolicyResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetManagedScalingPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetManagedScalingPolicyResponse) String() string {
	return dara.Prettify(s)
}

func (s GetManagedScalingPolicyResponse) GoString() string {
	return s.String()
}

func (s *GetManagedScalingPolicyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetManagedScalingPolicyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetManagedScalingPolicyResponse) GetBody() *GetManagedScalingPolicyResponseBody {
	return s.Body
}

func (s *GetManagedScalingPolicyResponse) SetHeaders(v map[string]*string) *GetManagedScalingPolicyResponse {
	s.Headers = v
	return s
}

func (s *GetManagedScalingPolicyResponse) SetStatusCode(v int32) *GetManagedScalingPolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *GetManagedScalingPolicyResponse) SetBody(v *GetManagedScalingPolicyResponseBody) *GetManagedScalingPolicyResponse {
	s.Body = v
	return s
}

func (s *GetManagedScalingPolicyResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
