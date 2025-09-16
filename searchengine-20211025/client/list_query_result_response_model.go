// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListQueryResultResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListQueryResultResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListQueryResultResponse
	GetStatusCode() *int32
	SetBody(v *ListQueryResultResponseBody) *ListQueryResultResponse
	GetBody() *ListQueryResultResponseBody
}

type ListQueryResultResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListQueryResultResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListQueryResultResponse) String() string {
	return dara.Prettify(s)
}

func (s ListQueryResultResponse) GoString() string {
	return s.String()
}

func (s *ListQueryResultResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListQueryResultResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListQueryResultResponse) GetBody() *ListQueryResultResponseBody {
	return s.Body
}

func (s *ListQueryResultResponse) SetHeaders(v map[string]*string) *ListQueryResultResponse {
	s.Headers = v
	return s
}

func (s *ListQueryResultResponse) SetStatusCode(v int32) *ListQueryResultResponse {
	s.StatusCode = &v
	return s
}

func (s *ListQueryResultResponse) SetBody(v *ListQueryResultResponseBody) *ListQueryResultResponse {
	s.Body = v
	return s
}

func (s *ListQueryResultResponse) Validate() error {
	return dara.Validate(s)
}
