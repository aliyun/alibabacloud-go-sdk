// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyApiGroupNetworkPolicyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyApiGroupNetworkPolicyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyApiGroupNetworkPolicyResponse
	GetStatusCode() *int32
	SetBody(v *ModifyApiGroupNetworkPolicyResponseBody) *ModifyApiGroupNetworkPolicyResponse
	GetBody() *ModifyApiGroupNetworkPolicyResponseBody
}

type ModifyApiGroupNetworkPolicyResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyApiGroupNetworkPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyApiGroupNetworkPolicyResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyApiGroupNetworkPolicyResponse) GoString() string {
	return s.String()
}

func (s *ModifyApiGroupNetworkPolicyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyApiGroupNetworkPolicyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyApiGroupNetworkPolicyResponse) GetBody() *ModifyApiGroupNetworkPolicyResponseBody {
	return s.Body
}

func (s *ModifyApiGroupNetworkPolicyResponse) SetHeaders(v map[string]*string) *ModifyApiGroupNetworkPolicyResponse {
	s.Headers = v
	return s
}

func (s *ModifyApiGroupNetworkPolicyResponse) SetStatusCode(v int32) *ModifyApiGroupNetworkPolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyApiGroupNetworkPolicyResponse) SetBody(v *ModifyApiGroupNetworkPolicyResponseBody) *ModifyApiGroupNetworkPolicyResponse {
	s.Body = v
	return s
}

func (s *ModifyApiGroupNetworkPolicyResponse) Validate() error {
	return dara.Validate(s)
}
