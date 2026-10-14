// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteOptimizeTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteOptimizeTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteOptimizeTaskResponse
	GetStatusCode() *int32
	SetBody(v *DeleteOptimizeTaskResponseBody) *DeleteOptimizeTaskResponse
	GetBody() *DeleteOptimizeTaskResponseBody
}

type DeleteOptimizeTaskResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteOptimizeTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteOptimizeTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteOptimizeTaskResponse) GoString() string {
	return s.String()
}

func (s *DeleteOptimizeTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteOptimizeTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteOptimizeTaskResponse) GetBody() *DeleteOptimizeTaskResponseBody {
	return s.Body
}

func (s *DeleteOptimizeTaskResponse) SetHeaders(v map[string]*string) *DeleteOptimizeTaskResponse {
	s.Headers = v
	return s
}

func (s *DeleteOptimizeTaskResponse) SetStatusCode(v int32) *DeleteOptimizeTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteOptimizeTaskResponse) SetBody(v *DeleteOptimizeTaskResponseBody) *DeleteOptimizeTaskResponse {
	s.Body = v
	return s
}

func (s *DeleteOptimizeTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
