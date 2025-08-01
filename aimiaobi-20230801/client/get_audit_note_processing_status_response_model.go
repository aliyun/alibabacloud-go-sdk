// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAuditNoteProcessingStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetAuditNoteProcessingStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetAuditNoteProcessingStatusResponse
	GetStatusCode() *int32
	SetBody(v *GetAuditNoteProcessingStatusResponseBody) *GetAuditNoteProcessingStatusResponse
	GetBody() *GetAuditNoteProcessingStatusResponseBody
}

type GetAuditNoteProcessingStatusResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAuditNoteProcessingStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAuditNoteProcessingStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s GetAuditNoteProcessingStatusResponse) GoString() string {
	return s.String()
}

func (s *GetAuditNoteProcessingStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetAuditNoteProcessingStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetAuditNoteProcessingStatusResponse) GetBody() *GetAuditNoteProcessingStatusResponseBody {
	return s.Body
}

func (s *GetAuditNoteProcessingStatusResponse) SetHeaders(v map[string]*string) *GetAuditNoteProcessingStatusResponse {
	s.Headers = v
	return s
}

func (s *GetAuditNoteProcessingStatusResponse) SetStatusCode(v int32) *GetAuditNoteProcessingStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAuditNoteProcessingStatusResponse) SetBody(v *GetAuditNoteProcessingStatusResponseBody) *GetAuditNoteProcessingStatusResponse {
	s.Body = v
	return s
}

func (s *GetAuditNoteProcessingStatusResponse) Validate() error {
	return dara.Validate(s)
}
