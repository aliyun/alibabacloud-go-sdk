// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDataTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteDataTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteDataTaskResponse
	GetStatusCode() *int32
	SetBody(v *DeleteDataTaskResponseBody) *DeleteDataTaskResponse
	GetBody() *DeleteDataTaskResponseBody
}

type DeleteDataTaskResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteDataTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteDataTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteDataTaskResponse) GoString() string {
	return s.String()
}

func (s *DeleteDataTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteDataTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteDataTaskResponse) GetBody() *DeleteDataTaskResponseBody {
	return s.Body
}

func (s *DeleteDataTaskResponse) SetHeaders(v map[string]*string) *DeleteDataTaskResponse {
	s.Headers = v
	return s
}

func (s *DeleteDataTaskResponse) SetStatusCode(v int32) *DeleteDataTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteDataTaskResponse) SetBody(v *DeleteDataTaskResponseBody) *DeleteDataTaskResponse {
	s.Body = v
	return s
}

func (s *DeleteDataTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
