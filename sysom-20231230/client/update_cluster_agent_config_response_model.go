// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateClusterAgentConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateClusterAgentConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateClusterAgentConfigResponse
	GetStatusCode() *int32
	SetBody(v *UpdateClusterAgentConfigResponseBody) *UpdateClusterAgentConfigResponse
	GetBody() *UpdateClusterAgentConfigResponseBody
}

type UpdateClusterAgentConfigResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateClusterAgentConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateClusterAgentConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateClusterAgentConfigResponse) GoString() string {
	return s.String()
}

func (s *UpdateClusterAgentConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateClusterAgentConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateClusterAgentConfigResponse) GetBody() *UpdateClusterAgentConfigResponseBody {
	return s.Body
}

func (s *UpdateClusterAgentConfigResponse) SetHeaders(v map[string]*string) *UpdateClusterAgentConfigResponse {
	s.Headers = v
	return s
}

func (s *UpdateClusterAgentConfigResponse) SetStatusCode(v int32) *UpdateClusterAgentConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateClusterAgentConfigResponse) SetBody(v *UpdateClusterAgentConfigResponseBody) *UpdateClusterAgentConfigResponse {
	s.Body = v
	return s
}

func (s *UpdateClusterAgentConfigResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
