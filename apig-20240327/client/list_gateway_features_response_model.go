// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListGatewayFeaturesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListGatewayFeaturesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListGatewayFeaturesResponse
	GetStatusCode() *int32
	SetBody(v *ListGatewayFeaturesResponseBody) *ListGatewayFeaturesResponse
	GetBody() *ListGatewayFeaturesResponseBody
}

type ListGatewayFeaturesResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListGatewayFeaturesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListGatewayFeaturesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListGatewayFeaturesResponse) GoString() string {
	return s.String()
}

func (s *ListGatewayFeaturesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListGatewayFeaturesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListGatewayFeaturesResponse) GetBody() *ListGatewayFeaturesResponseBody {
	return s.Body
}

func (s *ListGatewayFeaturesResponse) SetHeaders(v map[string]*string) *ListGatewayFeaturesResponse {
	s.Headers = v
	return s
}

func (s *ListGatewayFeaturesResponse) SetStatusCode(v int32) *ListGatewayFeaturesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListGatewayFeaturesResponse) SetBody(v *ListGatewayFeaturesResponseBody) *ListGatewayFeaturesResponse {
	s.Body = v
	return s
}

func (s *ListGatewayFeaturesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
