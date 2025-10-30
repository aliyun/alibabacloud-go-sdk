// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPocCreateTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *PocCreateTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *PocCreateTaskResponse
	GetStatusCode() *int32
	SetBody(v *PocCreateTaskResponseBody) *PocCreateTaskResponse
	GetBody() *PocCreateTaskResponseBody
}

type PocCreateTaskResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PocCreateTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PocCreateTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s PocCreateTaskResponse) GoString() string {
	return s.String()
}

func (s *PocCreateTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *PocCreateTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *PocCreateTaskResponse) GetBody() *PocCreateTaskResponseBody {
	return s.Body
}

func (s *PocCreateTaskResponse) SetHeaders(v map[string]*string) *PocCreateTaskResponse {
	s.Headers = v
	return s
}

func (s *PocCreateTaskResponse) SetStatusCode(v int32) *PocCreateTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *PocCreateTaskResponse) SetBody(v *PocCreateTaskResponseBody) *PocCreateTaskResponse {
	s.Body = v
	return s
}

func (s *PocCreateTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
