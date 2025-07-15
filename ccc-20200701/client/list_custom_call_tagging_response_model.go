// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCustomCallTaggingResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListCustomCallTaggingResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListCustomCallTaggingResponse
	GetStatusCode() *int32
	SetBody(v *ListCustomCallTaggingResponseBody) *ListCustomCallTaggingResponse
	GetBody() *ListCustomCallTaggingResponseBody
}

type ListCustomCallTaggingResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListCustomCallTaggingResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListCustomCallTaggingResponse) String() string {
	return dara.Prettify(s)
}

func (s ListCustomCallTaggingResponse) GoString() string {
	return s.String()
}

func (s *ListCustomCallTaggingResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListCustomCallTaggingResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListCustomCallTaggingResponse) GetBody() *ListCustomCallTaggingResponseBody {
	return s.Body
}

func (s *ListCustomCallTaggingResponse) SetHeaders(v map[string]*string) *ListCustomCallTaggingResponse {
	s.Headers = v
	return s
}

func (s *ListCustomCallTaggingResponse) SetStatusCode(v int32) *ListCustomCallTaggingResponse {
	s.StatusCode = &v
	return s
}

func (s *ListCustomCallTaggingResponse) SetBody(v *ListCustomCallTaggingResponseBody) *ListCustomCallTaggingResponse {
	s.Body = v
	return s
}

func (s *ListCustomCallTaggingResponse) Validate() error {
	return dara.Validate(s)
}
