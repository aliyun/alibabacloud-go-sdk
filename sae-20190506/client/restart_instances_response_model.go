// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRestartInstancesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RestartInstancesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RestartInstancesResponse
	GetStatusCode() *int32
	SetBody(v *RestartInstancesResponseBody) *RestartInstancesResponse
	GetBody() *RestartInstancesResponseBody
}

type RestartInstancesResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RestartInstancesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RestartInstancesResponse) String() string {
	return dara.Prettify(s)
}

func (s RestartInstancesResponse) GoString() string {
	return s.String()
}

func (s *RestartInstancesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RestartInstancesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RestartInstancesResponse) GetBody() *RestartInstancesResponseBody {
	return s.Body
}

func (s *RestartInstancesResponse) SetHeaders(v map[string]*string) *RestartInstancesResponse {
	s.Headers = v
	return s
}

func (s *RestartInstancesResponse) SetStatusCode(v int32) *RestartInstancesResponse {
	s.StatusCode = &v
	return s
}

func (s *RestartInstancesResponse) SetBody(v *RestartInstancesResponseBody) *RestartInstancesResponse {
	s.Body = v
	return s
}

func (s *RestartInstancesResponse) Validate() error {
	return dara.Validate(s)
}
