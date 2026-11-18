// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateOssUploadResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateOssUploadResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateOssUploadResponse
	GetStatusCode() *int32
	SetBody(v *CreateOssUploadResponseBody) *CreateOssUploadResponse
	GetBody() *CreateOssUploadResponseBody
}

type CreateOssUploadResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateOssUploadResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateOssUploadResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateOssUploadResponse) GoString() string {
	return s.String()
}

func (s *CreateOssUploadResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateOssUploadResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateOssUploadResponse) GetBody() *CreateOssUploadResponseBody {
	return s.Body
}

func (s *CreateOssUploadResponse) SetHeaders(v map[string]*string) *CreateOssUploadResponse {
	s.Headers = v
	return s
}

func (s *CreateOssUploadResponse) SetStatusCode(v int32) *CreateOssUploadResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateOssUploadResponse) SetBody(v *CreateOssUploadResponseBody) *CreateOssUploadResponse {
	s.Body = v
	return s
}

func (s *CreateOssUploadResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
