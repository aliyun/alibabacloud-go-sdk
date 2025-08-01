// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAvailableAuditNotesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetAvailableAuditNotesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetAvailableAuditNotesResponse
	GetStatusCode() *int32
	SetBody(v *GetAvailableAuditNotesResponseBody) *GetAvailableAuditNotesResponse
	GetBody() *GetAvailableAuditNotesResponseBody
}

type GetAvailableAuditNotesResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAvailableAuditNotesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAvailableAuditNotesResponse) String() string {
	return dara.Prettify(s)
}

func (s GetAvailableAuditNotesResponse) GoString() string {
	return s.String()
}

func (s *GetAvailableAuditNotesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetAvailableAuditNotesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetAvailableAuditNotesResponse) GetBody() *GetAvailableAuditNotesResponseBody {
	return s.Body
}

func (s *GetAvailableAuditNotesResponse) SetHeaders(v map[string]*string) *GetAvailableAuditNotesResponse {
	s.Headers = v
	return s
}

func (s *GetAvailableAuditNotesResponse) SetStatusCode(v int32) *GetAvailableAuditNotesResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAvailableAuditNotesResponse) SetBody(v *GetAvailableAuditNotesResponseBody) *GetAvailableAuditNotesResponse {
	s.Body = v
	return s
}

func (s *GetAvailableAuditNotesResponse) Validate() error {
	return dara.Validate(s)
}
