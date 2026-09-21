// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iVerifyWorkspaceAgenticFsMountRamAuthorizationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *VerifyWorkspaceAgenticFsMountRamAuthorizationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *VerifyWorkspaceAgenticFsMountRamAuthorizationResponse
	GetStatusCode() *int32
	SetBody(v *VerifyWorkspaceAgenticFsMountRamAuthorizationResponseBody) *VerifyWorkspaceAgenticFsMountRamAuthorizationResponse
	GetBody() *VerifyWorkspaceAgenticFsMountRamAuthorizationResponseBody
}

type VerifyWorkspaceAgenticFsMountRamAuthorizationResponse struct {
	Headers    map[string]*string                                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *VerifyWorkspaceAgenticFsMountRamAuthorizationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s VerifyWorkspaceAgenticFsMountRamAuthorizationResponse) String() string {
	return dara.Prettify(s)
}

func (s VerifyWorkspaceAgenticFsMountRamAuthorizationResponse) GoString() string {
	return s.String()
}

func (s *VerifyWorkspaceAgenticFsMountRamAuthorizationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *VerifyWorkspaceAgenticFsMountRamAuthorizationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *VerifyWorkspaceAgenticFsMountRamAuthorizationResponse) GetBody() *VerifyWorkspaceAgenticFsMountRamAuthorizationResponseBody {
	return s.Body
}

func (s *VerifyWorkspaceAgenticFsMountRamAuthorizationResponse) SetHeaders(v map[string]*string) *VerifyWorkspaceAgenticFsMountRamAuthorizationResponse {
	s.Headers = v
	return s
}

func (s *VerifyWorkspaceAgenticFsMountRamAuthorizationResponse) SetStatusCode(v int32) *VerifyWorkspaceAgenticFsMountRamAuthorizationResponse {
	s.StatusCode = &v
	return s
}

func (s *VerifyWorkspaceAgenticFsMountRamAuthorizationResponse) SetBody(v *VerifyWorkspaceAgenticFsMountRamAuthorizationResponseBody) *VerifyWorkspaceAgenticFsMountRamAuthorizationResponse {
	s.Body = v
	return s
}

func (s *VerifyWorkspaceAgenticFsMountRamAuthorizationResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
