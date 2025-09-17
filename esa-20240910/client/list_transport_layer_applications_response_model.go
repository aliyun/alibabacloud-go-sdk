// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTransportLayerApplicationsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListTransportLayerApplicationsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListTransportLayerApplicationsResponse
	GetStatusCode() *int32
	SetBody(v *ListTransportLayerApplicationsResponseBody) *ListTransportLayerApplicationsResponse
	GetBody() *ListTransportLayerApplicationsResponseBody
}

type ListTransportLayerApplicationsResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListTransportLayerApplicationsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListTransportLayerApplicationsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListTransportLayerApplicationsResponse) GoString() string {
	return s.String()
}

func (s *ListTransportLayerApplicationsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListTransportLayerApplicationsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListTransportLayerApplicationsResponse) GetBody() *ListTransportLayerApplicationsResponseBody {
	return s.Body
}

func (s *ListTransportLayerApplicationsResponse) SetHeaders(v map[string]*string) *ListTransportLayerApplicationsResponse {
	s.Headers = v
	return s
}

func (s *ListTransportLayerApplicationsResponse) SetStatusCode(v int32) *ListTransportLayerApplicationsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListTransportLayerApplicationsResponse) SetBody(v *ListTransportLayerApplicationsResponseBody) *ListTransportLayerApplicationsResponse {
	s.Body = v
	return s
}

func (s *ListTransportLayerApplicationsResponse) Validate() error {
	return dara.Validate(s)
}
