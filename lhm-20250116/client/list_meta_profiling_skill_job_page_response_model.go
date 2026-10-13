// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMetaProfilingSkillJobPageResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListMetaProfilingSkillJobPageResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListMetaProfilingSkillJobPageResponse
	GetStatusCode() *int32
	SetBody(v *ListMetaProfilingSkillJobPageResponseBody) *ListMetaProfilingSkillJobPageResponse
	GetBody() *ListMetaProfilingSkillJobPageResponseBody
}

type ListMetaProfilingSkillJobPageResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListMetaProfilingSkillJobPageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListMetaProfilingSkillJobPageResponse) String() string {
	return dara.Prettify(s)
}

func (s ListMetaProfilingSkillJobPageResponse) GoString() string {
	return s.String()
}

func (s *ListMetaProfilingSkillJobPageResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListMetaProfilingSkillJobPageResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListMetaProfilingSkillJobPageResponse) GetBody() *ListMetaProfilingSkillJobPageResponseBody {
	return s.Body
}

func (s *ListMetaProfilingSkillJobPageResponse) SetHeaders(v map[string]*string) *ListMetaProfilingSkillJobPageResponse {
	s.Headers = v
	return s
}

func (s *ListMetaProfilingSkillJobPageResponse) SetStatusCode(v int32) *ListMetaProfilingSkillJobPageResponse {
	s.StatusCode = &v
	return s
}

func (s *ListMetaProfilingSkillJobPageResponse) SetBody(v *ListMetaProfilingSkillJobPageResponseBody) *ListMetaProfilingSkillJobPageResponse {
	s.Body = v
	return s
}

func (s *ListMetaProfilingSkillJobPageResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
