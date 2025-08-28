// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCallTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateCallTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateCallTaskResponse
	GetStatusCode() *int32
	SetBody(v *CreateCallTaskResponseBody) *CreateCallTaskResponse
	GetBody() *CreateCallTaskResponseBody
}

type CreateCallTaskResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateCallTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateCallTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateCallTaskResponse) GoString() string {
	return s.String()
}

func (s *CreateCallTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateCallTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateCallTaskResponse) GetBody() *CreateCallTaskResponseBody {
	return s.Body
}

func (s *CreateCallTaskResponse) SetHeaders(v map[string]*string) *CreateCallTaskResponse {
	s.Headers = v
	return s
}

func (s *CreateCallTaskResponse) SetStatusCode(v int32) *CreateCallTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateCallTaskResponse) SetBody(v *CreateCallTaskResponseBody) *CreateCallTaskResponse {
	s.Body = v
	return s
}

func (s *CreateCallTaskResponse) Validate() error {
	return dara.Validate(s)
}
