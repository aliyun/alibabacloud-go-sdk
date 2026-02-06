// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTerminateCallResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *TerminateCallResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *TerminateCallResponse
	GetStatusCode() *int32
	SetBody(v *TerminateCallResponseBody) *TerminateCallResponse
	GetBody() *TerminateCallResponseBody
}

type TerminateCallResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *TerminateCallResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s TerminateCallResponse) String() string {
	return dara.Prettify(s)
}

func (s TerminateCallResponse) GoString() string {
	return s.String()
}

func (s *TerminateCallResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *TerminateCallResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *TerminateCallResponse) GetBody() *TerminateCallResponseBody {
	return s.Body
}

func (s *TerminateCallResponse) SetHeaders(v map[string]*string) *TerminateCallResponse {
	s.Headers = v
	return s
}

func (s *TerminateCallResponse) SetStatusCode(v int32) *TerminateCallResponse {
	s.StatusCode = &v
	return s
}

func (s *TerminateCallResponse) SetBody(v *TerminateCallResponseBody) *TerminateCallResponse {
	s.Body = v
	return s
}

func (s *TerminateCallResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
