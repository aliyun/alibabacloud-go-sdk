// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListLayersResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListLayersResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListLayersResponse
	GetStatusCode() *int32
	SetBody(v *ListLayersResponseBody) *ListLayersResponse
	GetBody() *ListLayersResponseBody
}

type ListLayersResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListLayersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListLayersResponse) String() string {
	return dara.Prettify(s)
}

func (s ListLayersResponse) GoString() string {
	return s.String()
}

func (s *ListLayersResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListLayersResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListLayersResponse) GetBody() *ListLayersResponseBody {
	return s.Body
}

func (s *ListLayersResponse) SetHeaders(v map[string]*string) *ListLayersResponse {
	s.Headers = v
	return s
}

func (s *ListLayersResponse) SetStatusCode(v int32) *ListLayersResponse {
	s.StatusCode = &v
	return s
}

func (s *ListLayersResponse) SetBody(v *ListLayersResponseBody) *ListLayersResponse {
	s.Body = v
	return s
}

func (s *ListLayersResponse) Validate() error {
	return dara.Validate(s)
}
