// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCancelTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CancelTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CancelTaskResponse
	GetStatusCode() *int32
	SetBody(v *Task) *CancelTaskResponse
	GetBody() *Task
}

type CancelTaskResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *Task              `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CancelTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s CancelTaskResponse) GoString() string {
	return s.String()
}

func (s *CancelTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CancelTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CancelTaskResponse) GetBody() *Task {
	return s.Body
}

func (s *CancelTaskResponse) SetHeaders(v map[string]*string) *CancelTaskResponse {
	s.Headers = v
	return s
}

func (s *CancelTaskResponse) SetStatusCode(v int32) *CancelTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *CancelTaskResponse) SetBody(v *Task) *CancelTaskResponse {
	s.Body = v
	return s
}

func (s *CancelTaskResponse) Validate() error {
	return dara.Validate(s)
}
