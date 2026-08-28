// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateExternalAgentBootstrapTokenResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateExternalAgentBootstrapTokenResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateExternalAgentBootstrapTokenResponse
	GetStatusCode() *int32
	SetBody(v *CreateExternalAgentBootstrapTokenResponseBody) *CreateExternalAgentBootstrapTokenResponse
	GetBody() *CreateExternalAgentBootstrapTokenResponseBody
}

type CreateExternalAgentBootstrapTokenResponse struct {
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateExternalAgentBootstrapTokenResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateExternalAgentBootstrapTokenResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateExternalAgentBootstrapTokenResponse) GoString() string {
	return s.String()
}

func (s *CreateExternalAgentBootstrapTokenResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateExternalAgentBootstrapTokenResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateExternalAgentBootstrapTokenResponse) GetBody() *CreateExternalAgentBootstrapTokenResponseBody {
	return s.Body
}

func (s *CreateExternalAgentBootstrapTokenResponse) SetHeaders(v map[string]*string) *CreateExternalAgentBootstrapTokenResponse {
	s.Headers = v
	return s
}

func (s *CreateExternalAgentBootstrapTokenResponse) SetStatusCode(v int32) *CreateExternalAgentBootstrapTokenResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateExternalAgentBootstrapTokenResponse) SetBody(v *CreateExternalAgentBootstrapTokenResponseBody) *CreateExternalAgentBootstrapTokenResponse {
	s.Body = v
	return s
}

func (s *CreateExternalAgentBootstrapTokenResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
