// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteIntegrationsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteIntegrationsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteIntegrationsResponse
	GetStatusCode() *int32
	SetBody(v *DeleteIntegrationsResponseBody) *DeleteIntegrationsResponse
	GetBody() *DeleteIntegrationsResponseBody
}

type DeleteIntegrationsResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteIntegrationsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteIntegrationsResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteIntegrationsResponse) GoString() string {
	return s.String()
}

func (s *DeleteIntegrationsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteIntegrationsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteIntegrationsResponse) GetBody() *DeleteIntegrationsResponseBody {
	return s.Body
}

func (s *DeleteIntegrationsResponse) SetHeaders(v map[string]*string) *DeleteIntegrationsResponse {
	s.Headers = v
	return s
}

func (s *DeleteIntegrationsResponse) SetStatusCode(v int32) *DeleteIntegrationsResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteIntegrationsResponse) SetBody(v *DeleteIntegrationsResponseBody) *DeleteIntegrationsResponse {
	s.Body = v
	return s
}

func (s *DeleteIntegrationsResponse) Validate() error {
	return dara.Validate(s)
}
