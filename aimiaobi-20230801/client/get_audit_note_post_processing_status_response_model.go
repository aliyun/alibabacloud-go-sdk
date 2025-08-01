// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAuditNotePostProcessingStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetAuditNotePostProcessingStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetAuditNotePostProcessingStatusResponse
	GetStatusCode() *int32
	SetBody(v *GetAuditNotePostProcessingStatusResponseBody) *GetAuditNotePostProcessingStatusResponse
	GetBody() *GetAuditNotePostProcessingStatusResponseBody
}

type GetAuditNotePostProcessingStatusResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAuditNotePostProcessingStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAuditNotePostProcessingStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s GetAuditNotePostProcessingStatusResponse) GoString() string {
	return s.String()
}

func (s *GetAuditNotePostProcessingStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetAuditNotePostProcessingStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetAuditNotePostProcessingStatusResponse) GetBody() *GetAuditNotePostProcessingStatusResponseBody {
	return s.Body
}

func (s *GetAuditNotePostProcessingStatusResponse) SetHeaders(v map[string]*string) *GetAuditNotePostProcessingStatusResponse {
	s.Headers = v
	return s
}

func (s *GetAuditNotePostProcessingStatusResponse) SetStatusCode(v int32) *GetAuditNotePostProcessingStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAuditNotePostProcessingStatusResponse) SetBody(v *GetAuditNotePostProcessingStatusResponseBody) *GetAuditNotePostProcessingStatusResponse {
	s.Body = v
	return s
}

func (s *GetAuditNotePostProcessingStatusResponse) Validate() error {
	return dara.Validate(s)
}
