// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteTempFileTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteTempFileTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteTempFileTaskResponse
	GetStatusCode() *int32
	SetBody(v *DeleteTempFileTaskResponseBody) *DeleteTempFileTaskResponse
	GetBody() *DeleteTempFileTaskResponseBody
}

type DeleteTempFileTaskResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteTempFileTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteTempFileTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteTempFileTaskResponse) GoString() string {
	return s.String()
}

func (s *DeleteTempFileTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteTempFileTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteTempFileTaskResponse) GetBody() *DeleteTempFileTaskResponseBody {
	return s.Body
}

func (s *DeleteTempFileTaskResponse) SetHeaders(v map[string]*string) *DeleteTempFileTaskResponse {
	s.Headers = v
	return s
}

func (s *DeleteTempFileTaskResponse) SetStatusCode(v int32) *DeleteTempFileTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteTempFileTaskResponse) SetBody(v *DeleteTempFileTaskResponseBody) *DeleteTempFileTaskResponse {
	s.Body = v
	return s
}

func (s *DeleteTempFileTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
