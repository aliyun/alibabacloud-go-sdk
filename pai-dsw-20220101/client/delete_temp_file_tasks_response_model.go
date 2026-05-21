// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteTempFileTasksResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteTempFileTasksResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteTempFileTasksResponse
	GetStatusCode() *int32
	SetBody(v *DeleteTempFileTasksResponseBody) *DeleteTempFileTasksResponse
	GetBody() *DeleteTempFileTasksResponseBody
}

type DeleteTempFileTasksResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteTempFileTasksResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteTempFileTasksResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteTempFileTasksResponse) GoString() string {
	return s.String()
}

func (s *DeleteTempFileTasksResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteTempFileTasksResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteTempFileTasksResponse) GetBody() *DeleteTempFileTasksResponseBody {
	return s.Body
}

func (s *DeleteTempFileTasksResponse) SetHeaders(v map[string]*string) *DeleteTempFileTasksResponse {
	s.Headers = v
	return s
}

func (s *DeleteTempFileTasksResponse) SetStatusCode(v int32) *DeleteTempFileTasksResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteTempFileTasksResponse) SetBody(v *DeleteTempFileTasksResponseBody) *DeleteTempFileTasksResponse {
	s.Body = v
	return s
}

func (s *DeleteTempFileTasksResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
