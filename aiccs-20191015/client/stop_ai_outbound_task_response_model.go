// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopAiOutboundTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *StopAiOutboundTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *StopAiOutboundTaskResponse
	GetStatusCode() *int32
	SetBody(v *StopAiOutboundTaskResponseBody) *StopAiOutboundTaskResponse
	GetBody() *StopAiOutboundTaskResponseBody
}

type StopAiOutboundTaskResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StopAiOutboundTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StopAiOutboundTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s StopAiOutboundTaskResponse) GoString() string {
	return s.String()
}

func (s *StopAiOutboundTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *StopAiOutboundTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *StopAiOutboundTaskResponse) GetBody() *StopAiOutboundTaskResponseBody {
	return s.Body
}

func (s *StopAiOutboundTaskResponse) SetHeaders(v map[string]*string) *StopAiOutboundTaskResponse {
	s.Headers = v
	return s
}

func (s *StopAiOutboundTaskResponse) SetStatusCode(v int32) *StopAiOutboundTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *StopAiOutboundTaskResponse) SetBody(v *StopAiOutboundTaskResponseBody) *StopAiOutboundTaskResponse {
	s.Body = v
	return s
}

func (s *StopAiOutboundTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
