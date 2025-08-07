// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateRestoreTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateRestoreTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateRestoreTaskResponse
	GetStatusCode() *int32
	SetBody(v *CreateRestoreTaskResponseBody) *CreateRestoreTaskResponse
	GetBody() *CreateRestoreTaskResponseBody
}

type CreateRestoreTaskResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateRestoreTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateRestoreTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateRestoreTaskResponse) GoString() string {
	return s.String()
}

func (s *CreateRestoreTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateRestoreTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateRestoreTaskResponse) GetBody() *CreateRestoreTaskResponseBody {
	return s.Body
}

func (s *CreateRestoreTaskResponse) SetHeaders(v map[string]*string) *CreateRestoreTaskResponse {
	s.Headers = v
	return s
}

func (s *CreateRestoreTaskResponse) SetStatusCode(v int32) *CreateRestoreTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateRestoreTaskResponse) SetBody(v *CreateRestoreTaskResponseBody) *CreateRestoreTaskResponse {
	s.Body = v
	return s
}

func (s *CreateRestoreTaskResponse) Validate() error {
	return dara.Validate(s)
}
