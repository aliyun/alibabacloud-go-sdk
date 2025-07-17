// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSuspendTaskFlowInstanceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SuspendTaskFlowInstanceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SuspendTaskFlowInstanceResponse
	GetStatusCode() *int32
	SetBody(v *SuspendTaskFlowInstanceResponseBody) *SuspendTaskFlowInstanceResponse
	GetBody() *SuspendTaskFlowInstanceResponseBody
}

type SuspendTaskFlowInstanceResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SuspendTaskFlowInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SuspendTaskFlowInstanceResponse) String() string {
	return dara.Prettify(s)
}

func (s SuspendTaskFlowInstanceResponse) GoString() string {
	return s.String()
}

func (s *SuspendTaskFlowInstanceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SuspendTaskFlowInstanceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SuspendTaskFlowInstanceResponse) GetBody() *SuspendTaskFlowInstanceResponseBody {
	return s.Body
}

func (s *SuspendTaskFlowInstanceResponse) SetHeaders(v map[string]*string) *SuspendTaskFlowInstanceResponse {
	s.Headers = v
	return s
}

func (s *SuspendTaskFlowInstanceResponse) SetStatusCode(v int32) *SuspendTaskFlowInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *SuspendTaskFlowInstanceResponse) SetBody(v *SuspendTaskFlowInstanceResponseBody) *SuspendTaskFlowInstanceResponse {
	s.Body = v
	return s
}

func (s *SuspendTaskFlowInstanceResponse) Validate() error {
	return dara.Validate(s)
}
