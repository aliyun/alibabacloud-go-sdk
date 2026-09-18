// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iVerifyWorkspaceAcrRamAuthorizationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *VerifyWorkspaceAcrRamAuthorizationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *VerifyWorkspaceAcrRamAuthorizationResponse
	GetStatusCode() *int32
	SetBody(v *VerifyWorkspaceAcrRamAuthorizationResponseBody) *VerifyWorkspaceAcrRamAuthorizationResponse
	GetBody() *VerifyWorkspaceAcrRamAuthorizationResponseBody
}

type VerifyWorkspaceAcrRamAuthorizationResponse struct {
	Headers    map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *VerifyWorkspaceAcrRamAuthorizationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s VerifyWorkspaceAcrRamAuthorizationResponse) String() string {
	return dara.Prettify(s)
}

func (s VerifyWorkspaceAcrRamAuthorizationResponse) GoString() string {
	return s.String()
}

func (s *VerifyWorkspaceAcrRamAuthorizationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *VerifyWorkspaceAcrRamAuthorizationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *VerifyWorkspaceAcrRamAuthorizationResponse) GetBody() *VerifyWorkspaceAcrRamAuthorizationResponseBody {
	return s.Body
}

func (s *VerifyWorkspaceAcrRamAuthorizationResponse) SetHeaders(v map[string]*string) *VerifyWorkspaceAcrRamAuthorizationResponse {
	s.Headers = v
	return s
}

func (s *VerifyWorkspaceAcrRamAuthorizationResponse) SetStatusCode(v int32) *VerifyWorkspaceAcrRamAuthorizationResponse {
	s.StatusCode = &v
	return s
}

func (s *VerifyWorkspaceAcrRamAuthorizationResponse) SetBody(v *VerifyWorkspaceAcrRamAuthorizationResponseBody) *VerifyWorkspaceAcrRamAuthorizationResponse {
	s.Body = v
	return s
}

func (s *VerifyWorkspaceAcrRamAuthorizationResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
