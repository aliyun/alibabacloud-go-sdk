// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDISyncTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateDISyncTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateDISyncTaskResponse
	GetStatusCode() *int32
	SetBody(v *UpdateDISyncTaskResponseBody) *UpdateDISyncTaskResponse
	GetBody() *UpdateDISyncTaskResponseBody
}

type UpdateDISyncTaskResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateDISyncTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateDISyncTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateDISyncTaskResponse) GoString() string {
	return s.String()
}

func (s *UpdateDISyncTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateDISyncTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateDISyncTaskResponse) GetBody() *UpdateDISyncTaskResponseBody {
	return s.Body
}

func (s *UpdateDISyncTaskResponse) SetHeaders(v map[string]*string) *UpdateDISyncTaskResponse {
	s.Headers = v
	return s
}

func (s *UpdateDISyncTaskResponse) SetStatusCode(v int32) *UpdateDISyncTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateDISyncTaskResponse) SetBody(v *UpdateDISyncTaskResponseBody) *UpdateDISyncTaskResponse {
	s.Body = v
	return s
}

func (s *UpdateDISyncTaskResponse) Validate() error {
	return dara.Validate(s)
}
