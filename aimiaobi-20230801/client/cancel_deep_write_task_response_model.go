// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCancelDeepWriteTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CancelDeepWriteTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CancelDeepWriteTaskResponse
	GetStatusCode() *int32
	SetBody(v *CancelDeepWriteTaskResponseBody) *CancelDeepWriteTaskResponse
	GetBody() *CancelDeepWriteTaskResponseBody
}

type CancelDeepWriteTaskResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CancelDeepWriteTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CancelDeepWriteTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s CancelDeepWriteTaskResponse) GoString() string {
	return s.String()
}

func (s *CancelDeepWriteTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CancelDeepWriteTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CancelDeepWriteTaskResponse) GetBody() *CancelDeepWriteTaskResponseBody {
	return s.Body
}

func (s *CancelDeepWriteTaskResponse) SetHeaders(v map[string]*string) *CancelDeepWriteTaskResponse {
	s.Headers = v
	return s
}

func (s *CancelDeepWriteTaskResponse) SetStatusCode(v int32) *CancelDeepWriteTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *CancelDeepWriteTaskResponse) SetBody(v *CancelDeepWriteTaskResponseBody) *CancelDeepWriteTaskResponse {
	s.Body = v
	return s
}

func (s *CancelDeepWriteTaskResponse) Validate() error {
	return dara.Validate(s)
}
