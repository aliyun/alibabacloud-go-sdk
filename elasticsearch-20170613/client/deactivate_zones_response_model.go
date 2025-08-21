// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeactivateZonesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeactivateZonesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeactivateZonesResponse
	GetStatusCode() *int32
	SetBody(v *DeactivateZonesResponseBody) *DeactivateZonesResponse
	GetBody() *DeactivateZonesResponseBody
}

type DeactivateZonesResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeactivateZonesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeactivateZonesResponse) String() string {
	return dara.Prettify(s)
}

func (s DeactivateZonesResponse) GoString() string {
	return s.String()
}

func (s *DeactivateZonesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeactivateZonesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeactivateZonesResponse) GetBody() *DeactivateZonesResponseBody {
	return s.Body
}

func (s *DeactivateZonesResponse) SetHeaders(v map[string]*string) *DeactivateZonesResponse {
	s.Headers = v
	return s
}

func (s *DeactivateZonesResponse) SetStatusCode(v int32) *DeactivateZonesResponse {
	s.StatusCode = &v
	return s
}

func (s *DeactivateZonesResponse) SetBody(v *DeactivateZonesResponseBody) *DeactivateZonesResponse {
	s.Body = v
	return s
}

func (s *DeactivateZonesResponse) Validate() error {
	return dara.Validate(s)
}
