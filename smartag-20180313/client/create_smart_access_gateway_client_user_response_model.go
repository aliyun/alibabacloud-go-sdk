// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSmartAccessGatewayClientUserResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateSmartAccessGatewayClientUserResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateSmartAccessGatewayClientUserResponse
	GetStatusCode() *int32
	SetBody(v *CreateSmartAccessGatewayClientUserResponseBody) *CreateSmartAccessGatewayClientUserResponse
	GetBody() *CreateSmartAccessGatewayClientUserResponseBody
}

type CreateSmartAccessGatewayClientUserResponse struct {
	Headers    map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateSmartAccessGatewayClientUserResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateSmartAccessGatewayClientUserResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateSmartAccessGatewayClientUserResponse) GoString() string {
	return s.String()
}

func (s *CreateSmartAccessGatewayClientUserResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateSmartAccessGatewayClientUserResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateSmartAccessGatewayClientUserResponse) GetBody() *CreateSmartAccessGatewayClientUserResponseBody {
	return s.Body
}

func (s *CreateSmartAccessGatewayClientUserResponse) SetHeaders(v map[string]*string) *CreateSmartAccessGatewayClientUserResponse {
	s.Headers = v
	return s
}

func (s *CreateSmartAccessGatewayClientUserResponse) SetStatusCode(v int32) *CreateSmartAccessGatewayClientUserResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateSmartAccessGatewayClientUserResponse) SetBody(v *CreateSmartAccessGatewayClientUserResponseBody) *CreateSmartAccessGatewayClientUserResponse {
	s.Body = v
	return s
}

func (s *CreateSmartAccessGatewayClientUserResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
