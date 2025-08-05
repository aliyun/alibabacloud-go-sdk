// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDownloadTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteDownloadTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteDownloadTaskResponse
	GetStatusCode() *int32
	SetBody(v *DeleteDownloadTaskResponseBody) *DeleteDownloadTaskResponse
	GetBody() *DeleteDownloadTaskResponseBody
}

type DeleteDownloadTaskResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteDownloadTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteDownloadTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteDownloadTaskResponse) GoString() string {
	return s.String()
}

func (s *DeleteDownloadTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteDownloadTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteDownloadTaskResponse) GetBody() *DeleteDownloadTaskResponseBody {
	return s.Body
}

func (s *DeleteDownloadTaskResponse) SetHeaders(v map[string]*string) *DeleteDownloadTaskResponse {
	s.Headers = v
	return s
}

func (s *DeleteDownloadTaskResponse) SetStatusCode(v int32) *DeleteDownloadTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteDownloadTaskResponse) SetBody(v *DeleteDownloadTaskResponseBody) *DeleteDownloadTaskResponse {
	s.Body = v
	return s
}

func (s *DeleteDownloadTaskResponse) Validate() error {
	return dara.Validate(s)
}
