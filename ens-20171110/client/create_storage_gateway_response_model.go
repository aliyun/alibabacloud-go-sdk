// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateStorageGatewayResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateStorageGatewayResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateStorageGatewayResponse
	GetStatusCode() *int32
	SetBody(v *CreateStorageGatewayResponseBody) *CreateStorageGatewayResponse
	GetBody() *CreateStorageGatewayResponseBody
}

type CreateStorageGatewayResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateStorageGatewayResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateStorageGatewayResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateStorageGatewayResponse) GoString() string {
	return s.String()
}

func (s *CreateStorageGatewayResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateStorageGatewayResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateStorageGatewayResponse) GetBody() *CreateStorageGatewayResponseBody {
	return s.Body
}

func (s *CreateStorageGatewayResponse) SetHeaders(v map[string]*string) *CreateStorageGatewayResponse {
	s.Headers = v
	return s
}

func (s *CreateStorageGatewayResponse) SetStatusCode(v int32) *CreateStorageGatewayResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateStorageGatewayResponse) SetBody(v *CreateStorageGatewayResponseBody) *CreateStorageGatewayResponse {
	s.Body = v
	return s
}

func (s *CreateStorageGatewayResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
