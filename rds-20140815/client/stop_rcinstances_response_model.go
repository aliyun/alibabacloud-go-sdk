// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopRCInstancesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *StopRCInstancesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *StopRCInstancesResponse
	GetStatusCode() *int32
	SetBody(v *StopRCInstancesResponseBody) *StopRCInstancesResponse
	GetBody() *StopRCInstancesResponseBody
}

type StopRCInstancesResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StopRCInstancesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StopRCInstancesResponse) String() string {
	return dara.Prettify(s)
}

func (s StopRCInstancesResponse) GoString() string {
	return s.String()
}

func (s *StopRCInstancesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *StopRCInstancesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *StopRCInstancesResponse) GetBody() *StopRCInstancesResponseBody {
	return s.Body
}

func (s *StopRCInstancesResponse) SetHeaders(v map[string]*string) *StopRCInstancesResponse {
	s.Headers = v
	return s
}

func (s *StopRCInstancesResponse) SetStatusCode(v int32) *StopRCInstancesResponse {
	s.StatusCode = &v
	return s
}

func (s *StopRCInstancesResponse) SetBody(v *StopRCInstancesResponseBody) *StopRCInstancesResponse {
	s.Body = v
	return s
}

func (s *StopRCInstancesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
