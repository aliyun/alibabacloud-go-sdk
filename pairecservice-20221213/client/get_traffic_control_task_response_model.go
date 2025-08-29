// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTrafficControlTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetTrafficControlTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetTrafficControlTaskResponse
	GetStatusCode() *int32
	SetBody(v *GetTrafficControlTaskResponseBody) *GetTrafficControlTaskResponse
	GetBody() *GetTrafficControlTaskResponseBody
}

type GetTrafficControlTaskResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetTrafficControlTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetTrafficControlTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s GetTrafficControlTaskResponse) GoString() string {
	return s.String()
}

func (s *GetTrafficControlTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetTrafficControlTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetTrafficControlTaskResponse) GetBody() *GetTrafficControlTaskResponseBody {
	return s.Body
}

func (s *GetTrafficControlTaskResponse) SetHeaders(v map[string]*string) *GetTrafficControlTaskResponse {
	s.Headers = v
	return s
}

func (s *GetTrafficControlTaskResponse) SetStatusCode(v int32) *GetTrafficControlTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *GetTrafficControlTaskResponse) SetBody(v *GetTrafficControlTaskResponseBody) *GetTrafficControlTaskResponse {
	s.Body = v
	return s
}

func (s *GetTrafficControlTaskResponse) Validate() error {
	return dara.Validate(s)
}
