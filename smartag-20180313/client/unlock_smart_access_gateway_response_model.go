// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUnlockSmartAccessGatewayResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UnlockSmartAccessGatewayResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UnlockSmartAccessGatewayResponse
	GetStatusCode() *int32
	SetBody(v *UnlockSmartAccessGatewayResponseBody) *UnlockSmartAccessGatewayResponse
	GetBody() *UnlockSmartAccessGatewayResponseBody
}

type UnlockSmartAccessGatewayResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UnlockSmartAccessGatewayResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UnlockSmartAccessGatewayResponse) String() string {
	return dara.Prettify(s)
}

func (s UnlockSmartAccessGatewayResponse) GoString() string {
	return s.String()
}

func (s *UnlockSmartAccessGatewayResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UnlockSmartAccessGatewayResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UnlockSmartAccessGatewayResponse) GetBody() *UnlockSmartAccessGatewayResponseBody {
	return s.Body
}

func (s *UnlockSmartAccessGatewayResponse) SetHeaders(v map[string]*string) *UnlockSmartAccessGatewayResponse {
	s.Headers = v
	return s
}

func (s *UnlockSmartAccessGatewayResponse) SetStatusCode(v int32) *UnlockSmartAccessGatewayResponse {
	s.StatusCode = &v
	return s
}

func (s *UnlockSmartAccessGatewayResponse) SetBody(v *UnlockSmartAccessGatewayResponseBody) *UnlockSmartAccessGatewayResponse {
	s.Body = v
	return s
}

func (s *UnlockSmartAccessGatewayResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
