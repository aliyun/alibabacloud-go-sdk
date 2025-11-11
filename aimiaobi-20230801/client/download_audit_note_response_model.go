// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDownloadAuditNoteResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DownloadAuditNoteResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DownloadAuditNoteResponse
	GetStatusCode() *int32
	SetBody(v *DownloadAuditNoteResponseBody) *DownloadAuditNoteResponse
	GetBody() *DownloadAuditNoteResponseBody
}

type DownloadAuditNoteResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DownloadAuditNoteResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DownloadAuditNoteResponse) String() string {
	return dara.Prettify(s)
}

func (s DownloadAuditNoteResponse) GoString() string {
	return s.String()
}

func (s *DownloadAuditNoteResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DownloadAuditNoteResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DownloadAuditNoteResponse) GetBody() *DownloadAuditNoteResponseBody {
	return s.Body
}

func (s *DownloadAuditNoteResponse) SetHeaders(v map[string]*string) *DownloadAuditNoteResponse {
	s.Headers = v
	return s
}

func (s *DownloadAuditNoteResponse) SetStatusCode(v int32) *DownloadAuditNoteResponse {
	s.StatusCode = &v
	return s
}

func (s *DownloadAuditNoteResponse) SetBody(v *DownloadAuditNoteResponseBody) *DownloadAuditNoteResponse {
	s.Body = v
	return s
}

func (s *DownloadAuditNoteResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
