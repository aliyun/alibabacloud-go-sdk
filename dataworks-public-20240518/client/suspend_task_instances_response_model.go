// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSuspendTaskInstancesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SuspendTaskInstancesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SuspendTaskInstancesResponse
	GetStatusCode() *int32
	SetBody(v *SuspendTaskInstancesResponseBody) *SuspendTaskInstancesResponse
	GetBody() *SuspendTaskInstancesResponseBody
}

type SuspendTaskInstancesResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SuspendTaskInstancesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SuspendTaskInstancesResponse) String() string {
	return dara.Prettify(s)
}

func (s SuspendTaskInstancesResponse) GoString() string {
	return s.String()
}

func (s *SuspendTaskInstancesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SuspendTaskInstancesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SuspendTaskInstancesResponse) GetBody() *SuspendTaskInstancesResponseBody {
	return s.Body
}

func (s *SuspendTaskInstancesResponse) SetHeaders(v map[string]*string) *SuspendTaskInstancesResponse {
	s.Headers = v
	return s
}

func (s *SuspendTaskInstancesResponse) SetStatusCode(v int32) *SuspendTaskInstancesResponse {
	s.StatusCode = &v
	return s
}

func (s *SuspendTaskInstancesResponse) SetBody(v *SuspendTaskInstancesResponseBody) *SuspendTaskInstancesResponse {
	s.Body = v
	return s
}

func (s *SuspendTaskInstancesResponse) Validate() error {
	return dara.Validate(s)
}
