// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteInstanceEndpointAclPolicyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteInstanceEndpointAclPolicyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteInstanceEndpointAclPolicyResponse
	GetStatusCode() *int32
	SetBody(v *DeleteInstanceEndpointAclPolicyResponseBody) *DeleteInstanceEndpointAclPolicyResponse
	GetBody() *DeleteInstanceEndpointAclPolicyResponseBody
}

type DeleteInstanceEndpointAclPolicyResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteInstanceEndpointAclPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteInstanceEndpointAclPolicyResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteInstanceEndpointAclPolicyResponse) GoString() string {
	return s.String()
}

func (s *DeleteInstanceEndpointAclPolicyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteInstanceEndpointAclPolicyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteInstanceEndpointAclPolicyResponse) GetBody() *DeleteInstanceEndpointAclPolicyResponseBody {
	return s.Body
}

func (s *DeleteInstanceEndpointAclPolicyResponse) SetHeaders(v map[string]*string) *DeleteInstanceEndpointAclPolicyResponse {
	s.Headers = v
	return s
}

func (s *DeleteInstanceEndpointAclPolicyResponse) SetStatusCode(v int32) *DeleteInstanceEndpointAclPolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteInstanceEndpointAclPolicyResponse) SetBody(v *DeleteInstanceEndpointAclPolicyResponseBody) *DeleteInstanceEndpointAclPolicyResponse {
	s.Body = v
	return s
}

func (s *DeleteInstanceEndpointAclPolicyResponse) Validate() error {
	return dara.Validate(s)
}
