// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFileUploadResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *FileUploadResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *FileUploadResponse
	GetStatusCode() *int32
	SetBody(v *FileUploadResponseBody) *FileUploadResponse
	GetBody() *FileUploadResponseBody
}

type FileUploadResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *FileUploadResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s FileUploadResponse) String() string {
	return dara.Prettify(s)
}

func (s FileUploadResponse) GoString() string {
	return s.String()
}

func (s *FileUploadResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *FileUploadResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *FileUploadResponse) GetBody() *FileUploadResponseBody {
	return s.Body
}

func (s *FileUploadResponse) SetHeaders(v map[string]*string) *FileUploadResponse {
	s.Headers = v
	return s
}

func (s *FileUploadResponse) SetStatusCode(v int32) *FileUploadResponse {
	s.StatusCode = &v
	return s
}

func (s *FileUploadResponse) SetBody(v *FileUploadResponseBody) *FileUploadResponse {
	s.Body = v
	return s
}

func (s *FileUploadResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
