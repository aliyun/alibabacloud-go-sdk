// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopRecordTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *StopRecordTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *StopRecordTaskResponse
	GetStatusCode() *int32
	SetBody(v *StopRecordTaskResponseBody) *StopRecordTaskResponse
	GetBody() *StopRecordTaskResponseBody
}

type StopRecordTaskResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StopRecordTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StopRecordTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s StopRecordTaskResponse) GoString() string {
	return s.String()
}

func (s *StopRecordTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *StopRecordTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *StopRecordTaskResponse) GetBody() *StopRecordTaskResponseBody {
	return s.Body
}

func (s *StopRecordTaskResponse) SetHeaders(v map[string]*string) *StopRecordTaskResponse {
	s.Headers = v
	return s
}

func (s *StopRecordTaskResponse) SetStatusCode(v int32) *StopRecordTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *StopRecordTaskResponse) SetBody(v *StopRecordTaskResponseBody) *StopRecordTaskResponse {
	s.Body = v
	return s
}

func (s *StopRecordTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
