// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAlgorithmVersionsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListAlgorithmVersionsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListAlgorithmVersionsResponse
	GetStatusCode() *int32
	SetBody(v *ListAlgorithmVersionsResponseBody) *ListAlgorithmVersionsResponse
	GetBody() *ListAlgorithmVersionsResponseBody
}

type ListAlgorithmVersionsResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAlgorithmVersionsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAlgorithmVersionsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListAlgorithmVersionsResponse) GoString() string {
	return s.String()
}

func (s *ListAlgorithmVersionsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListAlgorithmVersionsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListAlgorithmVersionsResponse) GetBody() *ListAlgorithmVersionsResponseBody {
	return s.Body
}

func (s *ListAlgorithmVersionsResponse) SetHeaders(v map[string]*string) *ListAlgorithmVersionsResponse {
	s.Headers = v
	return s
}

func (s *ListAlgorithmVersionsResponse) SetStatusCode(v int32) *ListAlgorithmVersionsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAlgorithmVersionsResponse) SetBody(v *ListAlgorithmVersionsResponseBody) *ListAlgorithmVersionsResponse {
	s.Body = v
	return s
}

func (s *ListAlgorithmVersionsResponse) Validate() error {
	return dara.Validate(s)
}
