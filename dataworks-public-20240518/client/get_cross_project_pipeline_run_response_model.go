// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCrossProjectPipelineRunResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetCrossProjectPipelineRunResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetCrossProjectPipelineRunResponse
	GetStatusCode() *int32
	SetBody(v *GetCrossProjectPipelineRunResponseBody) *GetCrossProjectPipelineRunResponse
	GetBody() *GetCrossProjectPipelineRunResponseBody
}

type GetCrossProjectPipelineRunResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetCrossProjectPipelineRunResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetCrossProjectPipelineRunResponse) String() string {
	return dara.Prettify(s)
}

func (s GetCrossProjectPipelineRunResponse) GoString() string {
	return s.String()
}

func (s *GetCrossProjectPipelineRunResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetCrossProjectPipelineRunResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetCrossProjectPipelineRunResponse) GetBody() *GetCrossProjectPipelineRunResponseBody {
	return s.Body
}

func (s *GetCrossProjectPipelineRunResponse) SetHeaders(v map[string]*string) *GetCrossProjectPipelineRunResponse {
	s.Headers = v
	return s
}

func (s *GetCrossProjectPipelineRunResponse) SetStatusCode(v int32) *GetCrossProjectPipelineRunResponse {
	s.StatusCode = &v
	return s
}

func (s *GetCrossProjectPipelineRunResponse) SetBody(v *GetCrossProjectPipelineRunResponseBody) *GetCrossProjectPipelineRunResponse {
	s.Body = v
	return s
}

func (s *GetCrossProjectPipelineRunResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
