// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDatasetFileMetasResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteDatasetFileMetasResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteDatasetFileMetasResponse
	GetStatusCode() *int32
	SetBody(v *DeleteDatasetFileMetasResponseBody) *DeleteDatasetFileMetasResponse
	GetBody() *DeleteDatasetFileMetasResponseBody
}

type DeleteDatasetFileMetasResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteDatasetFileMetasResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteDatasetFileMetasResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteDatasetFileMetasResponse) GoString() string {
	return s.String()
}

func (s *DeleteDatasetFileMetasResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteDatasetFileMetasResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteDatasetFileMetasResponse) GetBody() *DeleteDatasetFileMetasResponseBody {
	return s.Body
}

func (s *DeleteDatasetFileMetasResponse) SetHeaders(v map[string]*string) *DeleteDatasetFileMetasResponse {
	s.Headers = v
	return s
}

func (s *DeleteDatasetFileMetasResponse) SetStatusCode(v int32) *DeleteDatasetFileMetasResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteDatasetFileMetasResponse) SetBody(v *DeleteDatasetFileMetasResponseBody) *DeleteDatasetFileMetasResponse {
	s.Body = v
	return s
}

func (s *DeleteDatasetFileMetasResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
