// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTranscodeTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListTranscodeTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListTranscodeTaskResponse
	GetStatusCode() *int32
	SetBody(v *ListTranscodeTaskResponseBody) *ListTranscodeTaskResponse
	GetBody() *ListTranscodeTaskResponseBody
}

type ListTranscodeTaskResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListTranscodeTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListTranscodeTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s ListTranscodeTaskResponse) GoString() string {
	return s.String()
}

func (s *ListTranscodeTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListTranscodeTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListTranscodeTaskResponse) GetBody() *ListTranscodeTaskResponseBody {
	return s.Body
}

func (s *ListTranscodeTaskResponse) SetHeaders(v map[string]*string) *ListTranscodeTaskResponse {
	s.Headers = v
	return s
}

func (s *ListTranscodeTaskResponse) SetStatusCode(v int32) *ListTranscodeTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *ListTranscodeTaskResponse) SetBody(v *ListTranscodeTaskResponseBody) *ListTranscodeTaskResponse {
	s.Body = v
	return s
}

func (s *ListTranscodeTaskResponse) Validate() error {
	return dara.Validate(s)
}
