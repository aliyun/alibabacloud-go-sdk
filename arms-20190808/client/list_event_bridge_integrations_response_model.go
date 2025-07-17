// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListEventBridgeIntegrationsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListEventBridgeIntegrationsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListEventBridgeIntegrationsResponse
	GetStatusCode() *int32
	SetBody(v *ListEventBridgeIntegrationsResponseBody) *ListEventBridgeIntegrationsResponse
	GetBody() *ListEventBridgeIntegrationsResponseBody
}

type ListEventBridgeIntegrationsResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListEventBridgeIntegrationsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListEventBridgeIntegrationsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListEventBridgeIntegrationsResponse) GoString() string {
	return s.String()
}

func (s *ListEventBridgeIntegrationsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListEventBridgeIntegrationsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListEventBridgeIntegrationsResponse) GetBody() *ListEventBridgeIntegrationsResponseBody {
	return s.Body
}

func (s *ListEventBridgeIntegrationsResponse) SetHeaders(v map[string]*string) *ListEventBridgeIntegrationsResponse {
	s.Headers = v
	return s
}

func (s *ListEventBridgeIntegrationsResponse) SetStatusCode(v int32) *ListEventBridgeIntegrationsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListEventBridgeIntegrationsResponse) SetBody(v *ListEventBridgeIntegrationsResponseBody) *ListEventBridgeIntegrationsResponse {
	s.Body = v
	return s
}

func (s *ListEventBridgeIntegrationsResponse) Validate() error {
	return dara.Validate(s)
}
