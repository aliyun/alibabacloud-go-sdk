// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDatasetFileMetasResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateDatasetFileMetasResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateDatasetFileMetasResponse
	GetStatusCode() *int32
	SetBody(v *UpdateDatasetFileMetasResponseBody) *UpdateDatasetFileMetasResponse
	GetBody() *UpdateDatasetFileMetasResponseBody
}

type UpdateDatasetFileMetasResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateDatasetFileMetasResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateDatasetFileMetasResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateDatasetFileMetasResponse) GoString() string {
	return s.String()
}

func (s *UpdateDatasetFileMetasResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateDatasetFileMetasResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateDatasetFileMetasResponse) GetBody() *UpdateDatasetFileMetasResponseBody {
	return s.Body
}

func (s *UpdateDatasetFileMetasResponse) SetHeaders(v map[string]*string) *UpdateDatasetFileMetasResponse {
	s.Headers = v
	return s
}

func (s *UpdateDatasetFileMetasResponse) SetStatusCode(v int32) *UpdateDatasetFileMetasResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateDatasetFileMetasResponse) SetBody(v *UpdateDatasetFileMetasResponseBody) *UpdateDatasetFileMetasResponse {
	s.Body = v
	return s
}

func (s *UpdateDatasetFileMetasResponse) Validate() error {
	return dara.Validate(s)
}
