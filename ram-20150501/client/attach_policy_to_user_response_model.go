// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAttachPolicyToUserResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AttachPolicyToUserResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AttachPolicyToUserResponse
	GetStatusCode() *int32
	SetBody(v *AttachPolicyToUserResponseBody) *AttachPolicyToUserResponse
	GetBody() *AttachPolicyToUserResponseBody
}

type AttachPolicyToUserResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AttachPolicyToUserResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AttachPolicyToUserResponse) String() string {
	return dara.Prettify(s)
}

func (s AttachPolicyToUserResponse) GoString() string {
	return s.String()
}

func (s *AttachPolicyToUserResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AttachPolicyToUserResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AttachPolicyToUserResponse) GetBody() *AttachPolicyToUserResponseBody {
	return s.Body
}

func (s *AttachPolicyToUserResponse) SetHeaders(v map[string]*string) *AttachPolicyToUserResponse {
	s.Headers = v
	return s
}

func (s *AttachPolicyToUserResponse) SetStatusCode(v int32) *AttachPolicyToUserResponse {
	s.StatusCode = &v
	return s
}

func (s *AttachPolicyToUserResponse) SetBody(v *AttachPolicyToUserResponseBody) *AttachPolicyToUserResponse {
	s.Body = v
	return s
}

func (s *AttachPolicyToUserResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
