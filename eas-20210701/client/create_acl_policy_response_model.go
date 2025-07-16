// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAclPolicyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateAclPolicyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateAclPolicyResponse
	GetStatusCode() *int32
	SetBody(v *CreateAclPolicyResponseBody) *CreateAclPolicyResponse
	GetBody() *CreateAclPolicyResponseBody
}

type CreateAclPolicyResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateAclPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateAclPolicyResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateAclPolicyResponse) GoString() string {
	return s.String()
}

func (s *CreateAclPolicyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateAclPolicyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateAclPolicyResponse) GetBody() *CreateAclPolicyResponseBody {
	return s.Body
}

func (s *CreateAclPolicyResponse) SetHeaders(v map[string]*string) *CreateAclPolicyResponse {
	s.Headers = v
	return s
}

func (s *CreateAclPolicyResponse) SetStatusCode(v int32) *CreateAclPolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateAclPolicyResponse) SetBody(v *CreateAclPolicyResponseBody) *CreateAclPolicyResponse {
	s.Body = v
	return s
}

func (s *CreateAclPolicyResponse) Validate() error {
	return dara.Validate(s)
}
