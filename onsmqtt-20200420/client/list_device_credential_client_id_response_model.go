// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDeviceCredentialClientIdResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListDeviceCredentialClientIdResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListDeviceCredentialClientIdResponse
	GetStatusCode() *int32
	SetBody(v *ListDeviceCredentialClientIdResponseBody) *ListDeviceCredentialClientIdResponse
	GetBody() *ListDeviceCredentialClientIdResponseBody
}

type ListDeviceCredentialClientIdResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListDeviceCredentialClientIdResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListDeviceCredentialClientIdResponse) String() string {
	return dara.Prettify(s)
}

func (s ListDeviceCredentialClientIdResponse) GoString() string {
	return s.String()
}

func (s *ListDeviceCredentialClientIdResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListDeviceCredentialClientIdResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListDeviceCredentialClientIdResponse) GetBody() *ListDeviceCredentialClientIdResponseBody {
	return s.Body
}

func (s *ListDeviceCredentialClientIdResponse) SetHeaders(v map[string]*string) *ListDeviceCredentialClientIdResponse {
	s.Headers = v
	return s
}

func (s *ListDeviceCredentialClientIdResponse) SetStatusCode(v int32) *ListDeviceCredentialClientIdResponse {
	s.StatusCode = &v
	return s
}

func (s *ListDeviceCredentialClientIdResponse) SetBody(v *ListDeviceCredentialClientIdResponseBody) *ListDeviceCredentialClientIdResponse {
	s.Body = v
	return s
}

func (s *ListDeviceCredentialClientIdResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
