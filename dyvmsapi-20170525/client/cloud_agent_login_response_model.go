// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCloudAgentLoginResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CloudAgentLoginResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CloudAgentLoginResponse
	GetStatusCode() *int32
	SetBody(v *CloudAgentLoginResponseBody) *CloudAgentLoginResponse
	GetBody() *CloudAgentLoginResponseBody
}

type CloudAgentLoginResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CloudAgentLoginResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CloudAgentLoginResponse) String() string {
	return dara.Prettify(s)
}

func (s CloudAgentLoginResponse) GoString() string {
	return s.String()
}

func (s *CloudAgentLoginResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CloudAgentLoginResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CloudAgentLoginResponse) GetBody() *CloudAgentLoginResponseBody {
	return s.Body
}

func (s *CloudAgentLoginResponse) SetHeaders(v map[string]*string) *CloudAgentLoginResponse {
	s.Headers = v
	return s
}

func (s *CloudAgentLoginResponse) SetStatusCode(v int32) *CloudAgentLoginResponse {
	s.StatusCode = &v
	return s
}

func (s *CloudAgentLoginResponse) SetBody(v *CloudAgentLoginResponseBody) *CloudAgentLoginResponse {
	s.Body = v
	return s
}

func (s *CloudAgentLoginResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
