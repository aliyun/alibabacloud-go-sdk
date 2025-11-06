// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryGatewayRegionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryGatewayRegionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryGatewayRegionResponse
	GetStatusCode() *int32
	SetBody(v *QueryGatewayRegionResponseBody) *QueryGatewayRegionResponse
	GetBody() *QueryGatewayRegionResponseBody
}

type QueryGatewayRegionResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryGatewayRegionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryGatewayRegionResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryGatewayRegionResponse) GoString() string {
	return s.String()
}

func (s *QueryGatewayRegionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryGatewayRegionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryGatewayRegionResponse) GetBody() *QueryGatewayRegionResponseBody {
	return s.Body
}

func (s *QueryGatewayRegionResponse) SetHeaders(v map[string]*string) *QueryGatewayRegionResponse {
	s.Headers = v
	return s
}

func (s *QueryGatewayRegionResponse) SetStatusCode(v int32) *QueryGatewayRegionResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryGatewayRegionResponse) SetBody(v *QueryGatewayRegionResponseBody) *QueryGatewayRegionResponse {
	s.Body = v
	return s
}

func (s *QueryGatewayRegionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
