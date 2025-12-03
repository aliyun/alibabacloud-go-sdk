// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPushRpaTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *PushRpaTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *PushRpaTaskResponse
	GetStatusCode() *int32
	SetBody(v *PushRpaTaskResponseBody) *PushRpaTaskResponse
	GetBody() *PushRpaTaskResponseBody
}

type PushRpaTaskResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PushRpaTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PushRpaTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s PushRpaTaskResponse) GoString() string {
	return s.String()
}

func (s *PushRpaTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *PushRpaTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *PushRpaTaskResponse) GetBody() *PushRpaTaskResponseBody {
	return s.Body
}

func (s *PushRpaTaskResponse) SetHeaders(v map[string]*string) *PushRpaTaskResponse {
	s.Headers = v
	return s
}

func (s *PushRpaTaskResponse) SetStatusCode(v int32) *PushRpaTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *PushRpaTaskResponse) SetBody(v *PushRpaTaskResponseBody) *PushRpaTaskResponse {
	s.Body = v
	return s
}

func (s *PushRpaTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
