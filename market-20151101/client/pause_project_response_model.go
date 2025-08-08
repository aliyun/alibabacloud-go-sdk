// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPauseProjectResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *PauseProjectResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *PauseProjectResponse
	GetStatusCode() *int32
	SetBody(v *PauseProjectResponseBody) *PauseProjectResponse
	GetBody() *PauseProjectResponseBody
}

type PauseProjectResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PauseProjectResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PauseProjectResponse) String() string {
	return dara.Prettify(s)
}

func (s PauseProjectResponse) GoString() string {
	return s.String()
}

func (s *PauseProjectResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *PauseProjectResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *PauseProjectResponse) GetBody() *PauseProjectResponseBody {
	return s.Body
}

func (s *PauseProjectResponse) SetHeaders(v map[string]*string) *PauseProjectResponse {
	s.Headers = v
	return s
}

func (s *PauseProjectResponse) SetStatusCode(v int32) *PauseProjectResponse {
	s.StatusCode = &v
	return s
}

func (s *PauseProjectResponse) SetBody(v *PauseProjectResponseBody) *PauseProjectResponse {
	s.Body = v
	return s
}

func (s *PauseProjectResponse) Validate() error {
	return dara.Validate(s)
}
