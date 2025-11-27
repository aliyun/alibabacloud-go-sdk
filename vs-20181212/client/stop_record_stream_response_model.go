// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopRecordStreamResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *StopRecordStreamResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *StopRecordStreamResponse
	GetStatusCode() *int32
	SetBody(v *StopRecordStreamResponseBody) *StopRecordStreamResponse
	GetBody() *StopRecordStreamResponseBody
}

type StopRecordStreamResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StopRecordStreamResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StopRecordStreamResponse) String() string {
	return dara.Prettify(s)
}

func (s StopRecordStreamResponse) GoString() string {
	return s.String()
}

func (s *StopRecordStreamResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *StopRecordStreamResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *StopRecordStreamResponse) GetBody() *StopRecordStreamResponseBody {
	return s.Body
}

func (s *StopRecordStreamResponse) SetHeaders(v map[string]*string) *StopRecordStreamResponse {
	s.Headers = v
	return s
}

func (s *StopRecordStreamResponse) SetStatusCode(v int32) *StopRecordStreamResponse {
	s.StatusCode = &v
	return s
}

func (s *StopRecordStreamResponse) SetBody(v *StopRecordStreamResponseBody) *StopRecordStreamResponse {
	s.Body = v
	return s
}

func (s *StopRecordStreamResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
