// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModelRouterQueryModelGroupClientsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModelRouterQueryModelGroupClientsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModelRouterQueryModelGroupClientsResponse
	GetStatusCode() *int32
	SetBody(v *ModelRouterQueryModelGroupClientsResponseBody) *ModelRouterQueryModelGroupClientsResponse
	GetBody() *ModelRouterQueryModelGroupClientsResponseBody
}

type ModelRouterQueryModelGroupClientsResponse struct {
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModelRouterQueryModelGroupClientsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModelRouterQueryModelGroupClientsResponse) String() string {
	return dara.Prettify(s)
}

func (s ModelRouterQueryModelGroupClientsResponse) GoString() string {
	return s.String()
}

func (s *ModelRouterQueryModelGroupClientsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModelRouterQueryModelGroupClientsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModelRouterQueryModelGroupClientsResponse) GetBody() *ModelRouterQueryModelGroupClientsResponseBody {
	return s.Body
}

func (s *ModelRouterQueryModelGroupClientsResponse) SetHeaders(v map[string]*string) *ModelRouterQueryModelGroupClientsResponse {
	s.Headers = v
	return s
}

func (s *ModelRouterQueryModelGroupClientsResponse) SetStatusCode(v int32) *ModelRouterQueryModelGroupClientsResponse {
	s.StatusCode = &v
	return s
}

func (s *ModelRouterQueryModelGroupClientsResponse) SetBody(v *ModelRouterQueryModelGroupClientsResponseBody) *ModelRouterQueryModelGroupClientsResponse {
	s.Body = v
	return s
}

func (s *ModelRouterQueryModelGroupClientsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
