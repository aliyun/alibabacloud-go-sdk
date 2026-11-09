// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDeviceCredentialResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetDeviceCredentialResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetDeviceCredentialResponse
	GetStatusCode() *int32
	SetBody(v *GetDeviceCredentialResponseBody) *GetDeviceCredentialResponse
	GetBody() *GetDeviceCredentialResponseBody
}

type GetDeviceCredentialResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetDeviceCredentialResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetDeviceCredentialResponse) String() string {
	return dara.Prettify(s)
}

func (s GetDeviceCredentialResponse) GoString() string {
	return s.String()
}

func (s *GetDeviceCredentialResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetDeviceCredentialResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetDeviceCredentialResponse) GetBody() *GetDeviceCredentialResponseBody {
	return s.Body
}

func (s *GetDeviceCredentialResponse) SetHeaders(v map[string]*string) *GetDeviceCredentialResponse {
	s.Headers = v
	return s
}

func (s *GetDeviceCredentialResponse) SetStatusCode(v int32) *GetDeviceCredentialResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDeviceCredentialResponse) SetBody(v *GetDeviceCredentialResponseBody) *GetDeviceCredentialResponse {
	s.Body = v
	return s
}

func (s *GetDeviceCredentialResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
