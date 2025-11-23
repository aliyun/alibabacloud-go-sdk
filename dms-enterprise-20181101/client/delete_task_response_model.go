// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteTaskResponse
	GetStatusCode() *int32
	SetBody(v *DeleteTaskResponseBody) *DeleteTaskResponse
	GetBody() *DeleteTaskResponseBody
}

type DeleteTaskResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteTaskResponse) GoString() string {
	return s.String()
}

func (s *DeleteTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteTaskResponse) GetBody() *DeleteTaskResponseBody {
	return s.Body
}

func (s *DeleteTaskResponse) SetHeaders(v map[string]*string) *DeleteTaskResponse {
	s.Headers = v
	return s
}

func (s *DeleteTaskResponse) SetStatusCode(v int32) *DeleteTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteTaskResponse) SetBody(v *DeleteTaskResponseBody) *DeleteTaskResponse {
	s.Body = v
	return s
}

func (s *DeleteTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
