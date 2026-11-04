// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopTriggeredRecordResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *StopTriggeredRecordResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *StopTriggeredRecordResponse
	GetStatusCode() *int32
	SetBody(v *StopTriggeredRecordResponseBody) *StopTriggeredRecordResponse
	GetBody() *StopTriggeredRecordResponseBody
}

type StopTriggeredRecordResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StopTriggeredRecordResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StopTriggeredRecordResponse) String() string {
	return dara.Prettify(s)
}

func (s StopTriggeredRecordResponse) GoString() string {
	return s.String()
}

func (s *StopTriggeredRecordResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *StopTriggeredRecordResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *StopTriggeredRecordResponse) GetBody() *StopTriggeredRecordResponseBody {
	return s.Body
}

func (s *StopTriggeredRecordResponse) SetHeaders(v map[string]*string) *StopTriggeredRecordResponse {
	s.Headers = v
	return s
}

func (s *StopTriggeredRecordResponse) SetStatusCode(v int32) *StopTriggeredRecordResponse {
	s.StatusCode = &v
	return s
}

func (s *StopTriggeredRecordResponse) SetBody(v *StopTriggeredRecordResponseBody) *StopTriggeredRecordResponse {
	s.Body = v
	return s
}

func (s *StopTriggeredRecordResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
