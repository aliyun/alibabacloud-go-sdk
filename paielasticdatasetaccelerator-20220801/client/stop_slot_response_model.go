// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopSlotResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *StopSlotResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *StopSlotResponse
	GetStatusCode() *int32
	SetBody(v *StopSlotResponseBody) *StopSlotResponse
	GetBody() *StopSlotResponseBody
}

type StopSlotResponse struct {
	Headers    map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StopSlotResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StopSlotResponse) String() string {
	return dara.Prettify(s)
}

func (s StopSlotResponse) GoString() string {
	return s.String()
}

func (s *StopSlotResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *StopSlotResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *StopSlotResponse) GetBody() *StopSlotResponseBody {
	return s.Body
}

func (s *StopSlotResponse) SetHeaders(v map[string]*string) *StopSlotResponse {
	s.Headers = v
	return s
}

func (s *StopSlotResponse) SetStatusCode(v int32) *StopSlotResponse {
	s.StatusCode = &v
	return s
}

func (s *StopSlotResponse) SetBody(v *StopSlotResponseBody) *StopSlotResponse {
	s.Body = v
	return s
}

func (s *StopSlotResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
