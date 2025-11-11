// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAuditNoteResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteAuditNoteResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteAuditNoteResponse
	GetStatusCode() *int32
	SetBody(v *DeleteAuditNoteResponseBody) *DeleteAuditNoteResponse
	GetBody() *DeleteAuditNoteResponseBody
}

type DeleteAuditNoteResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteAuditNoteResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteAuditNoteResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteAuditNoteResponse) GoString() string {
	return s.String()
}

func (s *DeleteAuditNoteResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteAuditNoteResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteAuditNoteResponse) GetBody() *DeleteAuditNoteResponseBody {
	return s.Body
}

func (s *DeleteAuditNoteResponse) SetHeaders(v map[string]*string) *DeleteAuditNoteResponse {
	s.Headers = v
	return s
}

func (s *DeleteAuditNoteResponse) SetStatusCode(v int32) *DeleteAuditNoteResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteAuditNoteResponse) SetBody(v *DeleteAuditNoteResponseBody) *DeleteAuditNoteResponse {
	s.Body = v
	return s
}

func (s *DeleteAuditNoteResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
