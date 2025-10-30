// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteBatchTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteBatchTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteBatchTaskResponse
	GetStatusCode() *int32
	SetBody(v *DeleteBatchTaskResponseBody) *DeleteBatchTaskResponse
	GetBody() *DeleteBatchTaskResponseBody
}

type DeleteBatchTaskResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteBatchTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteBatchTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteBatchTaskResponse) GoString() string {
	return s.String()
}

func (s *DeleteBatchTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteBatchTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteBatchTaskResponse) GetBody() *DeleteBatchTaskResponseBody {
	return s.Body
}

func (s *DeleteBatchTaskResponse) SetHeaders(v map[string]*string) *DeleteBatchTaskResponse {
	s.Headers = v
	return s
}

func (s *DeleteBatchTaskResponse) SetStatusCode(v int32) *DeleteBatchTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteBatchTaskResponse) SetBody(v *DeleteBatchTaskResponseBody) *DeleteBatchTaskResponse {
	s.Body = v
	return s
}

func (s *DeleteBatchTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
