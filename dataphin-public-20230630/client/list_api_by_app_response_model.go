// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListApiByAppResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListApiByAppResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListApiByAppResponse
	GetStatusCode() *int32
	SetBody(v *ListApiByAppResponseBody) *ListApiByAppResponse
	GetBody() *ListApiByAppResponseBody
}

type ListApiByAppResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListApiByAppResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListApiByAppResponse) String() string {
	return dara.Prettify(s)
}

func (s ListApiByAppResponse) GoString() string {
	return s.String()
}

func (s *ListApiByAppResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListApiByAppResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListApiByAppResponse) GetBody() *ListApiByAppResponseBody {
	return s.Body
}

func (s *ListApiByAppResponse) SetHeaders(v map[string]*string) *ListApiByAppResponse {
	s.Headers = v
	return s
}

func (s *ListApiByAppResponse) SetStatusCode(v int32) *ListApiByAppResponse {
	s.StatusCode = &v
	return s
}

func (s *ListApiByAppResponse) SetBody(v *ListApiByAppResponseBody) *ListApiByAppResponse {
	s.Body = v
	return s
}

func (s *ListApiByAppResponse) Validate() error {
	return dara.Validate(s)
}
