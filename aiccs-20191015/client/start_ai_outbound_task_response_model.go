// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartAiOutboundTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *StartAiOutboundTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *StartAiOutboundTaskResponse
	GetStatusCode() *int32
	SetBody(v *StartAiOutboundTaskResponseBody) *StartAiOutboundTaskResponse
	GetBody() *StartAiOutboundTaskResponseBody
}

type StartAiOutboundTaskResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StartAiOutboundTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StartAiOutboundTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s StartAiOutboundTaskResponse) GoString() string {
	return s.String()
}

func (s *StartAiOutboundTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *StartAiOutboundTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *StartAiOutboundTaskResponse) GetBody() *StartAiOutboundTaskResponseBody {
	return s.Body
}

func (s *StartAiOutboundTaskResponse) SetHeaders(v map[string]*string) *StartAiOutboundTaskResponse {
	s.Headers = v
	return s
}

func (s *StartAiOutboundTaskResponse) SetStatusCode(v int32) *StartAiOutboundTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *StartAiOutboundTaskResponse) SetBody(v *StartAiOutboundTaskResponseBody) *StartAiOutboundTaskResponse {
	s.Body = v
	return s
}

func (s *StartAiOutboundTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
