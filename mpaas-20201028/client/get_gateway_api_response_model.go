// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetGatewayApiResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetGatewayApiResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetGatewayApiResponse
	GetStatusCode() *int32
	SetBody(v *GetGatewayApiResponseBody) *GetGatewayApiResponse
	GetBody() *GetGatewayApiResponseBody
}

type GetGatewayApiResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetGatewayApiResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetGatewayApiResponse) String() string {
	return dara.Prettify(s)
}

func (s GetGatewayApiResponse) GoString() string {
	return s.String()
}

func (s *GetGatewayApiResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetGatewayApiResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetGatewayApiResponse) GetBody() *GetGatewayApiResponseBody {
	return s.Body
}

func (s *GetGatewayApiResponse) SetHeaders(v map[string]*string) *GetGatewayApiResponse {
	s.Headers = v
	return s
}

func (s *GetGatewayApiResponse) SetStatusCode(v int32) *GetGatewayApiResponse {
	s.StatusCode = &v
	return s
}

func (s *GetGatewayApiResponse) SetBody(v *GetGatewayApiResponseBody) *GetGatewayApiResponse {
	s.Body = v
	return s
}

func (s *GetGatewayApiResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
