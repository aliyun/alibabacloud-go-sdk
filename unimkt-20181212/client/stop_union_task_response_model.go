// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopUnionTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *StopUnionTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *StopUnionTaskResponse
	GetStatusCode() *int32
	SetBody(v *StopUnionTaskResponseBody) *StopUnionTaskResponse
	GetBody() *StopUnionTaskResponseBody
}

type StopUnionTaskResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StopUnionTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StopUnionTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s StopUnionTaskResponse) GoString() string {
	return s.String()
}

func (s *StopUnionTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *StopUnionTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *StopUnionTaskResponse) GetBody() *StopUnionTaskResponseBody {
	return s.Body
}

func (s *StopUnionTaskResponse) SetHeaders(v map[string]*string) *StopUnionTaskResponse {
	s.Headers = v
	return s
}

func (s *StopUnionTaskResponse) SetStatusCode(v int32) *StopUnionTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *StopUnionTaskResponse) SetBody(v *StopUnionTaskResponseBody) *StopUnionTaskResponse {
	s.Body = v
	return s
}

func (s *StopUnionTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
