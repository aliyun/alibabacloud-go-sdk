// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSuspendTrafficResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SuspendTrafficResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SuspendTrafficResponse
	GetStatusCode() *int32
	SetBody(v *SuspendTrafficResponseBody) *SuspendTrafficResponse
	GetBody() *SuspendTrafficResponseBody
}

type SuspendTrafficResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SuspendTrafficResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SuspendTrafficResponse) String() string {
	return dara.Prettify(s)
}

func (s SuspendTrafficResponse) GoString() string {
	return s.String()
}

func (s *SuspendTrafficResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SuspendTrafficResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SuspendTrafficResponse) GetBody() *SuspendTrafficResponseBody {
	return s.Body
}

func (s *SuspendTrafficResponse) SetHeaders(v map[string]*string) *SuspendTrafficResponse {
	s.Headers = v
	return s
}

func (s *SuspendTrafficResponse) SetStatusCode(v int32) *SuspendTrafficResponse {
	s.StatusCode = &v
	return s
}

func (s *SuspendTrafficResponse) SetBody(v *SuspendTrafficResponseBody) *SuspendTrafficResponse {
	s.Body = v
	return s
}

func (s *SuspendTrafficResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
