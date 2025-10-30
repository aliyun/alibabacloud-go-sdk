// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRunClusterInspectResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RunClusterInspectResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RunClusterInspectResponse
	GetStatusCode() *int32
	SetBody(v *RunClusterInspectResponseBody) *RunClusterInspectResponse
	GetBody() *RunClusterInspectResponseBody
}

type RunClusterInspectResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RunClusterInspectResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RunClusterInspectResponse) String() string {
	return dara.Prettify(s)
}

func (s RunClusterInspectResponse) GoString() string {
	return s.String()
}

func (s *RunClusterInspectResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RunClusterInspectResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RunClusterInspectResponse) GetBody() *RunClusterInspectResponseBody {
	return s.Body
}

func (s *RunClusterInspectResponse) SetHeaders(v map[string]*string) *RunClusterInspectResponse {
	s.Headers = v
	return s
}

func (s *RunClusterInspectResponse) SetStatusCode(v int32) *RunClusterInspectResponse {
	s.StatusCode = &v
	return s
}

func (s *RunClusterInspectResponse) SetBody(v *RunClusterInspectResponseBody) *RunClusterInspectResponse {
	s.Body = v
	return s
}

func (s *RunClusterInspectResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
