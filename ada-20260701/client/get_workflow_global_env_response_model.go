// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetWorkflowGlobalEnvResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetWorkflowGlobalEnvResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetWorkflowGlobalEnvResponse
	GetStatusCode() *int32
	SetBody(v *GetWorkflowGlobalEnvResponseBody) *GetWorkflowGlobalEnvResponse
	GetBody() *GetWorkflowGlobalEnvResponseBody
}

type GetWorkflowGlobalEnvResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetWorkflowGlobalEnvResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetWorkflowGlobalEnvResponse) String() string {
	return dara.Prettify(s)
}

func (s GetWorkflowGlobalEnvResponse) GoString() string {
	return s.String()
}

func (s *GetWorkflowGlobalEnvResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetWorkflowGlobalEnvResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetWorkflowGlobalEnvResponse) GetBody() *GetWorkflowGlobalEnvResponseBody {
	return s.Body
}

func (s *GetWorkflowGlobalEnvResponse) SetHeaders(v map[string]*string) *GetWorkflowGlobalEnvResponse {
	s.Headers = v
	return s
}

func (s *GetWorkflowGlobalEnvResponse) SetStatusCode(v int32) *GetWorkflowGlobalEnvResponse {
	s.StatusCode = &v
	return s
}

func (s *GetWorkflowGlobalEnvResponse) SetBody(v *GetWorkflowGlobalEnvResponseBody) *GetWorkflowGlobalEnvResponse {
	s.Body = v
	return s
}

func (s *GetWorkflowGlobalEnvResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
