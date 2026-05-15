// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSkillGroupServiceCapabilityResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetSkillGroupServiceCapabilityResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetSkillGroupServiceCapabilityResponse
	GetStatusCode() *int32
	SetBody(v *GetSkillGroupServiceCapabilityResponseBody) *GetSkillGroupServiceCapabilityResponse
	GetBody() *GetSkillGroupServiceCapabilityResponseBody
}

type GetSkillGroupServiceCapabilityResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetSkillGroupServiceCapabilityResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetSkillGroupServiceCapabilityResponse) String() string {
	return dara.Prettify(s)
}

func (s GetSkillGroupServiceCapabilityResponse) GoString() string {
	return s.String()
}

func (s *GetSkillGroupServiceCapabilityResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetSkillGroupServiceCapabilityResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetSkillGroupServiceCapabilityResponse) GetBody() *GetSkillGroupServiceCapabilityResponseBody {
	return s.Body
}

func (s *GetSkillGroupServiceCapabilityResponse) SetHeaders(v map[string]*string) *GetSkillGroupServiceCapabilityResponse {
	s.Headers = v
	return s
}

func (s *GetSkillGroupServiceCapabilityResponse) SetStatusCode(v int32) *GetSkillGroupServiceCapabilityResponse {
	s.StatusCode = &v
	return s
}

func (s *GetSkillGroupServiceCapabilityResponse) SetBody(v *GetSkillGroupServiceCapabilityResponseBody) *GetSkillGroupServiceCapabilityResponse {
	s.Body = v
	return s
}

func (s *GetSkillGroupServiceCapabilityResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
