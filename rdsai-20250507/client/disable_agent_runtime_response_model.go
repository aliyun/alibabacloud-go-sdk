// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableAgentRuntimeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DisableAgentRuntimeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DisableAgentRuntimeResponse
	GetStatusCode() *int32
	SetBody(v *DisableAgentRuntimeResponseBody) *DisableAgentRuntimeResponse
	GetBody() *DisableAgentRuntimeResponseBody
}

type DisableAgentRuntimeResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DisableAgentRuntimeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DisableAgentRuntimeResponse) String() string {
	return dara.Prettify(s)
}

func (s DisableAgentRuntimeResponse) GoString() string {
	return s.String()
}

func (s *DisableAgentRuntimeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DisableAgentRuntimeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DisableAgentRuntimeResponse) GetBody() *DisableAgentRuntimeResponseBody {
	return s.Body
}

func (s *DisableAgentRuntimeResponse) SetHeaders(v map[string]*string) *DisableAgentRuntimeResponse {
	s.Headers = v
	return s
}

func (s *DisableAgentRuntimeResponse) SetStatusCode(v int32) *DisableAgentRuntimeResponse {
	s.StatusCode = &v
	return s
}

func (s *DisableAgentRuntimeResponse) SetBody(v *DisableAgentRuntimeResponseBody) *DisableAgentRuntimeResponse {
	s.Body = v
	return s
}

func (s *DisableAgentRuntimeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
