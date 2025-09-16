// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListVectorQueryResultResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListVectorQueryResultResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListVectorQueryResultResponse
	GetStatusCode() *int32
	SetBody(v *ListVectorQueryResultResponseBody) *ListVectorQueryResultResponse
	GetBody() *ListVectorQueryResultResponseBody
}

type ListVectorQueryResultResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListVectorQueryResultResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListVectorQueryResultResponse) String() string {
	return dara.Prettify(s)
}

func (s ListVectorQueryResultResponse) GoString() string {
	return s.String()
}

func (s *ListVectorQueryResultResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListVectorQueryResultResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListVectorQueryResultResponse) GetBody() *ListVectorQueryResultResponseBody {
	return s.Body
}

func (s *ListVectorQueryResultResponse) SetHeaders(v map[string]*string) *ListVectorQueryResultResponse {
	s.Headers = v
	return s
}

func (s *ListVectorQueryResultResponse) SetStatusCode(v int32) *ListVectorQueryResultResponse {
	s.StatusCode = &v
	return s
}

func (s *ListVectorQueryResultResponse) SetBody(v *ListVectorQueryResultResponseBody) *ListVectorQueryResultResponse {
	s.Body = v
	return s
}

func (s *ListVectorQueryResultResponse) Validate() error {
	return dara.Validate(s)
}
