// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteGatewayAuthConsumerResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteGatewayAuthConsumerResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteGatewayAuthConsumerResponse
	GetStatusCode() *int32
	SetBody(v *DeleteGatewayAuthConsumerResponseBody) *DeleteGatewayAuthConsumerResponse
	GetBody() *DeleteGatewayAuthConsumerResponseBody
}

type DeleteGatewayAuthConsumerResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteGatewayAuthConsumerResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteGatewayAuthConsumerResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteGatewayAuthConsumerResponse) GoString() string {
	return s.String()
}

func (s *DeleteGatewayAuthConsumerResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteGatewayAuthConsumerResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteGatewayAuthConsumerResponse) GetBody() *DeleteGatewayAuthConsumerResponseBody {
	return s.Body
}

func (s *DeleteGatewayAuthConsumerResponse) SetHeaders(v map[string]*string) *DeleteGatewayAuthConsumerResponse {
	s.Headers = v
	return s
}

func (s *DeleteGatewayAuthConsumerResponse) SetStatusCode(v int32) *DeleteGatewayAuthConsumerResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteGatewayAuthConsumerResponse) SetBody(v *DeleteGatewayAuthConsumerResponseBody) *DeleteGatewayAuthConsumerResponse {
	s.Body = v
	return s
}

func (s *DeleteGatewayAuthConsumerResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
