// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSkillAddMetaProfilingJobResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SkillAddMetaProfilingJobResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SkillAddMetaProfilingJobResponse
	GetStatusCode() *int32
	SetBody(v *SkillAddMetaProfilingJobResponseBody) *SkillAddMetaProfilingJobResponse
	GetBody() *SkillAddMetaProfilingJobResponseBody
}

type SkillAddMetaProfilingJobResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SkillAddMetaProfilingJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SkillAddMetaProfilingJobResponse) String() string {
	return dara.Prettify(s)
}

func (s SkillAddMetaProfilingJobResponse) GoString() string {
	return s.String()
}

func (s *SkillAddMetaProfilingJobResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SkillAddMetaProfilingJobResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SkillAddMetaProfilingJobResponse) GetBody() *SkillAddMetaProfilingJobResponseBody {
	return s.Body
}

func (s *SkillAddMetaProfilingJobResponse) SetHeaders(v map[string]*string) *SkillAddMetaProfilingJobResponse {
	s.Headers = v
	return s
}

func (s *SkillAddMetaProfilingJobResponse) SetStatusCode(v int32) *SkillAddMetaProfilingJobResponse {
	s.StatusCode = &v
	return s
}

func (s *SkillAddMetaProfilingJobResponse) SetBody(v *SkillAddMetaProfilingJobResponseBody) *SkillAddMetaProfilingJobResponse {
	s.Body = v
	return s
}

func (s *SkillAddMetaProfilingJobResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
