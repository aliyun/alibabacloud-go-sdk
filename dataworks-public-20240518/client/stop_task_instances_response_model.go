// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopTaskInstancesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *StopTaskInstancesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *StopTaskInstancesResponse
	GetStatusCode() *int32
	SetBody(v *StopTaskInstancesResponseBody) *StopTaskInstancesResponse
	GetBody() *StopTaskInstancesResponseBody
}

type StopTaskInstancesResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StopTaskInstancesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StopTaskInstancesResponse) String() string {
	return dara.Prettify(s)
}

func (s StopTaskInstancesResponse) GoString() string {
	return s.String()
}

func (s *StopTaskInstancesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *StopTaskInstancesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *StopTaskInstancesResponse) GetBody() *StopTaskInstancesResponseBody {
	return s.Body
}

func (s *StopTaskInstancesResponse) SetHeaders(v map[string]*string) *StopTaskInstancesResponse {
	s.Headers = v
	return s
}

func (s *StopTaskInstancesResponse) SetStatusCode(v int32) *StopTaskInstancesResponse {
	s.StatusCode = &v
	return s
}

func (s *StopTaskInstancesResponse) SetBody(v *StopTaskInstancesResponseBody) *StopTaskInstancesResponse {
	s.Body = v
	return s
}

func (s *StopTaskInstancesResponse) Validate() error {
	return dara.Validate(s)
}
