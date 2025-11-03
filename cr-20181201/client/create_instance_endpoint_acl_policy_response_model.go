// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateInstanceEndpointAclPolicyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateInstanceEndpointAclPolicyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateInstanceEndpointAclPolicyResponse
	GetStatusCode() *int32
	SetBody(v *CreateInstanceEndpointAclPolicyResponseBody) *CreateInstanceEndpointAclPolicyResponse
	GetBody() *CreateInstanceEndpointAclPolicyResponseBody
}

type CreateInstanceEndpointAclPolicyResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateInstanceEndpointAclPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateInstanceEndpointAclPolicyResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateInstanceEndpointAclPolicyResponse) GoString() string {
	return s.String()
}

func (s *CreateInstanceEndpointAclPolicyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateInstanceEndpointAclPolicyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateInstanceEndpointAclPolicyResponse) GetBody() *CreateInstanceEndpointAclPolicyResponseBody {
	return s.Body
}

func (s *CreateInstanceEndpointAclPolicyResponse) SetHeaders(v map[string]*string) *CreateInstanceEndpointAclPolicyResponse {
	s.Headers = v
	return s
}

func (s *CreateInstanceEndpointAclPolicyResponse) SetStatusCode(v int32) *CreateInstanceEndpointAclPolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateInstanceEndpointAclPolicyResponse) SetBody(v *CreateInstanceEndpointAclPolicyResponseBody) *CreateInstanceEndpointAclPolicyResponse {
	s.Body = v
	return s
}

func (s *CreateInstanceEndpointAclPolicyResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
