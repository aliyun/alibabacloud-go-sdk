// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitAuditNoteResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SubmitAuditNoteResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SubmitAuditNoteResponse
	GetStatusCode() *int32
	SetBody(v *SubmitAuditNoteResponseBody) *SubmitAuditNoteResponse
	GetBody() *SubmitAuditNoteResponseBody
}

type SubmitAuditNoteResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SubmitAuditNoteResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SubmitAuditNoteResponse) String() string {
	return dara.Prettify(s)
}

func (s SubmitAuditNoteResponse) GoString() string {
	return s.String()
}

func (s *SubmitAuditNoteResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SubmitAuditNoteResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SubmitAuditNoteResponse) GetBody() *SubmitAuditNoteResponseBody {
	return s.Body
}

func (s *SubmitAuditNoteResponse) SetHeaders(v map[string]*string) *SubmitAuditNoteResponse {
	s.Headers = v
	return s
}

func (s *SubmitAuditNoteResponse) SetStatusCode(v int32) *SubmitAuditNoteResponse {
	s.StatusCode = &v
	return s
}

func (s *SubmitAuditNoteResponse) SetBody(v *SubmitAuditNoteResponseBody) *SubmitAuditNoteResponse {
	s.Body = v
	return s
}

func (s *SubmitAuditNoteResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
