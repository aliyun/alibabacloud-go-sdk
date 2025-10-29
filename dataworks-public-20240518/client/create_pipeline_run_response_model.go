// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePipelineRunResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreatePipelineRunResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreatePipelineRunResponse
	GetStatusCode() *int32
	SetBody(v *CreatePipelineRunResponseBody) *CreatePipelineRunResponse
	GetBody() *CreatePipelineRunResponseBody
}

type CreatePipelineRunResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreatePipelineRunResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreatePipelineRunResponse) String() string {
	return dara.Prettify(s)
}

func (s CreatePipelineRunResponse) GoString() string {
	return s.String()
}

func (s *CreatePipelineRunResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreatePipelineRunResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreatePipelineRunResponse) GetBody() *CreatePipelineRunResponseBody {
	return s.Body
}

func (s *CreatePipelineRunResponse) SetHeaders(v map[string]*string) *CreatePipelineRunResponse {
	s.Headers = v
	return s
}

func (s *CreatePipelineRunResponse) SetStatusCode(v int32) *CreatePipelineRunResponse {
	s.StatusCode = &v
	return s
}

func (s *CreatePipelineRunResponse) SetBody(v *CreatePipelineRunResponseBody) *CreatePipelineRunResponse {
	s.Body = v
	return s
}

func (s *CreatePipelineRunResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
