// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyIpControlPolicyItemResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyIpControlPolicyItemResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyIpControlPolicyItemResponse
	GetStatusCode() *int32
	SetBody(v *ModifyIpControlPolicyItemResponseBody) *ModifyIpControlPolicyItemResponse
	GetBody() *ModifyIpControlPolicyItemResponseBody
}

type ModifyIpControlPolicyItemResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyIpControlPolicyItemResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyIpControlPolicyItemResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyIpControlPolicyItemResponse) GoString() string {
	return s.String()
}

func (s *ModifyIpControlPolicyItemResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyIpControlPolicyItemResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyIpControlPolicyItemResponse) GetBody() *ModifyIpControlPolicyItemResponseBody {
	return s.Body
}

func (s *ModifyIpControlPolicyItemResponse) SetHeaders(v map[string]*string) *ModifyIpControlPolicyItemResponse {
	s.Headers = v
	return s
}

func (s *ModifyIpControlPolicyItemResponse) SetStatusCode(v int32) *ModifyIpControlPolicyItemResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyIpControlPolicyItemResponse) SetBody(v *ModifyIpControlPolicyItemResponseBody) *ModifyIpControlPolicyItemResponse {
	s.Body = v
	return s
}

func (s *ModifyIpControlPolicyItemResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
