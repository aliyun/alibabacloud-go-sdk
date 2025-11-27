// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopTransferStreamResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *StopTransferStreamResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *StopTransferStreamResponse
	GetStatusCode() *int32
	SetBody(v *StopTransferStreamResponseBody) *StopTransferStreamResponse
	GetBody() *StopTransferStreamResponseBody
}

type StopTransferStreamResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StopTransferStreamResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StopTransferStreamResponse) String() string {
	return dara.Prettify(s)
}

func (s StopTransferStreamResponse) GoString() string {
	return s.String()
}

func (s *StopTransferStreamResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *StopTransferStreamResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *StopTransferStreamResponse) GetBody() *StopTransferStreamResponseBody {
	return s.Body
}

func (s *StopTransferStreamResponse) SetHeaders(v map[string]*string) *StopTransferStreamResponse {
	s.Headers = v
	return s
}

func (s *StopTransferStreamResponse) SetStatusCode(v int32) *StopTransferStreamResponse {
	s.StatusCode = &v
	return s
}

func (s *StopTransferStreamResponse) SetBody(v *StopTransferStreamResponseBody) *StopTransferStreamResponse {
	s.Body = v
	return s
}

func (s *StopTransferStreamResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
