// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQuickAddTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QuickAddTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QuickAddTaskResponse
	GetStatusCode() *int32
	SetBody(v *QuickAddTaskResponseBody) *QuickAddTaskResponse
	GetBody() *QuickAddTaskResponseBody
}

type QuickAddTaskResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QuickAddTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QuickAddTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s QuickAddTaskResponse) GoString() string {
	return s.String()
}

func (s *QuickAddTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QuickAddTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QuickAddTaskResponse) GetBody() *QuickAddTaskResponseBody {
	return s.Body
}

func (s *QuickAddTaskResponse) SetHeaders(v map[string]*string) *QuickAddTaskResponse {
	s.Headers = v
	return s
}

func (s *QuickAddTaskResponse) SetStatusCode(v int32) *QuickAddTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *QuickAddTaskResponse) SetBody(v *QuickAddTaskResponseBody) *QuickAddTaskResponse {
	s.Body = v
	return s
}

func (s *QuickAddTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
