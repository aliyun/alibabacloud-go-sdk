// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAclPolicyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteAclPolicyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteAclPolicyResponse
	GetStatusCode() *int32
	SetBody(v *DeleteAclPolicyResponseBody) *DeleteAclPolicyResponse
	GetBody() *DeleteAclPolicyResponseBody
}

type DeleteAclPolicyResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteAclPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteAclPolicyResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteAclPolicyResponse) GoString() string {
	return s.String()
}

func (s *DeleteAclPolicyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteAclPolicyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteAclPolicyResponse) GetBody() *DeleteAclPolicyResponseBody {
	return s.Body
}

func (s *DeleteAclPolicyResponse) SetHeaders(v map[string]*string) *DeleteAclPolicyResponse {
	s.Headers = v
	return s
}

func (s *DeleteAclPolicyResponse) SetStatusCode(v int32) *DeleteAclPolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteAclPolicyResponse) SetBody(v *DeleteAclPolicyResponseBody) *DeleteAclPolicyResponse {
	s.Body = v
	return s
}

func (s *DeleteAclPolicyResponse) Validate() error {
	return dara.Validate(s)
}
