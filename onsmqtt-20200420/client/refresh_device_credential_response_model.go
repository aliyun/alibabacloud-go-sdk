// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRefreshDeviceCredentialResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RefreshDeviceCredentialResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RefreshDeviceCredentialResponse
	GetStatusCode() *int32
	SetBody(v *RefreshDeviceCredentialResponseBody) *RefreshDeviceCredentialResponse
	GetBody() *RefreshDeviceCredentialResponseBody
}

type RefreshDeviceCredentialResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RefreshDeviceCredentialResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RefreshDeviceCredentialResponse) String() string {
	return dara.Prettify(s)
}

func (s RefreshDeviceCredentialResponse) GoString() string {
	return s.String()
}

func (s *RefreshDeviceCredentialResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RefreshDeviceCredentialResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RefreshDeviceCredentialResponse) GetBody() *RefreshDeviceCredentialResponseBody {
	return s.Body
}

func (s *RefreshDeviceCredentialResponse) SetHeaders(v map[string]*string) *RefreshDeviceCredentialResponse {
	s.Headers = v
	return s
}

func (s *RefreshDeviceCredentialResponse) SetStatusCode(v int32) *RefreshDeviceCredentialResponse {
	s.StatusCode = &v
	return s
}

func (s *RefreshDeviceCredentialResponse) SetBody(v *RefreshDeviceCredentialResponseBody) *RefreshDeviceCredentialResponse {
	s.Body = v
	return s
}

func (s *RefreshDeviceCredentialResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
