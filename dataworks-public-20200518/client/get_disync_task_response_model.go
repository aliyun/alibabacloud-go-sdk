// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDISyncTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetDISyncTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetDISyncTaskResponse
	GetStatusCode() *int32
	SetBody(v *GetDISyncTaskResponseBody) *GetDISyncTaskResponse
	GetBody() *GetDISyncTaskResponseBody
}

type GetDISyncTaskResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetDISyncTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetDISyncTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s GetDISyncTaskResponse) GoString() string {
	return s.String()
}

func (s *GetDISyncTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetDISyncTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetDISyncTaskResponse) GetBody() *GetDISyncTaskResponseBody {
	return s.Body
}

func (s *GetDISyncTaskResponse) SetHeaders(v map[string]*string) *GetDISyncTaskResponse {
	s.Headers = v
	return s
}

func (s *GetDISyncTaskResponse) SetStatusCode(v int32) *GetDISyncTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDISyncTaskResponse) SetBody(v *GetDISyncTaskResponseBody) *GetDISyncTaskResponse {
	s.Body = v
	return s
}

func (s *GetDISyncTaskResponse) Validate() error {
	return dara.Validate(s)
}
