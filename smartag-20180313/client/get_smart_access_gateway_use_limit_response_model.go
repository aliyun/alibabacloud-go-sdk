// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSmartAccessGatewayUseLimitResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetSmartAccessGatewayUseLimitResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetSmartAccessGatewayUseLimitResponse
	GetStatusCode() *int32
	SetBody(v *GetSmartAccessGatewayUseLimitResponseBody) *GetSmartAccessGatewayUseLimitResponse
	GetBody() *GetSmartAccessGatewayUseLimitResponseBody
}

type GetSmartAccessGatewayUseLimitResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetSmartAccessGatewayUseLimitResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetSmartAccessGatewayUseLimitResponse) String() string {
	return dara.Prettify(s)
}

func (s GetSmartAccessGatewayUseLimitResponse) GoString() string {
	return s.String()
}

func (s *GetSmartAccessGatewayUseLimitResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetSmartAccessGatewayUseLimitResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetSmartAccessGatewayUseLimitResponse) GetBody() *GetSmartAccessGatewayUseLimitResponseBody {
	return s.Body
}

func (s *GetSmartAccessGatewayUseLimitResponse) SetHeaders(v map[string]*string) *GetSmartAccessGatewayUseLimitResponse {
	s.Headers = v
	return s
}

func (s *GetSmartAccessGatewayUseLimitResponse) SetStatusCode(v int32) *GetSmartAccessGatewayUseLimitResponse {
	s.StatusCode = &v
	return s
}

func (s *GetSmartAccessGatewayUseLimitResponse) SetBody(v *GetSmartAccessGatewayUseLimitResponseBody) *GetSmartAccessGatewayUseLimitResponse {
	s.Body = v
	return s
}

func (s *GetSmartAccessGatewayUseLimitResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
