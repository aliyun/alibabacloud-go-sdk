// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUnbindSmartAccessGatewayResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UnbindSmartAccessGatewayResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UnbindSmartAccessGatewayResponse
	GetStatusCode() *int32
	SetBody(v *UnbindSmartAccessGatewayResponseBody) *UnbindSmartAccessGatewayResponse
	GetBody() *UnbindSmartAccessGatewayResponseBody
}

type UnbindSmartAccessGatewayResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UnbindSmartAccessGatewayResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UnbindSmartAccessGatewayResponse) String() string {
	return dara.Prettify(s)
}

func (s UnbindSmartAccessGatewayResponse) GoString() string {
	return s.String()
}

func (s *UnbindSmartAccessGatewayResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UnbindSmartAccessGatewayResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UnbindSmartAccessGatewayResponse) GetBody() *UnbindSmartAccessGatewayResponseBody {
	return s.Body
}

func (s *UnbindSmartAccessGatewayResponse) SetHeaders(v map[string]*string) *UnbindSmartAccessGatewayResponse {
	s.Headers = v
	return s
}

func (s *UnbindSmartAccessGatewayResponse) SetStatusCode(v int32) *UnbindSmartAccessGatewayResponse {
	s.StatusCode = &v
	return s
}

func (s *UnbindSmartAccessGatewayResponse) SetBody(v *UnbindSmartAccessGatewayResponseBody) *UnbindSmartAccessGatewayResponse {
	s.Body = v
	return s
}

func (s *UnbindSmartAccessGatewayResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
