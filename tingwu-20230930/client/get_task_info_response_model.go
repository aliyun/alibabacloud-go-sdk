// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTaskInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetTaskInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetTaskInfoResponse
	GetStatusCode() *int32
	SetBody(v *GetTaskInfoResponseBody) *GetTaskInfoResponse
	GetBody() *GetTaskInfoResponseBody
}

type GetTaskInfoResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetTaskInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetTaskInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s GetTaskInfoResponse) GoString() string {
	return s.String()
}

func (s *GetTaskInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetTaskInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetTaskInfoResponse) GetBody() *GetTaskInfoResponseBody {
	return s.Body
}

func (s *GetTaskInfoResponse) SetHeaders(v map[string]*string) *GetTaskInfoResponse {
	s.Headers = v
	return s
}

func (s *GetTaskInfoResponse) SetStatusCode(v int32) *GetTaskInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *GetTaskInfoResponse) SetBody(v *GetTaskInfoResponseBody) *GetTaskInfoResponse {
	s.Body = v
	return s
}

func (s *GetTaskInfoResponse) Validate() error {
	return dara.Validate(s)
}
