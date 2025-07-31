// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetExecuteTimeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetExecuteTimeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetExecuteTimeResponse
	GetStatusCode() *int32
	SetBody(v *GetExecuteTimeResponseBody) *GetExecuteTimeResponse
	GetBody() *GetExecuteTimeResponseBody
}

type GetExecuteTimeResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetExecuteTimeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetExecuteTimeResponse) String() string {
	return dara.Prettify(s)
}

func (s GetExecuteTimeResponse) GoString() string {
	return s.String()
}

func (s *GetExecuteTimeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetExecuteTimeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetExecuteTimeResponse) GetBody() *GetExecuteTimeResponseBody {
	return s.Body
}

func (s *GetExecuteTimeResponse) SetHeaders(v map[string]*string) *GetExecuteTimeResponse {
	s.Headers = v
	return s
}

func (s *GetExecuteTimeResponse) SetStatusCode(v int32) *GetExecuteTimeResponse {
	s.StatusCode = &v
	return s
}

func (s *GetExecuteTimeResponse) SetBody(v *GetExecuteTimeResponseBody) *GetExecuteTimeResponse {
	s.Body = v
	return s
}

func (s *GetExecuteTimeResponse) Validate() error {
	return dara.Validate(s)
}
