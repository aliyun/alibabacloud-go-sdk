// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTerminateAiOutboundTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *TerminateAiOutboundTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *TerminateAiOutboundTaskResponse
	GetStatusCode() *int32
	SetBody(v *TerminateAiOutboundTaskResponseBody) *TerminateAiOutboundTaskResponse
	GetBody() *TerminateAiOutboundTaskResponseBody
}

type TerminateAiOutboundTaskResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *TerminateAiOutboundTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s TerminateAiOutboundTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s TerminateAiOutboundTaskResponse) GoString() string {
	return s.String()
}

func (s *TerminateAiOutboundTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *TerminateAiOutboundTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *TerminateAiOutboundTaskResponse) GetBody() *TerminateAiOutboundTaskResponseBody {
	return s.Body
}

func (s *TerminateAiOutboundTaskResponse) SetHeaders(v map[string]*string) *TerminateAiOutboundTaskResponse {
	s.Headers = v
	return s
}

func (s *TerminateAiOutboundTaskResponse) SetStatusCode(v int32) *TerminateAiOutboundTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *TerminateAiOutboundTaskResponse) SetBody(v *TerminateAiOutboundTaskResponseBody) *TerminateAiOutboundTaskResponse {
	s.Body = v
	return s
}

func (s *TerminateAiOutboundTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
