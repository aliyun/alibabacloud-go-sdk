// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyResourceCenterPolicyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyResourceCenterPolicyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyResourceCenterPolicyResponse
	GetStatusCode() *int32
	SetBody(v *ModifyResourceCenterPolicyResponseBody) *ModifyResourceCenterPolicyResponse
	GetBody() *ModifyResourceCenterPolicyResponseBody
}

type ModifyResourceCenterPolicyResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyResourceCenterPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyResourceCenterPolicyResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyResourceCenterPolicyResponse) GoString() string {
	return s.String()
}

func (s *ModifyResourceCenterPolicyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyResourceCenterPolicyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyResourceCenterPolicyResponse) GetBody() *ModifyResourceCenterPolicyResponseBody {
	return s.Body
}

func (s *ModifyResourceCenterPolicyResponse) SetHeaders(v map[string]*string) *ModifyResourceCenterPolicyResponse {
	s.Headers = v
	return s
}

func (s *ModifyResourceCenterPolicyResponse) SetStatusCode(v int32) *ModifyResourceCenterPolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyResourceCenterPolicyResponse) SetBody(v *ModifyResourceCenterPolicyResponseBody) *ModifyResourceCenterPolicyResponse {
	s.Body = v
	return s
}

func (s *ModifyResourceCenterPolicyResponse) Validate() error {
	return dara.Validate(s)
}
