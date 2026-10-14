// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAgentConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetAgentConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetAgentConfigResponse
	GetStatusCode() *int32
	SetBody(v *GetAgentConfigResponseBody) *GetAgentConfigResponse
	GetBody() *GetAgentConfigResponseBody
}

type GetAgentConfigResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAgentConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAgentConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s GetAgentConfigResponse) GoString() string {
	return s.String()
}

func (s *GetAgentConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetAgentConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetAgentConfigResponse) GetBody() *GetAgentConfigResponseBody {
	return s.Body
}

func (s *GetAgentConfigResponse) SetHeaders(v map[string]*string) *GetAgentConfigResponse {
	s.Headers = v
	return s
}

func (s *GetAgentConfigResponse) SetStatusCode(v int32) *GetAgentConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAgentConfigResponse) SetBody(v *GetAgentConfigResponseBody) *GetAgentConfigResponse {
	s.Body = v
	return s
}

func (s *GetAgentConfigResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
