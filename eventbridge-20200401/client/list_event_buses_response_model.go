// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListEventBusesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListEventBusesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListEventBusesResponse
	GetStatusCode() *int32
	SetBody(v *ListEventBusesResponseBody) *ListEventBusesResponse
	GetBody() *ListEventBusesResponseBody
}

type ListEventBusesResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListEventBusesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListEventBusesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListEventBusesResponse) GoString() string {
	return s.String()
}

func (s *ListEventBusesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListEventBusesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListEventBusesResponse) GetBody() *ListEventBusesResponseBody {
	return s.Body
}

func (s *ListEventBusesResponse) SetHeaders(v map[string]*string) *ListEventBusesResponse {
	s.Headers = v
	return s
}

func (s *ListEventBusesResponse) SetStatusCode(v int32) *ListEventBusesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListEventBusesResponse) SetBody(v *ListEventBusesResponseBody) *ListEventBusesResponse {
	s.Body = v
	return s
}

func (s *ListEventBusesResponse) Validate() error {
	return dara.Validate(s)
}
