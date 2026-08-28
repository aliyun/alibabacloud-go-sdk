// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAgentSpecResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteAgentSpecResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteAgentSpecResponse
	GetStatusCode() *int32
	SetBody(v *DeleteAgentSpecResponseBody) *DeleteAgentSpecResponse
	GetBody() *DeleteAgentSpecResponseBody
}

type DeleteAgentSpecResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteAgentSpecResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteAgentSpecResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteAgentSpecResponse) GoString() string {
	return s.String()
}

func (s *DeleteAgentSpecResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteAgentSpecResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteAgentSpecResponse) GetBody() *DeleteAgentSpecResponseBody {
	return s.Body
}

func (s *DeleteAgentSpecResponse) SetHeaders(v map[string]*string) *DeleteAgentSpecResponse {
	s.Headers = v
	return s
}

func (s *DeleteAgentSpecResponse) SetStatusCode(v int32) *DeleteAgentSpecResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteAgentSpecResponse) SetBody(v *DeleteAgentSpecResponseBody) *DeleteAgentSpecResponse {
	s.Body = v
	return s
}

func (s *DeleteAgentSpecResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
