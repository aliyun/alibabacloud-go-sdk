// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRebootSmartAccessGatewayResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RebootSmartAccessGatewayResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RebootSmartAccessGatewayResponse
	GetStatusCode() *int32
	SetBody(v *RebootSmartAccessGatewayResponseBody) *RebootSmartAccessGatewayResponse
	GetBody() *RebootSmartAccessGatewayResponseBody
}

type RebootSmartAccessGatewayResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RebootSmartAccessGatewayResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RebootSmartAccessGatewayResponse) String() string {
	return dara.Prettify(s)
}

func (s RebootSmartAccessGatewayResponse) GoString() string {
	return s.String()
}

func (s *RebootSmartAccessGatewayResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RebootSmartAccessGatewayResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RebootSmartAccessGatewayResponse) GetBody() *RebootSmartAccessGatewayResponseBody {
	return s.Body
}

func (s *RebootSmartAccessGatewayResponse) SetHeaders(v map[string]*string) *RebootSmartAccessGatewayResponse {
	s.Headers = v
	return s
}

func (s *RebootSmartAccessGatewayResponse) SetStatusCode(v int32) *RebootSmartAccessGatewayResponse {
	s.StatusCode = &v
	return s
}

func (s *RebootSmartAccessGatewayResponse) SetBody(v *RebootSmartAccessGatewayResponseBody) *RebootSmartAccessGatewayResponse {
	s.Body = v
	return s
}

func (s *RebootSmartAccessGatewayResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
