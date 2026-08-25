// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopServerIdeInstanceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *StopServerIdeInstanceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *StopServerIdeInstanceResponse
	GetStatusCode() *int32
	SetBody(v *StopServerIdeInstanceResponseBody) *StopServerIdeInstanceResponse
	GetBody() *StopServerIdeInstanceResponseBody
}

type StopServerIdeInstanceResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StopServerIdeInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StopServerIdeInstanceResponse) String() string {
	return dara.Prettify(s)
}

func (s StopServerIdeInstanceResponse) GoString() string {
	return s.String()
}

func (s *StopServerIdeInstanceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *StopServerIdeInstanceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *StopServerIdeInstanceResponse) GetBody() *StopServerIdeInstanceResponseBody {
	return s.Body
}

func (s *StopServerIdeInstanceResponse) SetHeaders(v map[string]*string) *StopServerIdeInstanceResponse {
	s.Headers = v
	return s
}

func (s *StopServerIdeInstanceResponse) SetStatusCode(v int32) *StopServerIdeInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *StopServerIdeInstanceResponse) SetBody(v *StopServerIdeInstanceResponseBody) *StopServerIdeInstanceResponse {
	s.Body = v
	return s
}

func (s *StopServerIdeInstanceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
