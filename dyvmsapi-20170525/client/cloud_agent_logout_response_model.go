// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCloudAgentLogoutResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CloudAgentLogoutResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CloudAgentLogoutResponse
	GetStatusCode() *int32
	SetBody(v *CloudAgentLogoutResponseBody) *CloudAgentLogoutResponse
	GetBody() *CloudAgentLogoutResponseBody
}

type CloudAgentLogoutResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CloudAgentLogoutResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CloudAgentLogoutResponse) String() string {
	return dara.Prettify(s)
}

func (s CloudAgentLogoutResponse) GoString() string {
	return s.String()
}

func (s *CloudAgentLogoutResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CloudAgentLogoutResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CloudAgentLogoutResponse) GetBody() *CloudAgentLogoutResponseBody {
	return s.Body
}

func (s *CloudAgentLogoutResponse) SetHeaders(v map[string]*string) *CloudAgentLogoutResponse {
	s.Headers = v
	return s
}

func (s *CloudAgentLogoutResponse) SetStatusCode(v int32) *CloudAgentLogoutResponse {
	s.StatusCode = &v
	return s
}

func (s *CloudAgentLogoutResponse) SetBody(v *CloudAgentLogoutResponseBody) *CloudAgentLogoutResponse {
	s.Body = v
	return s
}

func (s *CloudAgentLogoutResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
