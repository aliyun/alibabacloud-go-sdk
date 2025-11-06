// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRestartClusterResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RestartClusterResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RestartClusterResponse
	GetStatusCode() *int32
	SetBody(v *RestartClusterResponseBody) *RestartClusterResponse
	GetBody() *RestartClusterResponseBody
}

type RestartClusterResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RestartClusterResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RestartClusterResponse) String() string {
	return dara.Prettify(s)
}

func (s RestartClusterResponse) GoString() string {
	return s.String()
}

func (s *RestartClusterResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RestartClusterResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RestartClusterResponse) GetBody() *RestartClusterResponseBody {
	return s.Body
}

func (s *RestartClusterResponse) SetHeaders(v map[string]*string) *RestartClusterResponse {
	s.Headers = v
	return s
}

func (s *RestartClusterResponse) SetStatusCode(v int32) *RestartClusterResponse {
	s.StatusCode = &v
	return s
}

func (s *RestartClusterResponse) SetBody(v *RestartClusterResponseBody) *RestartClusterResponse {
	s.Body = v
	return s
}

func (s *RestartClusterResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
