// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteCustomConnectorResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteCustomConnectorResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteCustomConnectorResponse
	GetStatusCode() *int32
	SetBody(v *DeleteCustomConnectorResponseBody) *DeleteCustomConnectorResponse
	GetBody() *DeleteCustomConnectorResponseBody
}

type DeleteCustomConnectorResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteCustomConnectorResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteCustomConnectorResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteCustomConnectorResponse) GoString() string {
	return s.String()
}

func (s *DeleteCustomConnectorResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteCustomConnectorResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteCustomConnectorResponse) GetBody() *DeleteCustomConnectorResponseBody {
	return s.Body
}

func (s *DeleteCustomConnectorResponse) SetHeaders(v map[string]*string) *DeleteCustomConnectorResponse {
	s.Headers = v
	return s
}

func (s *DeleteCustomConnectorResponse) SetStatusCode(v int32) *DeleteCustomConnectorResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteCustomConnectorResponse) SetBody(v *DeleteCustomConnectorResponseBody) *DeleteCustomConnectorResponse {
	s.Body = v
	return s
}

func (s *DeleteCustomConnectorResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
