// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListVideoProcessingsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListVideoProcessingsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListVideoProcessingsResponse
	GetStatusCode() *int32
	SetBody(v *ListVideoProcessingsResponseBody) *ListVideoProcessingsResponse
	GetBody() *ListVideoProcessingsResponseBody
}

type ListVideoProcessingsResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListVideoProcessingsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListVideoProcessingsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListVideoProcessingsResponse) GoString() string {
	return s.String()
}

func (s *ListVideoProcessingsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListVideoProcessingsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListVideoProcessingsResponse) GetBody() *ListVideoProcessingsResponseBody {
	return s.Body
}

func (s *ListVideoProcessingsResponse) SetHeaders(v map[string]*string) *ListVideoProcessingsResponse {
	s.Headers = v
	return s
}

func (s *ListVideoProcessingsResponse) SetStatusCode(v int32) *ListVideoProcessingsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListVideoProcessingsResponse) SetBody(v *ListVideoProcessingsResponseBody) *ListVideoProcessingsResponse {
	s.Body = v
	return s
}

func (s *ListVideoProcessingsResponse) Validate() error {
	return dara.Validate(s)
}
