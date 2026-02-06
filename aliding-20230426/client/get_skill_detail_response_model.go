// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSkillDetailResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetSkillDetailResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetSkillDetailResponse
	GetStatusCode() *int32
	SetBody(v *GetSkillDetailResponseBody) *GetSkillDetailResponse
	GetBody() *GetSkillDetailResponseBody
}

type GetSkillDetailResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetSkillDetailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetSkillDetailResponse) String() string {
	return dara.Prettify(s)
}

func (s GetSkillDetailResponse) GoString() string {
	return s.String()
}

func (s *GetSkillDetailResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetSkillDetailResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetSkillDetailResponse) GetBody() *GetSkillDetailResponseBody {
	return s.Body
}

func (s *GetSkillDetailResponse) SetHeaders(v map[string]*string) *GetSkillDetailResponse {
	s.Headers = v
	return s
}

func (s *GetSkillDetailResponse) SetStatusCode(v int32) *GetSkillDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *GetSkillDetailResponse) SetBody(v *GetSkillDetailResponseBody) *GetSkillDetailResponse {
	s.Body = v
	return s
}

func (s *GetSkillDetailResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
