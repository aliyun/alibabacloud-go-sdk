// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDesktopAgentRuntimeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListDesktopAgentRuntimeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListDesktopAgentRuntimeResponse
	GetStatusCode() *int32
	SetBody(v *ListDesktopAgentRuntimeResponseBody) *ListDesktopAgentRuntimeResponse
	GetBody() *ListDesktopAgentRuntimeResponseBody
}

type ListDesktopAgentRuntimeResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListDesktopAgentRuntimeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListDesktopAgentRuntimeResponse) String() string {
	return dara.Prettify(s)
}

func (s ListDesktopAgentRuntimeResponse) GoString() string {
	return s.String()
}

func (s *ListDesktopAgentRuntimeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListDesktopAgentRuntimeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListDesktopAgentRuntimeResponse) GetBody() *ListDesktopAgentRuntimeResponseBody {
	return s.Body
}

func (s *ListDesktopAgentRuntimeResponse) SetHeaders(v map[string]*string) *ListDesktopAgentRuntimeResponse {
	s.Headers = v
	return s
}

func (s *ListDesktopAgentRuntimeResponse) SetStatusCode(v int32) *ListDesktopAgentRuntimeResponse {
	s.StatusCode = &v
	return s
}

func (s *ListDesktopAgentRuntimeResponse) SetBody(v *ListDesktopAgentRuntimeResponseBody) *ListDesktopAgentRuntimeResponse {
	s.Body = v
	return s
}

func (s *ListDesktopAgentRuntimeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
