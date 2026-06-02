// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iLaunchFlowResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *LaunchFlowResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *LaunchFlowResponse
	GetStatusCode() *int32
	SetBody(v *LaunchFlowResponseBody) *LaunchFlowResponse
	GetBody() *LaunchFlowResponseBody
}

type LaunchFlowResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *LaunchFlowResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s LaunchFlowResponse) String() string {
	return dara.Prettify(s)
}

func (s LaunchFlowResponse) GoString() string {
	return s.String()
}

func (s *LaunchFlowResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *LaunchFlowResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *LaunchFlowResponse) GetBody() *LaunchFlowResponseBody {
	return s.Body
}

func (s *LaunchFlowResponse) SetHeaders(v map[string]*string) *LaunchFlowResponse {
	s.Headers = v
	return s
}

func (s *LaunchFlowResponse) SetStatusCode(v int32) *LaunchFlowResponse {
	s.StatusCode = &v
	return s
}

func (s *LaunchFlowResponse) SetBody(v *LaunchFlowResponseBody) *LaunchFlowResponse {
	s.Body = v
	return s
}

func (s *LaunchFlowResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
