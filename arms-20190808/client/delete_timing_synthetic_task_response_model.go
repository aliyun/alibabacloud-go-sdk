// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteTimingSyntheticTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteTimingSyntheticTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteTimingSyntheticTaskResponse
	GetStatusCode() *int32
	SetBody(v *DeleteTimingSyntheticTaskResponseBody) *DeleteTimingSyntheticTaskResponse
	GetBody() *DeleteTimingSyntheticTaskResponseBody
}

type DeleteTimingSyntheticTaskResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteTimingSyntheticTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteTimingSyntheticTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteTimingSyntheticTaskResponse) GoString() string {
	return s.String()
}

func (s *DeleteTimingSyntheticTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteTimingSyntheticTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteTimingSyntheticTaskResponse) GetBody() *DeleteTimingSyntheticTaskResponseBody {
	return s.Body
}

func (s *DeleteTimingSyntheticTaskResponse) SetHeaders(v map[string]*string) *DeleteTimingSyntheticTaskResponse {
	s.Headers = v
	return s
}

func (s *DeleteTimingSyntheticTaskResponse) SetStatusCode(v int32) *DeleteTimingSyntheticTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteTimingSyntheticTaskResponse) SetBody(v *DeleteTimingSyntheticTaskResponseBody) *DeleteTimingSyntheticTaskResponse {
	s.Body = v
	return s
}

func (s *DeleteTimingSyntheticTaskResponse) Validate() error {
	return dara.Validate(s)
}
