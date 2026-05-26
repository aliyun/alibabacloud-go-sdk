// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSkillContentResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetSkillContentResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetSkillContentResponse
	GetStatusCode() *int32
	SetBody(v *GetSkillContentResponseBody) *GetSkillContentResponse
	GetBody() *GetSkillContentResponseBody
}

type GetSkillContentResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetSkillContentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetSkillContentResponse) String() string {
	return dara.Prettify(s)
}

func (s GetSkillContentResponse) GoString() string {
	return s.String()
}

func (s *GetSkillContentResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetSkillContentResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetSkillContentResponse) GetBody() *GetSkillContentResponseBody {
	return s.Body
}

func (s *GetSkillContentResponse) SetHeaders(v map[string]*string) *GetSkillContentResponse {
	s.Headers = v
	return s
}

func (s *GetSkillContentResponse) SetStatusCode(v int32) *GetSkillContentResponse {
	s.StatusCode = &v
	return s
}

func (s *GetSkillContentResponse) SetBody(v *GetSkillContentResponseBody) *GetSkillContentResponse {
	s.Body = v
	return s
}

func (s *GetSkillContentResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
