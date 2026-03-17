// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifySmartAccessGatewayResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifySmartAccessGatewayResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifySmartAccessGatewayResponse
	GetStatusCode() *int32
	SetBody(v *ModifySmartAccessGatewayResponseBody) *ModifySmartAccessGatewayResponse
	GetBody() *ModifySmartAccessGatewayResponseBody
}

type ModifySmartAccessGatewayResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifySmartAccessGatewayResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifySmartAccessGatewayResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifySmartAccessGatewayResponse) GoString() string {
	return s.String()
}

func (s *ModifySmartAccessGatewayResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifySmartAccessGatewayResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifySmartAccessGatewayResponse) GetBody() *ModifySmartAccessGatewayResponseBody {
	return s.Body
}

func (s *ModifySmartAccessGatewayResponse) SetHeaders(v map[string]*string) *ModifySmartAccessGatewayResponse {
	s.Headers = v
	return s
}

func (s *ModifySmartAccessGatewayResponse) SetStatusCode(v int32) *ModifySmartAccessGatewayResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifySmartAccessGatewayResponse) SetBody(v *ModifySmartAccessGatewayResponseBody) *ModifySmartAccessGatewayResponse {
	s.Body = v
	return s
}

func (s *ModifySmartAccessGatewayResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
