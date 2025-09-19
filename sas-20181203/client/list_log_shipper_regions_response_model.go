// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListLogShipperRegionsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListLogShipperRegionsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListLogShipperRegionsResponse
	GetStatusCode() *int32
	SetBody(v *ListLogShipperRegionsResponseBody) *ListLogShipperRegionsResponse
	GetBody() *ListLogShipperRegionsResponseBody
}

type ListLogShipperRegionsResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListLogShipperRegionsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListLogShipperRegionsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListLogShipperRegionsResponse) GoString() string {
	return s.String()
}

func (s *ListLogShipperRegionsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListLogShipperRegionsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListLogShipperRegionsResponse) GetBody() *ListLogShipperRegionsResponseBody {
	return s.Body
}

func (s *ListLogShipperRegionsResponse) SetHeaders(v map[string]*string) *ListLogShipperRegionsResponse {
	s.Headers = v
	return s
}

func (s *ListLogShipperRegionsResponse) SetStatusCode(v int32) *ListLogShipperRegionsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListLogShipperRegionsResponse) SetBody(v *ListLogShipperRegionsResponseBody) *ListLogShipperRegionsResponse {
	s.Body = v
	return s
}

func (s *ListLogShipperRegionsResponse) Validate() error {
	return dara.Validate(s)
}
