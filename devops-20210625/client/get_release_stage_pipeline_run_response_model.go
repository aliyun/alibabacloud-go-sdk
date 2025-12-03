// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetReleaseStagePipelineRunResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetReleaseStagePipelineRunResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetReleaseStagePipelineRunResponse
	GetStatusCode() *int32
	SetBody(v *GetReleaseStagePipelineRunResponseBody) *GetReleaseStagePipelineRunResponse
	GetBody() *GetReleaseStagePipelineRunResponseBody
}

type GetReleaseStagePipelineRunResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetReleaseStagePipelineRunResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetReleaseStagePipelineRunResponse) String() string {
	return dara.Prettify(s)
}

func (s GetReleaseStagePipelineRunResponse) GoString() string {
	return s.String()
}

func (s *GetReleaseStagePipelineRunResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetReleaseStagePipelineRunResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetReleaseStagePipelineRunResponse) GetBody() *GetReleaseStagePipelineRunResponseBody {
	return s.Body
}

func (s *GetReleaseStagePipelineRunResponse) SetHeaders(v map[string]*string) *GetReleaseStagePipelineRunResponse {
	s.Headers = v
	return s
}

func (s *GetReleaseStagePipelineRunResponse) SetStatusCode(v int32) *GetReleaseStagePipelineRunResponse {
	s.StatusCode = &v
	return s
}

func (s *GetReleaseStagePipelineRunResponse) SetBody(v *GetReleaseStagePipelineRunResponseBody) *GetReleaseStagePipelineRunResponse {
	s.Body = v
	return s
}

func (s *GetReleaseStagePipelineRunResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
