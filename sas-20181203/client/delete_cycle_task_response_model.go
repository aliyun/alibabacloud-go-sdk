// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteCycleTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteCycleTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteCycleTaskResponse
	GetStatusCode() *int32
	SetBody(v *DeleteCycleTaskResponseBody) *DeleteCycleTaskResponse
	GetBody() *DeleteCycleTaskResponseBody
}

type DeleteCycleTaskResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteCycleTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteCycleTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteCycleTaskResponse) GoString() string {
	return s.String()
}

func (s *DeleteCycleTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteCycleTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteCycleTaskResponse) GetBody() *DeleteCycleTaskResponseBody {
	return s.Body
}

func (s *DeleteCycleTaskResponse) SetHeaders(v map[string]*string) *DeleteCycleTaskResponse {
	s.Headers = v
	return s
}

func (s *DeleteCycleTaskResponse) SetStatusCode(v int32) *DeleteCycleTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteCycleTaskResponse) SetBody(v *DeleteCycleTaskResponseBody) *DeleteCycleTaskResponse {
	s.Body = v
	return s
}

func (s *DeleteCycleTaskResponse) Validate() error {
	return dara.Validate(s)
}
