// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateGatewayLabelResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateGatewayLabelResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateGatewayLabelResponse
	GetStatusCode() *int32
	SetBody(v *UpdateGatewayLabelResponseBody) *UpdateGatewayLabelResponse
	GetBody() *UpdateGatewayLabelResponseBody
}

type UpdateGatewayLabelResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateGatewayLabelResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateGatewayLabelResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateGatewayLabelResponse) GoString() string {
	return s.String()
}

func (s *UpdateGatewayLabelResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateGatewayLabelResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateGatewayLabelResponse) GetBody() *UpdateGatewayLabelResponseBody {
	return s.Body
}

func (s *UpdateGatewayLabelResponse) SetHeaders(v map[string]*string) *UpdateGatewayLabelResponse {
	s.Headers = v
	return s
}

func (s *UpdateGatewayLabelResponse) SetStatusCode(v int32) *UpdateGatewayLabelResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateGatewayLabelResponse) SetBody(v *UpdateGatewayLabelResponseBody) *UpdateGatewayLabelResponse {
	s.Body = v
	return s
}

func (s *UpdateGatewayLabelResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
