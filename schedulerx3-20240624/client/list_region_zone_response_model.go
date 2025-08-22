// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListRegionZoneResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListRegionZoneResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListRegionZoneResponse
	GetStatusCode() *int32
	SetBody(v *ListRegionZoneResponseBody) *ListRegionZoneResponse
	GetBody() *ListRegionZoneResponseBody
}

type ListRegionZoneResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListRegionZoneResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListRegionZoneResponse) String() string {
	return dara.Prettify(s)
}

func (s ListRegionZoneResponse) GoString() string {
	return s.String()
}

func (s *ListRegionZoneResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListRegionZoneResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListRegionZoneResponse) GetBody() *ListRegionZoneResponseBody {
	return s.Body
}

func (s *ListRegionZoneResponse) SetHeaders(v map[string]*string) *ListRegionZoneResponse {
	s.Headers = v
	return s
}

func (s *ListRegionZoneResponse) SetStatusCode(v int32) *ListRegionZoneResponse {
	s.StatusCode = &v
	return s
}

func (s *ListRegionZoneResponse) SetBody(v *ListRegionZoneResponseBody) *ListRegionZoneResponse {
	s.Body = v
	return s
}

func (s *ListRegionZoneResponse) Validate() error {
	return dara.Validate(s)
}
