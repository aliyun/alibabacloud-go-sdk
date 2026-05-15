// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSkillGroupLatitudeStateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetSkillGroupLatitudeStateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetSkillGroupLatitudeStateResponse
	GetStatusCode() *int32
	SetBody(v *GetSkillGroupLatitudeStateResponseBody) *GetSkillGroupLatitudeStateResponse
	GetBody() *GetSkillGroupLatitudeStateResponseBody
}

type GetSkillGroupLatitudeStateResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetSkillGroupLatitudeStateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetSkillGroupLatitudeStateResponse) String() string {
	return dara.Prettify(s)
}

func (s GetSkillGroupLatitudeStateResponse) GoString() string {
	return s.String()
}

func (s *GetSkillGroupLatitudeStateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetSkillGroupLatitudeStateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetSkillGroupLatitudeStateResponse) GetBody() *GetSkillGroupLatitudeStateResponseBody {
	return s.Body
}

func (s *GetSkillGroupLatitudeStateResponse) SetHeaders(v map[string]*string) *GetSkillGroupLatitudeStateResponse {
	s.Headers = v
	return s
}

func (s *GetSkillGroupLatitudeStateResponse) SetStatusCode(v int32) *GetSkillGroupLatitudeStateResponse {
	s.StatusCode = &v
	return s
}

func (s *GetSkillGroupLatitudeStateResponse) SetBody(v *GetSkillGroupLatitudeStateResponseBody) *GetSkillGroupLatitudeStateResponse {
	s.Body = v
	return s
}

func (s *GetSkillGroupLatitudeStateResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
