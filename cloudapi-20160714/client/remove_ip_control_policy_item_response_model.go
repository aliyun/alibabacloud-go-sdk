// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveIpControlPolicyItemResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RemoveIpControlPolicyItemResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RemoveIpControlPolicyItemResponse
	GetStatusCode() *int32
	SetBody(v *RemoveIpControlPolicyItemResponseBody) *RemoveIpControlPolicyItemResponse
	GetBody() *RemoveIpControlPolicyItemResponseBody
}

type RemoveIpControlPolicyItemResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RemoveIpControlPolicyItemResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RemoveIpControlPolicyItemResponse) String() string {
	return dara.Prettify(s)
}

func (s RemoveIpControlPolicyItemResponse) GoString() string {
	return s.String()
}

func (s *RemoveIpControlPolicyItemResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RemoveIpControlPolicyItemResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RemoveIpControlPolicyItemResponse) GetBody() *RemoveIpControlPolicyItemResponseBody {
	return s.Body
}

func (s *RemoveIpControlPolicyItemResponse) SetHeaders(v map[string]*string) *RemoveIpControlPolicyItemResponse {
	s.Headers = v
	return s
}

func (s *RemoveIpControlPolicyItemResponse) SetStatusCode(v int32) *RemoveIpControlPolicyItemResponse {
	s.StatusCode = &v
	return s
}

func (s *RemoveIpControlPolicyItemResponse) SetBody(v *RemoveIpControlPolicyItemResponseBody) *RemoveIpControlPolicyItemResponse {
	s.Body = v
	return s
}

func (s *RemoveIpControlPolicyItemResponse) Validate() error {
	return dara.Validate(s)
}
