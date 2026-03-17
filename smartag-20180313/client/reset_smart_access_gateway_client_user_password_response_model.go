// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResetSmartAccessGatewayClientUserPasswordResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ResetSmartAccessGatewayClientUserPasswordResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ResetSmartAccessGatewayClientUserPasswordResponse
	GetStatusCode() *int32
	SetBody(v *ResetSmartAccessGatewayClientUserPasswordResponseBody) *ResetSmartAccessGatewayClientUserPasswordResponse
	GetBody() *ResetSmartAccessGatewayClientUserPasswordResponseBody
}

type ResetSmartAccessGatewayClientUserPasswordResponse struct {
	Headers    map[string]*string                                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ResetSmartAccessGatewayClientUserPasswordResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ResetSmartAccessGatewayClientUserPasswordResponse) String() string {
	return dara.Prettify(s)
}

func (s ResetSmartAccessGatewayClientUserPasswordResponse) GoString() string {
	return s.String()
}

func (s *ResetSmartAccessGatewayClientUserPasswordResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ResetSmartAccessGatewayClientUserPasswordResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ResetSmartAccessGatewayClientUserPasswordResponse) GetBody() *ResetSmartAccessGatewayClientUserPasswordResponseBody {
	return s.Body
}

func (s *ResetSmartAccessGatewayClientUserPasswordResponse) SetHeaders(v map[string]*string) *ResetSmartAccessGatewayClientUserPasswordResponse {
	s.Headers = v
	return s
}

func (s *ResetSmartAccessGatewayClientUserPasswordResponse) SetStatusCode(v int32) *ResetSmartAccessGatewayClientUserPasswordResponse {
	s.StatusCode = &v
	return s
}

func (s *ResetSmartAccessGatewayClientUserPasswordResponse) SetBody(v *ResetSmartAccessGatewayClientUserPasswordResponseBody) *ResetSmartAccessGatewayClientUserPasswordResponse {
	s.Body = v
	return s
}

func (s *ResetSmartAccessGatewayClientUserPasswordResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
