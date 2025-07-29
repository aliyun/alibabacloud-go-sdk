// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPauseTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *PauseTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *PauseTaskResponse
	GetStatusCode() *int32
}

type PauseTaskResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
}

func (s PauseTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s PauseTaskResponse) GoString() string {
	return s.String()
}

func (s *PauseTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *PauseTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *PauseTaskResponse) SetHeaders(v map[string]*string) *PauseTaskResponse {
	s.Headers = v
	return s
}

func (s *PauseTaskResponse) SetStatusCode(v int32) *PauseTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *PauseTaskResponse) Validate() error {
	return dara.Validate(s)
}
