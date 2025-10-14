// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteMonitorGroupInstancesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteMonitorGroupInstancesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteMonitorGroupInstancesResponse
	GetStatusCode() *int32
	SetBody(v *DeleteMonitorGroupInstancesResponseBody) *DeleteMonitorGroupInstancesResponse
	GetBody() *DeleteMonitorGroupInstancesResponseBody
}

type DeleteMonitorGroupInstancesResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteMonitorGroupInstancesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteMonitorGroupInstancesResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteMonitorGroupInstancesResponse) GoString() string {
	return s.String()
}

func (s *DeleteMonitorGroupInstancesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteMonitorGroupInstancesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteMonitorGroupInstancesResponse) GetBody() *DeleteMonitorGroupInstancesResponseBody {
	return s.Body
}

func (s *DeleteMonitorGroupInstancesResponse) SetHeaders(v map[string]*string) *DeleteMonitorGroupInstancesResponse {
	s.Headers = v
	return s
}

func (s *DeleteMonitorGroupInstancesResponse) SetStatusCode(v int32) *DeleteMonitorGroupInstancesResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteMonitorGroupInstancesResponse) SetBody(v *DeleteMonitorGroupInstancesResponseBody) *DeleteMonitorGroupInstancesResponse {
	s.Body = v
	return s
}

func (s *DeleteMonitorGroupInstancesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
