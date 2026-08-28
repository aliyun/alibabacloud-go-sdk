// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchUploadSkillsViaOssResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *BatchUploadSkillsViaOssResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *BatchUploadSkillsViaOssResponse
	GetStatusCode() *int32
	SetBody(v *BatchUploadSkillsViaOssResponseBody) *BatchUploadSkillsViaOssResponse
	GetBody() *BatchUploadSkillsViaOssResponseBody
}

type BatchUploadSkillsViaOssResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *BatchUploadSkillsViaOssResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s BatchUploadSkillsViaOssResponse) String() string {
	return dara.Prettify(s)
}

func (s BatchUploadSkillsViaOssResponse) GoString() string {
	return s.String()
}

func (s *BatchUploadSkillsViaOssResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *BatchUploadSkillsViaOssResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *BatchUploadSkillsViaOssResponse) GetBody() *BatchUploadSkillsViaOssResponseBody {
	return s.Body
}

func (s *BatchUploadSkillsViaOssResponse) SetHeaders(v map[string]*string) *BatchUploadSkillsViaOssResponse {
	s.Headers = v
	return s
}

func (s *BatchUploadSkillsViaOssResponse) SetStatusCode(v int32) *BatchUploadSkillsViaOssResponse {
	s.StatusCode = &v
	return s
}

func (s *BatchUploadSkillsViaOssResponse) SetBody(v *BatchUploadSkillsViaOssResponseBody) *BatchUploadSkillsViaOssResponse {
	s.Body = v
	return s
}

func (s *BatchUploadSkillsViaOssResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
