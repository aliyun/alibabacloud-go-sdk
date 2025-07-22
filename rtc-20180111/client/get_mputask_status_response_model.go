// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMPUTaskStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetMPUTaskStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetMPUTaskStatusResponse
	GetStatusCode() *int32
	SetBody(v *GetMPUTaskStatusResponseBody) *GetMPUTaskStatusResponse
	GetBody() *GetMPUTaskStatusResponseBody
}

type GetMPUTaskStatusResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetMPUTaskStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetMPUTaskStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s GetMPUTaskStatusResponse) GoString() string {
	return s.String()
}

func (s *GetMPUTaskStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetMPUTaskStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetMPUTaskStatusResponse) GetBody() *GetMPUTaskStatusResponseBody {
	return s.Body
}

func (s *GetMPUTaskStatusResponse) SetHeaders(v map[string]*string) *GetMPUTaskStatusResponse {
	s.Headers = v
	return s
}

func (s *GetMPUTaskStatusResponse) SetStatusCode(v int32) *GetMPUTaskStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *GetMPUTaskStatusResponse) SetBody(v *GetMPUTaskStatusResponseBody) *GetMPUTaskStatusResponse {
	s.Body = v
	return s
}

func (s *GetMPUTaskStatusResponse) Validate() error {
	return dara.Validate(s)
}
