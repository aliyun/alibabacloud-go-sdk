// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAttachControlPolicyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AttachControlPolicyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AttachControlPolicyResponse
	GetStatusCode() *int32
	SetBody(v *AttachControlPolicyResponseBody) *AttachControlPolicyResponse
	GetBody() *AttachControlPolicyResponseBody
}

type AttachControlPolicyResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AttachControlPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AttachControlPolicyResponse) String() string {
	return dara.Prettify(s)
}

func (s AttachControlPolicyResponse) GoString() string {
	return s.String()
}

func (s *AttachControlPolicyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AttachControlPolicyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AttachControlPolicyResponse) GetBody() *AttachControlPolicyResponseBody {
	return s.Body
}

func (s *AttachControlPolicyResponse) SetHeaders(v map[string]*string) *AttachControlPolicyResponse {
	s.Headers = v
	return s
}

func (s *AttachControlPolicyResponse) SetStatusCode(v int32) *AttachControlPolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *AttachControlPolicyResponse) SetBody(v *AttachControlPolicyResponseBody) *AttachControlPolicyResponse {
	s.Body = v
	return s
}

func (s *AttachControlPolicyResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
