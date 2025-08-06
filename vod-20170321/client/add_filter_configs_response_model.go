// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddFilterConfigsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddFilterConfigsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddFilterConfigsResponse
	GetStatusCode() *int32
	SetBody(v *AddFilterConfigsResponseBody) *AddFilterConfigsResponse
	GetBody() *AddFilterConfigsResponseBody
}

type AddFilterConfigsResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddFilterConfigsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddFilterConfigsResponse) String() string {
	return dara.Prettify(s)
}

func (s AddFilterConfigsResponse) GoString() string {
	return s.String()
}

func (s *AddFilterConfigsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddFilterConfigsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddFilterConfigsResponse) GetBody() *AddFilterConfigsResponseBody {
	return s.Body
}

func (s *AddFilterConfigsResponse) SetHeaders(v map[string]*string) *AddFilterConfigsResponse {
	s.Headers = v
	return s
}

func (s *AddFilterConfigsResponse) SetStatusCode(v int32) *AddFilterConfigsResponse {
	s.StatusCode = &v
	return s
}

func (s *AddFilterConfigsResponse) SetBody(v *AddFilterConfigsResponseBody) *AddFilterConfigsResponse {
	s.Body = v
	return s
}

func (s *AddFilterConfigsResponse) Validate() error {
	return dara.Validate(s)
}
