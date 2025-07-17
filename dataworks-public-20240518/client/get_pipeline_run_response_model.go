// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPipelineRunResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetPipelineRunResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetPipelineRunResponse
	GetStatusCode() *int32
	SetBody(v *GetPipelineRunResponseBody) *GetPipelineRunResponse
	GetBody() *GetPipelineRunResponseBody
}

type GetPipelineRunResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetPipelineRunResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetPipelineRunResponse) String() string {
	return dara.Prettify(s)
}

func (s GetPipelineRunResponse) GoString() string {
	return s.String()
}

func (s *GetPipelineRunResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetPipelineRunResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetPipelineRunResponse) GetBody() *GetPipelineRunResponseBody {
	return s.Body
}

func (s *GetPipelineRunResponse) SetHeaders(v map[string]*string) *GetPipelineRunResponse {
	s.Headers = v
	return s
}

func (s *GetPipelineRunResponse) SetStatusCode(v int32) *GetPipelineRunResponse {
	s.StatusCode = &v
	return s
}

func (s *GetPipelineRunResponse) SetBody(v *GetPipelineRunResponseBody) *GetPipelineRunResponse {
	s.Body = v
	return s
}

func (s *GetPipelineRunResponse) Validate() error {
	return dara.Validate(s)
}
