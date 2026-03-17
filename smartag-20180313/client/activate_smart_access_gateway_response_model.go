// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iActivateSmartAccessGatewayResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ActivateSmartAccessGatewayResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ActivateSmartAccessGatewayResponse
	GetStatusCode() *int32
	SetBody(v *ActivateSmartAccessGatewayResponseBody) *ActivateSmartAccessGatewayResponse
	GetBody() *ActivateSmartAccessGatewayResponseBody
}

type ActivateSmartAccessGatewayResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ActivateSmartAccessGatewayResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ActivateSmartAccessGatewayResponse) String() string {
	return dara.Prettify(s)
}

func (s ActivateSmartAccessGatewayResponse) GoString() string {
	return s.String()
}

func (s *ActivateSmartAccessGatewayResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ActivateSmartAccessGatewayResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ActivateSmartAccessGatewayResponse) GetBody() *ActivateSmartAccessGatewayResponseBody {
	return s.Body
}

func (s *ActivateSmartAccessGatewayResponse) SetHeaders(v map[string]*string) *ActivateSmartAccessGatewayResponse {
	s.Headers = v
	return s
}

func (s *ActivateSmartAccessGatewayResponse) SetStatusCode(v int32) *ActivateSmartAccessGatewayResponse {
	s.StatusCode = &v
	return s
}

func (s *ActivateSmartAccessGatewayResponse) SetBody(v *ActivateSmartAccessGatewayResponseBody) *ActivateSmartAccessGatewayResponse {
	s.Body = v
	return s
}

func (s *ActivateSmartAccessGatewayResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
