// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDataConnectorResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateDataConnectorResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateDataConnectorResponse
	GetStatusCode() *int32
	SetBody(v *UpdateDataConnectorResponseBody) *UpdateDataConnectorResponse
	GetBody() *UpdateDataConnectorResponseBody
}

type UpdateDataConnectorResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateDataConnectorResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateDataConnectorResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateDataConnectorResponse) GoString() string {
	return s.String()
}

func (s *UpdateDataConnectorResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateDataConnectorResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateDataConnectorResponse) GetBody() *UpdateDataConnectorResponseBody {
	return s.Body
}

func (s *UpdateDataConnectorResponse) SetHeaders(v map[string]*string) *UpdateDataConnectorResponse {
	s.Headers = v
	return s
}

func (s *UpdateDataConnectorResponse) SetStatusCode(v int32) *UpdateDataConnectorResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateDataConnectorResponse) SetBody(v *UpdateDataConnectorResponseBody) *UpdateDataConnectorResponse {
	s.Body = v
	return s
}

func (s *UpdateDataConnectorResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
