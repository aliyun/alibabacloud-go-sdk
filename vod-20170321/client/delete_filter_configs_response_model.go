// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteFilterConfigsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteFilterConfigsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteFilterConfigsResponse
	GetStatusCode() *int32
	SetBody(v *DeleteFilterConfigsResponseBody) *DeleteFilterConfigsResponse
	GetBody() *DeleteFilterConfigsResponseBody
}

type DeleteFilterConfigsResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteFilterConfigsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteFilterConfigsResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteFilterConfigsResponse) GoString() string {
	return s.String()
}

func (s *DeleteFilterConfigsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteFilterConfigsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteFilterConfigsResponse) GetBody() *DeleteFilterConfigsResponseBody {
	return s.Body
}

func (s *DeleteFilterConfigsResponse) SetHeaders(v map[string]*string) *DeleteFilterConfigsResponse {
	s.Headers = v
	return s
}

func (s *DeleteFilterConfigsResponse) SetStatusCode(v int32) *DeleteFilterConfigsResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteFilterConfigsResponse) SetBody(v *DeleteFilterConfigsResponseBody) *DeleteFilterConfigsResponse {
	s.Body = v
	return s
}

func (s *DeleteFilterConfigsResponse) Validate() error {
	return dara.Validate(s)
}
