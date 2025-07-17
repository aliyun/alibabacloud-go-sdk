// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteEventBridgeIntegrationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteEventBridgeIntegrationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteEventBridgeIntegrationResponse
	GetStatusCode() *int32
	SetBody(v *DeleteEventBridgeIntegrationResponseBody) *DeleteEventBridgeIntegrationResponse
	GetBody() *DeleteEventBridgeIntegrationResponseBody
}

type DeleteEventBridgeIntegrationResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteEventBridgeIntegrationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteEventBridgeIntegrationResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteEventBridgeIntegrationResponse) GoString() string {
	return s.String()
}

func (s *DeleteEventBridgeIntegrationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteEventBridgeIntegrationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteEventBridgeIntegrationResponse) GetBody() *DeleteEventBridgeIntegrationResponseBody {
	return s.Body
}

func (s *DeleteEventBridgeIntegrationResponse) SetHeaders(v map[string]*string) *DeleteEventBridgeIntegrationResponse {
	s.Headers = v
	return s
}

func (s *DeleteEventBridgeIntegrationResponse) SetStatusCode(v int32) *DeleteEventBridgeIntegrationResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteEventBridgeIntegrationResponse) SetBody(v *DeleteEventBridgeIntegrationResponseBody) *DeleteEventBridgeIntegrationResponse {
	s.Body = v
	return s
}

func (s *DeleteEventBridgeIntegrationResponse) Validate() error {
	return dara.Validate(s)
}
