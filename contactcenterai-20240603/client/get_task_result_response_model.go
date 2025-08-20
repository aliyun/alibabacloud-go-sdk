// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTaskResultResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetTaskResultResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetTaskResultResponse
	GetStatusCode() *int32
	SetBody(v *GetTaskResultResponseBody) *GetTaskResultResponse
	GetBody() *GetTaskResultResponseBody
}

type GetTaskResultResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetTaskResultResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetTaskResultResponse) String() string {
	return dara.Prettify(s)
}

func (s GetTaskResultResponse) GoString() string {
	return s.String()
}

func (s *GetTaskResultResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetTaskResultResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetTaskResultResponse) GetBody() *GetTaskResultResponseBody {
	return s.Body
}

func (s *GetTaskResultResponse) SetHeaders(v map[string]*string) *GetTaskResultResponse {
	s.Headers = v
	return s
}

func (s *GetTaskResultResponse) SetStatusCode(v int32) *GetTaskResultResponse {
	s.StatusCode = &v
	return s
}

func (s *GetTaskResultResponse) SetBody(v *GetTaskResultResponseBody) *GetTaskResultResponse {
	s.Body = v
	return s
}

func (s *GetTaskResultResponse) Validate() error {
	return dara.Validate(s)
}
