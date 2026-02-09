// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateMcubeHotpatchRollbackTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateMcubeHotpatchRollbackTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateMcubeHotpatchRollbackTaskResponse
	GetStatusCode() *int32
	SetBody(v *CreateMcubeHotpatchRollbackTaskResponseBody) *CreateMcubeHotpatchRollbackTaskResponse
	GetBody() *CreateMcubeHotpatchRollbackTaskResponseBody
}

type CreateMcubeHotpatchRollbackTaskResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateMcubeHotpatchRollbackTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateMcubeHotpatchRollbackTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateMcubeHotpatchRollbackTaskResponse) GoString() string {
	return s.String()
}

func (s *CreateMcubeHotpatchRollbackTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateMcubeHotpatchRollbackTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateMcubeHotpatchRollbackTaskResponse) GetBody() *CreateMcubeHotpatchRollbackTaskResponseBody {
	return s.Body
}

func (s *CreateMcubeHotpatchRollbackTaskResponse) SetHeaders(v map[string]*string) *CreateMcubeHotpatchRollbackTaskResponse {
	s.Headers = v
	return s
}

func (s *CreateMcubeHotpatchRollbackTaskResponse) SetStatusCode(v int32) *CreateMcubeHotpatchRollbackTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateMcubeHotpatchRollbackTaskResponse) SetBody(v *CreateMcubeHotpatchRollbackTaskResponseBody) *CreateMcubeHotpatchRollbackTaskResponse {
	s.Body = v
	return s
}

func (s *CreateMcubeHotpatchRollbackTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
