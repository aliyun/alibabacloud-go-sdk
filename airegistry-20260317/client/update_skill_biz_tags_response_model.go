// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateSkillBizTagsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateSkillBizTagsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateSkillBizTagsResponse
	GetStatusCode() *int32
	SetBody(v *UpdateSkillBizTagsResponseBody) *UpdateSkillBizTagsResponse
	GetBody() *UpdateSkillBizTagsResponseBody
}

type UpdateSkillBizTagsResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateSkillBizTagsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateSkillBizTagsResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateSkillBizTagsResponse) GoString() string {
	return s.String()
}

func (s *UpdateSkillBizTagsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateSkillBizTagsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateSkillBizTagsResponse) GetBody() *UpdateSkillBizTagsResponseBody {
	return s.Body
}

func (s *UpdateSkillBizTagsResponse) SetHeaders(v map[string]*string) *UpdateSkillBizTagsResponse {
	s.Headers = v
	return s
}

func (s *UpdateSkillBizTagsResponse) SetStatusCode(v int32) *UpdateSkillBizTagsResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateSkillBizTagsResponse) SetBody(v *UpdateSkillBizTagsResponseBody) *UpdateSkillBizTagsResponse {
	s.Body = v
	return s
}

func (s *UpdateSkillBizTagsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
