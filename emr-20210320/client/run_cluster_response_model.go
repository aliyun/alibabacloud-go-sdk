// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRunClusterResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RunClusterResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RunClusterResponse
	GetStatusCode() *int32
	SetBody(v *RunClusterResponseBody) *RunClusterResponse
	GetBody() *RunClusterResponseBody
}

type RunClusterResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RunClusterResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RunClusterResponse) String() string {
	return dara.Prettify(s)
}

func (s RunClusterResponse) GoString() string {
	return s.String()
}

func (s *RunClusterResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RunClusterResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RunClusterResponse) GetBody() *RunClusterResponseBody {
	return s.Body
}

func (s *RunClusterResponse) SetHeaders(v map[string]*string) *RunClusterResponse {
	s.Headers = v
	return s
}

func (s *RunClusterResponse) SetStatusCode(v int32) *RunClusterResponse {
	s.StatusCode = &v
	return s
}

func (s *RunClusterResponse) SetBody(v *RunClusterResponseBody) *RunClusterResponse {
	s.Body = v
	return s
}

func (s *RunClusterResponse) Validate() error {
	return dara.Validate(s)
}
