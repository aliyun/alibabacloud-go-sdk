// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUploadSkillViaOssResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UploadSkillViaOssResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UploadSkillViaOssResponse
	GetStatusCode() *int32
	SetBody(v *UploadSkillViaOssResponseBody) *UploadSkillViaOssResponse
	GetBody() *UploadSkillViaOssResponseBody
}

type UploadSkillViaOssResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UploadSkillViaOssResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UploadSkillViaOssResponse) String() string {
	return dara.Prettify(s)
}

func (s UploadSkillViaOssResponse) GoString() string {
	return s.String()
}

func (s *UploadSkillViaOssResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UploadSkillViaOssResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UploadSkillViaOssResponse) GetBody() *UploadSkillViaOssResponseBody {
	return s.Body
}

func (s *UploadSkillViaOssResponse) SetHeaders(v map[string]*string) *UploadSkillViaOssResponse {
	s.Headers = v
	return s
}

func (s *UploadSkillViaOssResponse) SetStatusCode(v int32) *UploadSkillViaOssResponse {
	s.StatusCode = &v
	return s
}

func (s *UploadSkillViaOssResponse) SetBody(v *UploadSkillViaOssResponseBody) *UploadSkillViaOssResponse {
	s.Body = v
	return s
}

func (s *UploadSkillViaOssResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
