// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitCleanTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SubmitCleanTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SubmitCleanTaskResponse
	GetStatusCode() *int32
	SetBody(v *SubmitCleanTaskResponseBody) *SubmitCleanTaskResponse
	GetBody() *SubmitCleanTaskResponseBody
}

type SubmitCleanTaskResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SubmitCleanTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SubmitCleanTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s SubmitCleanTaskResponse) GoString() string {
	return s.String()
}

func (s *SubmitCleanTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SubmitCleanTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SubmitCleanTaskResponse) GetBody() *SubmitCleanTaskResponseBody {
	return s.Body
}

func (s *SubmitCleanTaskResponse) SetHeaders(v map[string]*string) *SubmitCleanTaskResponse {
	s.Headers = v
	return s
}

func (s *SubmitCleanTaskResponse) SetStatusCode(v int32) *SubmitCleanTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *SubmitCleanTaskResponse) SetBody(v *SubmitCleanTaskResponseBody) *SubmitCleanTaskResponse {
	s.Body = v
	return s
}

func (s *SubmitCleanTaskResponse) Validate() error {
	return dara.Validate(s)
}
