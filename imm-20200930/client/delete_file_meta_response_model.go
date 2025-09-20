// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteFileMetaResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteFileMetaResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteFileMetaResponse
	GetStatusCode() *int32
	SetBody(v *DeleteFileMetaResponseBody) *DeleteFileMetaResponse
	GetBody() *DeleteFileMetaResponseBody
}

type DeleteFileMetaResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteFileMetaResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteFileMetaResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteFileMetaResponse) GoString() string {
	return s.String()
}

func (s *DeleteFileMetaResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteFileMetaResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteFileMetaResponse) GetBody() *DeleteFileMetaResponseBody {
	return s.Body
}

func (s *DeleteFileMetaResponse) SetHeaders(v map[string]*string) *DeleteFileMetaResponse {
	s.Headers = v
	return s
}

func (s *DeleteFileMetaResponse) SetStatusCode(v int32) *DeleteFileMetaResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteFileMetaResponse) SetBody(v *DeleteFileMetaResponseBody) *DeleteFileMetaResponse {
	s.Body = v
	return s
}

func (s *DeleteFileMetaResponse) Validate() error {
	return dara.Validate(s)
}
