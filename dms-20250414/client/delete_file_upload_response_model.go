// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteFileUploadResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteFileUploadResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteFileUploadResponse
	GetStatusCode() *int32
	SetBody(v *DeleteFileUploadResponseBody) *DeleteFileUploadResponse
	GetBody() *DeleteFileUploadResponseBody
}

type DeleteFileUploadResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteFileUploadResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteFileUploadResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteFileUploadResponse) GoString() string {
	return s.String()
}

func (s *DeleteFileUploadResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteFileUploadResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteFileUploadResponse) GetBody() *DeleteFileUploadResponseBody {
	return s.Body
}

func (s *DeleteFileUploadResponse) SetHeaders(v map[string]*string) *DeleteFileUploadResponse {
	s.Headers = v
	return s
}

func (s *DeleteFileUploadResponse) SetStatusCode(v int32) *DeleteFileUploadResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteFileUploadResponse) SetBody(v *DeleteFileUploadResponseBody) *DeleteFileUploadResponse {
	s.Body = v
	return s
}

func (s *DeleteFileUploadResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
