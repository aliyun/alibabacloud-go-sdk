// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopTrafficControlTargetResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *StopTrafficControlTargetResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *StopTrafficControlTargetResponse
	GetStatusCode() *int32
	SetBody(v *StopTrafficControlTargetResponseBody) *StopTrafficControlTargetResponse
	GetBody() *StopTrafficControlTargetResponseBody
}

type StopTrafficControlTargetResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StopTrafficControlTargetResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StopTrafficControlTargetResponse) String() string {
	return dara.Prettify(s)
}

func (s StopTrafficControlTargetResponse) GoString() string {
	return s.String()
}

func (s *StopTrafficControlTargetResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *StopTrafficControlTargetResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *StopTrafficControlTargetResponse) GetBody() *StopTrafficControlTargetResponseBody {
	return s.Body
}

func (s *StopTrafficControlTargetResponse) SetHeaders(v map[string]*string) *StopTrafficControlTargetResponse {
	s.Headers = v
	return s
}

func (s *StopTrafficControlTargetResponse) SetStatusCode(v int32) *StopTrafficControlTargetResponse {
	s.StatusCode = &v
	return s
}

func (s *StopTrafficControlTargetResponse) SetBody(v *StopTrafficControlTargetResponseBody) *StopTrafficControlTargetResponse {
	s.Body = v
	return s
}

func (s *StopTrafficControlTargetResponse) Validate() error {
	return dara.Validate(s)
}
