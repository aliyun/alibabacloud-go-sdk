// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDefaultCollectorConfigurationsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListDefaultCollectorConfigurationsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListDefaultCollectorConfigurationsResponse
	GetStatusCode() *int32
	SetBody(v *ListDefaultCollectorConfigurationsResponseBody) *ListDefaultCollectorConfigurationsResponse
	GetBody() *ListDefaultCollectorConfigurationsResponseBody
}

type ListDefaultCollectorConfigurationsResponse struct {
	Headers    map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListDefaultCollectorConfigurationsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListDefaultCollectorConfigurationsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListDefaultCollectorConfigurationsResponse) GoString() string {
	return s.String()
}

func (s *ListDefaultCollectorConfigurationsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListDefaultCollectorConfigurationsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListDefaultCollectorConfigurationsResponse) GetBody() *ListDefaultCollectorConfigurationsResponseBody {
	return s.Body
}

func (s *ListDefaultCollectorConfigurationsResponse) SetHeaders(v map[string]*string) *ListDefaultCollectorConfigurationsResponse {
	s.Headers = v
	return s
}

func (s *ListDefaultCollectorConfigurationsResponse) SetStatusCode(v int32) *ListDefaultCollectorConfigurationsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListDefaultCollectorConfigurationsResponse) SetBody(v *ListDefaultCollectorConfigurationsResponseBody) *ListDefaultCollectorConfigurationsResponse {
	s.Body = v
	return s
}

func (s *ListDefaultCollectorConfigurationsResponse) Validate() error {
	return dara.Validate(s)
}
