// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSaveWorkflowGlobalEnvResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SaveWorkflowGlobalEnvResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SaveWorkflowGlobalEnvResponse
	GetStatusCode() *int32
	SetBody(v *SaveWorkflowGlobalEnvResponseBody) *SaveWorkflowGlobalEnvResponse
	GetBody() *SaveWorkflowGlobalEnvResponseBody
}

type SaveWorkflowGlobalEnvResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SaveWorkflowGlobalEnvResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SaveWorkflowGlobalEnvResponse) String() string {
	return dara.Prettify(s)
}

func (s SaveWorkflowGlobalEnvResponse) GoString() string {
	return s.String()
}

func (s *SaveWorkflowGlobalEnvResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SaveWorkflowGlobalEnvResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SaveWorkflowGlobalEnvResponse) GetBody() *SaveWorkflowGlobalEnvResponseBody {
	return s.Body
}

func (s *SaveWorkflowGlobalEnvResponse) SetHeaders(v map[string]*string) *SaveWorkflowGlobalEnvResponse {
	s.Headers = v
	return s
}

func (s *SaveWorkflowGlobalEnvResponse) SetStatusCode(v int32) *SaveWorkflowGlobalEnvResponse {
	s.StatusCode = &v
	return s
}

func (s *SaveWorkflowGlobalEnvResponse) SetBody(v *SaveWorkflowGlobalEnvResponseBody) *SaveWorkflowGlobalEnvResponse {
	s.Body = v
	return s
}

func (s *SaveWorkflowGlobalEnvResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
