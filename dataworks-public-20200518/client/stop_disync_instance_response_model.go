// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopDISyncInstanceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *StopDISyncInstanceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *StopDISyncInstanceResponse
	GetStatusCode() *int32
	SetBody(v *StopDISyncInstanceResponseBody) *StopDISyncInstanceResponse
	GetBody() *StopDISyncInstanceResponseBody
}

type StopDISyncInstanceResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StopDISyncInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StopDISyncInstanceResponse) String() string {
	return dara.Prettify(s)
}

func (s StopDISyncInstanceResponse) GoString() string {
	return s.String()
}

func (s *StopDISyncInstanceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *StopDISyncInstanceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *StopDISyncInstanceResponse) GetBody() *StopDISyncInstanceResponseBody {
	return s.Body
}

func (s *StopDISyncInstanceResponse) SetHeaders(v map[string]*string) *StopDISyncInstanceResponse {
	s.Headers = v
	return s
}

func (s *StopDISyncInstanceResponse) SetStatusCode(v int32) *StopDISyncInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *StopDISyncInstanceResponse) SetBody(v *StopDISyncInstanceResponseBody) *StopDISyncInstanceResponse {
	s.Body = v
	return s
}

func (s *StopDISyncInstanceResponse) Validate() error {
	return dara.Validate(s)
}
