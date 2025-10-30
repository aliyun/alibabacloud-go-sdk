// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableEnterpriseAcceleratePolicyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DisableEnterpriseAcceleratePolicyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DisableEnterpriseAcceleratePolicyResponse
	GetStatusCode() *int32
	SetBody(v *DisableEnterpriseAcceleratePolicyResponseBody) *DisableEnterpriseAcceleratePolicyResponse
	GetBody() *DisableEnterpriseAcceleratePolicyResponseBody
}

type DisableEnterpriseAcceleratePolicyResponse struct {
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DisableEnterpriseAcceleratePolicyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DisableEnterpriseAcceleratePolicyResponse) String() string {
	return dara.Prettify(s)
}

func (s DisableEnterpriseAcceleratePolicyResponse) GoString() string {
	return s.String()
}

func (s *DisableEnterpriseAcceleratePolicyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DisableEnterpriseAcceleratePolicyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DisableEnterpriseAcceleratePolicyResponse) GetBody() *DisableEnterpriseAcceleratePolicyResponseBody {
	return s.Body
}

func (s *DisableEnterpriseAcceleratePolicyResponse) SetHeaders(v map[string]*string) *DisableEnterpriseAcceleratePolicyResponse {
	s.Headers = v
	return s
}

func (s *DisableEnterpriseAcceleratePolicyResponse) SetStatusCode(v int32) *DisableEnterpriseAcceleratePolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *DisableEnterpriseAcceleratePolicyResponse) SetBody(v *DisableEnterpriseAcceleratePolicyResponseBody) *DisableEnterpriseAcceleratePolicyResponse {
	s.Body = v
	return s
}

func (s *DisableEnterpriseAcceleratePolicyResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
