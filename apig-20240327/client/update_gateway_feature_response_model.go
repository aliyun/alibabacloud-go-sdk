// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateGatewayFeatureResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateGatewayFeatureResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateGatewayFeatureResponse
	GetStatusCode() *int32
	SetBody(v *UpdateGatewayFeatureResponseBody) *UpdateGatewayFeatureResponse
	GetBody() *UpdateGatewayFeatureResponseBody
}

type UpdateGatewayFeatureResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateGatewayFeatureResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateGatewayFeatureResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateGatewayFeatureResponse) GoString() string {
	return s.String()
}

func (s *UpdateGatewayFeatureResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateGatewayFeatureResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateGatewayFeatureResponse) GetBody() *UpdateGatewayFeatureResponseBody {
	return s.Body
}

func (s *UpdateGatewayFeatureResponse) SetHeaders(v map[string]*string) *UpdateGatewayFeatureResponse {
	s.Headers = v
	return s
}

func (s *UpdateGatewayFeatureResponse) SetStatusCode(v int32) *UpdateGatewayFeatureResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateGatewayFeatureResponse) SetBody(v *UpdateGatewayFeatureResponseBody) *UpdateGatewayFeatureResponse {
	s.Body = v
	return s
}

func (s *UpdateGatewayFeatureResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
