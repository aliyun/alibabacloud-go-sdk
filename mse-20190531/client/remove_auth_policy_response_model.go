// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveAuthPolicyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RemoveAuthPolicyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RemoveAuthPolicyResponse
	GetStatusCode() *int32
	SetBody(v *RemoveAuthPolicyResponseBody) *RemoveAuthPolicyResponse
	GetBody() *RemoveAuthPolicyResponseBody
}

type RemoveAuthPolicyResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RemoveAuthPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RemoveAuthPolicyResponse) String() string {
	return dara.Prettify(s)
}

func (s RemoveAuthPolicyResponse) GoString() string {
	return s.String()
}

func (s *RemoveAuthPolicyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RemoveAuthPolicyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RemoveAuthPolicyResponse) GetBody() *RemoveAuthPolicyResponseBody {
	return s.Body
}

func (s *RemoveAuthPolicyResponse) SetHeaders(v map[string]*string) *RemoveAuthPolicyResponse {
	s.Headers = v
	return s
}

func (s *RemoveAuthPolicyResponse) SetStatusCode(v int32) *RemoveAuthPolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *RemoveAuthPolicyResponse) SetBody(v *RemoveAuthPolicyResponseBody) *RemoveAuthPolicyResponse {
	s.Body = v
	return s
}

func (s *RemoveAuthPolicyResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
