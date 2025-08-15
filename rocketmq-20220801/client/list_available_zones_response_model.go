// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAvailableZonesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListAvailableZonesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListAvailableZonesResponse
	GetStatusCode() *int32
	SetBody(v *ListAvailableZonesResponseBody) *ListAvailableZonesResponse
	GetBody() *ListAvailableZonesResponseBody
}

type ListAvailableZonesResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAvailableZonesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAvailableZonesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListAvailableZonesResponse) GoString() string {
	return s.String()
}

func (s *ListAvailableZonesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListAvailableZonesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListAvailableZonesResponse) GetBody() *ListAvailableZonesResponseBody {
	return s.Body
}

func (s *ListAvailableZonesResponse) SetHeaders(v map[string]*string) *ListAvailableZonesResponse {
	s.Headers = v
	return s
}

func (s *ListAvailableZonesResponse) SetStatusCode(v int32) *ListAvailableZonesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAvailableZonesResponse) SetBody(v *ListAvailableZonesResponseBody) *ListAvailableZonesResponse {
	s.Body = v
	return s
}

func (s *ListAvailableZonesResponse) Validate() error {
	return dara.Validate(s)
}
