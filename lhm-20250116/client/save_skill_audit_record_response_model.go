// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSaveSkillAuditRecordResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SaveSkillAuditRecordResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SaveSkillAuditRecordResponse
	GetStatusCode() *int32
	SetBody(v *SaveSkillAuditRecordResponseBody) *SaveSkillAuditRecordResponse
	GetBody() *SaveSkillAuditRecordResponseBody
}

type SaveSkillAuditRecordResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SaveSkillAuditRecordResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SaveSkillAuditRecordResponse) String() string {
	return dara.Prettify(s)
}

func (s SaveSkillAuditRecordResponse) GoString() string {
	return s.String()
}

func (s *SaveSkillAuditRecordResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SaveSkillAuditRecordResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SaveSkillAuditRecordResponse) GetBody() *SaveSkillAuditRecordResponseBody {
	return s.Body
}

func (s *SaveSkillAuditRecordResponse) SetHeaders(v map[string]*string) *SaveSkillAuditRecordResponse {
	s.Headers = v
	return s
}

func (s *SaveSkillAuditRecordResponse) SetStatusCode(v int32) *SaveSkillAuditRecordResponse {
	s.StatusCode = &v
	return s
}

func (s *SaveSkillAuditRecordResponse) SetBody(v *SaveSkillAuditRecordResponseBody) *SaveSkillAuditRecordResponse {
	s.Body = v
	return s
}

func (s *SaveSkillAuditRecordResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
