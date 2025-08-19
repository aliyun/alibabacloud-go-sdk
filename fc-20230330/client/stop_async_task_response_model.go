// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopAsyncTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *StopAsyncTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *StopAsyncTaskResponse
	GetStatusCode() *int32
}

type StopAsyncTaskResponse struct {
	Headers    map[string]*string `json:"headers" xml:"headers"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
}

func (s StopAsyncTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s StopAsyncTaskResponse) GoString() string {
	return s.String()
}

func (s *StopAsyncTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *StopAsyncTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *StopAsyncTaskResponse) SetHeaders(v map[string]*string) *StopAsyncTaskResponse {
	s.Headers = v
	return s
}

func (s *StopAsyncTaskResponse) SetStatusCode(v int32) *StopAsyncTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *StopAsyncTaskResponse) Validate() error {
	return dara.Validate(s)
}
