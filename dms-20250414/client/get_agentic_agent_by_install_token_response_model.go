// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAgenticAgentByInstallTokenResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetAgenticAgentByInstallTokenResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetAgenticAgentByInstallTokenResponse
	GetStatusCode() *int32
	SetBody(v *GetAgenticAgentByInstallTokenResponseBody) *GetAgenticAgentByInstallTokenResponse
	GetBody() *GetAgenticAgentByInstallTokenResponseBody
}

type GetAgenticAgentByInstallTokenResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAgenticAgentByInstallTokenResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAgenticAgentByInstallTokenResponse) String() string {
	return dara.Prettify(s)
}

func (s GetAgenticAgentByInstallTokenResponse) GoString() string {
	return s.String()
}

func (s *GetAgenticAgentByInstallTokenResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetAgenticAgentByInstallTokenResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetAgenticAgentByInstallTokenResponse) GetBody() *GetAgenticAgentByInstallTokenResponseBody {
	return s.Body
}

func (s *GetAgenticAgentByInstallTokenResponse) SetHeaders(v map[string]*string) *GetAgenticAgentByInstallTokenResponse {
	s.Headers = v
	return s
}

func (s *GetAgenticAgentByInstallTokenResponse) SetStatusCode(v int32) *GetAgenticAgentByInstallTokenResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAgenticAgentByInstallTokenResponse) SetBody(v *GetAgenticAgentByInstallTokenResponseBody) *GetAgenticAgentByInstallTokenResponse {
	s.Body = v
	return s
}

func (s *GetAgenticAgentByInstallTokenResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
