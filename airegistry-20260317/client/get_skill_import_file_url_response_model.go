// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSkillImportFileUrlResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetSkillImportFileUrlResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetSkillImportFileUrlResponse
	GetStatusCode() *int32
	SetBody(v *GetSkillImportFileUrlResponseBody) *GetSkillImportFileUrlResponse
	GetBody() *GetSkillImportFileUrlResponseBody
}

type GetSkillImportFileUrlResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetSkillImportFileUrlResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetSkillImportFileUrlResponse) String() string {
	return dara.Prettify(s)
}

func (s GetSkillImportFileUrlResponse) GoString() string {
	return s.String()
}

func (s *GetSkillImportFileUrlResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetSkillImportFileUrlResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetSkillImportFileUrlResponse) GetBody() *GetSkillImportFileUrlResponseBody {
	return s.Body
}

func (s *GetSkillImportFileUrlResponse) SetHeaders(v map[string]*string) *GetSkillImportFileUrlResponse {
	s.Headers = v
	return s
}

func (s *GetSkillImportFileUrlResponse) SetStatusCode(v int32) *GetSkillImportFileUrlResponse {
	s.StatusCode = &v
	return s
}

func (s *GetSkillImportFileUrlResponse) SetBody(v *GetSkillImportFileUrlResponseBody) *GetSkillImportFileUrlResponse {
	s.Body = v
	return s
}

func (s *GetSkillImportFileUrlResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
