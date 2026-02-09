// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateMcubeHotpatchTaskStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateMcubeHotpatchTaskStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateMcubeHotpatchTaskStatusResponse
	GetStatusCode() *int32
	SetBody(v *UpdateMcubeHotpatchTaskStatusResponseBody) *UpdateMcubeHotpatchTaskStatusResponse
	GetBody() *UpdateMcubeHotpatchTaskStatusResponseBody
}

type UpdateMcubeHotpatchTaskStatusResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateMcubeHotpatchTaskStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateMcubeHotpatchTaskStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateMcubeHotpatchTaskStatusResponse) GoString() string {
	return s.String()
}

func (s *UpdateMcubeHotpatchTaskStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateMcubeHotpatchTaskStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateMcubeHotpatchTaskStatusResponse) GetBody() *UpdateMcubeHotpatchTaskStatusResponseBody {
	return s.Body
}

func (s *UpdateMcubeHotpatchTaskStatusResponse) SetHeaders(v map[string]*string) *UpdateMcubeHotpatchTaskStatusResponse {
	s.Headers = v
	return s
}

func (s *UpdateMcubeHotpatchTaskStatusResponse) SetStatusCode(v int32) *UpdateMcubeHotpatchTaskStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateMcubeHotpatchTaskStatusResponse) SetBody(v *UpdateMcubeHotpatchTaskStatusResponseBody) *UpdateMcubeHotpatchTaskStatusResponse {
	s.Body = v
	return s
}

func (s *UpdateMcubeHotpatchTaskStatusResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
