// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSkillFileDetectResultResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetSkillFileDetectResultResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetSkillFileDetectResultResponse
	GetStatusCode() *int32
	SetBody(v *GetSkillFileDetectResultResponseBody) *GetSkillFileDetectResultResponse
	GetBody() *GetSkillFileDetectResultResponseBody
}

type GetSkillFileDetectResultResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetSkillFileDetectResultResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetSkillFileDetectResultResponse) String() string {
	return dara.Prettify(s)
}

func (s GetSkillFileDetectResultResponse) GoString() string {
	return s.String()
}

func (s *GetSkillFileDetectResultResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetSkillFileDetectResultResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetSkillFileDetectResultResponse) GetBody() *GetSkillFileDetectResultResponseBody {
	return s.Body
}

func (s *GetSkillFileDetectResultResponse) SetHeaders(v map[string]*string) *GetSkillFileDetectResultResponse {
	s.Headers = v
	return s
}

func (s *GetSkillFileDetectResultResponse) SetStatusCode(v int32) *GetSkillFileDetectResultResponse {
	s.StatusCode = &v
	return s
}

func (s *GetSkillFileDetectResultResponse) SetBody(v *GetSkillFileDetectResultResponseBody) *GetSkillFileDetectResultResponse {
	s.Body = v
	return s
}

func (s *GetSkillFileDetectResultResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
