// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCheckResultResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListCheckResultResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListCheckResultResponse
	GetStatusCode() *int32
	SetBody(v *ListCheckResultResponseBody) *ListCheckResultResponse
	GetBody() *ListCheckResultResponseBody
}

type ListCheckResultResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListCheckResultResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListCheckResultResponse) String() string {
	return dara.Prettify(s)
}

func (s ListCheckResultResponse) GoString() string {
	return s.String()
}

func (s *ListCheckResultResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListCheckResultResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListCheckResultResponse) GetBody() *ListCheckResultResponseBody {
	return s.Body
}

func (s *ListCheckResultResponse) SetHeaders(v map[string]*string) *ListCheckResultResponse {
	s.Headers = v
	return s
}

func (s *ListCheckResultResponse) SetStatusCode(v int32) *ListCheckResultResponse {
	s.StatusCode = &v
	return s
}

func (s *ListCheckResultResponse) SetBody(v *ListCheckResultResponseBody) *ListCheckResultResponse {
	s.Body = v
	return s
}

func (s *ListCheckResultResponse) Validate() error {
	return dara.Validate(s)
}
