// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateOrUpdateEventBridgeIntegrationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateOrUpdateEventBridgeIntegrationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateOrUpdateEventBridgeIntegrationResponse
	GetStatusCode() *int32
	SetBody(v *CreateOrUpdateEventBridgeIntegrationResponseBody) *CreateOrUpdateEventBridgeIntegrationResponse
	GetBody() *CreateOrUpdateEventBridgeIntegrationResponseBody
}

type CreateOrUpdateEventBridgeIntegrationResponse struct {
	Headers    map[string]*string                                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateOrUpdateEventBridgeIntegrationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateOrUpdateEventBridgeIntegrationResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateOrUpdateEventBridgeIntegrationResponse) GoString() string {
	return s.String()
}

func (s *CreateOrUpdateEventBridgeIntegrationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateOrUpdateEventBridgeIntegrationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateOrUpdateEventBridgeIntegrationResponse) GetBody() *CreateOrUpdateEventBridgeIntegrationResponseBody {
	return s.Body
}

func (s *CreateOrUpdateEventBridgeIntegrationResponse) SetHeaders(v map[string]*string) *CreateOrUpdateEventBridgeIntegrationResponse {
	s.Headers = v
	return s
}

func (s *CreateOrUpdateEventBridgeIntegrationResponse) SetStatusCode(v int32) *CreateOrUpdateEventBridgeIntegrationResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateOrUpdateEventBridgeIntegrationResponse) SetBody(v *CreateOrUpdateEventBridgeIntegrationResponseBody) *CreateOrUpdateEventBridgeIntegrationResponse {
	s.Body = v
	return s
}

func (s *CreateOrUpdateEventBridgeIntegrationResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
