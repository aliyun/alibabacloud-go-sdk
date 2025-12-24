// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteGatewayLabelResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteGatewayLabelResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteGatewayLabelResponse
	GetStatusCode() *int32
	SetBody(v *DeleteGatewayLabelResponseBody) *DeleteGatewayLabelResponse
	GetBody() *DeleteGatewayLabelResponseBody
}

type DeleteGatewayLabelResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteGatewayLabelResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteGatewayLabelResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteGatewayLabelResponse) GoString() string {
	return s.String()
}

func (s *DeleteGatewayLabelResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteGatewayLabelResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteGatewayLabelResponse) GetBody() *DeleteGatewayLabelResponseBody {
	return s.Body
}

func (s *DeleteGatewayLabelResponse) SetHeaders(v map[string]*string) *DeleteGatewayLabelResponse {
	s.Headers = v
	return s
}

func (s *DeleteGatewayLabelResponse) SetStatusCode(v int32) *DeleteGatewayLabelResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteGatewayLabelResponse) SetBody(v *DeleteGatewayLabelResponseBody) *DeleteGatewayLabelResponse {
	s.Body = v
	return s
}

func (s *DeleteGatewayLabelResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
