// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRegisterCustomConnectorResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RegisterCustomConnectorResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RegisterCustomConnectorResponse
	GetStatusCode() *int32
	SetBody(v *RegisterCustomConnectorResponseBody) *RegisterCustomConnectorResponse
	GetBody() *RegisterCustomConnectorResponseBody
}

type RegisterCustomConnectorResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RegisterCustomConnectorResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RegisterCustomConnectorResponse) String() string {
	return dara.Prettify(s)
}

func (s RegisterCustomConnectorResponse) GoString() string {
	return s.String()
}

func (s *RegisterCustomConnectorResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RegisterCustomConnectorResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RegisterCustomConnectorResponse) GetBody() *RegisterCustomConnectorResponseBody {
	return s.Body
}

func (s *RegisterCustomConnectorResponse) SetHeaders(v map[string]*string) *RegisterCustomConnectorResponse {
	s.Headers = v
	return s
}

func (s *RegisterCustomConnectorResponse) SetStatusCode(v int32) *RegisterCustomConnectorResponse {
	s.StatusCode = &v
	return s
}

func (s *RegisterCustomConnectorResponse) SetBody(v *RegisterCustomConnectorResponseBody) *RegisterCustomConnectorResponse {
	s.Body = v
	return s
}

func (s *RegisterCustomConnectorResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
