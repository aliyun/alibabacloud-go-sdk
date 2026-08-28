// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDownloadSkillVersionViaOssResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DownloadSkillVersionViaOssResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DownloadSkillVersionViaOssResponse
	GetStatusCode() *int32
	SetBody(v *DownloadSkillVersionViaOssResponseBody) *DownloadSkillVersionViaOssResponse
	GetBody() *DownloadSkillVersionViaOssResponseBody
}

type DownloadSkillVersionViaOssResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DownloadSkillVersionViaOssResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DownloadSkillVersionViaOssResponse) String() string {
	return dara.Prettify(s)
}

func (s DownloadSkillVersionViaOssResponse) GoString() string {
	return s.String()
}

func (s *DownloadSkillVersionViaOssResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DownloadSkillVersionViaOssResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DownloadSkillVersionViaOssResponse) GetBody() *DownloadSkillVersionViaOssResponseBody {
	return s.Body
}

func (s *DownloadSkillVersionViaOssResponse) SetHeaders(v map[string]*string) *DownloadSkillVersionViaOssResponse {
	s.Headers = v
	return s
}

func (s *DownloadSkillVersionViaOssResponse) SetStatusCode(v int32) *DownloadSkillVersionViaOssResponse {
	s.StatusCode = &v
	return s
}

func (s *DownloadSkillVersionViaOssResponse) SetBody(v *DownloadSkillVersionViaOssResponseBody) *DownloadSkillVersionViaOssResponse {
	s.Body = v
	return s
}

func (s *DownloadSkillVersionViaOssResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
