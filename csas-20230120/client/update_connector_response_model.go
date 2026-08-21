// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateConnectorResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateConnectorResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateConnectorResponse
	GetStatusCode() *int32
	SetBody(v *UpdateConnectorResponseBody) *UpdateConnectorResponse
	GetBody() *UpdateConnectorResponseBody
}

type UpdateConnectorResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateConnectorResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateConnectorResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateConnectorResponse) GoString() string {
	return s.String()
}

func (s *UpdateConnectorResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateConnectorResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateConnectorResponse) GetBody() *UpdateConnectorResponseBody {
	return s.Body
}

func (s *UpdateConnectorResponse) SetHeaders(v map[string]*string) *UpdateConnectorResponse {
	s.Headers = v
	return s
}

func (s *UpdateConnectorResponse) SetStatusCode(v int32) *UpdateConnectorResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateConnectorResponse) SetBody(v *UpdateConnectorResponseBody) *UpdateConnectorResponse {
	s.Body = v
	return s
}

func (s *UpdateConnectorResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
