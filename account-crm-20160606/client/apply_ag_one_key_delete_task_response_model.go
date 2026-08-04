// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iApplyAgOneKeyDeleteTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ApplyAgOneKeyDeleteTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ApplyAgOneKeyDeleteTaskResponse
	GetStatusCode() *int32
	SetBody(v *ApplyAgOneKeyDeleteTaskResponseBody) *ApplyAgOneKeyDeleteTaskResponse
	GetBody() *ApplyAgOneKeyDeleteTaskResponseBody
}

type ApplyAgOneKeyDeleteTaskResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ApplyAgOneKeyDeleteTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ApplyAgOneKeyDeleteTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s ApplyAgOneKeyDeleteTaskResponse) GoString() string {
	return s.String()
}

func (s *ApplyAgOneKeyDeleteTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ApplyAgOneKeyDeleteTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ApplyAgOneKeyDeleteTaskResponse) GetBody() *ApplyAgOneKeyDeleteTaskResponseBody {
	return s.Body
}

func (s *ApplyAgOneKeyDeleteTaskResponse) SetHeaders(v map[string]*string) *ApplyAgOneKeyDeleteTaskResponse {
	s.Headers = v
	return s
}

func (s *ApplyAgOneKeyDeleteTaskResponse) SetStatusCode(v int32) *ApplyAgOneKeyDeleteTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *ApplyAgOneKeyDeleteTaskResponse) SetBody(v *ApplyAgOneKeyDeleteTaskResponseBody) *ApplyAgOneKeyDeleteTaskResponse {
	s.Body = v
	return s
}

func (s *ApplyAgOneKeyDeleteTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
