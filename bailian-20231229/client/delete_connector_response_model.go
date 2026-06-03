// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteConnectorResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteConnectorResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteConnectorResponse
	GetStatusCode() *int32
	SetBody(v *DeleteConnectorResponseBody) *DeleteConnectorResponse
	GetBody() *DeleteConnectorResponseBody
}

type DeleteConnectorResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteConnectorResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteConnectorResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteConnectorResponse) GoString() string {
	return s.String()
}

func (s *DeleteConnectorResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteConnectorResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteConnectorResponse) GetBody() *DeleteConnectorResponseBody {
	return s.Body
}

func (s *DeleteConnectorResponse) SetHeaders(v map[string]*string) *DeleteConnectorResponse {
	s.Headers = v
	return s
}

func (s *DeleteConnectorResponse) SetStatusCode(v int32) *DeleteConnectorResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteConnectorResponse) SetBody(v *DeleteConnectorResponseBody) *DeleteConnectorResponse {
	s.Body = v
	return s
}

func (s *DeleteConnectorResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
