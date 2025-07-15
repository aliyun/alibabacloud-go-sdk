// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListRealtimeSkillGroupStatesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListRealtimeSkillGroupStatesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListRealtimeSkillGroupStatesResponse
	GetStatusCode() *int32
	SetBody(v *ListRealtimeSkillGroupStatesResponseBody) *ListRealtimeSkillGroupStatesResponse
	GetBody() *ListRealtimeSkillGroupStatesResponseBody
}

type ListRealtimeSkillGroupStatesResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListRealtimeSkillGroupStatesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListRealtimeSkillGroupStatesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListRealtimeSkillGroupStatesResponse) GoString() string {
	return s.String()
}

func (s *ListRealtimeSkillGroupStatesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListRealtimeSkillGroupStatesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListRealtimeSkillGroupStatesResponse) GetBody() *ListRealtimeSkillGroupStatesResponseBody {
	return s.Body
}

func (s *ListRealtimeSkillGroupStatesResponse) SetHeaders(v map[string]*string) *ListRealtimeSkillGroupStatesResponse {
	s.Headers = v
	return s
}

func (s *ListRealtimeSkillGroupStatesResponse) SetStatusCode(v int32) *ListRealtimeSkillGroupStatesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListRealtimeSkillGroupStatesResponse) SetBody(v *ListRealtimeSkillGroupStatesResponseBody) *ListRealtimeSkillGroupStatesResponse {
	s.Body = v
	return s
}

func (s *ListRealtimeSkillGroupStatesResponse) Validate() error {
	return dara.Validate(s)
}
