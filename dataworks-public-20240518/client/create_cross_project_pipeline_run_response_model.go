// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCrossProjectPipelineRunResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateCrossProjectPipelineRunResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateCrossProjectPipelineRunResponse
	GetStatusCode() *int32
	SetBody(v *CreateCrossProjectPipelineRunResponseBody) *CreateCrossProjectPipelineRunResponse
	GetBody() *CreateCrossProjectPipelineRunResponseBody
}

type CreateCrossProjectPipelineRunResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateCrossProjectPipelineRunResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateCrossProjectPipelineRunResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateCrossProjectPipelineRunResponse) GoString() string {
	return s.String()
}

func (s *CreateCrossProjectPipelineRunResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateCrossProjectPipelineRunResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateCrossProjectPipelineRunResponse) GetBody() *CreateCrossProjectPipelineRunResponseBody {
	return s.Body
}

func (s *CreateCrossProjectPipelineRunResponse) SetHeaders(v map[string]*string) *CreateCrossProjectPipelineRunResponse {
	s.Headers = v
	return s
}

func (s *CreateCrossProjectPipelineRunResponse) SetStatusCode(v int32) *CreateCrossProjectPipelineRunResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateCrossProjectPipelineRunResponse) SetBody(v *CreateCrossProjectPipelineRunResponseBody) *CreateCrossProjectPipelineRunResponse {
	s.Body = v
	return s
}

func (s *CreateCrossProjectPipelineRunResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
