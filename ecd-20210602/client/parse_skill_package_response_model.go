// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iParseSkillPackageResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ParseSkillPackageResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ParseSkillPackageResponse
	GetStatusCode() *int32
	SetBody(v *ParseSkillPackageResponseBody) *ParseSkillPackageResponse
	GetBody() *ParseSkillPackageResponseBody
}

type ParseSkillPackageResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ParseSkillPackageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ParseSkillPackageResponse) String() string {
	return dara.Prettify(s)
}

func (s ParseSkillPackageResponse) GoString() string {
	return s.String()
}

func (s *ParseSkillPackageResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ParseSkillPackageResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ParseSkillPackageResponse) GetBody() *ParseSkillPackageResponseBody {
	return s.Body
}

func (s *ParseSkillPackageResponse) SetHeaders(v map[string]*string) *ParseSkillPackageResponse {
	s.Headers = v
	return s
}

func (s *ParseSkillPackageResponse) SetStatusCode(v int32) *ParseSkillPackageResponse {
	s.StatusCode = &v
	return s
}

func (s *ParseSkillPackageResponse) SetBody(v *ParseSkillPackageResponseBody) *ParseSkillPackageResponse {
	s.Body = v
	return s
}

func (s *ParseSkillPackageResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
