// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFileUploadCallbackResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *FileUploadCallbackResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *FileUploadCallbackResponse
	GetStatusCode() *int32
	SetBody(v *FileUploadCallbackResponseBody) *FileUploadCallbackResponse
	GetBody() *FileUploadCallbackResponseBody
}

type FileUploadCallbackResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *FileUploadCallbackResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s FileUploadCallbackResponse) String() string {
	return dara.Prettify(s)
}

func (s FileUploadCallbackResponse) GoString() string {
	return s.String()
}

func (s *FileUploadCallbackResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *FileUploadCallbackResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *FileUploadCallbackResponse) GetBody() *FileUploadCallbackResponseBody {
	return s.Body
}

func (s *FileUploadCallbackResponse) SetHeaders(v map[string]*string) *FileUploadCallbackResponse {
	s.Headers = v
	return s
}

func (s *FileUploadCallbackResponse) SetStatusCode(v int32) *FileUploadCallbackResponse {
	s.StatusCode = &v
	return s
}

func (s *FileUploadCallbackResponse) SetBody(v *FileUploadCallbackResponseBody) *FileUploadCallbackResponse {
	s.Body = v
	return s
}

func (s *FileUploadCallbackResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
