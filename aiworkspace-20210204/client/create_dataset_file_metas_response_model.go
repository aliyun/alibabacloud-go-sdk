// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDatasetFileMetasResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateDatasetFileMetasResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateDatasetFileMetasResponse
	GetStatusCode() *int32
	SetBody(v *CreateDatasetFileMetasResponseBody) *CreateDatasetFileMetasResponse
	GetBody() *CreateDatasetFileMetasResponseBody
}

type CreateDatasetFileMetasResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateDatasetFileMetasResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateDatasetFileMetasResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateDatasetFileMetasResponse) GoString() string {
	return s.String()
}

func (s *CreateDatasetFileMetasResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateDatasetFileMetasResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateDatasetFileMetasResponse) GetBody() *CreateDatasetFileMetasResponseBody {
	return s.Body
}

func (s *CreateDatasetFileMetasResponse) SetHeaders(v map[string]*string) *CreateDatasetFileMetasResponse {
	s.Headers = v
	return s
}

func (s *CreateDatasetFileMetasResponse) SetStatusCode(v int32) *CreateDatasetFileMetasResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateDatasetFileMetasResponse) SetBody(v *CreateDatasetFileMetasResponseBody) *CreateDatasetFileMetasResponse {
	s.Body = v
	return s
}

func (s *CreateDatasetFileMetasResponse) Validate() error {
	return dara.Validate(s)
}
