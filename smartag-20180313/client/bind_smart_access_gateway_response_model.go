// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBindSmartAccessGatewayResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *BindSmartAccessGatewayResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *BindSmartAccessGatewayResponse
	GetStatusCode() *int32
	SetBody(v *BindSmartAccessGatewayResponseBody) *BindSmartAccessGatewayResponse
	GetBody() *BindSmartAccessGatewayResponseBody
}

type BindSmartAccessGatewayResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *BindSmartAccessGatewayResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s BindSmartAccessGatewayResponse) String() string {
	return dara.Prettify(s)
}

func (s BindSmartAccessGatewayResponse) GoString() string {
	return s.String()
}

func (s *BindSmartAccessGatewayResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *BindSmartAccessGatewayResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *BindSmartAccessGatewayResponse) GetBody() *BindSmartAccessGatewayResponseBody {
	return s.Body
}

func (s *BindSmartAccessGatewayResponse) SetHeaders(v map[string]*string) *BindSmartAccessGatewayResponse {
	s.Headers = v
	return s
}

func (s *BindSmartAccessGatewayResponse) SetStatusCode(v int32) *BindSmartAccessGatewayResponse {
	s.StatusCode = &v
	return s
}

func (s *BindSmartAccessGatewayResponse) SetBody(v *BindSmartAccessGatewayResponseBody) *BindSmartAccessGatewayResponse {
	s.Body = v
	return s
}

func (s *BindSmartAccessGatewayResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
