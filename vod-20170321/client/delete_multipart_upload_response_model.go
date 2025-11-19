// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteMultipartUploadResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteMultipartUploadResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteMultipartUploadResponse
	GetStatusCode() *int32
	SetBody(v *DeleteMultipartUploadResponseBody) *DeleteMultipartUploadResponse
	GetBody() *DeleteMultipartUploadResponseBody
}

type DeleteMultipartUploadResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteMultipartUploadResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteMultipartUploadResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteMultipartUploadResponse) GoString() string {
	return s.String()
}

func (s *DeleteMultipartUploadResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteMultipartUploadResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteMultipartUploadResponse) GetBody() *DeleteMultipartUploadResponseBody {
	return s.Body
}

func (s *DeleteMultipartUploadResponse) SetHeaders(v map[string]*string) *DeleteMultipartUploadResponse {
	s.Headers = v
	return s
}

func (s *DeleteMultipartUploadResponse) SetStatusCode(v int32) *DeleteMultipartUploadResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteMultipartUploadResponse) SetBody(v *DeleteMultipartUploadResponseBody) *DeleteMultipartUploadResponse {
	s.Body = v
	return s
}

func (s *DeleteMultipartUploadResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
