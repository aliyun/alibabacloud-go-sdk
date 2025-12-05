// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetPolicyApprovalConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SetPolicyApprovalConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SetPolicyApprovalConfigResponse
	GetStatusCode() *int32
	SetBody(v *SetPolicyApprovalConfigResponseBody) *SetPolicyApprovalConfigResponse
	GetBody() *SetPolicyApprovalConfigResponseBody
}

type SetPolicyApprovalConfigResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SetPolicyApprovalConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SetPolicyApprovalConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s SetPolicyApprovalConfigResponse) GoString() string {
	return s.String()
}

func (s *SetPolicyApprovalConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SetPolicyApprovalConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SetPolicyApprovalConfigResponse) GetBody() *SetPolicyApprovalConfigResponseBody {
	return s.Body
}

func (s *SetPolicyApprovalConfigResponse) SetHeaders(v map[string]*string) *SetPolicyApprovalConfigResponse {
	s.Headers = v
	return s
}

func (s *SetPolicyApprovalConfigResponse) SetStatusCode(v int32) *SetPolicyApprovalConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *SetPolicyApprovalConfigResponse) SetBody(v *SetPolicyApprovalConfigResponseBody) *SetPolicyApprovalConfigResponse {
	s.Body = v
	return s
}

func (s *SetPolicyApprovalConfigResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
