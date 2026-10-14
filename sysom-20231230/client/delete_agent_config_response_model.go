// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAgentConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteAgentConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteAgentConfigResponse
	GetStatusCode() *int32
	SetBody(v *DeleteAgentConfigResponseBody) *DeleteAgentConfigResponse
	GetBody() *DeleteAgentConfigResponseBody
}

type DeleteAgentConfigResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteAgentConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteAgentConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteAgentConfigResponse) GoString() string {
	return s.String()
}

func (s *DeleteAgentConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteAgentConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteAgentConfigResponse) GetBody() *DeleteAgentConfigResponseBody {
	return s.Body
}

func (s *DeleteAgentConfigResponse) SetHeaders(v map[string]*string) *DeleteAgentConfigResponse {
	s.Headers = v
	return s
}

func (s *DeleteAgentConfigResponse) SetStatusCode(v int32) *DeleteAgentConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteAgentConfigResponse) SetBody(v *DeleteAgentConfigResponseBody) *DeleteAgentConfigResponse {
	s.Body = v
	return s
}

func (s *DeleteAgentConfigResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
