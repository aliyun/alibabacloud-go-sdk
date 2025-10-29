// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopLivePullToPushResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *StopLivePullToPushResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *StopLivePullToPushResponse
	GetStatusCode() *int32
	SetBody(v *StopLivePullToPushResponseBody) *StopLivePullToPushResponse
	GetBody() *StopLivePullToPushResponseBody
}

type StopLivePullToPushResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StopLivePullToPushResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StopLivePullToPushResponse) String() string {
	return dara.Prettify(s)
}

func (s StopLivePullToPushResponse) GoString() string {
	return s.String()
}

func (s *StopLivePullToPushResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *StopLivePullToPushResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *StopLivePullToPushResponse) GetBody() *StopLivePullToPushResponseBody {
	return s.Body
}

func (s *StopLivePullToPushResponse) SetHeaders(v map[string]*string) *StopLivePullToPushResponse {
	s.Headers = v
	return s
}

func (s *StopLivePullToPushResponse) SetStatusCode(v int32) *StopLivePullToPushResponse {
	s.StatusCode = &v
	return s
}

func (s *StopLivePullToPushResponse) SetBody(v *StopLivePullToPushResponseBody) *StopLivePullToPushResponse {
	s.Body = v
	return s
}

func (s *StopLivePullToPushResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
