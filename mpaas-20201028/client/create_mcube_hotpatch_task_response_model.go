// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateMcubeHotpatchTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateMcubeHotpatchTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateMcubeHotpatchTaskResponse
	GetStatusCode() *int32
	SetBody(v *CreateMcubeHotpatchTaskResponseBody) *CreateMcubeHotpatchTaskResponse
	GetBody() *CreateMcubeHotpatchTaskResponseBody
}

type CreateMcubeHotpatchTaskResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateMcubeHotpatchTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateMcubeHotpatchTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateMcubeHotpatchTaskResponse) GoString() string {
	return s.String()
}

func (s *CreateMcubeHotpatchTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateMcubeHotpatchTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateMcubeHotpatchTaskResponse) GetBody() *CreateMcubeHotpatchTaskResponseBody {
	return s.Body
}

func (s *CreateMcubeHotpatchTaskResponse) SetHeaders(v map[string]*string) *CreateMcubeHotpatchTaskResponse {
	s.Headers = v
	return s
}

func (s *CreateMcubeHotpatchTaskResponse) SetStatusCode(v int32) *CreateMcubeHotpatchTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateMcubeHotpatchTaskResponse) SetBody(v *CreateMcubeHotpatchTaskResponseBody) *CreateMcubeHotpatchTaskResponse {
	s.Body = v
	return s
}

func (s *CreateMcubeHotpatchTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
