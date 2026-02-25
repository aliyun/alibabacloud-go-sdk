// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRestartNodesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RestartNodesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RestartNodesResponse
	GetStatusCode() *int32
	SetBody(v *RestartNodesResponseBody) *RestartNodesResponse
	GetBody() *RestartNodesResponseBody
}

type RestartNodesResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RestartNodesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RestartNodesResponse) String() string {
	return dara.Prettify(s)
}

func (s RestartNodesResponse) GoString() string {
	return s.String()
}

func (s *RestartNodesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RestartNodesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RestartNodesResponse) GetBody() *RestartNodesResponseBody {
	return s.Body
}

func (s *RestartNodesResponse) SetHeaders(v map[string]*string) *RestartNodesResponse {
	s.Headers = v
	return s
}

func (s *RestartNodesResponse) SetStatusCode(v int32) *RestartNodesResponse {
	s.StatusCode = &v
	return s
}

func (s *RestartNodesResponse) SetBody(v *RestartNodesResponseBody) *RestartNodesResponse {
	s.Body = v
	return s
}

func (s *RestartNodesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
