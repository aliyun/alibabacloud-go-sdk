// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListRestQueryResultResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListRestQueryResultResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListRestQueryResultResponse
	GetStatusCode() *int32
	SetBody(v *ListRestQueryResultResponseBody) *ListRestQueryResultResponse
	GetBody() *ListRestQueryResultResponseBody
}

type ListRestQueryResultResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListRestQueryResultResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListRestQueryResultResponse) String() string {
	return dara.Prettify(s)
}

func (s ListRestQueryResultResponse) GoString() string {
	return s.String()
}

func (s *ListRestQueryResultResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListRestQueryResultResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListRestQueryResultResponse) GetBody() *ListRestQueryResultResponseBody {
	return s.Body
}

func (s *ListRestQueryResultResponse) SetHeaders(v map[string]*string) *ListRestQueryResultResponse {
	s.Headers = v
	return s
}

func (s *ListRestQueryResultResponse) SetStatusCode(v int32) *ListRestQueryResultResponse {
	s.StatusCode = &v
	return s
}

func (s *ListRestQueryResultResponse) SetBody(v *ListRestQueryResultResponseBody) *ListRestQueryResultResponse {
	s.Body = v
	return s
}

func (s *ListRestQueryResultResponse) Validate() error {
	return dara.Validate(s)
}
