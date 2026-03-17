// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifySmartAccessGatewayClientUserResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifySmartAccessGatewayClientUserResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifySmartAccessGatewayClientUserResponse
	GetStatusCode() *int32
	SetBody(v *ModifySmartAccessGatewayClientUserResponseBody) *ModifySmartAccessGatewayClientUserResponse
	GetBody() *ModifySmartAccessGatewayClientUserResponseBody
}

type ModifySmartAccessGatewayClientUserResponse struct {
	Headers    map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifySmartAccessGatewayClientUserResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifySmartAccessGatewayClientUserResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifySmartAccessGatewayClientUserResponse) GoString() string {
	return s.String()
}

func (s *ModifySmartAccessGatewayClientUserResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifySmartAccessGatewayClientUserResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifySmartAccessGatewayClientUserResponse) GetBody() *ModifySmartAccessGatewayClientUserResponseBody {
	return s.Body
}

func (s *ModifySmartAccessGatewayClientUserResponse) SetHeaders(v map[string]*string) *ModifySmartAccessGatewayClientUserResponse {
	s.Headers = v
	return s
}

func (s *ModifySmartAccessGatewayClientUserResponse) SetStatusCode(v int32) *ModifySmartAccessGatewayClientUserResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifySmartAccessGatewayClientUserResponse) SetBody(v *ModifySmartAccessGatewayClientUserResponseBody) *ModifySmartAccessGatewayClientUserResponse {
	s.Body = v
	return s
}

func (s *ModifySmartAccessGatewayClientUserResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
