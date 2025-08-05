// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMmsTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetMmsTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetMmsTaskResponse
	GetStatusCode() *int32
	SetBody(v *GetMmsTaskResponseBody) *GetMmsTaskResponse
	GetBody() *GetMmsTaskResponseBody
}

type GetMmsTaskResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetMmsTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetMmsTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s GetMmsTaskResponse) GoString() string {
	return s.String()
}

func (s *GetMmsTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetMmsTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetMmsTaskResponse) GetBody() *GetMmsTaskResponseBody {
	return s.Body
}

func (s *GetMmsTaskResponse) SetHeaders(v map[string]*string) *GetMmsTaskResponse {
	s.Headers = v
	return s
}

func (s *GetMmsTaskResponse) SetStatusCode(v int32) *GetMmsTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *GetMmsTaskResponse) SetBody(v *GetMmsTaskResponseBody) *GetMmsTaskResponse {
	s.Body = v
	return s
}

func (s *GetMmsTaskResponse) Validate() error {
	return dara.Validate(s)
}
