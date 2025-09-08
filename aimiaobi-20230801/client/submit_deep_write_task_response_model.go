// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitDeepWriteTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SubmitDeepWriteTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SubmitDeepWriteTaskResponse
	GetStatusCode() *int32
	SetBody(v *SubmitDeepWriteTaskResponseBody) *SubmitDeepWriteTaskResponse
	GetBody() *SubmitDeepWriteTaskResponseBody
}

type SubmitDeepWriteTaskResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SubmitDeepWriteTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SubmitDeepWriteTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s SubmitDeepWriteTaskResponse) GoString() string {
	return s.String()
}

func (s *SubmitDeepWriteTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SubmitDeepWriteTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SubmitDeepWriteTaskResponse) GetBody() *SubmitDeepWriteTaskResponseBody {
	return s.Body
}

func (s *SubmitDeepWriteTaskResponse) SetHeaders(v map[string]*string) *SubmitDeepWriteTaskResponse {
	s.Headers = v
	return s
}

func (s *SubmitDeepWriteTaskResponse) SetStatusCode(v int32) *SubmitDeepWriteTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *SubmitDeepWriteTaskResponse) SetBody(v *SubmitDeepWriteTaskResponseBody) *SubmitDeepWriteTaskResponse {
	s.Body = v
	return s
}

func (s *SubmitDeepWriteTaskResponse) Validate() error {
	return dara.Validate(s)
}
