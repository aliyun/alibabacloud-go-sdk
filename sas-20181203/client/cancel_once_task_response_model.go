// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCancelOnceTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CancelOnceTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CancelOnceTaskResponse
	GetStatusCode() *int32
	SetBody(v *CancelOnceTaskResponseBody) *CancelOnceTaskResponse
	GetBody() *CancelOnceTaskResponseBody
}

type CancelOnceTaskResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CancelOnceTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CancelOnceTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s CancelOnceTaskResponse) GoString() string {
	return s.String()
}

func (s *CancelOnceTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CancelOnceTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CancelOnceTaskResponse) GetBody() *CancelOnceTaskResponseBody {
	return s.Body
}

func (s *CancelOnceTaskResponse) SetHeaders(v map[string]*string) *CancelOnceTaskResponse {
	s.Headers = v
	return s
}

func (s *CancelOnceTaskResponse) SetStatusCode(v int32) *CancelOnceTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *CancelOnceTaskResponse) SetBody(v *CancelOnceTaskResponseBody) *CancelOnceTaskResponse {
	s.Body = v
	return s
}

func (s *CancelOnceTaskResponse) Validate() error {
	return dara.Validate(s)
}
