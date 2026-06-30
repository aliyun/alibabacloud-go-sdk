// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdatePolarClawAgentToolsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdatePolarClawAgentToolsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdatePolarClawAgentToolsResponse
	GetStatusCode() *int32
	SetBody(v *UpdatePolarClawAgentToolsResponseBody) *UpdatePolarClawAgentToolsResponse
	GetBody() *UpdatePolarClawAgentToolsResponseBody
}

type UpdatePolarClawAgentToolsResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdatePolarClawAgentToolsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdatePolarClawAgentToolsResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdatePolarClawAgentToolsResponse) GoString() string {
	return s.String()
}

func (s *UpdatePolarClawAgentToolsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdatePolarClawAgentToolsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdatePolarClawAgentToolsResponse) GetBody() *UpdatePolarClawAgentToolsResponseBody {
	return s.Body
}

func (s *UpdatePolarClawAgentToolsResponse) SetHeaders(v map[string]*string) *UpdatePolarClawAgentToolsResponse {
	s.Headers = v
	return s
}

func (s *UpdatePolarClawAgentToolsResponse) SetStatusCode(v int32) *UpdatePolarClawAgentToolsResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdatePolarClawAgentToolsResponse) SetBody(v *UpdatePolarClawAgentToolsResponseBody) *UpdatePolarClawAgentToolsResponse {
	s.Body = v
	return s
}

func (s *UpdatePolarClawAgentToolsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
