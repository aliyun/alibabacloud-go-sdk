// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAgentSpecVersionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteAgentSpecVersionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteAgentSpecVersionResponse
	GetStatusCode() *int32
	SetBody(v *DeleteAgentSpecVersionResponseBody) *DeleteAgentSpecVersionResponse
	GetBody() *DeleteAgentSpecVersionResponseBody
}

type DeleteAgentSpecVersionResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteAgentSpecVersionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteAgentSpecVersionResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteAgentSpecVersionResponse) GoString() string {
	return s.String()
}

func (s *DeleteAgentSpecVersionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteAgentSpecVersionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteAgentSpecVersionResponse) GetBody() *DeleteAgentSpecVersionResponseBody {
	return s.Body
}

func (s *DeleteAgentSpecVersionResponse) SetHeaders(v map[string]*string) *DeleteAgentSpecVersionResponse {
	s.Headers = v
	return s
}

func (s *DeleteAgentSpecVersionResponse) SetStatusCode(v int32) *DeleteAgentSpecVersionResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteAgentSpecVersionResponse) SetBody(v *DeleteAgentSpecVersionResponseBody) *DeleteAgentSpecVersionResponse {
	s.Body = v
	return s
}

func (s *DeleteAgentSpecVersionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
