// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteStorageGatewayResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteStorageGatewayResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteStorageGatewayResponse
	GetStatusCode() *int32
	SetBody(v *DeleteStorageGatewayResponseBody) *DeleteStorageGatewayResponse
	GetBody() *DeleteStorageGatewayResponseBody
}

type DeleteStorageGatewayResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteStorageGatewayResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteStorageGatewayResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteStorageGatewayResponse) GoString() string {
	return s.String()
}

func (s *DeleteStorageGatewayResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteStorageGatewayResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteStorageGatewayResponse) GetBody() *DeleteStorageGatewayResponseBody {
	return s.Body
}

func (s *DeleteStorageGatewayResponse) SetHeaders(v map[string]*string) *DeleteStorageGatewayResponse {
	s.Headers = v
	return s
}

func (s *DeleteStorageGatewayResponse) SetStatusCode(v int32) *DeleteStorageGatewayResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteStorageGatewayResponse) SetBody(v *DeleteStorageGatewayResponseBody) *DeleteStorageGatewayResponse {
	s.Body = v
	return s
}

func (s *DeleteStorageGatewayResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
