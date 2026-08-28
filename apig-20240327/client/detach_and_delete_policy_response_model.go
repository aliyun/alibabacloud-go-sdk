// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDetachAndDeletePolicyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DetachAndDeletePolicyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DetachAndDeletePolicyResponse
	GetStatusCode() *int32
	SetBody(v *DetachAndDeletePolicyResponseBody) *DetachAndDeletePolicyResponse
	GetBody() *DetachAndDeletePolicyResponseBody
}

type DetachAndDeletePolicyResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DetachAndDeletePolicyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DetachAndDeletePolicyResponse) String() string {
	return dara.Prettify(s)
}

func (s DetachAndDeletePolicyResponse) GoString() string {
	return s.String()
}

func (s *DetachAndDeletePolicyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DetachAndDeletePolicyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DetachAndDeletePolicyResponse) GetBody() *DetachAndDeletePolicyResponseBody {
	return s.Body
}

func (s *DetachAndDeletePolicyResponse) SetHeaders(v map[string]*string) *DetachAndDeletePolicyResponse {
	s.Headers = v
	return s
}

func (s *DetachAndDeletePolicyResponse) SetStatusCode(v int32) *DetachAndDeletePolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *DetachAndDeletePolicyResponse) SetBody(v *DetachAndDeletePolicyResponseBody) *DetachAndDeletePolicyResponse {
	s.Body = v
	return s
}

func (s *DetachAndDeletePolicyResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
