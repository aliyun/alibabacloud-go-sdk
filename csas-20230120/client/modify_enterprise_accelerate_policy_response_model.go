// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyEnterpriseAcceleratePolicyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyEnterpriseAcceleratePolicyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyEnterpriseAcceleratePolicyResponse
	GetStatusCode() *int32
	SetBody(v *ModifyEnterpriseAcceleratePolicyResponseBody) *ModifyEnterpriseAcceleratePolicyResponse
	GetBody() *ModifyEnterpriseAcceleratePolicyResponseBody
}

type ModifyEnterpriseAcceleratePolicyResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyEnterpriseAcceleratePolicyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyEnterpriseAcceleratePolicyResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyEnterpriseAcceleratePolicyResponse) GoString() string {
	return s.String()
}

func (s *ModifyEnterpriseAcceleratePolicyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyEnterpriseAcceleratePolicyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyEnterpriseAcceleratePolicyResponse) GetBody() *ModifyEnterpriseAcceleratePolicyResponseBody {
	return s.Body
}

func (s *ModifyEnterpriseAcceleratePolicyResponse) SetHeaders(v map[string]*string) *ModifyEnterpriseAcceleratePolicyResponse {
	s.Headers = v
	return s
}

func (s *ModifyEnterpriseAcceleratePolicyResponse) SetStatusCode(v int32) *ModifyEnterpriseAcceleratePolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyEnterpriseAcceleratePolicyResponse) SetBody(v *ModifyEnterpriseAcceleratePolicyResponseBody) *ModifyEnterpriseAcceleratePolicyResponse {
	s.Body = v
	return s
}

func (s *ModifyEnterpriseAcceleratePolicyResponse) Validate() error {
	return dara.Validate(s)
}
