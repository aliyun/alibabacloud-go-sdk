// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSmartAccessGatewayResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateSmartAccessGatewayResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateSmartAccessGatewayResponse
	GetStatusCode() *int32
	SetBody(v *CreateSmartAccessGatewayResponseBody) *CreateSmartAccessGatewayResponse
	GetBody() *CreateSmartAccessGatewayResponseBody
}

type CreateSmartAccessGatewayResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateSmartAccessGatewayResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateSmartAccessGatewayResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateSmartAccessGatewayResponse) GoString() string {
	return s.String()
}

func (s *CreateSmartAccessGatewayResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateSmartAccessGatewayResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateSmartAccessGatewayResponse) GetBody() *CreateSmartAccessGatewayResponseBody {
	return s.Body
}

func (s *CreateSmartAccessGatewayResponse) SetHeaders(v map[string]*string) *CreateSmartAccessGatewayResponse {
	s.Headers = v
	return s
}

func (s *CreateSmartAccessGatewayResponse) SetStatusCode(v int32) *CreateSmartAccessGatewayResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateSmartAccessGatewayResponse) SetBody(v *CreateSmartAccessGatewayResponseBody) *CreateSmartAccessGatewayResponse {
	s.Body = v
	return s
}

func (s *CreateSmartAccessGatewayResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
