// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopOfflineTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *StopOfflineTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *StopOfflineTaskResponse
	GetStatusCode() *int32
	SetBody(v *StopOfflineTaskResponseBody) *StopOfflineTaskResponse
	GetBody() *StopOfflineTaskResponseBody
}

type StopOfflineTaskResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StopOfflineTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StopOfflineTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s StopOfflineTaskResponse) GoString() string {
	return s.String()
}

func (s *StopOfflineTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *StopOfflineTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *StopOfflineTaskResponse) GetBody() *StopOfflineTaskResponseBody {
	return s.Body
}

func (s *StopOfflineTaskResponse) SetHeaders(v map[string]*string) *StopOfflineTaskResponse {
	s.Headers = v
	return s
}

func (s *StopOfflineTaskResponse) SetStatusCode(v int32) *StopOfflineTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *StopOfflineTaskResponse) SetBody(v *StopOfflineTaskResponseBody) *StopOfflineTaskResponse {
	s.Body = v
	return s
}

func (s *StopOfflineTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
