// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddIpControlPolicyItemResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddIpControlPolicyItemResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddIpControlPolicyItemResponse
	GetStatusCode() *int32
	SetBody(v *AddIpControlPolicyItemResponseBody) *AddIpControlPolicyItemResponse
	GetBody() *AddIpControlPolicyItemResponseBody
}

type AddIpControlPolicyItemResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddIpControlPolicyItemResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddIpControlPolicyItemResponse) String() string {
	return dara.Prettify(s)
}

func (s AddIpControlPolicyItemResponse) GoString() string {
	return s.String()
}

func (s *AddIpControlPolicyItemResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddIpControlPolicyItemResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddIpControlPolicyItemResponse) GetBody() *AddIpControlPolicyItemResponseBody {
	return s.Body
}

func (s *AddIpControlPolicyItemResponse) SetHeaders(v map[string]*string) *AddIpControlPolicyItemResponse {
	s.Headers = v
	return s
}

func (s *AddIpControlPolicyItemResponse) SetStatusCode(v int32) *AddIpControlPolicyItemResponse {
	s.StatusCode = &v
	return s
}

func (s *AddIpControlPolicyItemResponse) SetBody(v *AddIpControlPolicyItemResponseBody) *AddIpControlPolicyItemResponse {
	s.Body = v
	return s
}

func (s *AddIpControlPolicyItemResponse) Validate() error {
	return dara.Validate(s)
}
