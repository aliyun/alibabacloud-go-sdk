// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iLogPipelineJobRunResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *LogPipelineJobRunResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *LogPipelineJobRunResponse
	GetStatusCode() *int32
	SetBody(v *LogPipelineJobRunResponseBody) *LogPipelineJobRunResponse
	GetBody() *LogPipelineJobRunResponseBody
}

type LogPipelineJobRunResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *LogPipelineJobRunResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s LogPipelineJobRunResponse) String() string {
	return dara.Prettify(s)
}

func (s LogPipelineJobRunResponse) GoString() string {
	return s.String()
}

func (s *LogPipelineJobRunResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *LogPipelineJobRunResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *LogPipelineJobRunResponse) GetBody() *LogPipelineJobRunResponseBody {
	return s.Body
}

func (s *LogPipelineJobRunResponse) SetHeaders(v map[string]*string) *LogPipelineJobRunResponse {
	s.Headers = v
	return s
}

func (s *LogPipelineJobRunResponse) SetStatusCode(v int32) *LogPipelineJobRunResponse {
	s.StatusCode = &v
	return s
}

func (s *LogPipelineJobRunResponse) SetBody(v *LogPipelineJobRunResponseBody) *LogPipelineJobRunResponse {
	s.Body = v
	return s
}

func (s *LogPipelineJobRunResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
