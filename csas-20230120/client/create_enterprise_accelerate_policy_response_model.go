// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateEnterpriseAcceleratePolicyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateEnterpriseAcceleratePolicyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateEnterpriseAcceleratePolicyResponse
	GetStatusCode() *int32
	SetBody(v *CreateEnterpriseAcceleratePolicyResponseBody) *CreateEnterpriseAcceleratePolicyResponse
	GetBody() *CreateEnterpriseAcceleratePolicyResponseBody
}

type CreateEnterpriseAcceleratePolicyResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateEnterpriseAcceleratePolicyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateEnterpriseAcceleratePolicyResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateEnterpriseAcceleratePolicyResponse) GoString() string {
	return s.String()
}

func (s *CreateEnterpriseAcceleratePolicyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateEnterpriseAcceleratePolicyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateEnterpriseAcceleratePolicyResponse) GetBody() *CreateEnterpriseAcceleratePolicyResponseBody {
	return s.Body
}

func (s *CreateEnterpriseAcceleratePolicyResponse) SetHeaders(v map[string]*string) *CreateEnterpriseAcceleratePolicyResponse {
	s.Headers = v
	return s
}

func (s *CreateEnterpriseAcceleratePolicyResponse) SetStatusCode(v int32) *CreateEnterpriseAcceleratePolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateEnterpriseAcceleratePolicyResponse) SetBody(v *CreateEnterpriseAcceleratePolicyResponseBody) *CreateEnterpriseAcceleratePolicyResponse {
	s.Body = v
	return s
}

func (s *CreateEnterpriseAcceleratePolicyResponse) Validate() error {
	return dara.Validate(s)
}
