// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iVerifyWorkspaceOssMountRamAuthorizationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *VerifyWorkspaceOssMountRamAuthorizationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *VerifyWorkspaceOssMountRamAuthorizationResponse
	GetStatusCode() *int32
	SetBody(v *VerifyWorkspaceOssMountRamAuthorizationResponseBody) *VerifyWorkspaceOssMountRamAuthorizationResponse
	GetBody() *VerifyWorkspaceOssMountRamAuthorizationResponseBody
}

type VerifyWorkspaceOssMountRamAuthorizationResponse struct {
	Headers    map[string]*string                                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *VerifyWorkspaceOssMountRamAuthorizationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s VerifyWorkspaceOssMountRamAuthorizationResponse) String() string {
	return dara.Prettify(s)
}

func (s VerifyWorkspaceOssMountRamAuthorizationResponse) GoString() string {
	return s.String()
}

func (s *VerifyWorkspaceOssMountRamAuthorizationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *VerifyWorkspaceOssMountRamAuthorizationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *VerifyWorkspaceOssMountRamAuthorizationResponse) GetBody() *VerifyWorkspaceOssMountRamAuthorizationResponseBody {
	return s.Body
}

func (s *VerifyWorkspaceOssMountRamAuthorizationResponse) SetHeaders(v map[string]*string) *VerifyWorkspaceOssMountRamAuthorizationResponse {
	s.Headers = v
	return s
}

func (s *VerifyWorkspaceOssMountRamAuthorizationResponse) SetStatusCode(v int32) *VerifyWorkspaceOssMountRamAuthorizationResponse {
	s.StatusCode = &v
	return s
}

func (s *VerifyWorkspaceOssMountRamAuthorizationResponse) SetBody(v *VerifyWorkspaceOssMountRamAuthorizationResponseBody) *VerifyWorkspaceOssMountRamAuthorizationResponse {
	s.Body = v
	return s
}

func (s *VerifyWorkspaceOssMountRamAuthorizationResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
