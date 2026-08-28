// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPrecheckSkillUploadViaOssResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *PrecheckSkillUploadViaOssResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *PrecheckSkillUploadViaOssResponse
	GetStatusCode() *int32
	SetBody(v *PrecheckSkillUploadViaOssResponseBody) *PrecheckSkillUploadViaOssResponse
	GetBody() *PrecheckSkillUploadViaOssResponseBody
}

type PrecheckSkillUploadViaOssResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PrecheckSkillUploadViaOssResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PrecheckSkillUploadViaOssResponse) String() string {
	return dara.Prettify(s)
}

func (s PrecheckSkillUploadViaOssResponse) GoString() string {
	return s.String()
}

func (s *PrecheckSkillUploadViaOssResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *PrecheckSkillUploadViaOssResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *PrecheckSkillUploadViaOssResponse) GetBody() *PrecheckSkillUploadViaOssResponseBody {
	return s.Body
}

func (s *PrecheckSkillUploadViaOssResponse) SetHeaders(v map[string]*string) *PrecheckSkillUploadViaOssResponse {
	s.Headers = v
	return s
}

func (s *PrecheckSkillUploadViaOssResponse) SetStatusCode(v int32) *PrecheckSkillUploadViaOssResponse {
	s.StatusCode = &v
	return s
}

func (s *PrecheckSkillUploadViaOssResponse) SetBody(v *PrecheckSkillUploadViaOssResponseBody) *PrecheckSkillUploadViaOssResponse {
	s.Body = v
	return s
}

func (s *PrecheckSkillUploadViaOssResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
