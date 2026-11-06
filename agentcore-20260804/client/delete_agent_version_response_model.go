// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAgentVersionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteAgentVersionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteAgentVersionResponse
	GetStatusCode() *int32
	SetBody(v *DeleteAgentVersionResponseBody) *DeleteAgentVersionResponse
	GetBody() *DeleteAgentVersionResponseBody
}

type DeleteAgentVersionResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteAgentVersionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteAgentVersionResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteAgentVersionResponse) GoString() string {
	return s.String()
}

func (s *DeleteAgentVersionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteAgentVersionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteAgentVersionResponse) GetBody() *DeleteAgentVersionResponseBody {
	return s.Body
}

func (s *DeleteAgentVersionResponse) SetHeaders(v map[string]*string) *DeleteAgentVersionResponse {
	s.Headers = v
	return s
}

func (s *DeleteAgentVersionResponse) SetStatusCode(v int32) *DeleteAgentVersionResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteAgentVersionResponse) SetBody(v *DeleteAgentVersionResponseBody) *DeleteAgentVersionResponse {
	s.Body = v
	return s
}

func (s *DeleteAgentVersionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
