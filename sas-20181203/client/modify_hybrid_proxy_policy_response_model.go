// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyHybridProxyPolicyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyHybridProxyPolicyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyHybridProxyPolicyResponse
	GetStatusCode() *int32
	SetBody(v *ModifyHybridProxyPolicyResponseBody) *ModifyHybridProxyPolicyResponse
	GetBody() *ModifyHybridProxyPolicyResponseBody
}

type ModifyHybridProxyPolicyResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyHybridProxyPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyHybridProxyPolicyResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyHybridProxyPolicyResponse) GoString() string {
	return s.String()
}

func (s *ModifyHybridProxyPolicyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyHybridProxyPolicyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyHybridProxyPolicyResponse) GetBody() *ModifyHybridProxyPolicyResponseBody {
	return s.Body
}

func (s *ModifyHybridProxyPolicyResponse) SetHeaders(v map[string]*string) *ModifyHybridProxyPolicyResponse {
	s.Headers = v
	return s
}

func (s *ModifyHybridProxyPolicyResponse) SetStatusCode(v int32) *ModifyHybridProxyPolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyHybridProxyPolicyResponse) SetBody(v *ModifyHybridProxyPolicyResponseBody) *ModifyHybridProxyPolicyResponse {
	s.Body = v
	return s
}

func (s *ModifyHybridProxyPolicyResponse) Validate() error {
	return dara.Validate(s)
}
