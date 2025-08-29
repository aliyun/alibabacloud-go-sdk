// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListParamsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListParamsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListParamsResponse
	GetStatusCode() *int32
	SetBody(v *ListParamsResponseBody) *ListParamsResponse
	GetBody() *ListParamsResponseBody
}

type ListParamsResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListParamsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListParamsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListParamsResponse) GoString() string {
	return s.String()
}

func (s *ListParamsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListParamsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListParamsResponse) GetBody() *ListParamsResponseBody {
	return s.Body
}

func (s *ListParamsResponse) SetHeaders(v map[string]*string) *ListParamsResponse {
	s.Headers = v
	return s
}

func (s *ListParamsResponse) SetStatusCode(v int32) *ListParamsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListParamsResponse) SetBody(v *ListParamsResponseBody) *ListParamsResponse {
	s.Body = v
	return s
}

func (s *ListParamsResponse) Validate() error {
	return dara.Validate(s)
}
