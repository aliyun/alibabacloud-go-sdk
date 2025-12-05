// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iConnectKmsInstanceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ConnectKmsInstanceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ConnectKmsInstanceResponse
	GetStatusCode() *int32
	SetBody(v *ConnectKmsInstanceResponseBody) *ConnectKmsInstanceResponse
	GetBody() *ConnectKmsInstanceResponseBody
}

type ConnectKmsInstanceResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ConnectKmsInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ConnectKmsInstanceResponse) String() string {
	return dara.Prettify(s)
}

func (s ConnectKmsInstanceResponse) GoString() string {
	return s.String()
}

func (s *ConnectKmsInstanceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ConnectKmsInstanceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ConnectKmsInstanceResponse) GetBody() *ConnectKmsInstanceResponseBody {
	return s.Body
}

func (s *ConnectKmsInstanceResponse) SetHeaders(v map[string]*string) *ConnectKmsInstanceResponse {
	s.Headers = v
	return s
}

func (s *ConnectKmsInstanceResponse) SetStatusCode(v int32) *ConnectKmsInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *ConnectKmsInstanceResponse) SetBody(v *ConnectKmsInstanceResponseBody) *ConnectKmsInstanceResponse {
	s.Body = v
	return s
}

func (s *ConnectKmsInstanceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
