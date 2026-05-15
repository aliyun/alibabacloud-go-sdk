// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSkillGroupServiceStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetSkillGroupServiceStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetSkillGroupServiceStatusResponse
	GetStatusCode() *int32
	SetBody(v *GetSkillGroupServiceStatusResponseBody) *GetSkillGroupServiceStatusResponse
	GetBody() *GetSkillGroupServiceStatusResponseBody
}

type GetSkillGroupServiceStatusResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetSkillGroupServiceStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetSkillGroupServiceStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s GetSkillGroupServiceStatusResponse) GoString() string {
	return s.String()
}

func (s *GetSkillGroupServiceStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetSkillGroupServiceStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetSkillGroupServiceStatusResponse) GetBody() *GetSkillGroupServiceStatusResponseBody {
	return s.Body
}

func (s *GetSkillGroupServiceStatusResponse) SetHeaders(v map[string]*string) *GetSkillGroupServiceStatusResponse {
	s.Headers = v
	return s
}

func (s *GetSkillGroupServiceStatusResponse) SetStatusCode(v int32) *GetSkillGroupServiceStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *GetSkillGroupServiceStatusResponse) SetBody(v *GetSkillGroupServiceStatusResponseBody) *GetSkillGroupServiceStatusResponse {
	s.Body = v
	return s
}

func (s *GetSkillGroupServiceStatusResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
