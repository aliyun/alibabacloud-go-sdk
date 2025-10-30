// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAutoscalingConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateAutoscalingConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateAutoscalingConfigResponse
	GetStatusCode() *int32
	SetBody(v *CreateAutoscalingConfigResponseBody) *CreateAutoscalingConfigResponse
	GetBody() *CreateAutoscalingConfigResponseBody
}

type CreateAutoscalingConfigResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateAutoscalingConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateAutoscalingConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateAutoscalingConfigResponse) GoString() string {
	return s.String()
}

func (s *CreateAutoscalingConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateAutoscalingConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateAutoscalingConfigResponse) GetBody() *CreateAutoscalingConfigResponseBody {
	return s.Body
}

func (s *CreateAutoscalingConfigResponse) SetHeaders(v map[string]*string) *CreateAutoscalingConfigResponse {
	s.Headers = v
	return s
}

func (s *CreateAutoscalingConfigResponse) SetStatusCode(v int32) *CreateAutoscalingConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateAutoscalingConfigResponse) SetBody(v *CreateAutoscalingConfigResponseBody) *CreateAutoscalingConfigResponse {
	s.Body = v
	return s
}

func (s *CreateAutoscalingConfigResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
