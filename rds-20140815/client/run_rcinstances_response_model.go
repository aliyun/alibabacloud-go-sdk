// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRunRCInstancesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RunRCInstancesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RunRCInstancesResponse
	GetStatusCode() *int32
	SetBody(v *RunRCInstancesResponseBody) *RunRCInstancesResponse
	GetBody() *RunRCInstancesResponseBody
}

type RunRCInstancesResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RunRCInstancesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RunRCInstancesResponse) String() string {
	return dara.Prettify(s)
}

func (s RunRCInstancesResponse) GoString() string {
	return s.String()
}

func (s *RunRCInstancesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RunRCInstancesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RunRCInstancesResponse) GetBody() *RunRCInstancesResponseBody {
	return s.Body
}

func (s *RunRCInstancesResponse) SetHeaders(v map[string]*string) *RunRCInstancesResponse {
	s.Headers = v
	return s
}

func (s *RunRCInstancesResponse) SetStatusCode(v int32) *RunRCInstancesResponse {
	s.StatusCode = &v
	return s
}

func (s *RunRCInstancesResponse) SetBody(v *RunRCInstancesResponseBody) *RunRCInstancesResponse {
	s.Body = v
	return s
}

func (s *RunRCInstancesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
