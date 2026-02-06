// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSuspendCallResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SuspendCallResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SuspendCallResponse
	GetStatusCode() *int32
	SetBody(v *SuspendCallResponseBody) *SuspendCallResponse
	GetBody() *SuspendCallResponseBody
}

type SuspendCallResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SuspendCallResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SuspendCallResponse) String() string {
	return dara.Prettify(s)
}

func (s SuspendCallResponse) GoString() string {
	return s.String()
}

func (s *SuspendCallResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SuspendCallResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SuspendCallResponse) GetBody() *SuspendCallResponseBody {
	return s.Body
}

func (s *SuspendCallResponse) SetHeaders(v map[string]*string) *SuspendCallResponse {
	s.Headers = v
	return s
}

func (s *SuspendCallResponse) SetStatusCode(v int32) *SuspendCallResponse {
	s.StatusCode = &v
	return s
}

func (s *SuspendCallResponse) SetBody(v *SuspendCallResponseBody) *SuspendCallResponse {
	s.Body = v
	return s
}

func (s *SuspendCallResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
