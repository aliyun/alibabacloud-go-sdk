// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iVerifyAgentCodePackageOssRamAuthorizationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *VerifyAgentCodePackageOssRamAuthorizationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *VerifyAgentCodePackageOssRamAuthorizationResponse
	GetStatusCode() *int32
	SetBody(v *VerifyAgentCodePackageOssRamAuthorizationResponseBody) *VerifyAgentCodePackageOssRamAuthorizationResponse
	GetBody() *VerifyAgentCodePackageOssRamAuthorizationResponseBody
}

type VerifyAgentCodePackageOssRamAuthorizationResponse struct {
	Headers    map[string]*string                                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *VerifyAgentCodePackageOssRamAuthorizationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s VerifyAgentCodePackageOssRamAuthorizationResponse) String() string {
	return dara.Prettify(s)
}

func (s VerifyAgentCodePackageOssRamAuthorizationResponse) GoString() string {
	return s.String()
}

func (s *VerifyAgentCodePackageOssRamAuthorizationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *VerifyAgentCodePackageOssRamAuthorizationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *VerifyAgentCodePackageOssRamAuthorizationResponse) GetBody() *VerifyAgentCodePackageOssRamAuthorizationResponseBody {
	return s.Body
}

func (s *VerifyAgentCodePackageOssRamAuthorizationResponse) SetHeaders(v map[string]*string) *VerifyAgentCodePackageOssRamAuthorizationResponse {
	s.Headers = v
	return s
}

func (s *VerifyAgentCodePackageOssRamAuthorizationResponse) SetStatusCode(v int32) *VerifyAgentCodePackageOssRamAuthorizationResponse {
	s.StatusCode = &v
	return s
}

func (s *VerifyAgentCodePackageOssRamAuthorizationResponse) SetBody(v *VerifyAgentCodePackageOssRamAuthorizationResponseBody) *VerifyAgentCodePackageOssRamAuthorizationResponse {
	s.Body = v
	return s
}

func (s *VerifyAgentCodePackageOssRamAuthorizationResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
