// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListLineageResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListLineageResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListLineageResponse
	GetStatusCode() *int32
	SetBody(v *ListLineageResponseBody) *ListLineageResponse
	GetBody() *ListLineageResponseBody
}

type ListLineageResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListLineageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListLineageResponse) String() string {
	return dara.Prettify(s)
}

func (s ListLineageResponse) GoString() string {
	return s.String()
}

func (s *ListLineageResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListLineageResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListLineageResponse) GetBody() *ListLineageResponseBody {
	return s.Body
}

func (s *ListLineageResponse) SetHeaders(v map[string]*string) *ListLineageResponse {
	s.Headers = v
	return s
}

func (s *ListLineageResponse) SetStatusCode(v int32) *ListLineageResponse {
	s.StatusCode = &v
	return s
}

func (s *ListLineageResponse) SetBody(v *ListLineageResponseBody) *ListLineageResponse {
	s.Body = v
	return s
}

func (s *ListLineageResponse) Validate() error {
	return dara.Validate(s)
}
