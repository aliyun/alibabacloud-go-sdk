// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iImportGatewayApisResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ImportGatewayApisResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ImportGatewayApisResponse
	GetStatusCode() *int32
	SetBody(v *ImportGatewayApisResponseBody) *ImportGatewayApisResponse
	GetBody() *ImportGatewayApisResponseBody
}

type ImportGatewayApisResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ImportGatewayApisResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ImportGatewayApisResponse) String() string {
	return dara.Prettify(s)
}

func (s ImportGatewayApisResponse) GoString() string {
	return s.String()
}

func (s *ImportGatewayApisResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ImportGatewayApisResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ImportGatewayApisResponse) GetBody() *ImportGatewayApisResponseBody {
	return s.Body
}

func (s *ImportGatewayApisResponse) SetHeaders(v map[string]*string) *ImportGatewayApisResponse {
	s.Headers = v
	return s
}

func (s *ImportGatewayApisResponse) SetStatusCode(v int32) *ImportGatewayApisResponse {
	s.StatusCode = &v
	return s
}

func (s *ImportGatewayApisResponse) SetBody(v *ImportGatewayApisResponseBody) *ImportGatewayApisResponse {
	s.Body = v
	return s
}

func (s *ImportGatewayApisResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
