// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDatasetFileMetasResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListDatasetFileMetasResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListDatasetFileMetasResponse
	GetStatusCode() *int32
	SetBody(v *ListDatasetFileMetasResponseBody) *ListDatasetFileMetasResponse
	GetBody() *ListDatasetFileMetasResponseBody
}

type ListDatasetFileMetasResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListDatasetFileMetasResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListDatasetFileMetasResponse) String() string {
	return dara.Prettify(s)
}

func (s ListDatasetFileMetasResponse) GoString() string {
	return s.String()
}

func (s *ListDatasetFileMetasResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListDatasetFileMetasResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListDatasetFileMetasResponse) GetBody() *ListDatasetFileMetasResponseBody {
	return s.Body
}

func (s *ListDatasetFileMetasResponse) SetHeaders(v map[string]*string) *ListDatasetFileMetasResponse {
	s.Headers = v
	return s
}

func (s *ListDatasetFileMetasResponse) SetStatusCode(v int32) *ListDatasetFileMetasResponse {
	s.StatusCode = &v
	return s
}

func (s *ListDatasetFileMetasResponse) SetBody(v *ListDatasetFileMetasResponseBody) *ListDatasetFileMetasResponse {
	s.Body = v
	return s
}

func (s *ListDatasetFileMetasResponse) Validate() error {
	return dara.Validate(s)
}
