// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopTaskFlowInstanceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *StopTaskFlowInstanceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *StopTaskFlowInstanceResponse
	GetStatusCode() *int32
	SetBody(v *StopTaskFlowInstanceResponseBody) *StopTaskFlowInstanceResponse
	GetBody() *StopTaskFlowInstanceResponseBody
}

type StopTaskFlowInstanceResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StopTaskFlowInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StopTaskFlowInstanceResponse) String() string {
	return dara.Prettify(s)
}

func (s StopTaskFlowInstanceResponse) GoString() string {
	return s.String()
}

func (s *StopTaskFlowInstanceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *StopTaskFlowInstanceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *StopTaskFlowInstanceResponse) GetBody() *StopTaskFlowInstanceResponseBody {
	return s.Body
}

func (s *StopTaskFlowInstanceResponse) SetHeaders(v map[string]*string) *StopTaskFlowInstanceResponse {
	s.Headers = v
	return s
}

func (s *StopTaskFlowInstanceResponse) SetStatusCode(v int32) *StopTaskFlowInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *StopTaskFlowInstanceResponse) SetBody(v *StopTaskFlowInstanceResponseBody) *StopTaskFlowInstanceResponse {
	s.Body = v
	return s
}

func (s *StopTaskFlowInstanceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
