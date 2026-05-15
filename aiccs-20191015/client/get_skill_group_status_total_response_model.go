// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSkillGroupStatusTotalResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetSkillGroupStatusTotalResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetSkillGroupStatusTotalResponse
	GetStatusCode() *int32
	SetBody(v *GetSkillGroupStatusTotalResponseBody) *GetSkillGroupStatusTotalResponse
	GetBody() *GetSkillGroupStatusTotalResponseBody
}

type GetSkillGroupStatusTotalResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetSkillGroupStatusTotalResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetSkillGroupStatusTotalResponse) String() string {
	return dara.Prettify(s)
}

func (s GetSkillGroupStatusTotalResponse) GoString() string {
	return s.String()
}

func (s *GetSkillGroupStatusTotalResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetSkillGroupStatusTotalResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetSkillGroupStatusTotalResponse) GetBody() *GetSkillGroupStatusTotalResponseBody {
	return s.Body
}

func (s *GetSkillGroupStatusTotalResponse) SetHeaders(v map[string]*string) *GetSkillGroupStatusTotalResponse {
	s.Headers = v
	return s
}

func (s *GetSkillGroupStatusTotalResponse) SetStatusCode(v int32) *GetSkillGroupStatusTotalResponse {
	s.StatusCode = &v
	return s
}

func (s *GetSkillGroupStatusTotalResponse) SetBody(v *GetSkillGroupStatusTotalResponseBody) *GetSkillGroupStatusTotalResponse {
	s.Body = v
	return s
}

func (s *GetSkillGroupStatusTotalResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
