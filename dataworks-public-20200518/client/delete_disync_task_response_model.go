// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDISyncTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteDISyncTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteDISyncTaskResponse
	GetStatusCode() *int32
	SetBody(v *DeleteDISyncTaskResponseBody) *DeleteDISyncTaskResponse
	GetBody() *DeleteDISyncTaskResponseBody
}

type DeleteDISyncTaskResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteDISyncTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteDISyncTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteDISyncTaskResponse) GoString() string {
	return s.String()
}

func (s *DeleteDISyncTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteDISyncTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteDISyncTaskResponse) GetBody() *DeleteDISyncTaskResponseBody {
	return s.Body
}

func (s *DeleteDISyncTaskResponse) SetHeaders(v map[string]*string) *DeleteDISyncTaskResponse {
	s.Headers = v
	return s
}

func (s *DeleteDISyncTaskResponse) SetStatusCode(v int32) *DeleteDISyncTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteDISyncTaskResponse) SetBody(v *DeleteDISyncTaskResponseBody) *DeleteDISyncTaskResponse {
	s.Body = v
	return s
}

func (s *DeleteDISyncTaskResponse) Validate() error {
	return dara.Validate(s)
}
