// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDISyncTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateDISyncTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateDISyncTaskResponse
	GetStatusCode() *int32
	SetBody(v *CreateDISyncTaskResponseBody) *CreateDISyncTaskResponse
	GetBody() *CreateDISyncTaskResponseBody
}

type CreateDISyncTaskResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateDISyncTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateDISyncTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateDISyncTaskResponse) GoString() string {
	return s.String()
}

func (s *CreateDISyncTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateDISyncTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateDISyncTaskResponse) GetBody() *CreateDISyncTaskResponseBody {
	return s.Body
}

func (s *CreateDISyncTaskResponse) SetHeaders(v map[string]*string) *CreateDISyncTaskResponse {
	s.Headers = v
	return s
}

func (s *CreateDISyncTaskResponse) SetStatusCode(v int32) *CreateDISyncTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateDISyncTaskResponse) SetBody(v *CreateDISyncTaskResponseBody) *CreateDISyncTaskResponse {
	s.Body = v
	return s
}

func (s *CreateDISyncTaskResponse) Validate() error {
	return dara.Validate(s)
}
