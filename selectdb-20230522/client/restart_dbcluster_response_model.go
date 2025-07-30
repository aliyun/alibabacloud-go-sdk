// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRestartDBClusterResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RestartDBClusterResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RestartDBClusterResponse
	GetStatusCode() *int32
	SetBody(v *RestartDBClusterResponseBody) *RestartDBClusterResponse
	GetBody() *RestartDBClusterResponseBody
}

type RestartDBClusterResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RestartDBClusterResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RestartDBClusterResponse) String() string {
	return dara.Prettify(s)
}

func (s RestartDBClusterResponse) GoString() string {
	return s.String()
}

func (s *RestartDBClusterResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RestartDBClusterResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RestartDBClusterResponse) GetBody() *RestartDBClusterResponseBody {
	return s.Body
}

func (s *RestartDBClusterResponse) SetHeaders(v map[string]*string) *RestartDBClusterResponse {
	s.Headers = v
	return s
}

func (s *RestartDBClusterResponse) SetStatusCode(v int32) *RestartDBClusterResponse {
	s.StatusCode = &v
	return s
}

func (s *RestartDBClusterResponse) SetBody(v *RestartDBClusterResponseBody) *RestartDBClusterResponse {
	s.Body = v
	return s
}

func (s *RestartDBClusterResponse) Validate() error {
	return dara.Validate(s)
}
