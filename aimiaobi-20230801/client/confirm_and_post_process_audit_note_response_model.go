// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iConfirmAndPostProcessAuditNoteResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ConfirmAndPostProcessAuditNoteResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ConfirmAndPostProcessAuditNoteResponse
	GetStatusCode() *int32
	SetBody(v *ConfirmAndPostProcessAuditNoteResponseBody) *ConfirmAndPostProcessAuditNoteResponse
	GetBody() *ConfirmAndPostProcessAuditNoteResponseBody
}

type ConfirmAndPostProcessAuditNoteResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ConfirmAndPostProcessAuditNoteResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ConfirmAndPostProcessAuditNoteResponse) String() string {
	return dara.Prettify(s)
}

func (s ConfirmAndPostProcessAuditNoteResponse) GoString() string {
	return s.String()
}

func (s *ConfirmAndPostProcessAuditNoteResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ConfirmAndPostProcessAuditNoteResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ConfirmAndPostProcessAuditNoteResponse) GetBody() *ConfirmAndPostProcessAuditNoteResponseBody {
	return s.Body
}

func (s *ConfirmAndPostProcessAuditNoteResponse) SetHeaders(v map[string]*string) *ConfirmAndPostProcessAuditNoteResponse {
	s.Headers = v
	return s
}

func (s *ConfirmAndPostProcessAuditNoteResponse) SetStatusCode(v int32) *ConfirmAndPostProcessAuditNoteResponse {
	s.StatusCode = &v
	return s
}

func (s *ConfirmAndPostProcessAuditNoteResponse) SetBody(v *ConfirmAndPostProcessAuditNoteResponseBody) *ConfirmAndPostProcessAuditNoteResponse {
	s.Body = v
	return s
}

func (s *ConfirmAndPostProcessAuditNoteResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
