// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSkipPipelineJobRunResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SkipPipelineJobRunResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SkipPipelineJobRunResponse
	GetStatusCode() *int32
	SetBody(v *SkipPipelineJobRunResponseBody) *SkipPipelineJobRunResponse
	GetBody() *SkipPipelineJobRunResponseBody
}

type SkipPipelineJobRunResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SkipPipelineJobRunResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SkipPipelineJobRunResponse) String() string {
	return dara.Prettify(s)
}

func (s SkipPipelineJobRunResponse) GoString() string {
	return s.String()
}

func (s *SkipPipelineJobRunResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SkipPipelineJobRunResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SkipPipelineJobRunResponse) GetBody() *SkipPipelineJobRunResponseBody {
	return s.Body
}

func (s *SkipPipelineJobRunResponse) SetHeaders(v map[string]*string) *SkipPipelineJobRunResponse {
	s.Headers = v
	return s
}

func (s *SkipPipelineJobRunResponse) SetStatusCode(v int32) *SkipPipelineJobRunResponse {
	s.StatusCode = &v
	return s
}

func (s *SkipPipelineJobRunResponse) SetBody(v *SkipPipelineJobRunResponseBody) *SkipPipelineJobRunResponse {
	s.Body = v
	return s
}

func (s *SkipPipelineJobRunResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
