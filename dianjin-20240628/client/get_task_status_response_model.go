// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTaskStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetTaskStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetTaskStatusResponse
	GetStatusCode() *int32
	SetBody(v *GetTaskStatusResponseBody) *GetTaskStatusResponse
	GetBody() *GetTaskStatusResponseBody
}

type GetTaskStatusResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetTaskStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetTaskStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s GetTaskStatusResponse) GoString() string {
	return s.String()
}

func (s *GetTaskStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetTaskStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetTaskStatusResponse) GetBody() *GetTaskStatusResponseBody {
	return s.Body
}

func (s *GetTaskStatusResponse) SetHeaders(v map[string]*string) *GetTaskStatusResponse {
	s.Headers = v
	return s
}

func (s *GetTaskStatusResponse) SetStatusCode(v int32) *GetTaskStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *GetTaskStatusResponse) SetBody(v *GetTaskStatusResponseBody) *GetTaskStatusResponse {
	s.Body = v
	return s
}

func (s *GetTaskStatusResponse) Validate() error {
	return dara.Validate(s)
}
