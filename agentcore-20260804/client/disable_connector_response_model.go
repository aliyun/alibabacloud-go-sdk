// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableConnectorResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DisableConnectorResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DisableConnectorResponse
	GetStatusCode() *int32
	SetBody(v *DisableConnectorResponseBody) *DisableConnectorResponse
	GetBody() *DisableConnectorResponseBody
}

type DisableConnectorResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DisableConnectorResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DisableConnectorResponse) String() string {
	return dara.Prettify(s)
}

func (s DisableConnectorResponse) GoString() string {
	return s.String()
}

func (s *DisableConnectorResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DisableConnectorResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DisableConnectorResponse) GetBody() *DisableConnectorResponseBody {
	return s.Body
}

func (s *DisableConnectorResponse) SetHeaders(v map[string]*string) *DisableConnectorResponse {
	s.Headers = v
	return s
}

func (s *DisableConnectorResponse) SetStatusCode(v int32) *DisableConnectorResponse {
	s.StatusCode = &v
	return s
}

func (s *DisableConnectorResponse) SetBody(v *DisableConnectorResponseBody) *DisableConnectorResponse {
	s.Body = v
	return s
}

func (s *DisableConnectorResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
