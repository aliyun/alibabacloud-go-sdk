// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSkillVersionDetailResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetSkillVersionDetailResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetSkillVersionDetailResponse
	GetStatusCode() *int32
	SetBody(v *GetSkillVersionDetailResponseBody) *GetSkillVersionDetailResponse
	GetBody() *GetSkillVersionDetailResponseBody
}

type GetSkillVersionDetailResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetSkillVersionDetailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetSkillVersionDetailResponse) String() string {
	return dara.Prettify(s)
}

func (s GetSkillVersionDetailResponse) GoString() string {
	return s.String()
}

func (s *GetSkillVersionDetailResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetSkillVersionDetailResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetSkillVersionDetailResponse) GetBody() *GetSkillVersionDetailResponseBody {
	return s.Body
}

func (s *GetSkillVersionDetailResponse) SetHeaders(v map[string]*string) *GetSkillVersionDetailResponse {
	s.Headers = v
	return s
}

func (s *GetSkillVersionDetailResponse) SetStatusCode(v int32) *GetSkillVersionDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *GetSkillVersionDetailResponse) SetBody(v *GetSkillVersionDetailResponseBody) *GetSkillVersionDetailResponse {
	s.Body = v
	return s
}

func (s *GetSkillVersionDetailResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
